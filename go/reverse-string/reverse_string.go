package reverse

import (
	"bytes"
)

//String - returns 'word' in it's reversed form
func String(word string) string {
	var reverse bytes.Buffer
	runes := []rune(word)

	for i := len(runes); i != 0; i-- {
		reverse.WriteRune(runes[i-1])
	}

	return reverse.String()
}
