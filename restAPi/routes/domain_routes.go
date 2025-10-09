package routes

import (
	_ "net/http"

	"github.com/SkyFlareInfra/SkyFlare/infra"
	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/controllers"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/middlewares"
	_ "github.com/gin-gonic/gin"
)

type DomainRoute struct {
	domainController controllers.DomainController
	domainRepo       repository.DomainRepositoryInterface
	log           pkg.LogService
	handler          infra.Router
}

func NewDomainRoute(
	domainController controllers.DomainController,
	domainRepo repository.DomainRepositoryInterface,
	log pkg.LogService,
	handler infra.Router,
) DomainRoute {
	return DomainRoute{
		domainController,
		domainRepo,
		log,
		handler,
	}
}

func (sr DomainRoute) Configure() {
	sr.log.LogInfo("setting up site routes")

	dynamicProxy := middlewares.NewDynamicProxy(sr.domainRepo)

	route := sr.handler.Group("/api/v4/domain")
	{
		route.GET("/proxy", dynamicProxy.DynamicProxy())

		route.POST("/add", sr.domainController.AddDomain)
	}
}
