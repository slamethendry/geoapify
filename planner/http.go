package planner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
