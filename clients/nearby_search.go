package clients

import (
  "github.com/aquinofb/location_service/clients/nhs"
  "github.com/aquinofb/location_service/clients/google"

  "github.com/aquinofb/location_service/services"
  "github.com/aquinofb/location_service/models"
)

const(
  GBIsoCode = "GB"
)

func NearbySearch(lat, lng float32, location_type string) ([]models.Location, error) {
  country_iso_code, err := services.CountryIsoCode(lat, lng)

  if err != nil {
    return nil, err
  }

  switch  country_iso_code {
  case GBIsoCode:
    return nhs.NearbySearch(lat, lng, location_type)
  default:
    return google.NearbySearch(lat, lng, location_type)
  }
}
