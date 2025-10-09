package repository

import (
	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
)

type NameserverInterface interface {
	CreateGeneratedNameserver(ns *models.GeneratedNameserver) error
}

type Nameserver struct {
	db  infra.DatabaseManager
	log pkg.LogService
}

func NewNameserverRepository(
	db infra.DatabaseManager,
	log pkg.LogService,
) NameserverInterface {
	return &Nameserver{
		db:     db,
		log: log,
	}
}

func (ns *Nameserver) CreateGeneratedNameserver(generatedNS *models.GeneratedNameserver) error {
	if err := ns.db.DB.Create(generatedNS).Error; err != nil {
		return err
	}
	return nil
}
