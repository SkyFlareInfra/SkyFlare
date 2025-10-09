package models

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestDNSRecord_BeforeCreate(t *testing.T) {
	type fields struct {
		ID           uint
		UUID         string
		DomainID     uint
		Type         string
		Name         string
		Value        string
		IP           string
		ProxyEnabled bool
		ProxyIP      string
		OriginIP     string
		MetaData     map[string]interface{}
		TTL          int
		Priority     int
		Weight       int
		Port         int
		SerialNumber int
		RefreshTime  int
		RetryTime    int
		ExpireTime   int
		MinimumTTL   int
		Flag         int16
		Flags        string
		Tag          string
		Usage        int16
		Selector     int16
		MatchingType int16
		Certificate  string
		Order        int16
		Preference   int16
		Service      string
		Regexp       string
		Replacement  string
		KeyTag       int16
		Algorithm    int16
		DigestType   int16
		Digest       string
		CreatedAt    time.Time
		UpdatedAt    time.Time
		DeletedAt    gorm.DeletedAt
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := &DNSRecord{
				ID:           tt.fields.ID,
				UUID:         tt.fields.UUID,
				DomainID:     tt.fields.DomainID,
				Type:         tt.fields.Type,
				Name:         tt.fields.Name,
				Value:        tt.fields.Value,
				IP:           tt.fields.IP,
				ProxyEnabled: tt.fields.ProxyEnabled,
				ProxyIP:      tt.fields.ProxyIP,
				OriginIP:     tt.fields.OriginIP,
				MetaData:     tt.fields.MetaData,
				TTL:          tt.fields.TTL,
				Priority:     tt.fields.Priority,
				Weight:       tt.fields.Weight,
				Port:         tt.fields.Port,
				SerialNumber: tt.fields.SerialNumber,
				RefreshTime:  tt.fields.RefreshTime,
				RetryTime:    tt.fields.RetryTime,
				ExpireTime:   tt.fields.ExpireTime,
				MinimumTTL:   tt.fields.MinimumTTL,
				Flag:         tt.fields.Flag,
				Flags:        tt.fields.Flags,
				Tag:          tt.fields.Tag,
				Usage:        tt.fields.Usage,
				Selector:     tt.fields.Selector,
				MatchingType: tt.fields.MatchingType,
				Certificate:  tt.fields.Certificate,
				Order:        tt.fields.Order,
				Preference:   tt.fields.Preference,
				Service:      tt.fields.Service,
				Regexp:       tt.fields.Regexp,
				Replacement:  tt.fields.Replacement,
				KeyTag:       tt.fields.KeyTag,
				Algorithm:    tt.fields.Algorithm,
				DigestType:   tt.fields.DigestType,
				Digest:       tt.fields.Digest,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				DeletedAt:    tt.fields.DeletedAt,
			}
			if err := dr.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("DNSRecord.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
