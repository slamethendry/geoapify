package geoapify

import (
	"fmt"
	"net/url"
)

type Area struct {
	GPS1 LonLat
	GPS2 LonLat
}

type Marker struct {
	LonLat
	Color       string // Named | #hex
	Type        string // material | awesome | circle
	Size        string // small | medium | large | x-large | xx-large
	Icon        string // Icon name
	IconSize    string
	IconType    string
	Text        string // Marker text to be used instead of icon
	TextSize    string
	WhiteCircle string // yes | no
	Shadow      string // auto | no
	ShadowColor string
	StrokeColor string
}

type GeoMetryStyle struct {
	LineColor   string  // word | #hex
	LineOpacity float32 // 0.1 .. 1.0
	LineWidth   uint8   // pixel
	LineStyle   string  // solid | dotted | dashed | longdash
	FillOpacity float32 // 0.1 .. 1.0; ignored in PolyLine
	FillColor   string  // Ignored in PolyLine
}

type Circle struct {
	LonLat
	Radius uint16 // pixel
	GeoMetryStyle
}

type Rectangle struct {
	Area
	GeoMetryStyle
}

type PolyLine struct {
	Coordinates []LonLat
	GeoMetryStyle
}

type Polygon struct {
	Coordinates []LonLat
	GeoMetryStyle
}

type StaticMap struct {
	Area
	Center      LonLat
	Style       string  // apidocs.geoapify.com/docs/maps/map-tiles/#about
	Width       uint16  // pixel
	Height      uint16  // pixel
	Format      string  // jpeg | png; default is jpeg
	Zoom        float32 // 1.0 .. 20.0
	Pitch       uint8   // 1 .. 60
	Bearing     uint16  // 1 .. 360
	ScaleFactor uint8   // 1 | 2
	Marker      []Marker
	Geometry    []interface{}
	// GeoJSON: not implemented
}

// String() converts types to Geoapify API string format.
// Always returns empty string on error.

func (ll LonLat) String() string {

	if ll == [2]float64{} {
		return ""
	}

	lon := fmt.Sprintf("%.6f", ll.Lon())
	lat := fmt.Sprintf("%.6f", ll.Lat())
	return fmt.Sprintf("%s,%s", lon, lat)
}

func (g GeoMetryStyle) String() string {

	lineColor := fmt.Sprintf("%s", g.LineColor)
	if lineColor != "" {
		lineColor = fmt.Sprintf(";linecolor:%s", url.PathEscape(lineColor))
	}

	lineOpacity := ""
	if g.LineOpacity > 0 && g.LineOpacity <= 1 {
		lineOpacity = fmt.Sprintf(";lineopacity:%.1f", g.LineOpacity)
	}

	lineWidth := fmt.Sprintf("%d", g.LineWidth)
	if lineWidth != "" {
		lineWidth = fmt.Sprintf(";linewidth:%s", lineWidth)
	}

	lineStyle := ""
	switch g.LineStyle {
	case "solid", "dotted", "dashed", "longdash":
		lineStyle = fmt.Sprintf(";linestyle:%s", g.LineStyle)
	}

	fillColor := fmt.Sprintf("%s", g.FillColor)
	if fillColor != "" {
		fillColor = fmt.Sprintf(";fillcolor:%s", url.PathEscape(fillColor))
	}

	fillOpacity := ""
	if g.FillOpacity > 0 && g.FillOpacity <= 1 {
		fillOpacity = fmt.Sprintf(";fillopacity:%.1f", g.FillOpacity)
	}

	return fmt.Sprintf("%s%s%s%s%s%s", lineColor, lineOpacity, lineWidth, lineStyle, fillColor, fillOpacity)
}

type gpsList []LonLat

func (coords gpsList) String() string {

	out := ""
	arrayLen := len(coords)
	if arrayLen == 0 {
		return ""
	}

	out = coords[0].String()
	if out == "" {
		return ""
	}

	for i := 1; i < arrayLen; i++ {
		gps := coords[i].String()
		if gps == "" {
			return ""
		}
		out = fmt.Sprintf("%s,%s", out, gps)
	}

	return out
}

func (c Circle) String() string {

	gps := c.LonLat.String()

	if gps == "" || c.Radius < 1 {
		return ""
	}

	return fmt.Sprintf("circle:%s,%d%s", gps, c.Radius, c.GeoMetryStyle.String())
}

func (a Area) String() string {

	return gpsList([]LonLat{a.GPS1, a.GPS2}).String()
}

func (r Rectangle) String() string {

	return fmt.Sprintf("rect:%s%s", r.Area.String(), r.GeoMetryStyle.String())
}

func (p PolyLine) String() string {

	coords := gpsList(p.Coordinates).String()
	if coords == "" {
		return ""
	}

	p.GeoMetryStyle.FillColor = ""
	p.GeoMetryStyle.FillOpacity = -1

	return fmt.Sprintf("polyline:%s%s", coords, p.GeoMetryStyle.String())
}

func (p Polygon) String() string {

	coords := gpsList(p.Coordinates).String()
	if coords == "" {
		return ""
	}

	return fmt.Sprintf("polygon:%s%s", coords, p.GeoMetryStyle.String())
}

