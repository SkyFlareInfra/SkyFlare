package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Domain struct {
	ID           uint        `gorm:"primarykey" json:"ID"`
	UUID         string      `gorm:"index" json:"uuid"`
	Domain       string      `gorm:"column:domain;unique;not null" json:"Domain"`
	Content      string      `gorm:"column:content" json:"Content"`
	DNSRecords   []DNSRecord `gorm:"foreignKey:DomainID;constraint:OnDelete:CASCADE" json:"DNSRecords"`
	// Nameservers  Nameserver  `gorm:"foreignKey:DomainID;constraint:OnDelete:CASCADE" json:"Nameserver"`   // stubbed this for now (this would need to be implemented by either Ade or me)
	// Certificate  Certificate `gorm:"foreignKey:DomainID;constraint:OnDelete:CASCADE" json:"Certificate"` // stubbed this for now (this would need to be implemented by either Ade or me)
	ZoneFile     string      `gorm:"column:zoneFile" json:"ZoneFile"`
	OriginServer string      `gorm:"column:OriginServer" json:"OriginServer"`
	CreatedAt    time.Time   `json:"CreatedAt"`
	UpdatedAt    time.Time   `json:"UpdatedAt"`
	DeletedAt    *time.Time  `gorm:"index" json:"-,omitempty"`
}

func (s *Domain) BeforeCreate(tx *gorm.DB) (err error) {
	s.UUID = uuid.NewString()
	return
}
