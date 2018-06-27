package isogram

import (
	"unicode"
)

//IsIsogram - determines whether or not 'word' is an isogram
func IsIsogram(word string) bool {
	isogram := true
	var values = map[rune]rune{}

	for _, runeValue := range word {
		runeValue = unicode.ToUpper(runeValue)
		if runeValue == 45 || runeValue == 32 {
			continue //skip " " (32) and "-" (45) as they don't matter in an isogram
		} else if _, ok := values[runeValue]; ok {
			isogram = false
		} else {
			values[runeValue] = runeValue
		}
	}
	return isogram
}
