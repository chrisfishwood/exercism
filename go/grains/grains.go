package grains

import (
	"errors"
)

// Square - Returns the number of grains on a given square on the chessboard
func Square(val int) (uint64, error) {

	var grains uint64

	if val < 1 || val > 64 {
		return 0, errors.New("number out of range (1-64)")
	}

	grains = 1 << uint64(val-1)
	return grains, nil
}

// Total - returns the total number of grains on the chessboard
func Total() uint64 {
	/*
		I chose the readable solution over the fast solution.

		`total = 1<<uint64(64) - 1` instead of the for loop had the following performance...

		~/projects/chrisfishwood/exercism/go/grains (master) $ go test --bench .
		goos: darwin
		goarch: amd64
		BenchmarkSquare-8       30000000                38.4 ns/op
		BenchmarkTotal-8        2000000000               0.28 ns/op

		...which was smoking fast, but I found it far less readable to someone coming into this code.
		I also like that `Total` uses `Square` as part of it's solution so a test for `Total` will also test `Square`.

		The for loop solution is definitely slower, but at ~80ns it's still pretty darn performant.

		~/projects/chrisfishwood/exercism/go/grains (master) $ go test --bench .
		goos: darwin
		goarch: amd64
		BenchmarkSquare-8       30000000                38.2 ns/op
		BenchmarkTotal-8        20000000                79.6 ns/op
	*/

	var val uint64
	var total uint64

	for i := 1; i <= 64; i++ {
		val, _ = Square(i)
		total = total + val
	}
	return total
}
