package planner_test

import (
	"encoding/json"
	"os"
	"testing"

	p "github.com/slamethendry/geoapify/planner"

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

	agent1 := p.Agent{
		Start: p.LonLat{106.72683176435947, -6.28476945},
		End:   p.LonLat{106.7457861328624, -6.29673725},
	}

	customers := []p.Job{
		{
			ID:       "order-xyz",
			Location: p.LonLat{106.7213665, -6.2895597},
			Duration: 300,
		},
		{
			ID:       "order-abc",
			Location: p.LonLat{106.7157832, -6.2853618},
			Duration: 300,
		},
	}

	r := p.Request{
		Mode:   "drive",
		Agents: []p.Agent{agent1},
		Jobs:   customers,
	}

	simpleJSON := `{
    "mode": "drive",
    "agents": [
        {
            "start_location": [
                106.72683176435947,
                -6.28476945
            ],
            "end_location": [
                106.7457861328624,
                -6.29673725
            ]
        }
    ],
    "jobs": [
        {
            "id": "order-xyz",
            "location": [
                106.7213665,
                -6.2895597
            ],
            "duration": 300
        },
        {
            "id": "order-abc",
            "location": [
                106.7157832,
                -6.2853618
            ],
            "duration": 300
        }
    ]
}`

	assert.True(t, json.Valid([]byte(simpleJSON)))
	reqFromJSON := new(p.Request)
	err := json.Unmarshal([]byte(simpleJSON), reqFromJSON)
	assert.Nil(t, err)

	assert.EqualValues(t, r, *reqFromJSON) // same request

	planFromJSON, err := reqFromJSON.Post(key)
	assert.Nil(t, err)
	assert.NotEmpty(t, planFromJSON)

	plan, err := r.Post(key)
	assert.Nil(t, err)

	assert.EqualValues(t, plan, planFromJSON) // same output
}
