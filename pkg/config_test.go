package pkg

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string
		want DatabaseConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfigManager(t *testing.T) {
	tests := []struct {
		name string
		want *ConfigManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_LoadConfiguration(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   *DatabaseConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			if got := cm.LoadConfiguration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigManager.LoadConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_GetConfiguration(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   *DatabaseConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			if got := cm.GetConfiguration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigManager.GetConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_getEnvironment(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			if got := cm.getEnvironment(); got != tt.want {
				t.Errorf("ConfigManager.getEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_isLocalEnvironment(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	type args struct {
		env string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			if got := cm.isLocalEnvironment(tt.args.env); got != tt.want {
				t.Errorf("ConfigManager.isLocalEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_loadLocalConfig(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   *DatabaseConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			if got := cm.loadLocalConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigManager.loadLocalConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_loadProductionConfig(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   *DatabaseConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			if got := cm.loadProductionConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigManager.loadProductionConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigManager_ValidateAndFreePort(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	type args struct {
		port int
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
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			cm.ValidateAndFreePort(tt.args.port)
		})
	}
}

func TestConfigManager_terminateProcessOnPort(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	type args struct {
		port int
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
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			cm.terminateProcessOnPort(tt.args.port)
		})
	}
}

func TestConfigManager_verifyPortAvailable(t *testing.T) {
	type fields struct {
		config *DatabaseConfig
	}
	type args struct {
		port int
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
			cm := &ConfigManager{
				config: tt.fields.config,
			}
			cm.verifyPortAvailable(tt.args.port)
		})
	}
}

func TestDatabaseConfig_CheckAndKillProcess(t *testing.T) {
	type fields struct {
		Debug       string
		ServerPort  string
		Environment string
		DBUsername  string
		DBPassword  string
		DBSslMode   string
		DBHost      string
		DBPort      string
		DBName      string
		DatabaseUrl string
		NS1         string
		NS2         string
		StartSeeder bool
		ProxyIPv4   string
		ProxyIPv6   string
	}
	type args struct {
		port int
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
			c := &DatabaseConfig{
				Debug:       tt.fields.Debug,
				ServerPort:  tt.fields.ServerPort,
				Environment: tt.fields.Environment,
				DBUsername:  tt.fields.DBUsername,
				DBPassword:  tt.fields.DBPassword,
				DBSslMode:   tt.fields.DBSslMode,
				DBHost:      tt.fields.DBHost,
				DBPort:      tt.fields.DBPort,
				DBName:      tt.fields.DBName,
				DatabaseUrl: tt.fields.DatabaseUrl,
				NS1:         tt.fields.NS1,
				NS2:         tt.fields.NS2,
				StartSeeder: tt.fields.StartSeeder,
				ProxyIPv4:   tt.fields.ProxyIPv4,
				ProxyIPv6:   tt.fields.ProxyIPv6,
			}
			c.CheckAndKillProcess(tt.args.port)
		})
	}
}

func TestGetConf(t *testing.T) {
	tests := []struct {
		name string
		want *DatabaseConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetConf(t *testing.T) {
	type args struct {
		config *DatabaseConfig
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetConf(tt.args.config)
		})
	}
}
