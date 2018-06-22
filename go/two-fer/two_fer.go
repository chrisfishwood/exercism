package twofer

import (
	"strings"
)

const saying string = "One for X, one for me."

// ShareWith  - returns a sharing string wth 'name'
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return strings.Replace(saying, "X", name, 1)
}
