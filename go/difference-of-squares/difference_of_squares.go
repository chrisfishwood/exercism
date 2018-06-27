package diffsquares

// SumOfSquares - returns the sum of squares for 'n'
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// SquareOfSums - returns the square of sums for 'n'
func SquareOfSums(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// Difference - returns the difference between the square of sums and the sum of squares of 'n'
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
