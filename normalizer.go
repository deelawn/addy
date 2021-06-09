package addy

import "strings"

func normalize(s string, lookup map[string]string) (string, bool) {

	s = strings.ToUpper(s)
	if normalized, ok := lookup[s]; ok {
		return normalized, true
	}

	return s, false
}

// NormalizeAddress1 will return an uppercased version of the input argument with all directional
// and suffix keywords normalized, along with a count of how many tokens were normalized.
func NormalizeAddress1(address string) (string, int) {

	var (
		tokensNormalized int
		wasNormalized    bool
	)

	tokens := tokenize(address)
	for _, token := range tokens {

		// No need to normalize separators.
		if token.isSeparator {
			continue
		}

		// There should never be any overlap between things that can be normalized either directionally or
		// by suffix; if not one, then try the other.
		token.value, wasNormalized = normalizeDirectional(token.value)
		if !wasNormalized {
			token.value, wasNormalized = normalizeSuffix(token.value)
		}

		if wasNormalized {
			tokensNormalized++
		}
	}

	// All tokens have been normalized, or at least uppercased if no normalizations were applied.
	// Join them all back together and return them.
	return tokens.join(), tokensNormalized
}

// NormalizeAddress2 will return an uppercased version of the input argument with all secondary
// unit keywords normalized, along with a count of how many tokens were normalized.
func NormalizeAddress2(address string) (string, int) {

	var (
		tokensNormalized int
		wasNormalized    bool
	)

	tokens := tokenize(address)
	for _, token := range tokens {

		// No need to normalize separators.
		if token.isSeparator {
			continue
		}

		if token.value, wasNormalized = normalizeSecondaryUnit(token.value); wasNormalized {
			tokensNormalized++
		}
	}

	// All tokens have been normalized, or at least uppercased if no normalizations were applied.
	// Join them all back together and return them.
	return tokens.join(), tokensNormalized
}
