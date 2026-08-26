// Package textnorm provides small, deterministic text normalization helpers.
package textnorm

import "strings"

// Normalize trims and collapses runs of ASCII whitespace into a single space.
func Normalize(input string) string {
	var output strings.Builder
	output.Grow(len(input))

	wroteText := false
	pendingSpace := false
	for _, character := range input {
		if isASCIIWhitespace(character) {
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

func isASCIIWhitespace(character rune) bool {
	switch character {
	case ' ', '\t', '\n', '\r', '\f', '\v':
		return true
	default:
		return false
	}
}
