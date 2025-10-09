package routes

import (
	_ "net/http"
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/controllers"
	_ "github.com/gin-gonic/gin"
)

func TestNewDomainRoute(t *testing.T) {
	type args struct {
		domainController controllers.DomainController
		domainRepo       repository.DomainRepositoryInterface
		log              pkg.LogService
		handler          infra.Router
	}
	tests := []struct {
		name string
		args args
		want DomainRoute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDomainRoute(tt.args.domainController, tt.args.domainRepo, tt.args.log, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDomainRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainRoute_Configure(t *testing.T) {
	type fields struct {
		domainController controllers.DomainController
		domainRepo       repository.DomainRepositoryInterface
		log              pkg.LogService
		handler          infra.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := DomainRoute{
				domainController: tt.fields.domainController,
				domainRepo:       tt.fields.domainRepo,
				log:              tt.fields.log,
				handler:          tt.fields.handler,
			}
			sr.Configure()
		})
	}
}
