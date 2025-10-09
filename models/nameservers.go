package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneratedNameserver struct {
	UUID      string         `gorm:"index;unique;not null" json:"uid"`
	DomainID  uint           `gorm:"index"`
	NS1       string         `gorm:"size:253;not null"`
	NS2       string         `gorm:"size:253;not null"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Nameserver struct {
	ID                uint           `gorm:"primarykey" json:"ID"`
	UUID              string         `gorm:"index;null" json:"uid"`
	DomainID          *uint          `gorm:"index;not null" json:"DomainID"`
	GeneratedNSUUID   string         `gorm:"index" json:"generated_ns_uid"`
	NS1               string         `gorm:"column:ns1" json:"ns1"`
	NS2               string         `gorm:"column:ns2" json:"ns2"`
	IsVerified        bool           `gorm:"column:is_verified" json:"is_verified"`
	VerificationToken string         `gorm:"column:verification_token" json:"-"`
	CreatedAt         time.Time      `json:"CreatedAt"`
	UpdatedAt         time.Time      `json:"UpdatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

type ScannedNameserver struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	DomainID    *uint     `gorm:"index;not null" json:"DomainID"`
	UUID        string    `gorm:"index;null" json:"uid"`
	Nameservers []string  `gorm:"type:json" json:"nameservers"`
	PrimaryNS   string    `gorm:"-" json:"primary_ns,omitempty"`
	SecondaryNS string    `gorm:"-" json:"secondary_ns,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (s *Nameserver) BeforeCreate(tx *gorm.DB) (err error) {
	s.UUID = uuid.NewString()
	return
}
