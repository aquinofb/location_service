package services

import (
  "fmt"
  "io"
  "strings"
  "net/http"
  "io/ioutil"
)

const Username = "aquinofb"

func CountryIsoCode(lat, lng float32) string {
  resp, err := http.Get(buildCountryIsoCodeUri(lat, lng, Username))
  if (err != nil) {
    fmt.Errorf("Read body: %v", err)
  }

  isoCode := extratCountryIsoCode(resp.Body)
  resp.Body.Close()

  return isoCode
}

func buildCountryIsoCodeUri(lat, lng float32, username string) string {
  return fmt.Sprintf("http://api.geonames.org/countryCode?lat=%g&lng=%g&username=%s", lat, lng, username)
}

func extratCountryIsoCode(responseBody io.ReadCloser) string {
  data, err := ioutil.ReadAll(responseBody)
  if err != nil {
    fmt.Errorf("Read body: %v", err)
  }

  return strings.TrimSpace(string(data))
}
