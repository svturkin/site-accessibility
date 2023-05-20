package controllers

import (
	"net/http"
	"site-accessibility/modules/site-accessibility-check/services"

	"github.com/gin-gonic/gin"
)

type SiteAccessibilityController struct {
	siteService services.SiteAccessibilityInterface
}

func SiteAccessibilityControllerHandler() SiteAccessibilityController {
	handler := SiteAccessibilityController{
		siteService: services.SiteAccessibilityServiceHandler(),
	}

	return handler
}

func (ctrl *SiteAccessibilityController) GetUrlInfo(c *gin.Context) {
	if url := c.Query("url"); url != "" {
		siteInfo, err := ctrl.siteService.GetInfoByUrl(url)
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "site not found",
				"data":    nil,
				"message": http.StatusBadRequest,
			})
			return
		}

		if siteInfo == 0 {
			c.JSON(200, gin.H{
				"status":  "success",
				"data":    siteInfo,
				"message": "site is down!",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "success",
			"data":    siteInfo,
			"message": "success get data",
		})

	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func (ctrl *SiteAccessibilityController) GetFastestUrl(c *gin.Context) {
	url, _ := ctrl.siteService.GetFastestUrl()
	c.JSON(200, gin.H{
		"status":  "success",
		"data":    url,
		"message": "success get data",
	})
}

func (ctrl *SiteAccessibilityController) GetSlowestUrl(c *gin.Context) {
	url, _ := ctrl.siteService.GetSlowestUrl()
	c.JSON(200, gin.H{
		"status":  "success",
		"data":    url,
		"message": "success get data",
	})
}
