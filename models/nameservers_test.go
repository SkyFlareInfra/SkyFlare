package models

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNameserver_BeforeCreate(t *testing.T) {
	type fields struct {
		ID                uint
		UUID              string
		DomainID          *uint
		GeneratedNSUUID   string
		NS1               string
		NS2               string
		IsVerified        bool
		VerificationToken string
		CreatedAt         time.Time
		UpdatedAt         time.Time
		DeletedAt         gorm.DeletedAt
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
			s := &Nameserver{
				ID:                tt.fields.ID,
				UUID:              tt.fields.UUID,
				DomainID:          tt.fields.DomainID,
				GeneratedNSUUID:   tt.fields.GeneratedNSUUID,
				NS1:               tt.fields.NS1,
				NS2:               tt.fields.NS2,
				IsVerified:        tt.fields.IsVerified,
				VerificationToken: tt.fields.VerificationToken,
				CreatedAt:         tt.fields.CreatedAt,
				UpdatedAt:         tt.fields.UpdatedAt,
				DeletedAt:         tt.fields.DeletedAt,
			}
			if err := s.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Nameserver.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
