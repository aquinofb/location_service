package nhs

import (
  "os"
  "fmt"
  "strings"
  "encoding/xml"
  "github.com/aquinofb/location_service/models"
)

type OrganisationXML struct {
  Id string `xml:"OrganisationId"`
  Name string `xml:"Name"`
  Lat string `xml:"Latitude"`
  Lng string `xml:"Longitude"`
}

type ResultFromPostcodeXML struct {
  Entries []ResultEntryXML `xml:"entry"`
}

type ResultEntryXML struct {
  Id string `xml:"id"`
  Name string `xml:"content>organisationSummary>name"`
  Lat string `xml:"content>organisationSummary>geographicCoordinates>latitude"`
  Lng string `xml:"content>organisationSummary>geographicCoordinates>longitude"`
}

func OrganisationsFinder(organisationType, organisationId string) models.Location {
  data := responseBody(buildNHSOrganisationsUri(organisationType, organisationId))

  result := OrganisationXML{}
  if err := xml.Unmarshal(data, &result); err != nil {
    fmt.Printf("error: %v", err)
  }

  return models.Location{
    Id: result.Id,
    Name: result.Name,
    Lat: result.Lat,
    Lng: result.Lng,
    Types: []string{organisationType},
  }
}

func OrganisationsByPostcode(organisationType, postcode string) []models.Location {
  data := responseBody(buildNHSOrganisationsPostcodeUri(organisationType, postcode))

  result := ResultFromPostcodeXML{}
  if err := xml.Unmarshal(data, &result); err != nil {
    fmt.Printf("error: %v", err)
  }

  var locations []models.Location
  for _, entry := range result.Entries {
    location := 
      models.Location{
        Id: extractOrganisationIdFromUrl(entry.Id),
        Name: entry.Name,
        Lat: entry.Lat,
        Lng: entry.Lng,
        Types: []string{organisationType},
      }
    locations = append(locations, location)
  }

  return locations
}

func buildNHSOrganisationsUri(organisationType, organisationId string) string {
  return fmt.Sprintf("%s/organisations/%s/%s.xml?apikey=%s", NHSBaseAPI, organisationType, organisationId, os.Getenv("NHS_API_KEY"))
}

func buildNHSOrganisationsPostcodeUri(organisationType, postcode string) string {
  return fmt.Sprintf("%s/organisations/%s/postcode/%s.xml?apikey=%s", NHSBaseAPI, organisationType, postcode, os.Getenv("NHS_API_KEY"))
}

func extractOrganisationIdFromUrl(url string) string {
  urlSplited := strings.Split(url, "/")
  return urlSplited[len(urlSplited) - 1]
}

func extractOrganisationTypeFromUrl(url string) string {
  urlSplited := strings.Split(url, "/")
  return urlSplited[len(urlSplited) - 2]
}
