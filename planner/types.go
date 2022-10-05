package planner

// We may need to use json.Number instead of uint to represent number, because
// ",omitempty" will treat value of 0 as "empty" and not output it in the JSON
// but index of 0 is a valid data.
// https://stackoverflow.com/questions/38486564/unmarshal-marshal-json-with-int-set-to-0-does-not-seem-to-work
// Unfortunately json.Number is string, so there will be extra work elsewhere
// The jury is still out whether we need to convert to json.Number

// LonLat stores the GPS longitude and latitude coordinates.
// Some APIs use lonlat and others use latlon, so helpers functions can be
// used: Lon(), Lat(), SetLon(), SetLat().
type LonLat [2]float64

// Lon gets the longitude coordinate value
func (ll LonLat) Lon() float64 { return ll[0] }

// / Lat gets the latitude coordinate value
func (ll LonLat) Lat() float64 { return ll[1] }

// SetLon sets the longitude coordinate value
func (ll *LonLat) SetLon(lon float64) { ll[0] = lon }

// SetLat sets the latitude coordinate value
func (ll *LonLat) SetLat(lat float64) { ll[1] = lat }

type Location struct {
	ID     string `json:"id,omitempty"`
	LonLat `json:"location"`
}

type TimeWindow [2]uint

type Break struct {
	Duration    uint         `json:"duration,omitempty"`
	TimeWindows []TimeWindow `json:"time_windows,omitempty"`
}

type Agent struct {
	Start          LonLat       `json:"start_location,omitempty"`
	StartLocIndex  uint         `json:"start_location_index,omitempty"`
	End            LonLat       `json:"end_location,omitempty"`
	EndLocIndex    uint         `json:"end_location_index,omitempty"`
	PickupCapacity uint         `json:"pickup_capacity,omitempty"`
	DlvyCapacity   uint         `json:"delivery_capacity,omitempty"`
	Capabilities   []string     `json:"capabilities,omitempty"`
	Breaks         []Break      `json:"breaks,omitempty"`
	ID             string       `json:"id,omitempty"`
	Description    string       `json:"description,omitempty"`
	TimeWindows    []TimeWindow `json:"time_windows,omitempty"`
}

type Job struct {
	LocIndex     uint         `json:"location_index,omitempty"`
	Priority     uint8        `json:"priority,omitempty"` // 0..100: 0 lowest
	Duration     uint         `json:"duration,omitempty"`
	PickupAmount uint         `json:"pickup_amount,omitempty"`
	DlvyAmount   uint         `json:"delivery_amount,omitempty"`
	Requirements []string     `json:"requirements,omitempty"`
	ID           string       `json:"id,omitempty"`
	Description  string       `json:"description,omitempty"`
	TimeWindows  []TimeWindow `json:"time_windows,omitempty"`
	Location     LonLat       `json:"location,omitempty"`
}

type Destination struct {
	LonLat      `json:"location,omitempty"`
	LocIndex    uint         `json:"location_index,omitempty"`
	Duration    uint         `json:"duration,omitempty"`
	TimeWindows []TimeWindow `json:"time_windows,omitempty"`
}

type Shipment struct {
	ID           string      `json:"id,omitempty"`
	Pickup       Destination `json:"pickup,omitempty"`
	Delivery     Destination `json:"delivery,omitempty"`
	Requirements []string    `json:"requirements,omitempty"`
	Priority     uint8       `json:"priority,omitempty"` // 0..100: 0 lowest
	Description  string      `json:"description,omitempty"`
}

type Avoid struct {
	Type       string `json:"type"` // e.g. tolls
	Importance uint8  `json:"importance,omitempty"`
}

// Request is the API to describe the target routes that need to be planned.
// Mode must be one of [walk, hike, scooter, motorcycle, drive, truck,
// light_truck, medium_truck, truck_dangerous_goods, heavy_truck,
// long_truck, bicycle, mountain_bike, road_bike, bus, drive_shortest,
// drive_traffic_approximated, truck_traffic_approximated]
type Request struct {
	Mode      string     `json:"mode,omitempty"`
	Agents    []Agent    `json:"agents,omitempty"`
	Jobs      []Job      `json:"jobs,omitempty"`
	Shipments []Shipment `json:"shipments,omitempty"`
	Locations []Location `json:"locations,omitempty"`
	Traffic   string     `json:"traffic,omitempty"` // approximated
	Avoid     []Avoid    `json:"avoid,omitempty"`
}

