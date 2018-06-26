package isogram

import (
	"strings"
)

//IsIsogram - determines whether or not 'word' is an isogram
func IsIsogram(word string) bool {
	isogram := true
	var values = map[string]string{}

	for i := 0; i < len(word); i++ {
		letter := strings.ToLower(string(word[i]))
		if letter == " " || letter == "-" {
			continue //skip them as they don't matter in an isogram
		} else if _, ok := values[letter]; ok {
			isogram = false
		} else {
			values[letter] = letter
		}
	}
	return isogram
}
