package services

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/common"
	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/client"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/types"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/utility"
	"github.com/SkyFlareInfra/SkyFlare/service"
	"github.com/gin-gonic/gin"
)

func TestNewDomainService(t *testing.T) {
	type args struct {
		domainRepo    repository.DomainRepositoryInterface
		dnsLookup     client.DNSLookUpInterface
		dnsUtility    utility.DNSUtilityInterface
		nameserverSrv service.NameserverServiceInterface
		log           pkg.LogService
		config        pkg.DatabaseConfig
		restErr       *common.ErrorBuilder
	}
	tests := []struct {
		name string
		args args
		want DomainServicesInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDomainService(tt.args.domainRepo, tt.args.dnsLookup, tt.args.dnsUtility, tt.args.nameserverSrv, tt.args.log, tt.args.config, tt.args.restErr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDomainService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainService_AddDomain(t *testing.T) {
	type fields struct {
		domainRepo    repository.DomainRepositoryInterface
		dnsLookup     client.DNSLookUpInterface
		dnsUtility    utility.DNSUtilityInterface
		nameserverSrv service.NameserverServiceInterface
		log           pkg.LogService
		config        pkg.DatabaseConfig
		restErr       *common.ErrorBuilder
	}
	type args struct {
		c       *gin.Context
		payload types.DomainInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
		want1  *common.APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &DomainService{
				domainRepo:    tt.fields.domainRepo,
				dnsLookup:     tt.fields.dnsLookup,
				dnsUtility:    tt.fields.dnsUtility,
				nameserverSrv: tt.fields.nameserverSrv,
				log:           tt.fields.log,
				config:        tt.fields.config,
				restErr:       tt.fields.restErr,
			}
			got, got1 := ss.AddDomain(tt.args.c, tt.args.payload)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainService.AddDomain() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DomainService.AddDomain() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_filterOutNSRecords(t *testing.T) {
	type args struct {
		records []models.DNSRecord
	}
	tests := []struct {
		name string
		args args
		want []models.DNSRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterOutNSRecords(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOutNSRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainService_getCurrentNameservers(t *testing.T) {
	type fields struct {
		domainRepo    repository.DomainRepositoryInterface
		dnsLookup     client.DNSLookUpInterface
		dnsUtility    utility.DNSUtilityInterface
		nameserverSrv service.NameserverServiceInterface
		log           pkg.LogService
		config        pkg.DatabaseConfig
		restErr       *common.ErrorBuilder
	}
	type args struct {
		domain   string
		domainID uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.ScannedNameserver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &DomainService{
				domainRepo:    tt.fields.domainRepo,
				dnsLookup:     tt.fields.dnsLookup,
				dnsUtility:    tt.fields.dnsUtility,
				nameserverSrv: tt.fields.nameserverSrv,
				log:           tt.fields.log,
				config:        tt.fields.config,
				restErr:       tt.fields.restErr,
			}
			if got := ss.getCurrentNameservers(tt.args.domain, tt.args.domainID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainService.getCurrentNameservers() = %v, want %v", got, tt.want)
			}
		})
	}
}
