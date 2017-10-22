package http_client

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return data, nil
}
