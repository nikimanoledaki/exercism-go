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

		digit := n

		var secondDigit = (i%2 != 0)
		if secondDigit {
			digit += n
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}
	return sum%10 == 0
}
