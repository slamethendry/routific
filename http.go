package routific

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// post performs http POST, specifying auth token and JSON type
func post(visits interface{}, url string, token string) ([]byte, error) {

	v, err := json.Marshal(visits)
	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(v)))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode == 200 || res.StatusCode == 202 {
		return ioutil.ReadAll(res.Body)
	}

	e := fmt.Sprintf("Status Code %d", res.StatusCode)
	return []byte{}, errors.New(e)
}

// get performs http GET, specifying auth token
func get(url string, token string) ([]byte, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Authorization", "bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode == 200 || res.StatusCode == 202 {
		return ioutil.ReadAll(res.Body)
	}

	e := fmt.Sprintf("Status Code %d", res.StatusCode)
	return []byte{}, errors.New(e)
}
