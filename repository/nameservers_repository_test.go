package repository

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/models"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
)

func TestNewNameserverRepository(t *testing.T) {
	type args struct {
		db  infra.DatabaseManager
		log pkg.LogService
	}
	tests := []struct {
		name string
		args args
		want NameserverInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNameserverRepository(tt.args.db, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNameserverRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameserver_CreateGeneratedNameserver(t *testing.T) {
	type fields struct {
		db  infra.DatabaseManager
		log pkg.LogService
	}
	type args struct {
		generatedNS *models.GeneratedNameserver
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
			ns := &Nameserver{
				db:  tt.fields.db,
				log: tt.fields.log,
			}
			if err := ns.CreateGeneratedNameserver(tt.args.generatedNS); (err != nil) != tt.wantErr {
				t.Errorf("Nameserver.CreateGeneratedNameserver() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
