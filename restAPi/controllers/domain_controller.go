package controllers

import (
	"net/http"

	"github.com/SkyFlareInfra/SkyFlare/restAPi/services"
	"github.com/SkyFlareInfra/SkyFlare/restAPi/types"
	"github.com/gin-gonic/gin"
)

type DomainController struct {
	domainService services.DomainServicesInterface
}

func NewDomainController(
	domainService services.DomainServicesInterface,
) DomainController {
	return DomainController{
		domainService: domainService,
	}
}

func (dc DomainController) AddDomain(c *gin.Context) {
	var domain types.DomainInput

	if err := c.ShouldBindJSON(&domain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	resp, err := dc.domainService.AddDomain(c, domain)
	if err != nil {
		c.JSON(err.StatusCode, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "domain added successfully",
		"data":    resp,
	})
}
