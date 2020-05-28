package diffsquares

// SquareOfSum is the square of the sum of the first N natural numbers.
func SquareOfSum(number int) int {
	var result int
	for n := 1; n <= number; n++ {
		result += n
	}
	result *= result
	return result
}

// SumOfSquares is the sum of the squares of the first N natural numbers.
func SumOfSquares(number int) int {
	var result int
	for n := 1; n <= number; n++ {
		nSquared := n * n
		result += nSquared
	}
	return result
}

// Difference is the difference between SquareOfSum and SumOfSquares.
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}

// SquareDifference returns the difference between the square of the sum and the sum of the squares of natural numbers.
func SquareDifference(number int) int {
	difference := Difference(number)
	return difference * difference
}
