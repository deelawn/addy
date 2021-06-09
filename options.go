package addy

// options affect the output of the NormalizeX functions.
type option int

const (
	// OptionUpperCase will return an uppercased version of each token.
	OptionUpperCase option = iota
	// OptionPreserveCase will preserve the case that matches the original token whenever a normalization
	// occurs. If no normalization occurs, the same token will be returned.
	OptionPreserveCase
	// OptionLowerCase will return a lowercased version of each token.
	OptionLowerCase
	// OptionTitleCase will return a titlecased version of each token.
	OptionTitleCase
	// OptionExcludePOCasing will, if true, not apply the casing option to a PO token, instead always returning
	// it as an uppercased token, as is the common way of it being displayed.
	OptionExcludePOCasing
)

// normalizationOptions makes is easier to pass around a summarized version of the options that get
// passed in to the normalization functions.
type normalizationOptions struct {
	casingOption   option
	exludePOCasing bool
}

// isCasingOption returns true if the option will affect the case of the returned normalization.
func isCasingOption(o option) bool {
	return o < OptionExcludePOCasing
}

// parseOptions will take list of options and return an instance of normalizationOptions.
// If no options are provided, the normalizationOptions default values will be applied.
func parseOptions(options ...option) (nrmOpts normalizationOptions) {

	var casingOptionApplied bool
	for _, o := range options {

		// Only apply the first casing option encountered.
		if isCasingOption(o) && !casingOptionApplied {
			nrmOpts.casingOption = o
			casingOptionApplied = true
			continue
		}

		if o == OptionExcludePOCasing {
			nrmOpts.exludePOCasing = true
		}
	}

	return
}
