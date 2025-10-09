package middlewares

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/gin-gonic/gin"
)

func TestNewDynamicProxy(t *testing.T) {
	type args struct {
		domainRepo repository.DomainRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want DynamicProxy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDynamicProxy(tt.args.domainRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDynamicProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicProxy_DynamicProxy(t *testing.T) {
	type fields struct {
		domainRepo repository.DomainRepositoryInterface
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := DynamicProxy{
				domainRepo: tt.fields.domainRepo,
			}
			if got := dp.DynamicProxy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynamicProxy.DynamicProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
