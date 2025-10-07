package infra

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseManager struct {
	DB        *gorm.DB
	config    pkg.DatabaseConfig
	logger    pkg.LogService
}

type ConnectionConfig struct {
	MaxRetries   int
	RetryDelay   time.Duration
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

type DatabaseCredentials struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func NewDatabase(logService pkg.LogService, config pkg.DatabaseConfig) DatabaseManager {
	dbManager := NewDatabaseManager(logService, config)
	dbManager.Initialize()
	return dbManager
}

func NewDatabaseManager(logService pkg.LogService, config pkg.DatabaseConfig) DatabaseManager {
	return DatabaseManager{
		logger: logService,
		config: config,
	}
}

func (dm *DatabaseManager) Initialize() *gorm.DB {
	connConfig := ConnectionConfig{
		MaxRetries:   3,
		RetryDelay:   2 * time.Second,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxLifetime:  time.Hour,
	}

	credentials := dm.extractCredentials()
	connectionString := dm.buildConnectionString(credentials)

	dm.logger.LogInfoFormat("Connecting to database at %s@%s:%s", 
		credentials.Username, credentials.Host, credentials.Port)

	db := dm.establishConnection(connectionString, connConfig)
	dm.configureConnectionPool(db, connConfig)
	dm.setupDatabaseLogger(db)
	
	dm.performMigrations(db)
	dm.ensureIndexes(db)

	dm.DB = db
	return db
}

func (dm *DatabaseManager) extractCredentials() DatabaseCredentials {
	if dm.config.DatabaseUrl != "" {
		return dm.parseDatabaseURL(dm.config.DatabaseUrl)
	}

	return DatabaseCredentials{
		Username: dm.config.DBUsername,
		Password: dm.config.DBPassword,
		Host:     dm.config.DBHost,
		Port:     dm.config.DBPort,
		Name:     dm.config.DBName,
		SSLMode:  dm.config.DBSslMode,
	}
}

func (dm *DatabaseManager) parseDatabaseURL(dbURL string) DatabaseCredentials {
	parsed, err := url.Parse(dbURL)
	if err != nil {
		dm.logger.LogError("Failed to parse database URL", err)
		panic(fmt.Sprintf("Invalid database URL: %v", err))
	}

	username := parsed.User.Username()
	password, _ := parsed.User.Password()
	host := parsed.Hostname()
	port := parsed.Port()
	if port == "" {
		port = "5432"
	}

	dbName := strings.TrimPrefix(parsed.Path, "/")
	sslMode := "disable"

	if parsed.RawQuery != "" {
		queryParams, _ := url.ParseQuery(parsed.RawQuery)
		if mode := queryParams.Get("sslmode"); mode != "" {
			sslMode = mode
		}
	}

	return DatabaseCredentials{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     dbName,
		SSLMode:  sslMode,
	}
}

func (dm *DatabaseManager) buildConnectionString(creds DatabaseCredentials) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		creds.Host,
		creds.Port,
		creds.Username,
		creds.Name,
		url.QueryEscape(creds.Password),
		creds.SSLMode)
}

func (dm *DatabaseManager) establishConnection(dsn string, config ConnectionConfig) *gorm.DB {
	gormLogger := createGormLogger()

	for attempt := 1; attempt <= config.MaxRetries; attempt++ {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			Logger: gormLogger,
		})

		if err == nil {
			return db
		}

		if attempt < config.MaxRetries {
			dm.logger.LogErrorFormat("Connection attempt %d failed, retrying: %v", attempt, err)
			time.Sleep(config.RetryDelay)
			continue
		}

		dm.logger.LogError("All connection attempts failed", err)
		panic(fmt.Sprintf("Database connection failed: %v", err))
	}

	return nil
}

func (dm *DatabaseManager) configureConnectionPool(db *gorm.DB, config ConnectionConfig) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get SQL DB: %v", err))
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.MaxLifetime)
}

func (dm *DatabaseManager) setupDatabaseLogger(db *gorm.DB) {
	db.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Warn,
			SlowThreshold:             200 * time.Millisecond,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
}

func (dm *DatabaseManager) performMigrations(db *gorm.DB) {
	models := []interface{}{
		// we add all models here
	}

	migrationService := NewMigrationService(dm.logger)
	migrationService.RunMigrations(db, models)
}

func (dm *DatabaseManager) ensureIndexes(db *gorm.DB) {
	indexService := NewIndexService(dm.logger)
	indexService.CreateRequiredIndexes(db)
}

type MigrationService struct {
	logger pkg.LogService
}

func NewMigrationService(logger pkg.LogService) *MigrationService {
	return &MigrationService{logger: logger}
}

func (ms *MigrationService) RunMigrations(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		modelName := fmt.Sprintf("%T", model)
		
		if !db.Migrator().HasTable(model) {
			ms.logger.LogInfoFormat("Creating table for: %s", modelName)
			if err := db.Migrator().CreateTable(model); err != nil {
				ms.handleMigrationError(err, "create table", modelName)
			}
		} else {
			ms.logger.LogInfoFormat("Migrating table for: %s", modelName)
			if err := db.AutoMigrate(model); err != nil {
				ms.handleMigrationError(err, "migrate", modelName)
			}
		}
	}
}

func (ms *MigrationService) handleMigrationError(err error, operation, modelName string) {
	ms.logger.LogErrorFormat("%s failed for %s: %v", operation, modelName, err)
	sentry.CaptureException(fmt.Errorf("%s failed for %s: %w", operation, modelName, err))
	panic(fmt.Errorf("%s failed for %s: %w", operation, modelName, err))
}

type IndexService struct {
	logger pkg.LogService
}

func NewIndexService(logger pkg.LogService) *IndexService {
	return &IndexService{logger: logger}
}

func (is *IndexService) CreateRequiredIndexes(db *gorm.DB) error {
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_generated_nameservers_domain_id ON generated_nameservers(domain_id);",
		"CREATE INDEX IF NOT EXISTS idx_generated_nameservers_uuid ON generated_nameservers(uuid);",
		"CREATE INDEX IF NOT EXISTS idx_nameservers_generated_ns_uuid ON nameservers(generated_ns_uuid);",
		"CREATE INDEX IF NOT EXISTS idx_nameservers_domain_id ON nameservers(domain_id);",
	}

	for _, index := range indexes {
		if err := db.Exec(index).Error; err != nil {
			return fmt.Errorf("failed to create index: %w", err)
		}
	}
	return nil
}

func createGormLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
}
