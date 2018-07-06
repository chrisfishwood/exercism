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
	var reversedString bytes.Buffer

	for i := 0; i < len(word); i++ {
		stringDigit := string(word[i])
		if isOdd && i%2 == 0 {
			reversedString.WriteString(stringDigit)
			continue
		}
		intDouble, _ := strconv.Atoi(stringDigit)
		intDouble = intDouble * 2
		if intDouble > 9 {
			intDouble = intDouble - 9
		}
		reversedString.WriteString(strconv.FormatInt(int64(intDouble), 10))
	}

	word = reversedString.String()
	sum := 0
	for i := 0; i < len(word); i++ {
		sumInt, _ := strconv.Atoi(string(word[i]))
		sum = sum + sumInt
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
