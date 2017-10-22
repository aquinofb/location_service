package nhs

import (
	"github.com/aquinofb/location_service/models"
	"github.com/aquinofb/location_service/services"
)

const (
	NHSBaseAPI = "https://v1.syndication.nhschoices.nhs.uk"
)

var mapTypes = map[string]string{
	"accident_and_emergency": "srv0001",
	"sexual_health_services": "srv0137",
	"pharmacies":             "pharmacies",
}

func NearbySearch(lat, lng float32, locationType string) ([]models.Location, error) {
	postcode, err := services.PostcodeIO(lat, lng)

	if err != nil {
		return nil, err
	}

	if locationType == "accident_and_emergency" || locationType == "sexual_health_services" {
		return ServicesByPostcode(mapTypes[locationType], postcode)
	} else {
		return OrganisationsByPostcode(mapTypes[locationType], postcode)
	}
}
