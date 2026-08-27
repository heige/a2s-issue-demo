// Package textnorm provides small, deterministic text normalization helpers.
package textnorm

import (
	"strings"
	"unicode"
)

// Normalize trims and collapses runs of whitespace into a single space.
func Normalize(input string) string {
	var output strings.Builder
	output.Grow(len(input))

	wroteText := false
	pendingSpace := false
	for _, character := range input {
		if isWhitespace(character) {
			if wroteText {
				pendingSpace = true
			}
			continue
		}

		if pendingSpace {
			output.WriteByte(' ')
			pendingSpace = false
		}
		output.WriteRune(character)
		wroteText = true
	}

	return output.String()
}

// isWhitespace reports whether character is Unicode whitespace per unicode.IsSpace,
// which includes NBSP (U+00A0) and EM SPACE (U+2003) in addition to ASCII whitespace.
func isWhitespace(character rune) bool {
	return unicode.IsSpace(character)
}
