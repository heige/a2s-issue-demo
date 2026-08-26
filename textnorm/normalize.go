// Package textnorm provides small, deterministic text normalization helpers.
package textnorm

import "strings"

// Normalize trims and collapses runs of whitespace into a single space.
// In addition to ASCII whitespace (space, tab, newline, carriage return,
// form feed, vertical tab), the Unicode non-breaking space (U+00A0) and
// em space (U+2003) are also folded into a single plain space. Existing
// ASCII trim and collapse behavior is preserved.
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

func isWhitespace(character rune) bool {
	switch character {
	case ' ', '\t', '\n', '\r', '\f', '\v', '\u00a0', '\u2003':
		return true
	default:
		return false
	}
}
