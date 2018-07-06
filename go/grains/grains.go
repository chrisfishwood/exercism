package grains

import (
	"errors"
)

// Square -
func Square(val int) (uint64, error) {
	var previous uint64 = 1
	var grains uint64
	if val <= 0 || val > 64 {
		return 0, errors.New("number out of range (1-64)")
	} else if val == 1 {
		return 1, nil
	} else {
		for i := 2; i <= val; i++ {
			grains = 2 * previous
			previous = grains
		}
	}
	return uint64(grains), nil
}

// Total -
func Total() uint64 {
	var val uint64 = 1
	var total uint64
	for i := 0; i <= 64; i++ {
		val, _ = Square(i)
		total = total + val
	}
	return total
}
