package staticmap

import (
	"fmt"
	"net/url"
)

// types2string converts types to Geoapify API string format.
// Always returns empty string on error.

func (c Coordinate) string() string {

	// For routing, output of 6 decimal places is plenty accurate
	// Ref: https://en.wikipedia.org/wiki/Decimal_degrees
	return fmt.Sprintf("%.6f", c)
}

func (g GPS) string() string {

	/* lon := g.Lon.String()
	lat := g.Lat.String()

	if lon == "" || lat == "" {
		return ""
	} */

	empty := GPS{}
	if g == empty {
		return ""
	}

	//return fmt.Sprintf("%s,%s", lon, lat)
	return fmt.Sprintf("%s,%s", g.Lon.string(), g.Lat.string())
}

func (g GeoMetryStyle) string() string {

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

type gpsList []GPS

func (coords gpsList) string() string {

	out := ""
	arrayLen := len(coords)
	if arrayLen == 0 {
		return ""
	}

	out = coords[0].string()
	if out == "" {
		return ""
	}

	for i := 1; i < arrayLen; i++ {
		gps := coords[i].string()
		if gps == "" {
			return ""
		}
		out = fmt.Sprintf("%s,%s", out, gps)
	}

	return out
}

func (c Circle) string() string {

	gps := c.GPS.string()

	if gps == "" || c.Radius < 1 {
		return ""
	}

	return fmt.Sprintf("circle:%s,%d%s", gps, c.Radius, c.GeoMetryStyle.string())
}

func (a Area) string() string {

	return gpsList([]GPS{a.GPS1, a.GPS2}).string()
}

func (r Rectangle) string() string {

	return fmt.Sprintf("rect:%s%s", r.Area.string(), r.GeoMetryStyle.string())
}

func (p PolyLine) string() string {

	coords := gpsList(p.Coordinates).string()
	if coords == "" {
		return ""
	}

	p.GeoMetryStyle.FillColor = ""
	p.GeoMetryStyle.FillOpacity = -1

	return fmt.Sprintf("polyline:%s%s", coords, p.GeoMetryStyle.string())
}

func (p Polygon) string() string {

	coords := gpsList(p.Coordinates).string()
	if coords == "" {
		return ""
	}

	return fmt.Sprintf("polygon:%s%s", coords, p.GeoMetryStyle.string())
}

func (m Marker) string() string {

	gps := m.GPS.string()
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
		return (t.(Circle)).string()
	case Rectangle:
		return (t.(Rectangle)).string()
	case PolyLine:
		return (t.(PolyLine)).string()
	case Polygon:
		return (t.(Polygon)).string()
	default:
		return ""
	}
}

func (s StaticMap) string() string {

	/* If we use url.Values.Encode(), the output looks slightly different
	   than Geoapify spec. So we go custom. */

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

	if center := s.Center.string(); center != "" {
		v["center"] = fmt.Sprintf("lonlat:%s", center)
	}

	if area := s.Area.string(); area != "" { // ??
		v["area"] = fmt.Sprintf("rect:%s", area)
	}

	marker := ""
	if numMarker := len(s.Marker); numMarker > 0 {
		marker = s.Marker[0].string()
		for i := 1; i < numMarker; i++ {
			marker = fmt.Sprintf("%s|%s", marker, s.Marker[i].string())
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
