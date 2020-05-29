package diffsquares

// SquareOfSum is the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	return ((n * n) * (n + 1) * (n + 1)) / 4
}

// SumOfSquares is the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference is the difference between SquareOfSum and SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareDifference returns the difference between the square of the sum and the sum of the squares of natural numbers.
func SquareDifference(number int) int {
	difference := Difference(number)
	return difference * difference
}
