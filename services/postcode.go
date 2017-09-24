package services

import (
  "fmt"
  "strings"
  "io"
  "io/ioutil"
  "net/http"
  "github.com/tidwall/gjson"
)

func Postcode(lat, lng float32) string {
  resp, err := http.Get(buildPostcodeUri(lat, lng))
  if (err != nil) {
    fmt.Errorf("Read body: %v", err)
  }

  postcode := extractPostcode(resp.Body)
  resp.Body.Close()
  
  return removeWhiteSpaces(postcode)
}

func buildPostcodeUri(lat, lng float32) string {
  return fmt.Sprintf("https://api.postcodes.io/postcodes?lat=%g&lon=%g", lat, lng)
}

func extractPostcode(responseBody io.ReadCloser) string {
  data, err := ioutil.ReadAll(responseBody)
  if err != nil {
    fmt.Errorf("Read body: %v", err)
  }

  return gjson.Get(string(data), "result.1.postcode").String()
}

func removeWhiteSpaces(postcode string) string {
  return strings.Replace(postcode, " ", "", -1)
}
