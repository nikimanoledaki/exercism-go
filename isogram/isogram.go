package isogram

import "unicode"

// IsIsogram returns true if inputted word is an isogram.
func IsIsogram(word string) bool {
	chars := make(map[rune]int)
	for _, c := range word {
		lowerC := unicode.ToLower(c)
		if unicode.IsLetter(c) {
			chars[lowerC]++
		}
		if chars[lowerC] > 1 {
			return false
		}
	}
	return true
}
