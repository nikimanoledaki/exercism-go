package grains

import (
	"errors"
	"math"
)

// Square returns the current chessboard square as a uint64.
func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, errors.New("Input cannot be less than or equal to 0 or greater than 64")
	}
	o := float64(square)
	p := math.Pow(2, o-1)
	q := uint64(p)
	return q, nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	return 18446744073709551615
}
