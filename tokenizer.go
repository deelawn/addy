package addy

import (
	"strings"
)

// separators are the runes we will use to tokenize.
var separators = map[rune]bool{
	' ':  true,
	',':  true,
	'-':  true,
	'\t': true,
	'.':  true,
}

// token represents a string defined between separators or a separator itself.
type token struct {
	value       string
	isSeparator bool
}

// tokens are a token slice.
type tokens []*token

// join will construct a string from a slice of tokens.
func (t tokens) join() string {

	var joined string
	for _, tk := range t {
		joined += tk.value
	}

	return joined
}

// tokenize will take an address string break it into string tokens
func tokenize(address string) (tokens tokens) {

	address = strings.TrimSpace(address)
	reader := strings.NewReader(address)
	var (
		currentTokenValue string
		rn                rune
		err               error
	)

	for {

		// Read the next rune. If an error is returned, we've reached the end of the string.
		// Add the remaining token to the slice of tokens if we have one and exit the tokenizing loop.
		rn, _, err = reader.ReadRune()
		if err != nil {
			if len(currentTokenValue) != 0 {
				tokens = append(tokens, &token{value: currentTokenValue})
			}
			break
		}

		// If it isn't a separator, add the rune to the current token and move to the next rune.
		if !separators[rn] {
			currentTokenValue += string(rn)
			continue
		}

		// Current rune is a separator. Save and reset the token if we have one.
		if len(currentTokenValue) != 0 {
			tokens = append(tokens, &token{value: currentTokenValue})
			currentTokenValue = ""
		}

		// Add the separator to the tokens list if the previous toke wasn't the same separator.
		// This helps to remove thing like duplicate spaces, etc.
		if len(tokens) > 0 && tokens[len(tokens)-1].value != string(rn) {
			tokens = append(tokens, &token{value: string(rn), isSeparator: true})
		}
	}

	return
}
