package nhs

import (
	"encoding/xml"
	"fmt"
	"github.com/aquinofb/location_service/http_client"
	"github.com/aquinofb/location_service/models"
	"os"
	"strings"
)

type Result struct {
	Entry Entry `xml:"entry"`
}

type Entry struct {
	Id      string `xml:"id"`
	Content struct {
		Service struct {
			Name string `xml:"deliverer"`
			Type struct {
				Code string `xml:"code,attr"`
			} `xml:"type"`
			GeographicCoordinates struct {
				Lat string `xml:"latitude"`
				Lng string `xml:"longitude"`
			} `xml:"geographicCoordinates"`
		} `xml:"service"`
	} `xml:"content"`
}

type ResultPostcodeXML struct {
	Entries []Entry `xml:"entry"`
}

func ServicesFinder(serviceType, serviceId string) models.Location {
	data := http_client.Get(buildNHSServicesUri(serviceType, serviceId))

	result := Result{}
	if err := xml.Unmarshal(data, &result); err != nil {
		fmt.Printf("error: %v", err)
	}

	service := result.Entry.Content.Service

	return models.Location{
		Id:    serviceId,
		Name:  service.Name,
		Lat:   service.GeographicCoordinates.Lat,
		Lng:   service.GeographicCoordinates.Lng,
		Types: []string{strings.ToLower(service.Type.Code)},
	}
}

func ServicesByPostcode(serviceType, postcode string) []models.Location {
	data := http_client.Get(buildNHSServicesPostcodeUri(serviceType, postcode))

	result := ResultPostcodeXML{}
	if err := xml.Unmarshal(data, &result); err != nil {
		fmt.Printf("error: %v", err)
	}

	var locations []models.Location
	for _, entry := range result.Entries {
		location := ServicesFinder(serviceType, extractServiceIdFromUrl(entry.Id))
		locations = append(locations, location)
	}

	return locations
}

func buildNHSServicesUri(serviceType, serviceId string) string {
	return fmt.Sprintf("%s/services/types/%s/%s.xml?apikey=%s", NHSBaseAPI, serviceType, serviceId, os.Getenv("NHS_API_KEY"))
}

func buildNHSServicesPostcodeUri(serviceType, postcode string) string {
	return fmt.Sprintf("%s/services/types/%s/postcode/%s.xml?apikey=%s", NHSBaseAPI, serviceType, postcode, os.Getenv("NHS_API_KEY"))
}

func extractServiceIdFromUrl(url string) string {
	urlSplited := strings.Split(url, "/")
	return urlSplited[len(urlSplited)-1]
}
