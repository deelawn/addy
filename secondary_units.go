package addy

import "strings"

// secondary units are the standard USPS abbreviations:
// https://pe.usps.com/text/pub28/28apc_003.htm

const (
	secondaryUnitApartment  string = "APT"
	secondaryUnitBasement   string = "BSMT"
	secondaryUnitBuilding   string = "BLDG"
	secondaryUnitDepartment string = "DEPT"
	secondaryUnitFloor      string = "FL"
	secondaryUnitFront      string = "FRNT"
	secondaryUnitHanger     string = "HNGR"
	secondaryUnitKey        string = "KEY"
	secondaryUnitLobby      string = "LBBY"
	secondaryUnitLot        string = "LOT"
	secondaryUnitLower      string = "LOWR"
	secondaryUnitOffice     string = "OFC"
	secondaryUnitPenthouse  string = "PH"
	secondaryUnitPier       string = "PIER"
	secondaryUnitRear       string = "REAR"
	secondaryUnitRoom       string = "RM"
	secondaryUnitSide       string = "SIDE"
	secondaryUnitSlip       string = "SLIP"
	secondaryUnitSpace      string = "SPC"
	secondaryUnitStop       string = "STOP"
	secondaryUnitSuite      string = "STE"
	secondaryUnitTrailer    string = "TRLR"
	secondaryUnitUnit       string = "UNIT"
	secondaryUnitUpper      string = "UPPR"

	// A couple extra not officially specified as secondary units.
	secondaryUnitPO  string = "PO"
	secondaryUnitBox string = "BOX"
)

var secondaryUnitNormalizations = map[string]string{
	secondaryUnitApartment:  secondaryUnitApartment,
	"APARTMENT":             secondaryUnitApartment,
	secondaryUnitBasement:   secondaryUnitBasement,
	"BASEMENT":              secondaryUnitBasement,
	secondaryUnitBuilding:   secondaryUnitBuilding,
	"BUILDING":              secondaryUnitBuilding,
	secondaryUnitDepartment: secondaryUnitDepartment,
	"DEPARTMENT":            secondaryUnitDepartment,
	secondaryUnitFloor:      secondaryUnitFloor,
	"FLOOR":                 secondaryUnitFloor,
	secondaryUnitFront:      secondaryUnitFront,
	"FRONT":                 secondaryUnitFront,
	secondaryUnitHanger:     secondaryUnitHanger,
	"HANGER":                secondaryUnitHanger,
	secondaryUnitKey:        secondaryUnitKey,
	secondaryUnitLobby:      secondaryUnitLobby,
	"LOBBY":                 secondaryUnitLobby,
	secondaryUnitLot:        secondaryUnitLot,
	secondaryUnitLower:      secondaryUnitLower,
	"LOWER":                 secondaryUnitLower,
	secondaryUnitOffice:     secondaryUnitOffice,
	"OFFICE":                secondaryUnitOffice,
	secondaryUnitPenthouse:  secondaryUnitPenthouse,
	"PENTHOUSE":             secondaryUnitPenthouse,
	secondaryUnitPier:       secondaryUnitPier,
	secondaryUnitRear:       secondaryUnitRear,
	secondaryUnitRoom:       secondaryUnitRoom,
	"ROOM":                  secondaryUnitRoom,
	secondaryUnitSide:       secondaryUnitSide,
	secondaryUnitSlip:       secondaryUnitSlip,
	secondaryUnitSpace:      secondaryUnitSpace,
	"SPACE":                 secondaryUnitSpace,
	secondaryUnitStop:       secondaryUnitStop,
	secondaryUnitSuite:      secondaryUnitSuite,
	"SUITE":                 secondaryUnitSuite,
	secondaryUnitTrailer:    secondaryUnitTrailer,
	"TRAILER":               secondaryUnitTrailer,
	secondaryUnitUnit:       secondaryUnitUnit,
	secondaryUnitUpper:      secondaryUnitUpper,
	"UPPER":                 secondaryUnitUpper,

	// A couple extra not officially specified as secondary units.
	secondaryUnitPO:  secondaryUnitPO,
	"P.O.":           secondaryUnitPO,
	secondaryUnitBox: secondaryUnitBox,
}

// normalizeSecondaryUnit will normalize any string it determines to be a secondary unit requiring normalization.
// Otherwise it will just return what was passed in. The boolean it returns is true if the string
// was found in the lookup table.
func normalizeSecondaryUnit(s string, normOpts normalizationOptions) (string, bool) {

	normalized, foundInLookup := normalize(s, secondaryUnitNormalizations, normOpts)

	// Special case for the PO secondary unit. It should normally always be returned as PO unless
	// the override option is specified, meaning that it should adhere to the casing option.
	// So, if the option is not specified, override whatever casing was applied to it so it is
	// reverted to uppercase.
	if !normOpts.exludePOCasing && strings.ToUpper(normalized) == secondaryUnitPO {
		normalized = secondaryUnitPO
	}

	return normalized, foundInLookup
}
