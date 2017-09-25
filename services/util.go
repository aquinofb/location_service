package services

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func responseBody(url string) []byte {
  resp, err := http.Get(url)
  if (err != nil) {
    fmt.Errorf("Read body: %v", err)
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Errorf("Read body: %v", err)
  }

  resp.Body.Close()

  return data
}
