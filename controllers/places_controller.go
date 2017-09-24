package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"

  "github.com/aquinofb/location_service/clients"
)

type QueryParams struct {
  Type string `form:"type"`
  Lat float32 `form:"lat"`
  Lng float32 `form:"lng"`
}

func PlacesIndex(c *gin.Context) {
  // ##.###
  // "http://ws.geonames.org/countryCodeJSON?lat=52.2&lng=-2.29&username=aquinofb"

  var params QueryParams
  c.Bind(&params)

  locations := clients.LocationFinder(params.Lat, params.Lng, params.Type)
  c.JSON(http.StatusOK, gin.H{
    "results": locations,
  })
}

// func buildUri(lat float32, lng float32) string {
//   return fmt.Sprintf("https://api.postcodes.io/postcodes?lat=%g&lon=%g", lat, lng)
// }
