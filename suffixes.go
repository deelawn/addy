package addy

// suffixes are the standard USPS suffix abbreviations:
// https://pe.usps.com/text/pub28/28apc_002.htm

const (
	suffixAlley      string = "ALY"
	suffixAnex       string = "ANX"
	suffixArcade     string = "ARC"
	suffixAvenue     string = "AVE"
	suffixBayou      string = "BYU"
	suffixBeach      string = "BCH"
	suffixBend       string = "BND"
	suffixBluff      string = "BLF"
	suffixBluffs     string = "BLFS"
	suffixBottom     string = "BTM"
	suffixBoulevard  string = "BLVD"
	suffixBranch     string = "BR"
	suffixBridge     string = "BRG"
	suffixBrook      string = "BRK"
	suffixBrooks     string = "BRKS"
	suffixBurg       string = "BG"
	suffixBurgs      string = "BGS"
	suffixBypass     string = "BYP"
	suffixCamp       string = "CP"
	suffixCanyon     string = "CYN"
	suffixCape       string = "CPE"
	suffixCauseway   string = "CSWY"
	suffixCenter     string = "CTR"
	suffixCenters    string = "CTRS"
	suffixCircle     string = "CIR"
	suffixCircles    string = "CIRS"
	suffixCliff      string = "CLF"
	suffixCliffs     string = "CLFS"
	suffixClub       string = "CLB"
	suffixCommon     string = "CMN"
	suffixCommons    string = "CMNS"
	suffixCorner     string = "COR"
	suffixCorners    string = "CORS"
	suffixCourse     string = "CRSE"
	suffixCourt      string = "CT"
	suffixCourts     string = "CTS"
	suffixCove       string = "CV"
	suffixCoves      string = "CVS"
	suffixCreek      string = "CRK"
	suffixCrescent   string = "CRES"
	suffixCrest      string = "CRST"
	suffixCrossing   string = "XING"
	suffixCrossroad  string = "XRD"
	suffixCrossroads string = "XRDS"
	suffixCurve      string = "CURV"
	suffixDale       string = "DL"
	suffixDam        string = "DM"
	suffixDivide     string = "DV"
	suffixDrive      string = "DR"
	suffixDrives     string = "DRS"
	suffixEstate     string = "EST"
	suffixEstates    string = "ESTS"
	suffixExpressway string = "EXPY"
	suffixExtension  string = "EXT"
	suffixExtensions string = "EXTS"
	suffixFall       string = "FALL"
	suffixFalls      string = "FLS"
	suffixFerry      string = "FRY"
	suffixField      string = "FLD"
	suffixFields     string = "FLDS"
	suffixFlat       string = "FLT"
	suffixFlats      string = "FLTS"
	suffixFord       string = "FRD"
	suffixFords      string = "FRDS"
	suffixForest     string = "FRST"
	suffixForge      string = "FRG"
	suffixForges     string = "FRGS"
	suffixFork       string = "FRK"
	suffixForks      string = "FRKS"
	suffixFort       string = "FT"
	suffixFreeway    string = "FWY"
	suffixGarden     string = "GDN"
	suffixGardens    string = "GDNS"
	suffixGateway    string = "GTWY"
	suffixGlen       string = "GLN"
	suffixGlens      string = "GLNS"
	suffixGreen      string = "GRN"
	suffixGreens     string = "GRNS"
	suffixGrove      string = "GRV"
	suffixGroves     string = "GRVS"
	suffixHarbor     string = "HBR"
	suffixHarbors    string = "HBRS"
	suffixHaven      string = "HVN"
	suffixHeights    string = "HTS"
	suffixHighway    string = "HWY"
	suffixHill       string = "HL"
	suffixHills      string = "HLS"
	suffixHollow     string = "HOLW"
	suffixInlet      string = "INLT"
	suffixIsland     string = "IS"
	suffixIslands    string = "ISS"
	suffixIsle       string = "ISLE"
	suffixJunction   string = "JCT"
	suffixJunctions  string = "JCTS"
	suffixKey        string = "KY"
	suffixKeys       string = "KYS"
	suffixKnoll      string = "KNL"
	suffixKnolls     string = "KNLS"
	suffixLake       string = "LK"
	suffixLakes      string = "LKS"
	suffixLand       string = "LAND"
	suffixLanding    string = "LNDG"
	suffixLane       string = "LN"
	suffixLight      string = "LGT"
	suffixLights     string = "LGTS"
	suffixLoaf       string = "LF"
	suffixLock       string = "LCK"
	suffixLocks      string = "LCKS"
	suffixLodge      string = "LDG"
	suffixLoop       string = "LOOP"
	suffixMall       string = "MALL"
	suffixManor      string = "MNR"
	suffixManors     string = "MNRS"
	suffixMeadow     string = "MDW"
	suffixMeadows    string = "MDWS"
	suffixMews       string = "MEWS"
	suffixMill       string = "ML"
	suffixMills      string = "MLS"
	suffixMission    string = "MSN"
	suffixMotorway   string = "MTWY"
	suffixMount      string = "MT"
	suffixMountain   string = "MTN"
	suffixMountains  string = "MTNS"
	suffixNeck       string = "NCK"
	suffixOrchard    string = "ORCH"
	suffixOval       string = "OVAL"
	suffixOverpass   string = "OPAS"
	suffixPark       string = "PARK"
	suffixParkway    string = "PKWY"
	suffixPass       string = "PASS"
	suffixPassage    string = "PSGE"
	suffixPath       string = "PATH"
	suffixPike       string = "PIKE"
	suffixPine       string = "PNE"
	suffixPines      string = "PNES"
	suffixPlace      string = "PL"
	suffixPlain      string = "PLN"
	suffixPlains     string = "PLNS"
	suffixPlaza      string = "PLZ"
	suffixPoint      string = "PT"
	suffixPoints     string = "PTS"
	suffixPort       string = "PRT"
	suffixPorts      string = "PRTS"
	suffixPrairie    string = "PR"
	suffixRadial     string = "RADL"
	suffixRamp       string = "RAMP"
	suffixRanch      string = "RNCH"
	suffixRapid      string = "RPD"
	suffixRapids     string = "RPDS"
	suffixRest       string = "RST"
	suffixRidge      string = "RDG"
	suffixRidges     string = "RDGS"
	suffixRiver      string = "RIV"
	suffixRoad       string = "RD"
	suffixRoads      string = "RDS"
	suffixRoute      string = "RTE"
	suffixRow        string = "ROW"
	suffixRue        string = "RUE"
	suffixRun        string = "RUN"
	suffixShoal      string = "SHL"
	suffixShoals     string = "SHLS"
	suffixShore      string = "SHR"
	suffixShores     string = "SHRS"
	suffixSkyway     string = "SKWY"
	suffixSpring     string = "SPG"
	suffixSprings    string = "SPGS"
	suffixSpur       string = "SPUR"
	suffixSquare     string = "SQ"
	suffixSquares    string = "SQS"
	suffixStation    string = "STA"
	suffixStravenue  string = "STRA"
	suffixStream     string = "STRM"
	suffixStreet     string = "ST"
	suffixStreets    string = "STS"
	suffixSummit     string = "SMT"
	suffixTerrace    string = "TER"
	suffixThroughway string = "TRWY"
	suffixTrace      string = "TRCE"
	suffixTrack      string = "TRAK"
	suffixTrafficway string = "TRFY"
	suffixTrail      string = "TRL"
	suffixTrailer    string = "TRLR"
	suffixTunnel     string = "TUNL"
	suffixTurnpike   string = "TPKE"
	suffixUnderpass  string = "UPAS"
	suffixUnion      string = "UN"
	suffixUnions     string = "UNS"
	suffixValley     string = "VLY"
	suffixValleys    string = "VLYS"
	suffixViaduct    string = "VIA"
	suffixView       string = "VW"
	suffixViews      string = "VWS"
	suffixVillage    string = "VLG"
	suffixVillages   string = "VLGS"
	suffixVille      string = "VL"
	suffixVista      string = "VIS"
	suffixWalk       string = "WALK"
	suffixWall       string = "WALL"
	suffixWay        string = "WAY"
	suffixWays       string = "WAYS"
	suffixWell       string = "WL"
	suffixWells      string = "WLS"
)

