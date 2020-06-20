package grains

import "math"

// Square returns the current chessboard square as a uint64.
func Square(square int) (uint64, error) {
	if square == 1 {
		return 1, nil
	}
	o := float64(square)
	p := math.Pow(2, o-1)
	q := uint64(p)
	return q, nil
}

// Total says hi.
func Total() uint64 {
	// p := math.Pow(64, 2)
	// q := uint64(p)
	return 0
}
