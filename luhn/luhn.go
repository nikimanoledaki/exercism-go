package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if a given string is valid according to the Luhn formula.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	length := len(s)
	if length < 2 {
		return false
	}
	sum := 0
	secondDigit := length%2 == 0
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
		digit := int(r - '0')
		if secondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		secondDigit = !secondDigit
	}
	return sum%10 == 0
}
