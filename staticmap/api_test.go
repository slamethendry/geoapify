package staticmap_test

import (
	"net/url"
	"os"
	"testing"

	smap "github.com/slamethendry/geoapify/staticmap"
	"github.com/stretchr/testify/assert"
)

func TestStaticMapWithMarkersURL(t *testing.T) {

	key := os.Getenv("GEOAPIFY_KEY")
	expected, _ := url.ParseRequestURI("https://maps.geoapify.com/v1/staticmap?style=osm-bright-grey&width=600&height=400&center=lonlat:-122.670651,45.522488&zoom=14.8&marker=lonlat:-122.676481,45.524460;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:1;whitecircle:no|lonlat:-122.671296,45.523095;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:2;whitecircle:no|lonlat:-122.664446,45.522964;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:3;whitecircle:no&apiKey=" + key)

	s := smap.StaticMap{
		Style:  "osm-bright-grey",
		Width:  600,
		Height: 400,
		Center: smap.GPS{Lon: -122.670651, Lat: 45.522488},
		Zoom:   14.8,
		Marker: []smap.Marker{
			{
				GPS:         smap.GPS{Lon: -122.676481, Lat: 45.524460},
				Type:        "material",
				Color:       "#1f63e6",
				Size:        "x-large",
				Icon:        "cloud",
				IconType:    "awesome",
				Text:        "1",
				WhiteCircle: "no",
			},
			{
				GPS:         smap.GPS{Lon: -122.671296, Lat: 45.523095},
				Type:        "material",
				Color:       "#1f63e6",
				Size:        "x-large",
				Icon:        "cloud",
				IconType:    "awesome",
				Text:        "2",
				WhiteCircle: "no",
			},
			{
				GPS:         smap.GPS{Lon: -122.664446, Lat: 45.522964},
				Type:        "material",
				Color:       "#1f63e6",
				Size:        "x-large",
				Icon:        "cloud",
				IconType:    "awesome",
				Text:        "3",
				WhiteCircle: "no",
			},
		},
	}

	out, _ := url.ParseRequestURI(s.URL())

	assert.NotEmpty(t, out)
	assert.Equal(t, expected.Scheme, out.Scheme)
	assert.Equal(t, expected.Host, out.Host)
	assert.Equal(t, expected.Path, out.Path)
	assert.Equal(t, expected.Query(), out.Query())
	// assert.Empty(t, s.URL())
}
