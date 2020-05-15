// Package hamming calculates the hamming distance.
package hamming

import "errors"

// Distance compares two strings and counts the differences.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length of two strings must be equal")
	}

	var difference int
	for index := range a {
		if a[index] != b[index] {
			difference++
		}
	}

	return difference, nil
}
