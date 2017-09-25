package services

import (
	"fmt"
	"github.com/aquinofb/location_service/http_client"
	"github.com/tidwall/gjson"
	"strings"
)

func Postcode(lat, lng float32) string {
	data := http_client.Get(buildPostcodeUri(lat, lng))

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
