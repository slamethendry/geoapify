// Package geoapify/planner is a wrapper for Geoapify
// [Route Planner](https://apidocs.geoapify.com/docs/route-planner/) for use in
// the backend server.
//
// Prerequisite: API key is assumed to be available in the environment
// variable GEOAPIFY_KEY.
package planner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const v1URL = "https://api.geoapify.com/v1/routeplanner"

func (r Request) Post(key string) (Plan, error) {

	apiURL := fmt.Sprintf("%s?apiKey=%s", v1URL, key)

	reqJSON, err := json.Marshal(r)
	if err != nil {
		return Plan{}, err
	}

	client := new(http.Client)
	reader := bytes.NewReader(reqJSON)
	res, err := client.Post(apiURL, "application/json", reader)
	if err != nil {
		return Plan{}, err
	}

	if res.StatusCode != 200 {
		e := fmt.Sprintf("Expecting 200, but got status: %d", res.StatusCode)
		return Plan{}, errors.New(e)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
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

/*
func (r Request) String() string {

	rJSON, err := json.Marshal(r)

	if err != nil {
		return ""
	}

	return string(rJSON)
}
*/
