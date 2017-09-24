package main

import (
  "github.com/aquinofb/location_service/config"
)

func main() {
  config.Router().Run(":4000")
}

