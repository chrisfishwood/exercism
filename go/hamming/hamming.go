package hamming

import "errors"

//Distance - computes the hamming disstance between 'a' and 'b'
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("shit's on fire yo (and the strings should be the same length)")
	}

	ctr := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ctr++
		}
	}
	return ctr, nil
}