var suffixNormalizations = map[string]string{
	suffixAlley:      suffixAlley,
	"ALLEY":          suffixAlley,
	"ALLEE":          suffixAlley,
	"ALLY":           suffixAlley,
	suffixAnex:       suffixAnex,
	"ANEX":           suffixAnex,
	"ANNEX":          suffixAnex,
	"ANNX":           suffixAnex,
	suffixArcade:     suffixArcade,
	"ARCADE":         suffixArcade,
	suffixAvenue:     suffixAvenue,
	"AV":             suffixAvenue,
	"AVEN":           suffixAvenue,
	"AVENU":          suffixAvenue,
	"AVENUE":         suffixAvenue,
	"AVN":            suffixAvenue,
	"AVNUE":          suffixAvenue,
	suffixBayou:      suffixBayou,
	"BAYOO":          suffixBayou,
	"BAYOU":          suffixBayou,
	suffixBeach:      suffixBeach,
	"BEACH":          suffixBeach,
	suffixBend:       suffixBend,
	"BEND":           suffixBend,
	suffixBluff:      suffixBluff,
	"BLUF":           suffixBluff,
	"BLUFF":          suffixBluff,
	suffixBluffs:     suffixBluffs,
	"BLUFFS":         suffixBluffs,
	suffixBottom:     suffixBottom,
	"BOT":            suffixBottom,
	"BOTTM":          suffixBottom,
	"BOTTOM":         suffixBottom,
	suffixBoulevard:  suffixBoulevard,
	"BOUL":           suffixBoulevard,
	"BOULEVARD":      suffixBoulevard,
	"BOULV":          suffixBoulevard,
	suffixBranch:     suffixBranch,
	"BRNCH":          suffixBranch,
	"BRANCH":         suffixBranch,
	suffixBridge:     suffixBridge,
	"BRDGE":          suffixBridge,
	"BRIDGE":         suffixBridge,
	suffixBrook:      suffixBrook,
	"BROOK":          suffixBrook,
	suffixBrooks:     suffixBrooks,
	"BROOKS":         suffixBrooks,
	suffixBurg:       suffixBurg,
	"BURG":           suffixBurg,
	suffixBurgs:      suffixBurgs,
	"BURGS":          suffixBurgs,
	suffixBypass:     suffixBypass,
	"BYPA":           suffixBypass,
	"BYPAS":          suffixBypass,
	"BYPASS":         suffixBypass,
	"BYPS":           suffixBypass,
	suffixCamp:       suffixCamp,
	"CAMP":           suffixCamp,
	"CMP":            suffixCamp,
	suffixCanyon:     suffixCanyon,
	"CANYN":          suffixCanyon,
	"CANYON":         suffixCanyon,
	"CNYN":           suffixCanyon,
	suffixCape:       suffixCape,
	"CAPE":           suffixCape,
	suffixCauseway:   suffixCauseway,
	"CAUSEWAY":       suffixCauseway,
	"CAUSWA":         suffixCauseway,
	suffixCenter:     suffixCenter,
	"CEN":            suffixCenter,
	"CENT":           suffixCenter,
	"CENTER":         suffixCenter,
	"CENTR":          suffixCenter,
	"CNTER":          suffixCenter,
	"CNTR":           suffixCenter,
	suffixCenters:    suffixCenters,
	"CENTERS":        suffixCenters,
	suffixCircle:     suffixCircle,
	"CIRC":           suffixCircle,
	"CIRCL":          suffixCircle,
	"CIRCLE":         suffixCircle,
	"CRCL":           suffixCircle,
	"CRCLE":          suffixCircle,
	suffixCircles:    suffixCircles,
	"CIRCLES":        suffixCircles,
	suffixCliff:      suffixCliff,
	"CLIFF":          suffixCliff,
	suffixCliffs:     suffixCliffs,
	"CLIFFS":         suffixCliffs,
	suffixClub:       suffixClub,
	"CLUB":           suffixClub,
	suffixCommon:     suffixCommon,
	"COMMON":         suffixCommon,
	suffixCommons:    suffixCommons,
	"COMMONS":        suffixCommons,
	suffixCorner:     suffixCorner,
	"CORNER":         suffixCorner,
	suffixCorners:    suffixCorners,
	"CORNERS":        suffixCorners,
	suffixCourse:     suffixCourse,
	"COURSE":         suffixCourse,
	suffixCourt:      suffixCourt,
	"COURT":          suffixCourt,
	suffixCourts:     suffixCourts,
	"COURTS":         suffixCourts,
	suffixCove:       suffixCove,
	"COVE":           suffixCove,
	suffixCoves:      suffixCoves,
	"COVES":          suffixCoves,
	suffixCreek:      suffixCreek,
	"CREEK":          suffixCreek,
	suffixCrescent:   suffixCrescent,
	"CRESCENT":       suffixCrescent,
	"CRSENT":         suffixCrescent,
	"CRSNT":          suffixCrescent,
	suffixCrest:      suffixCrest,
	"CREST":          suffixCrest,
	suffixCrossing:   suffixCrossing,
	"CROSSING":       suffixCrossing,
	"CRSSNG":         suffixCrossing,
	suffixCrossroad:  suffixCrossroad,
	"CROSSROAD":      suffixCrossroad,
	suffixCrossroads: suffixCrossroads,
	"CROSSROADS":     suffixCrossroads,
	suffixCurve:      suffixCurve,
	"CURVE":          suffixCurve,
	suffixDale:       suffixDale,
	"DALE":           suffixDale,
	suffixDam:        suffixDam,
	"DAM":            suffixDam,
	suffixDivide:     suffixDivide,
	"DIV":            suffixDivide,
	"DIVIDE":         suffixDivide,
	"DVD":            suffixDivide,
	suffixDrive:      suffixDrive,
	"DRIV":           suffixDrive,
	"DRIVE":          suffixDrive,
	"DRV":            suffixDrive,
	suffixDrives:     suffixDrives,
	"DRIVES":         suffixDrives,
	suffixEstate:     suffixEstate,
	"ESTATE":         suffixEstate,
	suffixEstates:    suffixEstates,
	"ESTATES":        suffixEstates,
	suffixExpressway: suffixExpressway,
	"EXP":            suffixExpressway,
	"EXPR":           suffixExpressway,
	"EXPRESS":        suffixExpressway,
	"EXPRESSWAY":     suffixExpressway,
	"EXPW":           suffixExpressway,
	suffixExtension:  suffixExtension,
	"EXTENSION":      suffixExtension,
	"EXTN":           suffixExtension,
	"EXTNSN":         suffixExtension,
	suffixExtensions: suffixExtensions,
	"EXTENSIONS":     suffixExtensions,
	suffixFall:       suffixFall,
	suffixFalls:      suffixFalls,
	"FALLS":          suffixFalls,
	suffixFerry:      suffixFerry,
	"FERRY":          suffixFerry,
	"FRRY":           suffixFerry,
	suffixField:      suffixField,
	"FIELD":          suffixField,
	suffixFields:     suffixFields,
	"FIELDS":         suffixFields,
	suffixFlat:       suffixFlat,
	"FLAT":           suffixFlat,
	suffixFlats:      suffixFlats,
	"FLATS":          suffixFlats,
	suffixFord:       suffixFord,
	"FORD":           suffixFord,
	suffixFords:      suffixFords,
	"FORDS":          suffixFords,
	suffixForest:     suffixForest,
	"FOREST":         suffixForest,
	"FORESTS":        suffixForest,
	suffixForge:      suffixForge,
	"FORG":           suffixForge,
	"FORGE":          suffixForge,
	suffixForges:     suffixForges,
	"FORGES":         suffixForges,
	suffixFork:       suffixFork,
	"FORK":           suffixFork,
	suffixForks:      suffixForks,
	"FORKS":          suffixForks,
	suffixFort:       suffixFort,
	"FORT":           suffixFort,
	"FRT":            suffixFort,
	suffixFreeway:    suffixFreeway,
	"FREEWAY":        suffixFreeway,
	"FREEWY":         suffixFreeway,
	"FRWAY":          suffixFreeway,
	suffixGarden:     suffixGarden,
	"GARDEN":         suffixGarden,
	"GARDN":          suffixGarden,
	"GRDEN":          suffixGarden,
	"GRDN":           suffixGarden,
	suffixGardens:    suffixGardens,
	"GARDENS":        suffixGardens,
	"GRDNS":          suffixGardens,
	suffixGateway:    suffixGateway,
	"GATEWAY":        suffixGateway,
	"GATEWY":         suffixGateway,
	"GATWAY":         suffixGateway,
	"GTWAY":          suffixGateway,
	suffixGlen:       suffixGlen,
	"GLEN":           suffixGlen,
	suffixGlens:      suffixGlens,
	"GLENS":          suffixGlens,
	suffixGreen:      suffixGreen,
	"GREEN":          suffixGreen,
	suffixGreens:     suffixGreens,
	"GREENS":         suffixGreens,
	suffixGrove:      suffixGrove,
	"GROV":           suffixGrove,
	"GROVE":          suffixGrove,
	suffixGroves:     suffixGroves,
	"GROVES":         suffixGroves,
	suffixHarbor:     suffixHarbor,
	"HARB":           suffixHarbor,
	"HARBOR":         suffixHarbor,
	"HARBR":          suffixHarbor,
	"HRBOR":          suffixHarbor,
	suffixHarbors:    suffixHarbors,
	"HARBORS":        suffixHarbors,
	suffixHaven:      suffixHaven,
	"HAVEN":          suffixHaven,
	suffixHeights:    suffixHeights,
	"HT":             suffixHeights,
	"HEIGHTS":        suffixHeights,
	suffixHighway:    suffixHighway,
	"HIGHWAY":        suffixHighway,
	"HIGHWY":         suffixHighway,
	"HIWAY":          suffixHighway,
	"HIWY":           suffixHighway,
	"HWAY":           suffixHighway,
	suffixHill:       suffixHill,
	"HILL":           suffixHill,
	suffixHills:      suffixHills,
	"HILLS":          suffixHills,
	suffixHollow:     suffixHollow,
	"HLLW":           suffixHollow,
	"HOLLOW":         suffixHollow,
	"HOLLOWS":        suffixHollow,
	"HOLWS":          suffixHollow,
	suffixInlet:      suffixInlet,
	"INLET":          suffixInlet,
	suffixIsland:     suffixIsland,
	"ISLAND":         suffixIsland,
	"ISLND":          suffixIsland,
	suffixIslands:    suffixIslands,
	"ISLANDS":        suffixIslands,
	"ISLNDS":         suffixIslands,
	suffixIsle:       suffixIsle,
	"ISLES":          suffixIsle,
	suffixJunction:   suffixJunction,
	"JCTION":         suffixJunction,
	"JCTN":           suffixJunction,
	"JUNCTION":       suffixJunction,
	"JUNCTN":         suffixJunction,
	"JUNCTON":        suffixJunction,
	suffixJunctions:  suffixJunctions,
	"JCTNS":          suffixJunctions,
	"JUNCTIONS":      suffixJunctions,
	suffixKey:        suffixKey,
	"KEY":            suffixKey,
	suffixKeys:       suffixKeys,
	"KEYS":           suffixKeys,
	suffixKnoll:      suffixKnoll,
	"KNOL":           suffixKnoll,
	"KNOLL":          suffixKnoll,
	suffixKnolls:     suffixKnolls,
	"KNOLLS":         suffixKnolls,
	suffixLake:       suffixLake,
	"LAKE":           suffixLake,
	suffixLakes:      suffixLakes,
	"LAKES":          suffixLakes,
	suffixLand:       suffixLand,
	suffixLanding:    suffixLanding,
	"LANDING":        suffixLanding,
	"LNDNG":          suffixLanding,
	suffixLane:       suffixLane,
	"LANE":           suffixLane,
	suffixLight:      suffixLight,
	"LIGHT":          suffixLight,
	suffixLights:     suffixLights,
	"LIGHTS":         suffixLights,
	suffixLoaf:       suffixLoaf,
	"LOAF":           suffixLoaf,
	suffixLock:       suffixLock,
	"LOCK":           suffixLock,
	suffixLocks:      suffixLocks,
	"LOCKS":          suffixLocks,
	suffixLodge:      suffixLodge,
	"LDGE":           suffixLodge,
	"LODG":           suffixLodge,
	"LODGE":          suffixLodge,
	suffixLoop:       suffixLoop,
	"LOOPS":          suffixLoop,
	suffixMall:       suffixMall,
	suffixManor:      suffixManor,
	"MANOR":          suffixManor,
	suffixManors:     suffixManors,
	"MANORS":         suffixManors,
	suffixMeadow:     suffixMeadow,
	"MEADOW":         suffixMeadow,
	suffixMeadows:    suffixMeadows,
	"MEADOWS":        suffixMeadows,
	"MEDOWS":         suffixMeadows,
	suffixMews:       suffixMews,
	suffixMill:       suffixMill,
	"MILL":           suffixMill,
	suffixMills:      suffixMills,
	"MILLS":          suffixMills,
	suffixMission:    suffixMission,
	"MISSN":          suffixMission,
	"MSSN":           suffixMission,
	"MISSION":        suffixMission,
	suffixMotorway:   suffixMotorway,
	"MOTORWAY":       suffixMotorway,
	suffixMount:      suffixMount,
	"MNT":            suffixMount,
	"MOUNT":          suffixMount,
	suffixMountain:   suffixMountain,
	"MNTAIN":         suffixMountain,
	"MNTN":           suffixMountain,
	"MOUNTAIN":       suffixMountain,
	"MOUNTIN":        suffixMountain,
	"MTIN":           suffixMountain,
	suffixMountains:  suffixMountains,
	"MNTS":           suffixMountains,
	"MOUNTAINS":      suffixMountains,
	suffixNeck:       suffixNeck,
	"NECK":           suffixNeck,
	suffixOrchard:    suffixOrchard,
	"ORCHARD":        suffixOrchard,
	"ORCHRD":         suffixOrchard,
	suffixOval:       suffixOval,
	"OVL":            suffixOval,
	suffixOverpass:   suffixOverpass,
	"OVERPASS":       suffixOverpass,
	suffixPark:       suffixPark,
	"PRK":            suffixPark,
	"PARKS":          suffixPark,
	suffixParkway:    suffixParkway,
	"PARKWAY":        suffixParkway,
	"PARKWY":         suffixParkway,
	"PKWAY":          suffixParkway,
	"PKY":            suffixParkway,
	"PARKWAYS":       suffixParkway,
	"PKWYS":          suffixParkway,
	suffixPass:       suffixPass,
	suffixPassage:    suffixPassage,
	"PASSAGE":        suffixPassage,
	suffixPath:       suffixPath,
	"PATHS":          suffixPath,
	suffixPike:       suffixPike,
	"PIKES":          suffixPike,
	suffixPine:       suffixPine,
	"PINE":           suffixPine,
	suffixPines:      suffixPines,
	"PINES":          suffixPines,
	suffixPlace:      suffixPlace,
	"PLACE":          suffixPlace,
	suffixPlain:      suffixPlain,
	"PLAIN":          suffixPlain,
	suffixPlains:     suffixPlains,
	"PLAINS":         suffixPlains,
	suffixPlaza:      suffixPlaza,
	"PLAZA":          suffixPlaza,
	"PLZA":           suffixPlaza,
	suffixPoint:      suffixPoint,
	"POINT":          suffixPoint,
	suffixPoints:     suffixPoints,
	"POINTS":         suffixPoints,
	suffixPort:       suffixPort,
	"PORT":           suffixPort,
	suffixPorts:      suffixPorts,
	"PORTS":          suffixPorts,
	suffixPrairie:    suffixPrairie,
	"PRAIRIE":        suffixPrairie,
	"PRR":            suffixPrairie,
	suffixRadial:     suffixRadial,
	"RAD":            suffixRadial,
	"RADIAL":         suffixRadial,
	"RADIEL":         suffixRadial,
	suffixRamp:       suffixRamp,
	suffixRanch:      suffixRanch,
	"RANCH":          suffixRanch,
	"RANCHES":        suffixRanch,
	"RNCHS":          suffixRanch,
	suffixRapid:      suffixRapid,
	"RAPID":          suffixRapid,
	suffixRapids:     suffixRapids,
	"RAPIDS":         suffixRapids,
	suffixRest:       suffixRest,
	"REST":           suffixRest,
	suffixRidge:      suffixRidge,
	"RDGE":           suffixRidge,
	"RIDGE":          suffixRidge,
	suffixRidges:     suffixRidges,
	"RIDGES":         suffixRidges,
	suffixRiver:      suffixRiver,
	"RIVER":          suffixRiver,
	"RVR":            suffixRiver,
	"RIVR":           suffixRiver,
	suffixRoad:       suffixRoad,
	"ROAD":           suffixRoad,
	suffixRoads:      suffixRoads,
	"ROADS":          suffixRoads,
	suffixRoute:      suffixRoute,
	"ROUTE":          suffixRoute,
	suffixRow:        suffixRow,
	suffixRue:        suffixRue,
	suffixRun:        suffixRun,
	suffixShoal:      suffixShoal,
	"SHOAL":          suffixShoal,
	suffixShoals:     suffixShoals,
	"SHOALS":         suffixShoals,
	suffixShore:      suffixShore,
	"SHOAR":          suffixShore,
	"SHORE":          suffixShore,
	suffixShores:     suffixShores,
	"SHOARS":         suffixShores,
	"SHORES":         suffixShores,
	suffixSkyway:     suffixSkyway,
	"SKYWAY":         suffixSkyway,
	suffixSpring:     suffixSpring,
	"SPNG":           suffixSpring,
	"SPRING":         suffixSpring,
	"SPRNG":          suffixSpring,
	suffixSprings:    suffixSprings,
	"SPNGS":          suffixSprings,
	"SPRINGS":        suffixSprings,
	"SPRNGS":         suffixSprings,
	suffixSpur:       suffixSpur,
	"SPURS":          suffixSpur,
	suffixSquare:     suffixSquare,
	"SQR":            suffixSquare,
	"SQRE":           suffixSquare,
	"SQU":            suffixSquare,
	"SQUARE":         suffixSquare,
	suffixSquares:    suffixSquares,
	"SQRS":           suffixSquares,
	"SQUARES":        suffixSquares,
	suffixStation:    suffixStation,
	"STATION":        suffixStation,
	"STATN":          suffixStation,
	"STN":            suffixStation,
	suffixStravenue:  suffixStravenue,
	"STRAV":          suffixStravenue,
	"STRAVEN":        suffixStravenue,
	"STRAVENUE":      suffixStravenue,
	"STRAVN":         suffixStravenue,
	"STRVN":          suffixStravenue,
	"STRVNUE":        suffixStravenue,
	suffixStream:     suffixStream,
	"STREAM":         suffixStream,
	"STREME":         suffixStream,
	suffixStreet:     suffixStreet,
	"STREET":         suffixStreet,
	"STRT":           suffixStreet,
	"STR":            suffixStreet,
	suffixStreets:    suffixStreets,
	"STREETS":        suffixStreets,
	suffixSummit:     suffixSummit,
	"SUMIT":          suffixSummit,
	"SUMITT":         suffixSummit,
	"SUMMIT":         suffixSummit,
	suffixTerrace:    suffixTerrace,
	"TERR":           suffixTerrace,
	"TERRACE":        suffixTerrace,
	suffixThroughway: suffixThroughway,
	"THROUGHWAY":     suffixThroughway,
	suffixTrace:      suffixTrace,
	"TRACE":          suffixTrace,
	"TRACES":         suffixTrace,
	suffixTrack:      suffixTrack,
	"TRACK":          suffixTrack,
	"TRACKS":         suffixTrack,
	"TRK":            suffixTrack,
	"TRKS":           suffixTrack,
	suffixTrafficway: suffixTrafficway,
	"TRAFFICWAY":     suffixTrafficway,
	suffixTrail:      suffixTrail,
	"TRAIL":          suffixTrail,
	"TRAILS":         suffixTrail,
	"TRLS":           suffixTrail,
	suffixTrailer:    suffixTrailer,
	"TRAILER":        suffixTrailer,
	"TRLRS":          suffixTrailer,
	suffixTunnel:     suffixTunnel,
	"TUNEL":          suffixTunnel,
	"TUNLS":          suffixTunnel,
	"TUNNEL":         suffixTunnel,
	"TUNNELS":        suffixTunnel,
	"TUNNL":          suffixTunnel,
	suffixTurnpike:   suffixTurnpike,
	"TRNPK":          suffixTurnpike,
	"TURNPIKE":       suffixTurnpike,
	"TURNPK":         suffixTurnpike,
	suffixUnderpass:  suffixUnderpass,
	"UNDERPASS":      suffixUnderpass,
	suffixUnion:      suffixUnion,
	"UNION":          suffixUnion,
	suffixUnions:     suffixUnions,
	"UNIONS":         suffixUnions,
	suffixValley:     suffixValley,
	"VALLEY":         suffixValley,
	"VALLY":          suffixValley,
	"VLLY":           suffixValley,
	suffixValleys:    suffixValleys,
	"VALLEYS":        suffixValleys,
	suffixViaduct:    suffixViaduct,
	"VDCT":           suffixViaduct,
	"VIADCT":         suffixViaduct,
	"VIADUCT":        suffixViaduct,
	suffixView:       suffixView,
	"VIEW":           suffixView,
	suffixViews:      suffixViews,
	"VIEWS":          suffixViews,
	suffixVillage:    suffixVillage,
	"VILL":           suffixVillage,
	"VILLAG":         suffixVillage,
	"VILLAGE":        suffixVillage,
	"VILLG":          suffixVillage,
	"VILLIAGE":       suffixVillage,
	suffixVillages:   suffixVillages,
	"VILLAGES":       suffixVillages,
	suffixVille:      suffixVille,
	"VILLE":          suffixVille,
	suffixVista:      suffixVista,
	"VIST":           suffixVista,
	"VISTA":          suffixVista,
	"VST":            suffixVista,
	"VSTA":           suffixVista,
	suffixWalk:       suffixWalk,
	"WALKS":          suffixWalk,
	suffixWall:       suffixWall,
	suffixWay:        suffixWay,
	"WY":             suffixWay,
	suffixWays:       suffixWays,
	suffixWell:       suffixWell,
	"WELL":           suffixWell,
	suffixWells:      suffixWells,
	"WELLS":          suffixWells,
}

// normalizeSuffix will normalize any string it determines to be a suffix requiring normalization.
// Otherwise it will just return what was passed in. The boolean it returns is true if the string
// was found in the lookup table.
func normalizeSuffix(s string, normOpts normalizationOptions) (string, bool) {
	return normalize(s, suffixNormalizations, normOpts)
}
