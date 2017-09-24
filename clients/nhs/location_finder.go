package nhs

// const (
//   ServiceTypes = ["a_and_e", "b_and_c"]
//   OrganisationTypes = ["pharmacies"]
// )

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/aquinofb/location_service/services"
  "github.com/aquinofb/location_service/models"
)

const (
  NHSBaseAPI = "https://v1.syndication.nhschoices.nhs.uk"
)

var mapTypes = map[string]string{
  "accident_and_emergency": "srv0001",
  "sexual_health_services": "srv0137",
  "pharmacies": "pharmacies",
}

func LocationFinder(lat, lng float32, locationType string) []models.Location {
  postcode := services.Postcode(lat, lng)

  if locationType == "accident_and_emergency" || locationType == "sexual_health_services" {
    return ServicesByPostcode(mapTypes[locationType], postcode)
  } else {
    return OrganisationsByPostcode(mapTypes[locationType], postcode)
  }
}

func responseBody(uri string) []byte {
  resp, err := http.Get(uri)
  if (err != nil) {
    fmt.Errorf("Read body: %v", err)
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Errorf("Read body: %v", err)
  }

  resp.Body.Close()

  return data
}

