package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a given string is valid according to the Luhn formula.
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
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
		if i%2 == 0 {
			sum += n
		} else {
			double := n * 2
			if double > 9 {
				double -= 9
			}
			sum += double
		}
	}
	return sum%10 == 0
}
