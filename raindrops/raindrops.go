package raindrops

import (
	"strconv"
)

// Convert converts a number intro a string.
func Convert(number int) string {
	var sounds string
	if number%3 == 0 {
		sounds += "Pling"
	}

	if number%5 == 0 {
		sounds += "Plang"
	}

	if number%7 == 0 {
		sounds += "Plong"
	}

	if sounds == "" {
		return strconv.Itoa(number)
	}

	return sounds
}
