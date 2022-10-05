package staticmap_test

import (
	"net/url"
	"testing"

	smap "github.com/slamethendry/geoapify"
	"github.com/stretchr/testify/assert"
)

var gps1 = smap.LonLat{-74.044724, 40.693664}

var gps2 = smap.LonLat{-74.041724, 40.692664}

var gps3 = smap.LonLat{-74.055555, 40.777777}

var style = smap.GeoMetryStyle{
	LineColor:   "#ff6600",
	LineOpacity: 0.3,
	LineWidth:   5,
	LineStyle:   "solid",
	FillColor:   "#6600ff",
	FillOpacity: 0.8,
}

var c = smap.Circle{
	LonLat:        gps1,
	Radius:        40,
	GeoMetryStyle: style,
}

var rect = smap.Rectangle{
	Area:          smap.Area{GPS1: gps1, GPS2: gps2},
	GeoMetryStyle: style,
}

var polyline = smap.PolyLine{
	Coordinates:   []smap.LonLat{gps1, gps2, gps3},
	GeoMetryStyle: style,
}

var polygon = smap.Polygon{
	Coordinates:   []smap.LonLat{gps1, gps2, gps3},
	GeoMetryStyle: style,
}

func TestGPS(t *testing.T) {

	area := smap.Area{GPS1: gps1, GPS2: gps2}
	assert.Equal(t, "-74.044724,40.693664,-74.041724,40.692664", area.String())
	area = smap.Area{}
	assert.Equal(t, "", area.String())
}

func TestGeometryStyle(t *testing.T) {

	assert.Equal(t, ";linecolor:%23ff6600;lineopacity:0.3;linewidth:5;linestyle:solid;fillcolor:%236600ff;fillopacity:0.8", style.String())
}

func TestCircle(t *testing.T) {

	assert.Equal(t, "circle:-74.044724,40.693664,40;linecolor:%23ff6600;lineopacity:0.3;linewidth:5;linestyle:solid;fillcolor:%236600ff;fillopacity:0.8", c.String())
}

func TestRectangle(t *testing.T) {

	assert.Equal(t, "rect:-74.044724,40.693664,-74.041724,40.692664;linecolor:%23ff6600;lineopacity:0.3;linewidth:5;linestyle:solid;fillcolor:%236600ff;fillopacity:0.8", rect.String())
}

func TestPolyline(t *testing.T) {

	assert.Equal(t, "polyline:-74.044724,40.693664,-74.041724,40.692664,-74.055555,40.777777;linecolor:%23ff6600;lineopacity:0.3;linewidth:5;linestyle:solid", polyline.String())
}

func TestPolygon(t *testing.T) {

	assert.Equal(t, "polygon:-74.044724,40.693664,-74.041724,40.692664,-74.055555,40.777777;linecolor:%23ff6600;lineopacity:0.3;linewidth:5;linestyle:solid;fillcolor:%236600ff;fillopacity:0.8", polygon.String())
}

func TestMarker(t *testing.T) {

	markerStyle := smap.Marker{
		Type:        "material",
		Color:       "#1f63e6",
		Size:        "x-large",
		Icon:        "cloud",
		IconType:    "awesome",
		WhiteCircle: "no",
	}

	m1 := markerStyle
	m1.LonLat = smap.LonLat{-122.676481, 45.524460}
	m1.Text = "1"
	m2 := markerStyle
	m2.LonLat = smap.LonLat{-122.671296, 45.523095}
	m2.Text = "2"
	m3 := markerStyle
	m3.LonLat = smap.LonLat{-122.664446, 45.522964}
	m3.Text = "3"

	assert.Equal(t, "lonlat:-122.676481,45.524460;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:1;whitecircle:no|lonlat:-122.671296,45.523095;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:2;whitecircle:no|lonlat:-122.664446,45.522964;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:3;whitecircle:no", m1.String()+"|"+m2.String()+"|"+m3.String())
}

func TestStaticMap(t *testing.T) {

	s := smap.StaticMap{
		Center:      smap.LonLat{-122.68464, 45.510378},
		Style:       "osm-liberty",
		Zoom:        15.8,
		Bearing:     54,
		Pitch:       52,
		Width:       600,
		Height:      400,
		ScaleFactor: 2,
	}

	expected, _ := url.ParseQuery("style=osm-liberty&width=600&height=400&center=lonlat:-122.684640,45.510378&zoom=15.8&pitch=52&bearing=54&scaleFactor=2")

	out, _ := url.ParseQuery(s.String())
	assert.Equal(t, expected, out)
}

func TestStaticMapWithMarkers(t *testing.T) {

	expected, _ := url.ParseQuery("style=osm-bright-grey&width=600&height=400&center=lonlat:-122.670651,45.522488&zoom=14.8&marker=lonlat:-122.676481,45.524460;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:1;whitecircle:no|lonlat:-122.671296,45.523095;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:2;whitecircle:no|lonlat:-122.664446,45.522964;type:material;color:%231f63e6;size:x-large;icon:cloud;icontype:awesome;text:3;whitecircle:no")

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

	out, _ := url.ParseQuery(s.String())
	assert.Equal(t, expected, out)
}

func TestStaticMapWithGeometry(t *testing.T) {

	expected, _ := url.ParseQuery("style=osm-carto&width=600&height=400&zoom=8.9&geometry=circle:-74.044724,40.693664,50;linewidth:5;linecolor:%23ff6600;fillcolor:%236600ff;lineopacity:0.3;fillopacity:0.8%7Ccircle:-74.04372450744129, 40.61366453643252,50;linewidth:5;linecolor:%23ff6600;fillcolor:%236600ff;lineopacity:0.3;fillopacity:0.6")

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

	out, _ := url.ParseQuery(s.String())
	assert.Equal(t, expected, out)
}
