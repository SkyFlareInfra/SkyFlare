package pkg

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Debug       string `mapstructure:"DEBUG"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	DBUsername  string `mapstructure:"DB_USERNAME"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBSslMode   string `mapstructure:"DB_SSL_MODE"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	NS1         string `mapstructure:"NS_1"`
	NS2         string `mapstructure:"NS_2"`
	StartSeeder bool   `mapstructure:"START_SEEDER"`
}

type ConfigManager struct {
	config *DatabaseConfig
}

var configManager *ConfigManager

func NewConfigManager() *ConfigManager {
	if configManager == nil {
		configManager = &ConfigManager{}
	}
	return configManager
}

func (cm *ConfigManager) LoadConfiguration() *DatabaseConfig {
	env := cm.getEnvironment()

	if cm.isLocalEnvironment(env) {
		cm.config = cm.loadLocalConfig()
	} else {
		cm.config = cm.loadProductionConfig()
	}

	return cm.config
}

func (cm *ConfigManager) GetConfiguration() *DatabaseConfig {
	if cm.config == nil {
		return cm.LoadConfiguration()
	}
	return cm.config
}

func (cm *ConfigManager) getEnvironment() string {
	return strings.ToUpper(strings.TrimSpace(os.Getenv("ENVIRONMENT")))
}

func (cm *ConfigManager) isLocalEnvironment(env string) bool {
	return env == "LOCAL" || env == ""
}

func (cm *ConfigManager) loadLocalConfig() *DatabaseConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.SetDefault("START_SEEDER", true)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read local config: ", err)
	}

	var config DatabaseConfig
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Failed to unmarshal local config: ", err)
	}

	return &config
}

func (cm *ConfigManager) loadProductionConfig() *DatabaseConfig {
	config := &DatabaseConfig{}

	stringFieldMappings := []struct {
		envVar   string
		fieldPtr *string
		required bool
	}{
		{"DEBUG", &config.Debug, false},
		{"SERVER_PORT", &config.ServerPort, true},
		{"ENVIRONMENT", &config.Environment, true},
		{"DB_USERNAME", &config.DBUsername, true},
		{"DB_PASSWORD", &config.DBPassword, true},
		{"DB_SSL_MODE", &config.DBSslMode, false},
		{"DB_HOST", &config.DBHost, true},
		{"DB_PORT", &config.DBPort, true},
		{"DB_NAME", &config.DBName, true},
		{"DATABASE_URL", &config.DatabaseUrl, false},
		{"NS_1", &config.NS1, false},
		{"NS_2", &config.NS2, false},
	}

	var missingFields []string

	for _, mapping := range stringFieldMappings {
		value := os.Getenv(mapping.envVar)
		if mapping.required && value == "" {
			missingFields = append(missingFields, mapping.envVar)
			continue
		}
		*mapping.fieldPtr = value
	}

	startSeederStr := os.Getenv("START_SEEDER")
	if startSeederStr != "" {
		if startSeeder, err := strconv.ParseBool(startSeederStr); err == nil {
			config.StartSeeder = startSeeder
		} else {
			log.Printf("Invalid boolean value for START_SEEDER: %s, using default: false", startSeederStr)
			config.StartSeeder = false
		}
	} else {
		config.StartSeeder = false // default value for production (main)
	}

	if len(missingFields) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingFields)
	}

	return config
}

func (cm *ConfigManager) ValidateAndFreePort(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		listener.Close()
		return
	}

	cm.terminateProcessOnPort(port)
	cm.verifyPortAvailable(port)
}

func (cm *ConfigManager) terminateProcessOnPort(port int) {
	fmt.Printf("Port %d is occupied. Attempting to free it...\n", port)

	cmd := exec.Command("lsof", "-t", fmt.Sprintf("-i:%d", port))
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to identify processes on port %d: %v", port, err)
	}

	pids := strings.Fields(string(output))
	if len(pids) == 0 {
		log.Fatalf("Port %d appears blocked but no processes found", port)
	}

	for _, pid := range pids {
		if err := exec.Command("kill", "-9", pid).Run(); err != nil {
			log.Printf("Failed to terminate process %s: %v", pid, err)
		} else {
			fmt.Printf("Terminated process %s\n", pid)
		}
	}
}

func (cm *ConfigManager) verifyPortAvailable(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Port %d still unavailable after cleanup", port)
	}
	listener.Close()
	fmt.Printf("Port %d is now available\n", port)
}

func (c *DatabaseConfig) CheckAndKillProcess(port int) {
	configMgr := NewConfigManager()
	configMgr.ValidateAndFreePort(port)
}

// Convenience functions for backward compatibility
func LoadConfig() DatabaseConfig {
	return *NewConfigManager().LoadConfiguration()
}

func GetConf() *DatabaseConfig {
	return NewConfigManager().GetConfiguration()
}

func SetConf(config *DatabaseConfig) {
	NewConfigManager().config = config
}
