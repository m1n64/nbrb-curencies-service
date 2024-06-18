package services

import (
	"io/ioutil"
	"net/http"
	"time"
)

// GetRequest retrieves the content of the specified URL.
//
// url - the URL to make the GET request to.
// []byte - the response body as a byte slice.
func GetRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, err
}
