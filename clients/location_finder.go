package clients

import (
  "github.com/aquinofb/location_service/clients/nhs"
  "github.com/aquinofb/location_service/clients/google"

  "github.com/aquinofb/location_service/services"
  "github.com/aquinofb/location_service/models"
)

func LocationFinder(lat, lng float32, location_type string) []models.Location {
  switch services.CountryIsoCode(lat, lng) {
  case "GB":
    return nhs.LocationFinder(lat, lng, location_type)
  default:
    return google.LocationFinder(lat, lng, location_type)
  }
}
