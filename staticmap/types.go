package staticmap

// Coordinate cannot use float32 due to rounding inconsistency at the 6th digit

type LonLat [2]float64

type Area struct {
	GPS1 LonLat
	GPS2 LonLat
}

type Marker struct {
	LonLat
	Color       string // Named or in hex form ("#" should be encoded as "%23")
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
