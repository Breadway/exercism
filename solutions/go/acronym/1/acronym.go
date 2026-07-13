package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate seperates the words of a string, and creates an acronym from the first letter of each word.
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})
	acronym := strings.Builder{}
	for _, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			acronym.WriteRune(unicode.ToUpper(runes[0]))
		}
	}
	return acronym.String()
}
