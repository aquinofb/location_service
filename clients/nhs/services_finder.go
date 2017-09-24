package nhs

import (
  "os"
  "fmt"
  "strings"
  "encoding/xml"
  "github.com/aquinofb/location_service/models"
)

type Result struct {
  Entry Entry `xml:"entry"`
}

type Entry struct {
  Id string `xml:"id"`
  Name string `xml:"content>service>deliverer"`
  Lat string `xml:"content>service>geographicCoordinates>latitude"`
  Lng string `xml:"content>service>geographicCoordinates>longitude"`
  Type struct {
    Code string `xml:"code,attr"`
  } `xml:"content>service>type"`
}

type ResultPostcodeXML struct {
  Entries []Entry `xml:"entry"`
}

func ServicesFinder(serviceType, serviceId string) models.Location {
  data := responseBody(buildNHSServicesUri(serviceType, serviceId))

  result := Result{}
  if err := xml.Unmarshal(data, &result); err != nil {
    fmt.Printf("error: %v", err)
  }

  return models.Location{
    Id: serviceId,
    Name: result.Entry.Name,
    Lat: result.Entry.Lat,
    Lng: result.Entry.Lng,
    Types: []string{strings.ToLower(result.Entry.Type.Code)},
  }
}

func ServicesByPostcode(serviceType, postcode string) []models.Location {
  data := responseBody(buildNHSServicesPostcodeUri(serviceType, postcode))
  
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
  return urlSplited[len(urlSplited) - 1]
}
