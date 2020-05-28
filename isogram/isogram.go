package isogram

import "unicode"

// IsIsogram returns true if inputted word is an isogram.
func IsIsogram(word string) bool {
	chars := make(map[rune]int)
	for _, c := range word {
		lowerC := unicode.ToLower(c)
		if chars[lowerC] == 1 {
			return false
		} else if unicode.IsLetter(c) {
			chars[lowerC]++
		}
	}
	return true
}
