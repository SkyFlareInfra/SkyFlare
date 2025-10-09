package application

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SkyFlareInfra/SkyFlare/common"
	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/client"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/controllers"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/routes"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/services"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/utility"
	"github.com/SkyFlareInfra/SkyFlare/service"
	_ "github.com/getsentry/sentry-go"
	"go.uber.org/fx"
)

var Module = fx.Options(
	common.Module,
	infra.Module,
	pkg.Module,
	routes.Module,
	controllers.Module,
	services.Module,
	service.Module,
	client.Module,
	utility.Module,
	repository.Module,
	fx.Invoke(NewApp),
)

type App struct {
	log       pkg.LogService
	configEnv pkg.DatabaseConfig
	database  infra.DatabaseManager
	handler   infra.Router
	routes    routes.RouteRegistry
}

func NewApp(
	lifecycle fx.Lifecycle,
	log pkg.LogService,
	configEnv pkg.DatabaseConfig,
	database infra.DatabaseManager,
	handler infra.Router,
	routes routes.RouteRegistry,
) *App {
	app := &App{
		log:       log,
		configEnv: configEnv,
		database:  database,
		handler:   handler,
		routes:    routes,
	}

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
				╔══════════════════════════════════════════════════╗
				║                  Starting Application            ║
				╠══════════════════════════════════════════════════╣
				║                                                  ║
				║                    S K Y F L A R E               ║
				║             CloudService Controller              ║
				║                     v1.0                         ║
				║                                                  ║
				╚══════════════════════════════════════════════════╝
			`
	fmt.Print(banner)
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
	a.routes.SetupRoutes()
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
