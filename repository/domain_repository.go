package repository

import (
	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/models"
	"gorm.io/gorm"
)

type DomainRepositoryInterface interface {
	Create(site *models.Domain) error
	GetAll(filter models.Domain) ([]models.Domain, error)
	GetDomainByName(domain string) (*models.Domain, error)
	GetByUID(uid string) (*models.Domain, error)
	GetByID(id uint) (*models.Domain, error)
	Get(filter *models.Domain) (*models.Domain, error)
	Update(filter models.Domain, data *models.Domain) error
	Delete(id uint) error
	DeleteByUID(uid string) error
	CreateInTransaction(tx *gorm.DB, site *models.Domain) error
	GetDomainNameByDomainID(domainID uint) (string, error)
}

type DomainRepository struct {
	db infra.DatabaseManager
}

func NewDomainRepository(
	db infra.DatabaseManager,
) DomainRepositoryInterface {
	return &DomainRepository{
		db: db,
	}
}

func (ss *DomainRepository) CreateInTransaction(tx *gorm.DB, site *models.Domain) error {
	result := tx.Create(site)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sr *DomainRepository) Create(site *models.Domain) error {
	result := sr.db.DB.Create(&site)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sr *DomainRepository) GetAll(filter models.Domain) ([]models.Domain, error) {
	var domain []models.Domain
	if err := sr.db.DB.Where(filter).Find(&domain).Error; err != nil {
		return nil, err
	}

	return domain, nil
}

func (sr *DomainRepository) GetDomainByName(name string) (*models.Domain, error) {
	var domain models.Domain
	if err := sr.db.DB.Where("domain = ?", name).First(&domain).Error; err != nil {
		return nil, err
	}
	return &domain, nil
}

func (sr *DomainRepository) GetByUID(uuid string) (*models.Domain, error) {
	var domain models.Domain

	err := sr.db.DB.
		Preload("DNSRecords").
		Preload("Nameservers").
		Preload("Certificate").
		Where("uuid = ?", uuid).
		First(&domain).
		Error

	if err != nil {
		return nil, err
	}
	return &domain, nil
}

func (sr *DomainRepository) GetByID(id uint) (*models.Domain, error) {
	var domain models.Domain
	if err := sr.db.DB.Where("id = ?", id).First(&domain).Error; err != nil {
		return nil, err
	}

	return &domain, nil
}

func (sr *DomainRepository) Get(filter *models.Domain) (*models.Domain, error) {
	var domain *models.Domain

	result := sr.db.DB.Where(filter).Last(&domain)
	if result.Error != nil {
		return nil, result.Error
	}

	return domain, nil
}

func (sr *DomainRepository) Update(filter models.Domain, data *models.Domain) error {
	return sr.db.DB.Where(&filter).Updates(&data).Error
}

func (sr *DomainRepository) Delete(id uint) error {
	var domain models.Domain
	if err := sr.db.DB.Where("id = ?", id).First(&domain).Error; err != nil {
		return err
	}

	if err := sr.db.DB.Delete(&domain).Error; err != nil {
		return err
	}

	return nil
}

func (sr *DomainRepository) DeleteByUID(uid string) error {
	return sr.db.DB.Transaction(func(tx *gorm.DB) error {
		var domain models.Domain
		if err := tx.Where("uuid = ?", uid).First(&domain).Error; err != nil {
			return err
		}

		return tx.Delete(&domain).Error
	})
}

func (dr *DomainRepository) GetDomainNameByDomainID(domainID uint) (string, error) {
	var domain models.Domain
	err := dr.db.DB.Model(&models.Domain{}).Where("id = ?", domainID).First(&domain).Error
	if err != nil {
		return "", err
	}

	return domain.Domain, nil
}
