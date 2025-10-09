package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DNSRecord struct {
	ID           uint                   `gorm:"primarykey" json:"ID"`
	UUID         string                 `gorm:"index;null" json:"uid"`
	DomainID     uint                   `gorm:"index;not null" json:"DomainID"`
	Type         string                 `gorm:"type:varchar(10);not null" json:"Type"`   // e.g., A, AAAA, CNAME
	Name         string                 `gorm:"type:varchar(255);not null" json:"Name"`  // Subdomain or root ('@')
	Value        string                 `gorm:"type:varchar(255);not null" json:"Value"` // IP address, domain, or content
	IP           string                 `gorm:"type:varchar(255);" json:"IP,omitempty"`  // Only populated for A/AAAA records
	ProxyEnabled bool                   `gorm:"default:false" json:"proxy_enabled"`
	ProxyIP      string                 `gorm:"type:varchar(45)" json:"proxy_ip,omitempty"`  // Cloudflare-like anycast IP
	OriginIP     string                 `gorm:"type:varchar(45)" json:"origin_ip,omitempty"` // For A/AAAA records
	MetaData     map[string]interface{} `gorm:"type:jsonb" json:"metadata,omitempty"`
	TTL          int                    `gorm:"type:int" json:"TTL"`                // Time to Live
	Priority     int                    `gorm:"type:int" json:"Priority,omitempty"` // MX record priority
	Weight       int                    `gorm:"type:int" json:"Weight,omitempty"`   // SRV record weight
	Port         int                    `gorm:"type:int" json:"Port,omitempty"`     // SRV record port
	SerialNumber int                    `gorm:"type:int" json:"SerialNumber,omitempty"`
	RefreshTime  int                    `gorm:"type:int" json:"RefreshTime,omitempty"`
	RetryTime    int                    `gorm:"type:int" json:"RetryTime,omitempty"`
	ExpireTime   int                    `gorm:"type:int" json:"ExpireTime,omitempty"`
	MinimumTTL   int                    `gorm:"type:int" json:"MinimumTTL,omitempty"`
	Flag         int16                  `gorm:"type:smallint" json:"Flag,omitempty"`
	Flags        string                 `gorm:"type:varchar(255);" json:"Flags,omitempty"`
	Tag          string                 `gorm:"type:varchar(255);" json:"Tag,omitempty"`
	Usage        int16                  `gorm:"type:smallint" json:"Usage,omitempty"`
	Selector     int16                  `gorm:"type:smallint" json:"Selector,omitempty"`
	MatchingType int16                  `gorm:"type:smallint" json:"MatchingType,omitempty"`
	Certificate  string                 `gorm:"type:varchar(255);" json:"Certificate,omitempty"`
	Order        int16                  `gorm:"type:smallint" json:"Order,omitempty"`
	Preference   int16                  `gorm:"type:smallint" json:"Preference,omitempty"`
	Service      string                 `gorm:"type:varchar(255);" json:"Service,omitempty"`
	Regexp       string                 `gorm:"type:varchar(255);" json:"Regexp,omitempty"`
	Replacement  string                 `gorm:"type:varchar(255);" json:"Replacement,omitempty"`
	KeyTag       int16                  `gorm:"type:smallint" json:"KeyTag,omitempty"`
	Algorithm    int16                  `gorm:"type:smallint" json:"Algorithm,omitempty"`
	DigestType   int16                  `gorm:"type:smallint" json:"DigestType,omitempty"`
	Digest       string                 `gorm:"type:varchar(255);" json:"Digest,omitempty"`
	CreatedAt    time.Time              `json:"CreatedAt"`
	UpdatedAt    time.Time              `json:"UpdatedAt"`
	DeletedAt    gorm.DeletedAt         `gorm:"index" json:"-"`
}

func (dr *DNSRecord) BeforeCreate(tx *gorm.DB) (err error) {
	dr.UUID = uuid.NewString()
	return
}
