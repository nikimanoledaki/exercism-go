package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if a given string is valid according to the Luhn formula.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	length, secondDigit := len(s), false
	if length%2 == 0 {
		secondDigit = true
	}
	sum := 0
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
	return length >= 2 && sum%10 == 0
}
