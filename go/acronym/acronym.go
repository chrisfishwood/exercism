package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate - Returns the acronyms of 's'
func Abbreviate(s string) string {
	var acro string
	wordsPattern := regexp.MustCompile("\\w+")
	words := wordsPattern.FindAllString(s, -1)
	for _, word := range words {
		acro += strings.ToUpper(string(word[0]))
	}
	return acro
}