func (m Marker) String() string {

	gps := m.LonLat.String()
	if gps == "" {
		return ""
	}

	color := ""
	if m.Color != "" {
		color = fmt.Sprintf(";color:%s", url.PathEscape(m.Color))
	}

	mType := ""
	switch m.Type {
	case "material", "awesome", "circle":
		mType = fmt.Sprintf(";type:%s", m.Type)
	}

	size := ""
	switch m.Size {
	case "small", "medium", "large", "x-large", "xx-large":
		size = fmt.Sprintf(";size:%s", m.Size)
	}

	icon := ""
	if m.Icon != "" {
		icon = fmt.Sprintf(";icon:%s", m.Icon)
	}

	iconSize := ""
	switch m.IconSize {
	case "small", "medium", "large":
		size = fmt.Sprintf(";iconsize:%s", m.IconSize)
	}

	iconType := ""
	switch m.IconType {
	case "material", "awesome":
		iconType = fmt.Sprintf(";icontype:%s", m.IconType)
	}

	text := ""
	if m.Text != "" {
		text = fmt.Sprintf(";text:%s", url.PathEscape(m.Text))
	}

	textSize := ""
	switch m.TextSize {
	case "small", "medium", "large":
		textSize = fmt.Sprintf(";textsize:%s", m.TextSize)
	}

	whiteCircle := ""
	switch m.WhiteCircle {
	case "yes", "no":
		whiteCircle = fmt.Sprintf(";whitecircle:%s", m.WhiteCircle)
	}

	shadow := ""
	switch m.Shadow {
	case "auto", "no":
		shadow = fmt.Sprintf(";shadow:%s", m.Shadow)
	}

	shadowColor := ""
	if m.ShadowColor != "" {
		shadowColor = fmt.Sprintf(";shadowcolor:%s", url.PathEscape(m.ShadowColor))
	}

	strokeColor := ""
	if m.StrokeColor != "" {
		strokeColor = fmt.Sprintf(";strokecolor:%s", url.PathEscape(m.StrokeColor))
	}

	return fmt.Sprintf("lonlat:%s%s%s%s%s%s%s%s%s%s%s%s%s", gps, mType, color, size, icon, iconSize, iconType, text, textSize, whiteCircle, shadow, shadowColor, strokeColor)
}

func stringOf(t interface{}) string {

	switch t.(type) {
	case Circle:
		return (t.(Circle)).String()
	case Rectangle:
		return (t.(Rectangle)).String()
	case PolyLine:
		return (t.(PolyLine)).String()
	case Polygon:
		return (t.(Polygon)).String()
	default:
		return ""
	}
}

func (s StaticMap) String() string {

	// If we use url.Values.Encode(), the output looks slightly different
	// than Geoapify spec. So we go custom to minimise potential issue.

	out := ""
	v := s.values()
	i := 0
	for k := range v {
		if i == 0 {
			out = fmt.Sprintf("%s=%s", k, v[k])
		} else {
			out = fmt.Sprintf("%s&%s=%s", out, k, v[k])
		}
		i++
	}
	return out
}

func (s StaticMap) values() map[string]string {

	v := make(map[string]string)

	switch s.Style {
	case "osm-carto", "osm-bright", "osm-bright-grey", "osm-bright-smooth", "klokantech-basic", "osm-liberty", "maptiler-3d", "toner", "toner-grey", "positron", "positron-blue", "positron-red", "dark-matter", "dark-matter-brown", "dark-matter-dark-grey", "dark-matter-dark-purple", "dark-matter-purple-roads", "dark-matter-yellow-roads":
		v["style"] = s.Style
	default:
		v["style"] = "osm-bright"
	}

	if center := s.Center.String(); center != "" {
		v["center"] = fmt.Sprintf("lonlat:%s", center)
	}

	if area := s.Area.String(); area != "" { // ??
		v["area"] = fmt.Sprintf("rect:%s", area)
	}

	marker := ""
	if numMarker := len(s.Marker); numMarker > 0 {
		marker = s.Marker[0].String()
		for i := 1; i < numMarker; i++ {
			marker = fmt.Sprintf("%s|%s", marker, s.Marker[i].String())
		}
	}
	if marker != "" {
		v["marker"] = marker
	}

	geometry := ""
	if numGeo := len(s.Geometry); numGeo > 0 {
		geometry = fmt.Sprintf("%s", stringOf(s.Geometry[0]))
		for i := 1; i < numGeo; i++ {
			geometry = fmt.Sprintf("%s|%s", geometry, stringOf(s.Geometry[i]))
		}
	}
	if geometry != "" {
		v["geometry"] = geometry
	}

	switch s.Format {
	case "jpeg", "png":
		v["format"] = s.Format
	}

	if s.Width > 0 {
		v["width"] = fmt.Sprintf("%d", s.Width)
	}

	if s.Height > 0 {
		v["height"] = fmt.Sprintf("%d", s.Height)
	}

	if s.Zoom >= 1 && s.Zoom <= 20 {
		v["zoom"] = fmt.Sprintf("%.1f", s.Zoom)
	}

	if s.Pitch > 0 && s.Pitch <= 60 {
		v["pitch"] = fmt.Sprintf("%d", s.Pitch)
	}

	if s.Bearing > 0 && s.Bearing <= 360 {
		v["bearing"] = fmt.Sprintf("%d", s.Bearing)
	}

	if s.ScaleFactor == 1 || s.ScaleFactor == 2 {
		v["scaleFactor"] = fmt.Sprintf("%d", s.ScaleFactor)
	}

	return v
}
