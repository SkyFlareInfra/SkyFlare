package controllers

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/restAPi/services"
	"github.com/gin-gonic/gin"
)

func TestNewDomainController(t *testing.T) {
	type args struct {
		domainService services.DomainServicesInterface
	}
	tests := []struct {
		name string
		args args
		want DomainController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDomainController(tt.args.domainService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDomainController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainController_AddDomain(t *testing.T) {
	type fields struct {
		domainService services.DomainServicesInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := DomainController{
				domainService: tt.fields.domainService,
			}
			dc.AddDomain(tt.args.c)
		})
	}
}
