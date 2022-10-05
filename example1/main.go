package main

import (
	"fmt"
	"os"

	s "github.com/slamethendry/geoapify/staticmap"
)

func main() {

	apiKey := os.Getenv("GEOAPIFY_KEY")
	if apiKey == "" {
		panic("API key is missing")
	}

	// The request example is small, so we can use synchronous API directly
	// plan, err := request.Post(apiKey)

	// Or we can use batch request API regardless of job size
	response, err := batchRequest.Post(apiKey, 8, 60)
	if err != nil {
		panic(err)
	}

	plan := response.Results[0].Plan

	// Once we get the resulting route plan, we map it using static map API
	// with map markers with start location at the Office and ending at each
	// agent's end location.

	smap := s.StaticMap{Style: "osm-bright"}

	officeMarker := s.Marker{
		LonLat:      s.LonLat(office),
		Text:        "O",
		WhiteCircle: "no",
	}

	smap.Marker = append(smap.Marker, officeMarker)

	for _, feature := range plan.Features {

		fmt.Println("\nAgent: ", agents[feature.Properties.AgentIndex].ID)

		for i := 0; i < len(feature.Properties.Actions); i++ {

			action := feature.Properties.Actions[i]

			if action.Type == "job" {

				fmt.Println(action.Type, action.JobID)
				m := s.Marker{
					LonLat: s.LonLat(jobs[action.JobIndex].LonLat),
					Text:   fmt.Sprintf("%d", i), // Type "job" begins at i == 1
					Color:  agentColours[feature.Properties.AgentIndex],
				}
				smap.Marker = append(smap.Marker, m)
			}
		}

		end := s.Marker{
			LonLat:      s.LonLat(agents[feature.Properties.AgentIndex].End),
			Text:        "E",
			WhiteCircle: "no",
			Color:       agentColours[feature.Properties.AgentIndex],
		}
		smap.Marker = append(smap.Marker, end)

		fmt.Printf("%.1f hours\n", float32(feature.Properties.Time)/3600)
		fmt.Printf("%.1f km\n", float32(feature.Properties.Distance)/1000)

	}
	fmt.Printf("\n%s\n", smap.URL())
}
