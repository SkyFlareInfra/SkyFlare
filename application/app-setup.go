package application

import (
	"context"
	"strconv"

	"github.com/SkyFlareInfra/SkyFlare/errors"
	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	_ "github.com/getsentry/sentry-go"
	"go.uber.org/fx"
)

var Module = fx.Options(
	errors.Module,
	fx.Invoke(NewApp),
)

type App struct {
	log       pkg.LogService
	configEnv pkg.DatabaseConfig
	database  infra.DatabaseManager
	handler   infra.Router
}

func NewApp(
	lifecycle fx.Lifecycle,

) *App {
	app := &App{}

	lifecycle.Append(fx.Hook{
		OnStart: app.Start,
		OnStop:  app.Stop,
	})

	return app
}

func (a *App) Start(ctx context.Context) error {
	if err := a.validateConfig(); err != nil {
		return err
	}

	a.printStartupBanner()

	if err := a.setupDatabase(); err != nil {
		return err
	}

	a.startServerAsync()
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.log.LogInfo("Stopping Application")
	return a.closeDatabase()
}

func (a *App) validateConfig() error {
	port, err := strconv.Atoi(a.configEnv.ServerPort)
	if err != nil {
		a.log.LogErrorFormat("Invalid server port: %v", err)
		return err
	}
	a.configEnv.CheckAndKillProcess(port)
	return nil
}

func (a *App) printStartupBanner() {
	banner := `
			┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
			┃              Starting Application                ┃
			┡━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┩
		    ┃                                                  ┃
			┃              S K Y F L A R E                     ┃
			┃           CloudService Controller                ┃
			┃                   v1.0                           ┃
			┃                                                  ┃
			┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
			`
	a.log.LogInfo(banner)
}

func (a *App) setupDatabase() error {
	sqlDB, err := a.database.DB.DB()
	if err != nil {
		a.log.LogErrorFormat("Failed to get database connection: %v", err)
		return err
	}

	return sqlDB.Ping()
}

func (a *App) startServerAsync() {
	// a.routes.Setup() // commented out for now
	go func() {
		addr := ":" + a.configEnv.ServerPort
		if err := a.handler.Run(addr); err != nil {
			a.log.LogError("Error running server: ", err)
		}
	}()
}

func (a *App) closeDatabase() error {
	sqlDB, err := a.database.DB.DB()
	if err != nil {
		a.log.LogErrorFormat("Error retrieving the database connection: %v", err)
		return err
	}
	return sqlDB.Close()
}
