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

	locations, err := clients.NearbySearch(params.Lat, params.Lng, params.Type)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": locations,
		})
	}
}
