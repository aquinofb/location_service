package services

import (
	"fmt"
	"strings"
	// "github.com/aquinofb/location_service/http_client"
)

const Username = "aquinofb"

func CountryIsoCode(lat, lng float32) string {
	// data := http_client.Get(buildCountryIsoCodeUri(lat, lng, Username))

	// return extratCountryIsoCode(data)
	return "GB"
}

func buildCountryIsoCodeUri(lat, lng float32, username string) string {
	return fmt.Sprintf("http://api.geonames.org/countryCode?lat=%g&lng=%g&username=%s", lat, lng, username)
}

func extratCountryIsoCode(data []byte) string {
	return strings.TrimSpace(string(data))
}
