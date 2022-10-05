package geoapify

//package planner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const geoapify = "https://api.geoapify.com"
const plannerV1 = "/v1/routeplanner"
const batchV1 = "/v1/batch"

// Post a sychronous request to get a route plan.
// See https://apidocs.geoapify.com/docs/route-planner/#about
func (r Request) Post(apiKey string) (Plan, error) {

	apiURL := fmt.Sprintf("%s%s?apiKey=%s", geoapify, plannerV1, apiKey)

	body, err := postJSON(apiURL, r)
	if err != nil {
		return Plan{}, err
	}

	p := new(Plan)
	err = json.Unmarshal(body, p)
	if err != nil {
		return Plan{}, err
	}

	return *p, nil
}

// Post a batch request to get route plan(s) in async manner.
// See https://apidocs.geoapify.com/docs/batch/#about
// It will loop 'maxTry' x 'interval' seconds to wait for Geoapify to return
// the route plan(s).
func (b BatchRequest) Post(apiKey string, maxTry, interval uint) (BatchResponse, error) {

	b.API = plannerV1
	apiURL := fmt.Sprintf("%s%s?apiKey=%s", geoapify, batchV1, apiKey)

	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return BatchResponse{}, err
	}

	client := new(http.Client)
	reader := bytes.NewReader(jsonBytes)
	res, err := client.Post(apiURL, "application/json", reader)
	if err != nil {
		return BatchResponse{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return BatchResponse{}, err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		e := fmt.Sprintf("POST: Expecting 200 or 202, but got status %d\n%s",
			res.StatusCode,
			string(body))
		return BatchResponse{}, errors.New(e)
	}

	var job struct {
		ID string `json:"id"`
	}
	err = json.Unmarshal(body, &job)
	if err != nil {
		return BatchResponse{}, err
	}

	done := false
	apiURL = fmt.Sprintf("%s&id=%s", apiURL, job.ID)

	for try := uint(0); try < maxTry && !done; try++ {
		time.Sleep(time.Duration(interval) * time.Second)

		status, body, err := getJSON(apiURL)
		if err != nil {
			return BatchResponse{}, err
		}

		if status == 200 {
			done = true

			var output BatchResponse
			if err = json.Unmarshal(body, &output); err != nil {
				return output, err
			}
			return output, nil
		}
	}

	if !done { // maxTry reached
		e := fmt.Sprintf("Timed out after %d x %d seconds", maxTry, interval)
		return BatchResponse{}, errors.New(e)
	}

	return BatchResponse{}, errors.New("Processing failed")
}

// helper functions

func postJSON(apiURL string, data interface{}) ([]byte, error) {

	reqJSON, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	client := new(http.Client)
	reader := bytes.NewReader(reqJSON)
	res, err := client.Post(apiURL, "application/json", reader)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		e := fmt.Sprintf("POST: Expecting 200 or 202, but got status %d\n%s",
			res.StatusCode,
			string(body))
		return []byte{}, errors.New(e)
	}

	return body, nil
}

func getJSON(apiURL string) (int, []byte, error) {

	client := new(http.Client)
	res, err := client.Get(apiURL)
	if err != nil {
		return 0, []byte{}, err
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		return res.StatusCode, body, err
	}

	if res.StatusCode == 202 {
		return res.StatusCode, []byte{}, nil
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	e := fmt.Sprintf("GET: Expecting 200 or 202, but got status %d\n%s",
		res.StatusCode,
		string(body))
	return 0, []byte{}, errors.New(e)
}
