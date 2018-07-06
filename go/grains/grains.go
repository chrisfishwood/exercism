package grains

import (
	"errors"
	"fmt"
)

// Square -
func Square(val int) (uint64, error) {
	fmt.Printf("Square called with %d\n", val)
	var total uint64 = 1
	var previous uint64 = 1
	if val == 0 {
		return 0, errors.New("0 is bad, mkay")
	} else if val == 1 {
		return total, nil
	} else {
		for i := 2; i <= val; i++ {
			v := 2 * previous
			previous = v
			total = total + v
			fmt.Printf("Previous: %d, Total: %d\n", previous, total)
		}
	}
	return uint64(total), errors.New("Some error")
}

// Total -
func Total() uint64 {
	//var squerror error
	var total uint64 = 1
	total, _ = Square(64)
	return total
}
