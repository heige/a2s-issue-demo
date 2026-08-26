// Package textnorm provides small, deterministic text normalization helpers.
package textnorm

import "strings"

// Normalize trims and collapses runs of whitespace into a single space.
// Recognized whitespace includes ASCII whitespace plus NBSP (U+00A0) and
// EM SPACE (U+2003) per GitHub Issue #5.
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

// isWhitespace reports whether character is an ASCII whitespace rune or one of
// the Unicode whitespace runes NBSP (U+00A0) and EM SPACE (U+2003).
func isWhitespace(character rune) bool {
	switch character {
	case ' ', '\t', '\n', '\r', '\f', '\v', '\u00a0', '\u2003':
		return true
	default:
		return false
	}
}
