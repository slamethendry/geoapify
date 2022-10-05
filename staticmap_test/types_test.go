package staticmap_test

import (
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
