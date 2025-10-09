package client

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/pkg"
)

func TestNewDNSLookUp(t *testing.T) {
	type args struct {
		config pkg.DatabaseConfig
		log    pkg.LogService
	}
	tests := []struct {
		name string
		args args
		want DNSLookUpInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDNSLookUp(tt.args.config, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDNSLookUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSLookUp_IsRestrictedDomain(t *testing.T) {
	type fields struct {
		config pkg.DatabaseConfig
		log    pkg.LogService
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := DNSLookUp{
				config: tt.fields.config,
				log:    tt.fields.log,
			}
			got, err := dl.IsRestrictedDomain(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSLookUp.IsRestrictedDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DNSLookUp.IsRestrictedDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSLookUp_IsPopularDomain(t *testing.T) {
	type fields struct {
		config pkg.DatabaseConfig
		log    pkg.LogService
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := DNSLookUp{
				config: tt.fields.config,
				log:    tt.fields.log,
			}
			if got := dl.IsPopularDomain(tt.args.domain); got != tt.want {
				t.Errorf("DNSLookUp.IsPopularDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSLookUp_CheckDNSSECStatus(t *testing.T) {
	type fields struct {
		config pkg.DatabaseConfig
		log    pkg.LogService
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := DNSLookUp{
				config: tt.fields.config,
				log:    tt.fields.log,
			}
			got, err := dl.CheckDNSSECStatus(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSLookUp.CheckDNSSECStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DNSLookUp.CheckDNSSECStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSLookUp_GetNameservers(t *testing.T) {
	type fields struct {
		config pkg.DatabaseConfig
		log    pkg.LogService
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := DNSLookUp{
				config: tt.fields.config,
				log:    tt.fields.log,
			}
			got, err := dl.GetNameservers(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSLookUp.GetNameservers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DNSLookUp.GetNameservers() = %v, want %v", got, tt.want)
			}
		})
	}
}
