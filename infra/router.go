package infra

import (
	"net/http"
	"strings"
	"time"

	"github.com/SkyFlareInfra/SkyFlare/pkg"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/acme"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

type SecurityHeaders struct {
	XFrameOptions           string
	XContentTypeOptions     string
	XSSProtection           string
	StrictTransportSecurity string
}

type RouterConfig struct {
	Environment string
	Mode        string
	Origins     []string
}

type RouterBuilder struct {
	config    pkg.DatabaseConfig
	logger    pkg.LogService
	router    *gin.Engine
	routerCfg RouterConfig
	security  SecurityHeaders
}

func NewRouter(configEnv pkg.DatabaseConfig, logger pkg.LogService) Router {
	builder := &RouterBuilder{
		config: configEnv,
		logger: logger,
		security: SecurityHeaders{
			XFrameOptions:           "DENY",
			XContentTypeOptions:     "nosniff",
			XSSProtection:           "1; mode=block",
			StrictTransportSecurity: "max-age=31536000; includeSubDomains; preload",
		},
	}

	return builder.
		setupEnvironment().
		createEngine().
		configureSentry().
		configureCORS().
		configureSecurity().
		registerRoutes().
		build()
}

func (rb *RouterBuilder) setupEnvironment() *RouterBuilder {
	env := strings.ToLower(rb.config.Environment)

	rb.routerCfg = RouterConfig{
		Environment: env,
		Mode:        gin.ReleaseMode,
		Origins:     []string{"http://localhost:5500"},
	}

	if env == "local" || strings.HasPrefix(env, "stag") {
		rb.routerCfg.Mode = gin.DebugMode
		rb.routerCfg.Origins = append(rb.routerCfg.Origins,
			"http://localhost:3000", "http://localhost:5500",
		)
	}

	gin.SetMode(rb.routerCfg.Mode)
	return rb
}

func (rb *RouterBuilder) createEngine() *RouterBuilder {
	rb.router = gin.Default()
	return rb
}

func (rb *RouterBuilder) configureSentry() *RouterBuilder {
	if rb.routerCfg.Environment == "local" {
		return rb
	}

	if err := sentry.Init(sentry.ClientOptions{
		Environment: rb.routerCfg.Environment,
	}); err != nil {
		rb.logger.LogPanicFormat("Sentry initialization failed: %v", err)
	}

	rb.router.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	return rb
}

func (rb *RouterBuilder) configureCORS() *RouterBuilder {
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "x-idempotency", "content-security-policy", "Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     rb.routerCfg.Origins,
		AllowOriginFunc:  func(origin string) bool { return true },
		ExposeHeaders:    []string{"Set-Cookie"},
	}

	rb.router.Use(cors.New(corsConfig))
	return rb
}

func (rb *RouterBuilder) configureSecurity() *RouterBuilder {
	rb.router.Use(func(c *gin.Context) {
		c.Header("X-Frame-Options", rb.security.XFrameOptions)
		c.Header("X-Content-Type-Options", rb.security.XContentTypeOptions)
		c.Header("X-XSS-Protection", rb.security.XSSProtection)
		c.Header("Strict-Transport-Security", rb.security.StrictTransportSecurity)
		c.Header("Last-Modified", time.Now().Format(http.TimeFormat))
		c.Next()
	})
	return rb
}

func (rb *RouterBuilder) registerRoutes() *RouterBuilder {
	rb.router.GET("/health", rb.healthCheck)
	rb.router.GET("/.well-known/acme-challenge/:token", rb.acmeChallenge)
	return rb
}

func (rb *RouterBuilder) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API Up and Running"})
}

func (rb *RouterBuilder) acmeChallenge(c *gin.Context) {
	token := c.Param("token")
	if keyAuth, exists := acme.GetChallengeResponse(token); exists {
		c.Header("Content-Type", "text/plain")
		c.String(http.StatusOK, keyAuth)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (rb *RouterBuilder) build() Router {
	return Router{
		Engine: rb.router,
	}
}
