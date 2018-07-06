package luhn

import (
	"bytes"
	"strconv"
)

// Valid - return whether or not 'word' is a valid Luhn account number
func Valid(word string) bool {
	word = trimAllSpace(word)

	if len(word) <= 1 {
		return false
	}

	isOdd := len(word)%2 == 1
	sum := 0

	for i := 0; i < len(word); i++ {
		digit, _ := strconv.Atoi(string(word[i]))

		if isOdd && i%2 == 0 {
			sum = sum + digit
			continue
		}

		digit = digit * 2
		if digit > 9 {
			digit = digit - 9
		}
		sum = sum + digit
	}

	return sum%10 == 0
}

// removes all spaces (" ") from 'word'
func trimAllSpace(word string) string {
	var noSpaces bytes.Buffer

	for i := 0; i < len(word); i++ {
		char := string(string(word[i]))
		if char != " " {
			noSpaces.WriteString(char)
		}
	}

	return noSpaces.String()
}
