package geoapify

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const staticMapURL = "https://maps.geoapify.com/v1/staticmap"

// URL produces the url to a map for the given StaticMap object.
// It does not fetch the map, but the url can be used by other function
// to fetch the map image file.
// On error, it gives "".
func (s StaticMap) URL(apiKey string) string {

	if apiKey == "" {
		return ""
	}

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

	return fmt.Sprintf("%s?%s&apiKey=%s", staticMapURL, out, apiKey)
}

// Get fetches a map image file content (as []byte) for the given StaticMap
// object. The content can then be served by a web server or saved as a file.
func (s StaticMap) Get(apiKey string) ([]byte, error) {

	if apiKey == "" {
		return []byte{}, errors.New("Error: Missing API key")
	}

	resp, err := http.Get(s.URL(apiKey))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (s StaticMap) values() map[string]string {

	// If we use url.Values.Encode(), the output looks slightly different
	// than Geoapify spec. So we go custom to minimise potential issue.

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
