package models

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestDomain_BeforeCreate(t *testing.T) {
	type fields struct {
		ID           uint
		UUID         string
		Domain       string
		Content      string
		DNSRecords   []DNSRecord
		ZoneFile     string
		OriginServer string
		CreatedAt    time.Time
		UpdatedAt    time.Time
		DeletedAt    *time.Time
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
			s := &Domain{
				ID:           tt.fields.ID,
				UUID:         tt.fields.UUID,
				Domain:       tt.fields.Domain,
				Content:      tt.fields.Content,
				DNSRecords:   tt.fields.DNSRecords,
				ZoneFile:     tt.fields.ZoneFile,
				OriginServer: tt.fields.OriginServer,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				DeletedAt:    tt.fields.DeletedAt,
			}
			if err := s.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Domain.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
