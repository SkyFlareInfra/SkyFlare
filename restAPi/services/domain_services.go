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

const (
	failedTransferMsg = "; Transfer failed.\n"
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

func (ds *DomainService) AddDomain(c *gin.Context, payload types.DomainInput) (map[string]interface{}, *common.APIError) {
	if err := ds.validateDomainCreation(payload); err != nil {
		return nil, err
	}

	domain := &models.Domain{
		Domain:   payload.Domain,
		ZoneFile: payload.ZoneFile,
	}

	dnsRecords, apiErr := ds.processDNSRecords(payload)
	if apiErr != nil {
		return nil, apiErr
	}
	domain.DNSRecords = dnsRecords

	if payload.ZoneFile == "" {
		zoneFile, _ := ds.dnsUtility.GetZoneFile(payload.Domain)
		if strings.TrimSpace(zoneFile) != "" && zoneFile != failedTransferMsg {
			domain.ZoneFile = zoneFile
		}
	}

	dnssecStatus, err := ds.dnsLookup.CheckDNSSECStatus(payload.Domain)
	dnssecInstructions := ""
	if err != nil {
		ds.log.LogError("error checking DNSSEC status", err)
	} else if dnssecStatus {
		dnssecInstructions = "Please note that your domain is configured to use DNSSEC. To ensure proper DNS resolution, please make sure to configure DNSSEC properly at your registrar."
	}

	if err := ds.domainRepo.Create(domain); err != nil {
		ds.log.LogError("error creating domain", err)
		return nil, ds.restErr.InternalServerError(fmt.Sprintf("failed to create domain for domain: %s", payload.Domain))
	}

	scannedNameserverObj := ds.getCurrentNameservers(payload.Domain, domain.ID)

	generatedNameservers, err := ds.nameserverSrv.GenerateNameServers(domain.ID)
	if err != nil {
		ds.log.LogError("error generating nameservers", err)
		return nil, ds.restErr.InternalServerError("failed to generate nameservers")
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
			ds.log.LogInfo("domain already uses generated nameservers",
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

func (ds *DomainService) validateDomainCreation(payload types.DomainInput) *common.APIError {
	if isRestricted, err := ds.dnsLookup.IsRestrictedDomain(payload.Domain); err != nil {
		ds.log.LogError("error checking if domain is restricted", err)
		return ds.restErr.InternalServerError("failed to check domain restriction")
	} else if isRestricted {
		return ds.restErr.BadRequest(fmt.Sprintf("%s is a popular domain, you cannot create this domain", payload.Domain))
	}

	existingDomain, err := ds.domainRepo.GetDomainByName(payload.Domain)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ds.log.LogError("error fetching domain", err)
		return ds.restErr.InternalServerError(fmt.Sprintf("failed to create domain for domain: %s", payload.Domain))
	}
	if existingDomain != nil {
		return ds.restErr.BadRequest("a domain with the same domain name has already been created")
	}

	return nil
}

func (ds *DomainService) processDNSRecords(payload types.DomainInput) ([]models.DNSRecord, *common.APIError) {
	switch {
	case len(payload.DNSRecords) > 0:
		return filterOutNSRecords(ds.dnsUtility.CreateDNSRecordsFromInput(payload.DNSRecords)), nil
	case payload.ZoneFile != "":
		parsedRecords, err := ds.dnsUtility.ParseZoneFile(payload.ZoneFile)
		if err != nil {
			ds.log.LogError("error parsing zone file", err)
			return nil, ds.restErr.BadRequest("Invalid zone file format")
		}
		return filterOutNSRecords(parsedRecords), nil
	default:
		scannedRecords, scanErr := ds.dnsUtility.ScanDNSRecords(payload.Domain)
		if scanErr != nil {
			ds.log.LogError("error scanning DNS records", scanErr)
		}
		if len(scannedRecords) > 0 {
			return filterOutNSRecords(scannedRecords), nil
		}
		return []models.DNSRecord{}, nil
	}
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

func (ds *DomainService) getCurrentNameservers(domain string, domainID uint) *models.ScannedNameserver {
	nameservers, err := ds.dnsLookup.GetNameservers(domain)
	if err != nil {
		ds.log.LogError("failed to get nameservers", err)
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
