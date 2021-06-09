package addy

// directionals are the standard USPS abbreviations:
// https://pe.usps.com/text/pub28/28api_007.htm

const (
	directionalNorth     string = "N"
	directionalNortheast string = "NE"
	directionalNorthwest string = "NW"
	directionalSouth     string = "S"
	directionalSoutheast string = "SE"
	directionalSouthwest string = "SW"
	directionalEast      string = "E"
	directionalWest      string = "W"
)

var directionalNormalizations = map[string]string{
	directionalNorth:     directionalNorth,
	"NORTH":              directionalNorth,
	"NORTE":              directionalNorth,
	directionalNortheast: directionalNortheast,
	"NORTHEAST":          directionalNortheast,
	"NORESTE":            directionalNortheast,
	directionalNorthwest: directionalNorthwest,
	"NORTHWEST":          directionalNorthwest,
	"NOROESTE":           directionalNorthwest,
	directionalSouth:     directionalSouth,
	"SOUTH":              directionalSouth,
	"SUR":                directionalSouth,
	directionalSoutheast: directionalSoutheast,
	"SOUTHEAST":          directionalSoutheast,
	"SURESTE":            directionalSoutheast,
	directionalSouthwest: directionalSouthwest,
	"SOUTHWEST":          directionalSouthwest,
	"SUROESTE":           directionalSouthwest,
	directionalEast:      directionalEast,
	"EAST":               directionalEast,
	"ESTE":               directionalEast,
	directionalWest:      directionalWest,
	"WEST":               directionalWest,
	"OESTE":              directionalWest,
}

// normalizeDirectional will normalize any string it determines to be a direction requiring normalization.
// Otherwise it will just return what was passed in. The boolean it returns is true if the string
// was found in the lookup table.
func normalizeDirectional(s string, normOpts normalizationOptions) (string, bool) {
	return normalize(s, directionalNormalizations, normOpts)
}
