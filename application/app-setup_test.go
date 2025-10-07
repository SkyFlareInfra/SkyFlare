package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	_ "github.com/getsentry/sentry-go"
	"go.uber.org/fx"
)

func TestNewApp(t *testing.T) {
	type args struct {
		lifecycle fx.Lifecycle
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApp(tt.args.lifecycle, tt.args.log, tt.args.configEnv, tt.args.database, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_Start(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	type args struct {
		ctx context.Context
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
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			if err := a.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("App.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_Stop(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	type args struct {
		ctx context.Context
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
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			if err := a.Stop(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("App.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_validateConfig(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			if err := a.validateConfig(); (err != nil) != tt.wantErr {
				t.Errorf("App.validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_printStartupBanner(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			a.printStartupBanner()
		})
	}
}

func TestApp_setupDatabase(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			if err := a.setupDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("App.setupDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_startServerAsync(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			a.startServerAsync()
		})
	}
}

func TestApp_closeDatabase(t *testing.T) {
	type fields struct {
		log       pkg.LogService
		configEnv pkg.DatabaseConfig
		database  infra.DatabaseManager
		handler   infra.Router
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				log:       tt.fields.log,
				configEnv: tt.fields.configEnv,
				database:  tt.fields.database,
				handler:   tt.fields.handler,
			}
			if err := a.closeDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("App.closeDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
