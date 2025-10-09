package utility

import (
	"reflect"
	"sync"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/types"
	"github.com/jellydator/ttlcache/v3"
	"github.com/miekg/dns"
)

func TestNewDNSLookUp(t *testing.T) {
	type args struct {
		log    pkg.LogService
		config pkg.DatabaseConfig
	}
	tests := []struct {
		name string
		args args
		want DNSUtilityInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDNSLookUp(tt.args.log, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDNSLookUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSUtility_ParseZoneFile(t *testing.T) {
	type fields struct {
		log     pkg.LogService
		cache   *ttlcache.Cache[string, []models.DNSRecord]
		errorMu sync.Mutex
		config  pkg.DatabaseConfig
	}
	type args struct {
		zoneFile string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.DNSRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := &DNSUtility{
				log:     tt.fields.log,
				cache:   tt.fields.cache,
				errorMu: tt.fields.errorMu,
				config:  tt.fields.config,
			}
			got, err := du.ParseZoneFile(tt.args.zoneFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSUtility.ParseZoneFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DNSUtility.ParseZoneFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSUtility_ScanDNSRecords(t *testing.T) {
	type fields struct {
		log     pkg.LogService
		cache   *ttlcache.Cache[string, []models.DNSRecord]
		errorMu sync.Mutex
		config  pkg.DatabaseConfig
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.DNSRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := &DNSUtility{
				log:     tt.fields.log,
				cache:   tt.fields.cache,
				errorMu: tt.fields.errorMu,
				config:  tt.fields.config,
			}
			got, err := du.ScanDNSRecords(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSUtility.ScanDNSRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DNSUtility.ScanDNSRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDNSAnswers(t *testing.T) {
	type args struct {
		domain  string
		answers []dns.RR
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
			if got := parseDNSAnswers(tt.args.domain, tt.args.answers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDNSAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deduplicateRecords(t *testing.T) {
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
			if got := deduplicateRecords(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deduplicateRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSUtility_CreateDNSRecordsFromInput(t *testing.T) {
	type fields struct {
		log     pkg.LogService
		cache   *ttlcache.Cache[string, []models.DNSRecord]
		errorMu sync.Mutex
		config  pkg.DatabaseConfig
	}
	type args struct {
		inputs []types.DNSRecordInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.DNSRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := &DNSUtility{
				log:     tt.fields.log,
				cache:   tt.fields.cache,
				errorMu: tt.fields.errorMu,
				config:  tt.fields.config,
			}
			if got := du.CreateDNSRecordsFromInput(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DNSUtility.CreateDNSRecordsFromInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSUtility_GetZoneFile(t *testing.T) {
	type fields struct {
		log     pkg.LogService
		cache   *ttlcache.Cache[string, []models.DNSRecord]
		errorMu sync.Mutex
		config  pkg.DatabaseConfig
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := &DNSUtility{
				log:     tt.fields.log,
				cache:   tt.fields.cache,
				errorMu: tt.fields.errorMu,
				config:  tt.fields.config,
			}
			got, err := du.GetZoneFile(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSUtility.GetZoneFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DNSUtility.GetZoneFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNSUtility_GetDNSResponse(t *testing.T) {
	type fields struct {
		log     pkg.LogService
		cache   *ttlcache.Cache[string, []models.DNSRecord]
		errorMu sync.Mutex
		config  pkg.DatabaseConfig
	}
	type args struct {
		domain     string
		recordType uint16
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.DNSRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := &DNSUtility{
				log:     tt.fields.log,
				cache:   tt.fields.cache,
				errorMu: tt.fields.errorMu,
				config:  tt.fields.config,
			}
			got, err := du.GetDNSResponse(tt.args.domain, tt.args.recordType)
			if (err != nil) != tt.wantErr {
				t.Errorf("DNSUtility.GetDNSResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DNSUtility.GetDNSResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
