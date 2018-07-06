package grains

import (
	"errors"
	"fmt"
)

// Square -
func Square(val int) (uint64, error) {
	fmt.Printf("Square called with %d\n", val)
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
	//var squerror error
	var total uint64 = 1
	fmt.Println("calling total")
	total, _ = Square(64)
	return total
}
