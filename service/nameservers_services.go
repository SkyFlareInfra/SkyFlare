package service

import (
	"fmt"
	"sync"

	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/client"
	"github.com/google/uuid"
)

type NameserverServiceInterface interface {
	GenerateNameServers(domainID uint) (*models.GeneratedNameserver, error)
}

type NameserverService struct {
	nameserverRepo repository.NameserverInterface
	domainRepo     repository.DomainRepositoryInterface
	dnsLookup      client.DNSLookUpInterface
	log            pkg.LogService
	config         pkg.DatabaseConfig
	cache          *sync.Map
}

func NewNameserverService(
	nameserverRepo repository.NameserverInterface,
	domainRepo repository.DomainRepositoryInterface,
	dnsLookup client.DNSLookUpInterface,
	log pkg.LogService,
	config pkg.DatabaseConfig,
) NameserverServiceInterface {
	return &NameserverService{
		nameserverRepo: nameserverRepo,
		domainRepo:     domainRepo,
		dnsLookup:      dnsLookup,
		log:            log,
		config:         config,
		cache:          &sync.Map{},
	}
}

func (ns *NameserverService) GenerateNameServers(domainID uint) (*models.GeneratedNameserver, error) {
	if ns.config.NS1 == "" || ns.config.NS2 == "" {
		return nil, fmt.Errorf("nameserver configuration is incomplete")
	}

	generatedNS := &models.GeneratedNameserver{
		DomainID: domainID,
		UUID:     uuid.NewString(),
		NS1:      ns.config.NS1,
		NS2:      ns.config.NS2,
	}

	if err := ns.nameserverRepo.CreateGeneratedNameserver(generatedNS); err != nil {
		ns.log.LogError("failed to create nameserver record", err)
		return nil, fmt.Errorf("failed to generate nameservers: %w", err)
	}

	return generatedNS, nil
}
