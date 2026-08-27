// Package textnorm provides small, deterministic text normalization helpers.
package textnorm

import (
	"strings"
	"unicode"
)

// Normalize trims and collapses runs of Unicode whitespace into a single space.
func Normalize(input string) string {
	var output strings.Builder
	output.Grow(len(input))

	wroteText := false
	pendingSpace := false
	for _, character := range input {
		if unicode.IsSpace(character) {
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
