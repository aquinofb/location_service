package clients

import (
  "fmt"
  "github.com/aquinofb/location_service/clients/nhs"
  "github.com/aquinofb/location_service/clients/google"

  "github.com/aquinofb/location_service/services"
  "github.com/aquinofb/location_service/models"
)

func LocationFinder(lat, lng float32, location_type string) []models.Location {
  switch services.CountryIsoCode(lat, lng) {
  case "GB":
    fmt.Println("It's GB!!!")
    return nhs.LocationFinder(lat, lng, location_type)
  default:
    fmt.Println("It's not GB :'(")
    return google.LocationFinder(lat, lng, location_type)
  }
  
  // if (service?) {
  //   getpostcode
  //   NHSService.by_postcode()
  // } else if (organisation?)
}
