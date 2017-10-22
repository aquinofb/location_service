package google

import (
	"encoding/json"
	"fmt"
	"github.com/aquinofb/location_service/http_client"
	"github.com/aquinofb/location_service/models"
	"os"
)

var mapTypes = map[string]string{
	"accident_and_emergency": "hospital",
	"sexual_health_services": "hospital",
	"pharmacies":             "pharmacy",
}

type Result struct {
	Places []Place `json:"results"`
}

type Place struct {
	Id       string `json:"id"`
	PlaceId  string `json:"place_id"`
	Name     string `json:"name"`
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
}

func (place Place) toLocation(locationType string) models.Location {
	return models.Location{
		Id:    place.Id,
		Name:  place.Name,
		Types: []string{locationType},
		Lat:   fmt.Sprintf("%g", place.Geometry.Location.Lat),
		Lng:   fmt.Sprintf("%g", place.Geometry.Location.Lng),
	}
}

func LocationDetails(reference string) models.Location {
	return models.Location{}
}

func NearbySearch(lat, lng float32, locationType string) ([]models.Location, error) {
	data, _ := http_client.Get(nearbySearchURI(lat, lng, mapTypes[locationType]))

	result := Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var locations []models.Location
	for _, place := range result.Places {
		locations = append(locations, place.toLocation(locationType))
	}

	return locations, nil
}

func nearbySearchURI(lat, lng float32, locationType string) string {
	return fmt.Sprintf("%s/nearbysearch/json?location=%g,%g&radius=3000&type=%s&key=%s", baseAPI(), lat, lng, locationType, os.Getenv("GOOGLE_API_KEY"))
}

func baseAPI() string {
	return "https://maps.googleapis.com/maps/api/place"
}
