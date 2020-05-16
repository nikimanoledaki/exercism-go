// Package hamming calculates the hamming distance.
package hamming

import (
	"errors"
	"strings"
)

// Distance compares two strings and counts the differences.
func Distance(a, b string) (int, error) {
	charsA, charsB := strings.Split(a, ""), strings.Split(b, "")
	if len(charsA) != len(charsB) {
		return 0, errors.New("length of two strings must be equal")
	}

	var difference int
	for index := range charsA {
		if charsA[index] != charsB[index] {
			difference++
		}
	}

	return difference, nil
}
