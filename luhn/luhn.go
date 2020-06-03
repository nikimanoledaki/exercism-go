package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a given string is valid according to the Luhn formula.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	length := len(s)
	if length < 2 {
		return false
	}

	sum := 0
	for i := 0; i < length; i++ {
		n, err := strconv.Atoi(string(s[length-1-i]))

		if err != nil {
			return false
		}

		sum += n

		if i%2 != 0 {
			sum += n

			double := n * 2
			if double > 9 {
				sum -= 9
			}
		}
	}
	return sum%10 == 0
}
