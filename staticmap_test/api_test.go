package staticmap_test

import (
	"net/url"
	"os"
	"testing"

	smap "github.com/slamethendry/geoapify"
	"github.com/stretchr/testify/assert"
)

func TestSimpleStaticMap(t *testing.T) {

	key := os.Getenv("GEOAPIFY_KEY")
	s := smap.StaticMap{
		Center:      smap.LonLat{-122.684640, 45.510378},
		Style:       "osm-liberty",
		Zoom:        15.8,
		Bearing:     54,
		Pitch:       52,
		Width:       600,
		Height:      400,
		ScaleFactor: 2,
	}

	expected, _ := url.ParseRequestURI("https://maps.geoapify.com/v1/staticmap?style=osm-liberty&width=600&height=400&center=lonlat:-122.684640,45.510378&zoom=15.8&pitch=52&bearing=54&scaleFactor=2&apiKey=" + key)

	out, _ := url.ParseRequestURI(s.URL(key))

	assert.NotEmpty(t, out)
	assert.Equal(t, expected.Scheme, out.Scheme)
	assert.Equal(t, expected.Host, out.Host)
	assert.Equal(t, expected.Path, out.Path)
	assert.Equal(t, expected.Query(), out.Query())
}

func TestStaticMapWithGeometry(t *testing.T) {

	key := os.Getenv("GEOAPIFY_KEY")
	expected, _ := url.ParseRequestURI("https://maps.geoapify.com/v1/staticmap?style=osm-carto&width=600&height=400&zoom=8.9&geometry=circle:-74.044724,40.693664,50;linewidth:5;linecolor:%23ff6600;fillcolor:%236600ff;lineopacity:0.3;fillopacity:0.8%7Ccircle:-74.04372450744129, 40.61366453643252,50;linewidth:5;linecolor:%23ff6600;fillcolor:%236600ff;lineopacity:0.3;fillopacity:0.6&apiKey=" + key)

	geometry := []interface{}{
		smap.Circle{
			LonLat: smap.LonLat{-74.044724, 40.693664},
			Radius: 50,
			GeoMetryStyle: smap.GeoMetryStyle{
				LineWidth:   5,
				LineColor:   "#ff6600",
				FillColor:   "#6600ff",
				LineOpacity: 0.8,
				FillOpacity: 0.8,
			},
		},
		smap.Circle{
			LonLat: smap.LonLat{-74.043725, 40.613665},
			Radius: 50,
			GeoMetryStyle: smap.GeoMetryStyle{
				LineWidth:   5,
				LineColor:   "#ff6600",
				FillColor:   "#6600ff",
				LineOpacity: 0.3,
				FillOpacity: 0.6,
			},
		},
	}

	s := smap.StaticMap{
		Style:    "osm-carto",
		Width:    600,
		Height:   400,
		Zoom:     8.9,
		Geometry: geometry,
	}

	out, _ := url.ParseRequestURI(s.URL(key))

	assert.NotEmpty(t, out)
	assert.Equal(t, expected.Scheme, out.Scheme)
	assert.Equal(t, expected.Host, out.Host)
	assert.Equal(t, expected.Path, out.Path)
	assert.Equal(t, expected.Query(), out.Query())
}

func TestStaticMapWithMarkers(t *testing.T) {

	key := os.Getenv("GEOAPIFY_KEY")
	expected, _ := url.ParseRequestURI("https://maps.geoapify.com/v1/staticmap?style=osm-bright-grey&width=600&height=400&center=lonlat:-122.670651,45.522488&zoom=14.8&marker=lonlat:-122.676481,45.524460;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:1;whitecircle:no|lonlat:-122.671296,45.523095;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:2;whitecircle:no|lonlat:-122.664446,45.522964;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:3;whitecircle:no&apiKey=" + key)

	s := smap.StaticMap{
		Style:  "osm-bright-grey",
		Width:  600,
		Height: 400,
		Center: smap.LonLat{-122.670651, 45.522488},
		Zoom:   14.8,
		Marker: []smap.Marker{
			{
				LonLat:      smap.LonLat{-122.676481, 45.524460},
				Type:        "material",
				Color:       "#1f63e6",
				Size:        "x-large",
				Icon:        "cloud",
				IconType:    "awesome",
				Text:        "1",
				WhiteCircle: "no",
			},
			{
				LonLat:      smap.LonLat{-122.671296, 45.523095},
				Type:        "material",
				Color:       "#1f63e6",
				Size:        "x-large",
				Icon:        "cloud",
				IconType:    "awesome",
				Text:        "2",
				WhiteCircle: "no",
			},
			{
				LonLat:      smap.LonLat{-122.664446, 45.522964},
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

	out, _ := url.ParseRequestURI(s.URL(key))

	assert.NotEmpty(t, out)
	assert.Equal(t, expected.Scheme, out.Scheme)
	assert.Equal(t, expected.Host, out.Host)
	assert.Equal(t, expected.Path, out.Path)
	assert.Equal(t, expected.Query(), out.Query())
}
