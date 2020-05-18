package raindrops

import "fmt"

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

	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		sounds += fmt.Sprintf("%d", number)
	}

	return sounds
}
