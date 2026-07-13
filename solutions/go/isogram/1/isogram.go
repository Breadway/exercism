package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]struct{})
	// Use runes for safety
	for _, r := range strings.ToLower(word) {
		// Skip spaces and hyphens immediately
		if r == ' ' || r == '-' {
			continue
		}

		// If the rune is already in our map, it's not an isogram
		if _, exists := seen[r]; exists {
			return false
		}

		// Record that we've seen this character
		seen[r] = struct{}{}
	}

	return true
}
