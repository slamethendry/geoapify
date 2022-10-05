package geoapify

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const staticMapURL = "https://maps.geoapify.com/v1/staticmap"

// URL produces the url to a map for the given StaticMap object.
// It does not fetch the map, but the url can be used by other function
// to fetch the map image file.
// On error, it gives "".
func (s StaticMap) URL() string {

	if key := os.Getenv("GEOAPIFY_KEY"); key != "" {
		return fmt.Sprintf("%s?%s&apiKey=%s", staticMapURL, s.String(), key)
	}

	return ""
}

// Get fetches a map image file content (as []byte) for the given StaticMap
// object. The content can then be served by a web server or saved as a file.
func (s StaticMap) Get() ([]byte, error) {

	url := s.URL()
	if url == "" {
		return []byte{}, errors.New("Error: Empty URL")
	}

	http.Get(url)
	resp, err := http.Get(url)
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
