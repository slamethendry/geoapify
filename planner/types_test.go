package planner_test

import (
	"encoding/json"
	"testing"

	p "github.com/slamethendry/geoapify/planner"
	"github.com/stretchr/testify/assert"
)

func TestLonLat(t *testing.T) {

	const lat = 1.123456789
	const lon = 5.987654321
	var ll p.LonLat
	ll.SetLat(lat)
	ll.SetLon(lon)

	assert.Equal(t, lat, ll.Lat())
	assert.Equal(t, lon, ll.Lon())
}

func TestResponse(t *testing.T) {

	assert.True(t, json.Valid(resJSON))

	jsonRes := new(p.Plan)
	err := json.Unmarshal(resJSON, jsonRes)
	assert.Nil(t, err)

	assert.Equal(t, res.Type, jsonRes.Type)
	assert.Equal(t, res.Params, jsonRes.Params)

	assert.EqualValues(t, res.Features[0], jsonRes.Features[0])
}

func TestRequest(t *testing.T) {

	assert.True(t, json.Valid(reqJSON))

	reqFromJSON := new(p.Request)
	err := json.Unmarshal(reqJSON, reqFromJSON)
	assert.Nil(t, err)

	assert.EqualValues(t, *reqFromJSON, req)
}
