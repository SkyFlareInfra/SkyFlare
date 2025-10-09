package routes

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
)

func TestNewRouteBuilder(t *testing.T) {
	type args struct {
		logger pkg.LogService
		router infra.Router
	}
	tests := []struct {
		name string
		args args
		want *RouteBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouteBuilder(tt.args.logger, tt.args.router); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouteBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRegistry(t *testing.T) {
	type args struct {
		siteRoute DomainRoute
	}
	tests := []struct {
		name string
		args args
		want RouteRegistry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRegistry(tt.args.siteRoute); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteRegistry_SetupRoutes(t *testing.T) {
	tests := []struct {
		name string
		rr   RouteRegistry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rr.SetupRoutes()
		})
	}
}
