package planner_test

import (
	"encoding/json"
	"os"
	"testing"

	p "github.com/slamethendry/geoapify"

	"github.com/stretchr/testify/assert"
)

func TestPostExampleRequest(t *testing.T) {

	key := os.Getenv("GEOAPIFY_KEY")
	assert.NotEmpty(t, key)

	plan, err := req.Post(key) // req: refer to types_data_test.go
	assert.Nil(t, err)
	assert.NotEmpty(t, plan)

	assert.True(t, json.Valid(reqJSON)) // reqJSON: refer to types_data_test.go
	reqFromJSON := new(p.Request)
	err = json.Unmarshal(reqJSON, reqFromJSON)
	assert.Nil(t, err)

	assert.EqualValues(t, *reqFromJSON, req) // same request

	planFromJSON, err := reqFromJSON.Post(key)
	assert.Nil(t, err)
	assert.EqualValues(t, plan, planFromJSON) // same output
}

func TestPostSimpleRequest(t *testing.T) {

	key := os.Getenv("GEOAPIFY_KEY")
	assert.NotEmpty(t, key)

	assert.True(t, json.Valid([]byte(simpleJSON))) // see types_data_test.go
	reqFromJSON := new(p.Request)
	err := json.Unmarshal([]byte(simpleJSON), reqFromJSON)
	assert.Nil(t, err)

	planFromJSON, err := reqFromJSON.Post(key)
	assert.Nil(t, err)
	assert.NotEmpty(t, planFromJSON)

	plan, err := simpleRequest.Post(key) // see types_data_test.go
	assert.Nil(t, err)
	assert.NotEmpty(t, plan)

	assert.EqualValues(t, simpleRequest, *reqFromJSON) // same request
	assert.EqualValues(t, plan, planFromJSON)          // same output

	br := p.BatchRequest{
		Inputs: []p.BatchInput{
			{
				Body: simpleRequest,
			},
		},
	}
	res, err := br.Post(key, 6, 60)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	var planFromBatch p.Plan = res.Results[0].Plan
	jsonPlan, err := json.MarshalIndent(planFromBatch, "", " ")
	assert.NotEmpty(t, string(jsonPlan))
	assert.EqualValues(t, plan, planFromBatch) // same output
}
