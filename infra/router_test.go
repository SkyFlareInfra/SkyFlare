package infra

import (
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/gin-gonic/gin"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		configEnv pkg.DatabaseConfig
		logger    pkg.LogService
	}
	tests := []struct {
		name string
		args args
		want Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.configEnv, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_setupEnvironment(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   *RouterBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.setupEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.setupEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_createEngine(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   *RouterBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.createEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.createEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_configureSentry(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   *RouterBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.configureSentry(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.configureSentry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_configureCORS(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   *RouterBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.configureCORS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.configureCORS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_configureSecurity(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   *RouterBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.configureSecurity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.configureSecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_registerRoutes(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   *RouterBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.registerRoutes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.registerRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterBuilder_healthCheck(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
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
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			rb.healthCheck(tt.args.c)
		})
	}
}

func TestRouterBuilder_acmeChallenge(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
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
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			rb.acmeChallenge(tt.args.c)
		})
	}
}

func TestRouterBuilder_build(t *testing.T) {
	type fields struct {
		config    pkg.DatabaseConfig
		logger    pkg.LogService
		router    *gin.Engine
		routerCfg RouterConfig
		security  SecurityHeaders
	}
	tests := []struct {
		name   string
		fields fields
		want   Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := &RouterBuilder{
				config:    tt.fields.config,
				logger:    tt.fields.logger,
				router:    tt.fields.router,
				routerCfg: tt.fields.routerCfg,
				security:  tt.fields.security,
			}
			if got := rb.build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterBuilder.build() = %v, want %v", got, tt.want)
			}
		})
	}
}
