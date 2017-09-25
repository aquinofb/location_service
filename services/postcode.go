package services

import (
  "fmt"
  "strings"
  "github.com/tidwall/gjson"
)

func Postcode(lat, lng float32) string {
  data := responseBody(buildPostcodeUri(lat, lng))

  postcode := extractPostcode(data)

  return removeWhiteSpaces(postcode)
}

func buildPostcodeUri(lat, lng float32) string {
  return fmt.Sprintf("https://api.postcodes.io/postcodes?lat=%g&lon=%g", lat, lng)
}

func extractPostcode(data []byte) string {
  return gjson.Get(string(data), "result.1.postcode").String()
}

func removeWhiteSpaces(postcode string) string {
  return strings.Replace(postcode, " ", "", -1)
}
