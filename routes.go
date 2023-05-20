package main

import (
	"github.com/gin-gonic/gin"

	"site-accessibility/modules/site-accessibility-check/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	handler := controllers.SiteAccessibilityControllerHandler()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	userRoutes := router.Group("/url")
	{
		userRoutes.GET("/info", handler.GetUrlInfo)
		userRoutes.GET("/fastest", handler.GetFastestUrl)
		userRoutes.GET("/slowest", handler.GetSlowestUrl)
	}

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.GET("/counter")
	}

	return router
}
