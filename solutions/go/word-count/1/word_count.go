package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})
	freq := make(Frequency)
	for _, word := range words {
		cleanedWord := strings.Trim(word, "'")
		if cleanedWord != "" {
			freq[cleanedWord]++
		}
	}
	return freq
}
