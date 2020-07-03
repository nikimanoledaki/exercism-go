package grains

import (
	"errors"
)

// Square returns the current chessboard square as a uint64.
func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, errors.New("Input cannot be less than or equal to 0 or greater than 64")
	}
	return (1 << (square - 1)), nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
