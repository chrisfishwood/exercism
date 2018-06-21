package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate - Returns the acronyms of 's'
func Abbreviate(s string) string {
	wordsPattern := regexp.MustCompile("\\w+")
	words := wordsPattern.FindAllString(s, -1)
	return BuildAcro(words, "")
}

// BuildAcro - recursively builds an acronym from the the words in 'words'
func BuildAcro(words []string, acro string) string {
	if len(words) == 0 {
		return acro
	}
	first, rest := words[0], words[1:]
	acro += strings.ToUpper(string(first[0]))

	return BuildAcro(rest, acro)
}
