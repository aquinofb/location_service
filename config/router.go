package config

import (
	"github.com/aquinofb/location_service/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", controllers.HomeIndex)

	api := router.Group("/api")
	{
		api.GET("/places", controllers.PlacesIndex)
		api.GET("/places/:reference", controllers.HomeShow)
	}

	return router
}
