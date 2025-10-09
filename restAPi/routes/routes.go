package routes

import (
	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
        NewRouteBuilder,
        CreateRegistry,
		NewDomainRoute,
    ),
)

// RouteInitializer defines the setup behavior
type RouteInitializer interface {
	Configure()
}

// RouteRegistry holds all route initializers
type RouteRegistry []RouteInitializer

// RouteBuilder constructs routes with required services
type RouteBuilder struct {
	Logger  pkg.LogService
	Router  infra.Router
}

// NewRouteBuilder creates a route builder instance
func NewRouteBuilder(logger pkg.LogService, router infra.Router) *RouteBuilder {
	return &RouteBuilder{
		Logger: logger,
		Router: router,
	}
}

// CreateRegistry collects all route initializers
func CreateRegistry(
    siteRoute DomainRoute,
) RouteRegistry {
    return RouteRegistry{
        siteRoute,
    }
}

// SetupRoutes configures all registered routes
func (rr RouteRegistry) SetupRoutes() {
	for _, initializer := range rr {
		initializer.Configure()
	}
}
