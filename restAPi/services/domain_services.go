package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/SkyFlareInfra/SkyFlare/common"
	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/client"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/types"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/utility"
	"github.com/SkyFlareInfra/SkyFlare/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DomainServicesInterface interface {
	AddDomain(c *gin.Context, payload types.DomainInput) (map[string]interface{}, *common.APIError)
}

type DomainService struct {
	domainRepo    repository.DomainRepositoryInterface
	dnsLookup     client.DNSLookUpInterface
	dnsUtility    utility.DNSUtilityInterface
	nameserverSrv service.NameserverServiceInterface
	log           pkg.LogService
	config        pkg.DatabaseConfig
	restErr       *common.ErrorBuilder
}

func NewDomainService(
	domainRepo repository.DomainRepositoryInterface,
	dnsLookup client.DNSLookUpInterface,
	dnsUtility utility.DNSUtilityInterface,
	nameserverSrv service.NameserverServiceInterface,
	log pkg.LogService,
	config pkg.DatabaseConfig,
	restErr *common.ErrorBuilder,
) DomainServicesInterface {
	return &DomainService{
		domainRepo:    domainRepo,
		dnsLookup:     dnsLookup,
		dnsUtility:    dnsUtility,
		nameserverSrv: nameserverSrv,
		log:           log,
		config:        config,
		restErr:       restErr,
	}
}

func (ss *DomainService) AddDomain(c *gin.Context, payload types.DomainInput) (map[string]interface{}, *common.APIError) {
	if isRestricted, err := ss.dnsLookup.IsRestrictedDomain(payload.Domain); err != nil {
		ss.log.LogError("error checking if domain is restricted", err)
		return nil, ss.restErr.InternalServerError("failed to check domain restriction")
	} else if isRestricted {
		return nil, ss.restErr.BadRequest(fmt.Sprintf("%s is a popular domain, you cannot create this site", payload.Domain))
	}

	existingDomain, err := ss.domainRepo.GetDomainByName(payload.Domain)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ss.log.LogError("error fetching domain", err)
		return nil, ss.restErr.InternalServerError("failed to create site")
	}
	if existingDomain != nil {
		return nil, ss.restErr.BadRequest("a site with the same domain name has already been created")
	}

	domain := &models.Domain{
		Domain:   payload.Domain,
		ZoneFile: payload.ZoneFile,
	}

	if len(payload.DNSRecords) > 0 {
		domain.DNSRecords = filterOutNSRecords(ss.dnsUtility.CreateDNSRecordsFromInput(payload.DNSRecords))
	} else if payload.ZoneFile != "" {
		parsedRecords, parseErr := ss.dnsUtility.ParseZoneFile(payload.ZoneFile)
		if parseErr != nil {
			ss.log.LogError("error parsing zone file", parseErr)
			return nil, ss.restErr.BadRequest("Invalid zone file format")
		}
		domain.DNSRecords = filterOutNSRecords(parsedRecords)
	} else {
		scannedRecords, scanErr := ss.dnsUtility.ScanDNSRecords(payload.Domain)
		if scanErr != nil {
			ss.log.LogError("error scanning DNS records", scanErr)
		}
		if len(scannedRecords) > 0 {
			domain.DNSRecords = filterOutNSRecords(scannedRecords)
		}
	}

	// scannedNameserverObj := ss.getCurrentNameservers(payload.Domain)
	scannedNameserverObj := ss.getCurrentNameservers(payload.Domain, domain.ID)

	if payload.ZoneFile == "" {
		zoneFile, _ := ss.dnsUtility.GetZoneFile(payload.Domain)
		if strings.TrimSpace(zoneFile) != "" && zoneFile != "; Transfer failed.\n" {
			domain.ZoneFile = zoneFile
		}
	}

	dnssecStatus, err := ss.dnsLookup.CheckDNSSECStatus(payload.Domain)
	dnssecInstructions := ""
	if err != nil {
		ss.log.LogError("error checking DNSSEC status", err)
	} else if dnssecStatus {
		dnssecInstructions = "Please note that your domain is configured to use DNSSEC. To ensure proper DNS resolution, please make sure to configure DNSSEC properly at your registrar."
	}

	if err := ss.domainRepo.Create(domain); err != nil {
		ss.log.LogError("error creating site", err)
		return nil, ss.restErr.InternalServerError("failed to create site")
	}

	generatedNameservers, err := ss.nameserverSrv.GenerateNameServers(domain.ID)
	if err != nil {
		ss.log.LogError("error generating nameservers", err)
		return nil, ss.restErr.InternalServerError("failed to generate nameservers")
	}

	if scannedNameserverObj != nil && len(scannedNameserverObj.Nameservers) > 0 {
		generatedNS := []string{
			generatedNameservers.NS1,
			generatedNameservers.NS2,
		}
		scannedSet := make(map[string]struct{})
		for _, ns := range scannedNameserverObj.Nameservers {
			scannedSet[ns] = struct{}{}
		}

		match := true
		for _, ns := range generatedNS {
			if _, exists := scannedSet[ns]; !exists {
				match = false
				break
			}
		}

		if match {
			ss.log.LogInfo("domain already uses generated nameservers",
				zap.String("domain", payload.Domain))
		}
	}

	return map[string]interface{}{
		"message":              "Please update your nameservers at your registrar",
		"domain":               domain,
		"generatedNameServers": generatedNameservers,
		"scannedNameServers":   scannedNameserverObj,
		"dnssecInstructions":   dnssecInstructions,
	}, nil
}

func filterOutNSRecords(records []models.DNSRecord) []models.DNSRecord {
	var filtered []models.DNSRecord
	for _, r := range records {
		if r.Type != "NS" {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

func (ss *DomainService) getCurrentNameservers(domain string, domainID uint) *models.ScannedNameserver {
	nameservers, err := ss.dnsLookup.GetNameservers(domain)
	if err != nil {
		ss.log.LogError("failed to get nameservers", err)
		return &models.ScannedNameserver{
			DomainID: &domainID,
			UUID:     uuid.New().String(),
		}
	}

	if len(nameservers) >= 2 {
		return &models.ScannedNameserver{
			DomainID:    &domainID,
			UUID:        uuid.New().String(),
			Nameservers: nameservers,
			PrimaryNS:   nameservers[0],
			SecondaryNS: nameservers[1],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}
	return &models.ScannedNameserver{
		ID:          0,
		DomainID:    &domainID,
		UUID:        uuid.New().String(),
		Nameservers: nameservers,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
