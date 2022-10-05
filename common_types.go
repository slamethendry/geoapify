// Geoapify is a wrapper for Geoapify.com's API. For now it only covers:
//
// [Static Maps](https://apidocs.geoapify.com/docs/maps/static/) and
//
// [Route Planner](https://apidocs.geoapify.com/docs/route-planner)
//
// Prerequisite: API key
package geoapify

// LonLat stores the GPS longitude and latitude coordinates.
// Some APIs use lonlat and others use latlon, so helpers functions can be
// used: Lon(), Lat(), SetLon(), SetLat().
type LonLat [2]float64

// Coordinate cannot use float32 due to rounding inconsistency at the 6th digit

// Lon gets the longitude coordinate value
func (ll LonLat) Lon() float64 { return ll[0] }

// / Lat gets the latitude coordinate value
func (ll LonLat) Lat() float64 { return ll[1] }

// SetLon sets the longitude coordinate value
func (ll *LonLat) SetLon(lon float64) { ll[0] = lon }

// SetLat sets the latitude coordinate value
func (ll *LonLat) SetLat(lat float64) { ll[1] = lat }
