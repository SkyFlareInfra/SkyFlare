package middlewares

import (
	"net/http/httputil"
	"net/url"

	"github.com/SkyFlareInfra/SkyFlare/repository"
	"github.com/gin-gonic/gin"
)

type DynamicProxy struct {
	domainRepo repository.DomainRepositoryInterface
}

func NewDynamicProxy(
	domainRepo repository.DomainRepositoryInterface,
) DynamicProxy {
	return DynamicProxy{
		domainRepo,
	}
}

func (dp DynamicProxy) DynamicProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		site, err := dp.domainRepo.GetDomainByName(host)
		if err != nil || site == nil {
			c.JSON(404, gin.H{
				"error": "Site not found",
			})
			return
		}

		if site.OriginServer != "" {
			origin, err := url.Parse(site.OriginServer)
			if err != nil {
				c.JSON(500, gin.H{
					"error": "Invalid origin server URL",
				})
				return
			}
			proxy := httputil.NewSingleHostReverseProxy(origin)
			proxy.ServeHTTP(c.Writer, c.Request)
		} else {
			c.String(200, site.Content)
		}
	}
}