/* ************************ */

type Action struct {
	Type          string `json:"type"` // start|end|pickup|delivery
	StartTime     uint   `json:"start_time,omitempty"`
	Duration      uint   `json:"duration,omitempty"`
	ShipmentIndex uint   `json:"shipment_index,omitempty"`
	ShipmentID    string `json:"shipment_id,omitempty"`
	JobID         string `json:"job_id,omitempty"`
	WaypointIndex uint   `json:"waypoint_index,omitempty"`
	Index         uint   `json:"index,omitempty"`
	JobIndex      uint   `json:"job_index,omitempty"`
}

type Waypoint struct {
	OriginalLoc      LonLat   `json:"original_location,omitempty"`
	OriginalLocIndex uint     `json:"original_location_index,omitempty"`
	OriginalLocID    string   `json:"original_location_id,omitempty"`
	Location         LonLat   `json:"location"`
	StartTime        uint     `json:"start_time,omitempty"`
	Duration         uint     `json:"duration,omitempty"`
	Actions          []Action `json:"actions,omitempty"`
	PrevLegIndex     uint     `json:"prev_leg_index,omitempty"`
	NextLegIndex     uint     `json:"next_leg_index,omitempty"`
}

type Step struct {
	Distance  uint `json:"distance,omitempty"` // meters
	Time      uint `json:"time,omitempty"`     // seconds
	FromIndex uint `json:"from_index,omitempty"`
	ToIndex   uint `json:"to_index,omitempty"`
}

type Leg struct {
	Distance          uint   `json:"distance,omitempty"` // meters
	Time              uint   `json:"time,omitempty"`     // seconds
	Steps             []Step `json:"steps,omitempty"`
	FromWaypointIndex uint   `json:"from_waypoint_index,omitempty"`
	ToWaypointIndex   uint   `json:"to_waypoint_index,omitempty"`
}

type AgentPlan struct {
	AgentIndex uint       `json:"agent_index"`
	Distance   uint       `json:"distance,omitempty"` // meters
	Time       uint       `json:"time,omitempty"`     // seconds
	TotalTime  uint       `json:"total_time,omitempty"`
	StartTime  uint       `json:"start_time,omitempty"`
	EndTime    uint       `json:"end_time,omitempty"`
	Mode       string     `json:"mode,omitempty"`
	Actions    []Action   `json:"actions,omitempty"`
	Waypoints  []Waypoint `json:"waypoints,omitempty"`
	Legs       []Leg      `json:"legs,omitempty"`
}

type Params struct {
	Mode   string  `json:"mode,omitempty"`
	Agents []Agent `json:"agents,omitempty"`
	Jobs   []Job   `json:"jobs,omitempty"`
}

type Issues struct {
	Shipments []uint `json:"unassigned_shipments,omitempty"`
	Agents    []uint `json:"unassigned_agents,omitempty"`
}

type Properties struct {
	Mode   string `json:"mode,omitempty"`
	Params `json:"params,omitempty"`
	Issues `json:"issues,omitempty"`
}

type Line []LonLat

type Geometry struct {
	Type        string `json:"type"` // MultiLineString
	Coordinates []Line
}

type Feature struct {
	Geometry   `json:"geometry"`
	Type       string    `json:"type"` // Feature
	Properties AgentPlan `json:"properties"`
}

// Plan is the route(s) output as planned by Geoapify based on the given input
type Plan struct {
	Type       string    `json:"type"` // FeatureCollection
	Features   []Feature `json:"features,omitempty"`
	Properties `json:"properties,omitempty"`
}

/* ************************ */

type BatchInput struct {
	ID     string  `json:"id,omitempty"`
	Body   Request `json:"body,omitempty"`
	Params Params  `json:"params"`
}

// BatchRequest is the input for long-running / batch job request
// Max number of inputs / len(BatchInput) is 1000
type BatchRequest struct {
	API    string       `json:"api"`
	Inputs []BatchInput `json:"inputs"`
	Params Params       `json:"params,omitempty"`
}

type Result struct {
	Params Params  `json:"params,omitempty"`
	Body   Request `json:"body,omitemtpy"`
	Plan   `json:"result,omitempty"`
}

// BatchResponse is the output for long-running / batch job request
type BatchResponse struct {
	ID      string   `json:"id"`
	API     string   `json:"api"`
	Params  Params   `json:"params,omitempty"`
	Results []Result `json:"results,omitempty"`
}
