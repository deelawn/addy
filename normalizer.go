package addy

import "strings"

func normalize(s string, lookup map[string]string, normOpts normalizationOptions) (string, bool) {

	lookupKey := strings.ToUpper(s)
	normalized, foundInLookup := lookup[lookupKey]

	// If we didn't find it in the lookup, assign the normalized value to what was passed in.
	if !foundInLookup {
		normalized = s
	}

	// This is the casing that we will apply to the normalized string.
	casingToApply := normOpts.casingOption

	// If we have the preserve casing option, we need to figure out the casing of the string that
	// was passed in so we know how to case the normalized result.
	if casingToApply == OptionPreserveCase {

		// Fastpath: not found in the lookup so just return the original string.
		if !foundInLookup {
			return s, false
		}

		// This will be the default if they provided a token with some weird mixture
		// of upper and lowercased characers. Titlecase is used by default because it is
		// the only casing option that includes both upper and lowercase characters, so there
		// is a chance that titlecase was the intended casing.
		casingToApply = OptionTitleCase

		// If applying these casings is equal to the original value, then we know that is
		// the casing that is to be preserved.
		if applyNormalizationCasingOption(s, OptionUpperCase) == s {
			casingToApply = OptionUpperCase
		} else if applyNormalizationCasingOption(s, OptionLowerCase) == s {
			casingToApply = OptionLowerCase
		}
	}

	normalized = applyNormalizationCasingOption(normalized, casingToApply)
	return normalized, foundInLookup
}

func applyNormalizationCasingOption(s string, o option) (casingApplied string) {

	// Now use the options to convert to the appropriate case.
	switch o {
	case OptionUpperCase:
		casingApplied = strings.ToUpper(s)

	case OptionLowerCase:
		casingApplied = strings.ToLower(s)

	case OptionTitleCase:
		casingApplied = strings.Title(strings.ToLower(s))
	}

	return
}

// NormalizeAddress1 will return a version of the input argument with all directional
// and suffix keywords normalized, along with a count of how many tokens were normalized.
// The casing is uppercased by default, but can be specified using optional arguments.
// Only the first casing option encountered will be applied.
func NormalizeAddress1(address string, options ...option) (string, int) {

	var (
		tokensNormalized int
		wasNormalized    bool
	)

	normOpts := parseOptions(options...)
	tokens := tokenize(address)
	for _, token := range tokens {

		// No need to normalize separators.
		if token.isSeparator {
			continue
		}

		// There should never be any overlap between things that can be normalized either directionally or
		// by suffix; if not one, then try the other.
		token.value, wasNormalized = normalizeDirectional(token.value, normOpts)
		if !wasNormalized {
			token.value, wasNormalized = normalizeSuffix(token.value, normOpts)
		}

		if wasNormalized {
			tokensNormalized++
		}
	}

	// All tokens have been normalized, or at least uppercased if no normalizations were applied.
	// Join them all back together and return them.
	return tokens.join(), tokensNormalized
}

// NormalizeAddress2 will return a version of the input argument with all secondary
// unit keywords normalized, along with a count of how many tokens were normalized.
// The casing is uppercased by default, but can be specified using optional arguments.
// Only the first casing option encountered will be applied.
func NormalizeAddress2(address string, options ...option) (string, int) {

	var (
		tokensNormalized int
		wasNormalized    bool
	)

	normOpts := parseOptions(options...)
	tokens := tokenize(address)
	for _, token := range tokens {

		// No need to normalize separators.
		if token.isSeparator {
			continue
		}

		if token.value, wasNormalized = normalizeSecondaryUnit(token.value, normOpts); wasNormalized {
			tokensNormalized++
		}
	}

	// All tokens have been normalized, or at least uppercased if no normalizations were applied.
	// Join them all back together and return them.
	return tokens.join(), tokensNormalized
}
