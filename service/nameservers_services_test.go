package service

import (
	"reflect"
	"sync"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/client"
)

func TestNewNameserverService(t *testing.T) {
	type args struct {
		nameserverRepo repository.NameserverInterface
		domainRepo     repository.DomainRepositoryInterface
		dnsLookup      client.DNSLookUpInterface
		log            pkg.LogService
		config         pkg.DatabaseConfig
	}
	tests := []struct {
		name string
		args args
		want NameserverServiceInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNameserverService(tt.args.nameserverRepo, tt.args.domainRepo, tt.args.dnsLookup, tt.args.log, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNameserverService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameserverService_GenerateNameServers(t *testing.T) {
	type fields struct {
		nameserverRepo repository.NameserverInterface
		domainRepo     repository.DomainRepositoryInterface
		dnsLookup      client.DNSLookUpInterface
		log            pkg.LogService
		config         pkg.DatabaseConfig
		cache          *sync.Map
	}
	type args struct {
		domainID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.GeneratedNameserver
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &NameserverService{
				nameserverRepo: tt.fields.nameserverRepo,
				domainRepo:     tt.fields.domainRepo,
				dnsLookup:      tt.fields.dnsLookup,
				log:            tt.fields.log,
				config:         tt.fields.config,
				cache:          tt.fields.cache,
			}
			got, err := ns.GenerateNameServers(tt.args.domainID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameserverService.GenerateNameServers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NameserverService.GenerateNameServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
