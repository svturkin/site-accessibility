package main

import (
	"github.com/gin-gonic/gin"

	"site-accessibility/helpers"
	"site-accessibility/modules/site-accessibility-check/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	handler := controllers.SiteAccessibilityControllerHandler()

	router.Use(helpers.CountRequests())

	router.GET("/", handler.PingServer)

	userRoutes := router.Group("/url")
	{
		userRoutes.GET("/info", handler.GetUrlInfo)
		userRoutes.GET("/fastest", handler.GetFastestUrl)
		userRoutes.GET("/slowest", handler.GetSlowestUrl)
	}

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.GET("/counter", handler.GetCounter)
	}

	return router
}
