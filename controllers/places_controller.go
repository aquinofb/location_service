package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/aquinofb/location_service/clients"
)

type QueryParams struct {
	Type string  `form:"type"`
	Lat  float32 `form:"lat"`
	Lng  float32 `form:"lng"`
}

func PlacesIndex(c *gin.Context) {
	var params QueryParams
	c.Bind(&params)

	locations := clients.LocationFinder(params.Lat, params.Lng, params.Type)
	c.JSON(http.StatusOK, gin.H{
		"results": locations,
	})
}
