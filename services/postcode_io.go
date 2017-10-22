package services

import (
	"fmt"
	"github.com/aquinofb/location_service/http_client"
	"github.com/tidwall/gjson"
	"strings"
)

type NotFoundError struct {
  Message string
}

func (e NotFoundError) Error() string {
  return e.Message
}

func PostcodeIO(lat, lng float32) (string, error) {
	data, _ := http_client.Get(buildPostcodeUri(999.99, 999.99))

	postcode := extractPostcode(data)

	if len(postcode) == 0 {
		return "", &NotFoundError{"Postcode not found"}
	}

	return removeWhiteSpaces(postcode), nil
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
