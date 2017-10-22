package nhs

import (
	"encoding/xml"
	"fmt"
	"github.com/aquinofb/location_service/http_client"
	"github.com/aquinofb/location_service/models"
	"os"
	"strings"
)

type OrganisationXML struct {
	Id   string `xml:"OrganisationId"`
	Name string `xml:"Name"`
	Lat  string `xml:"Latitude"`
	Lng  string `xml:"Longitude"`
}

type ResultFromPostcodeXML struct {
	Entries []ResultEntryXML `xml:"entry"`
}

type ResultEntryXML struct {
	Id      string `xml:"id"`
	Content struct {
		OrganisationSummary struct {
			Name                  string `xml:"name"`
			GeographicCoordinates struct {
				Lat string `xml:"latitude"`
				Lng string `xml:"longitude"`
			} `xml:"geographicCoordinates"`
		} `xml:"organisationSummary"`
	} `xml:"content"`
}

func (entry ResultEntryXML) toLocation(id, organisationType string) models.Location {
	organisation := entry.Content.OrganisationSummary
	return models.Location{
		Id:    id,
		Name:  organisation.Name,
		Lat:   organisation.GeographicCoordinates.Lat,
		Lng:   organisation.GeographicCoordinates.Lng,
		Types: []string{organisationType},
	}
}

func OrganisationsFinder(organisationType, organisationId string) models.Location {
	data, _ := http_client.Get(buildNHSOrganisationsUri(organisationType, organisationId))

	result := OrganisationXML{}
	if err := xml.Unmarshal(data, &result); err != nil {
		fmt.Printf("error: %v", err)
	}

	return models.Location{
		Id:    result.Id,
		Name:  result.Name,
		Lat:   result.Lat,
		Lng:   result.Lng,
		Types: []string{organisationType},
	}
}

func OrganisationsByPostcode(organisationType, postcode string) ([]models.Location, error) {
	data, _ := http_client.Get(buildNHSOrganisationsPostcodeUri(organisationType, postcode))

	result := ResultFromPostcodeXML{}
	if err := xml.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var locations []models.Location
	for _, entry := range result.Entries {
		location :=
			entry.toLocation(
				extractOrganisationIdFromUrl(entry.Id),
				organisationType,
			)
		locations = append(locations, location)
	}

	return locations, nil
}

func buildNHSOrganisationsUri(organisationType, organisationId string) string {
	return fmt.Sprintf("%s/organisations/%s/%s.xml?apikey=%s", NHSBaseAPI, organisationType, organisationId, os.Getenv("NHS_API_KEY"))
}

func buildNHSOrganisationsPostcodeUri(organisationType, postcode string) string {
	return fmt.Sprintf("%s/organisations/%s/postcode/%s.xml?apikey=%s", NHSBaseAPI, organisationType, postcode, os.Getenv("NHS_API_KEY"))
}

func extractOrganisationIdFromUrl(url string) string {
	urlSplited := strings.Split(url, "/")
	return urlSplited[len(urlSplited)-1]
}

func extractOrganisationTypeFromUrl(url string) string {
	urlSplited := strings.Split(url, "/")
	return urlSplited[len(urlSplited)-2]
}
