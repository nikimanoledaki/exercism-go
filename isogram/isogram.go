package isogram

import "unicode"

// IsIsogram returns true if inputted word is an isogram.
func IsIsogram(word string) bool {
	chars := make(map[rune]int)
	for _, c := range word {
		if unicode.IsLetter(c) {
			lowerC := unicode.ToLower(c)
			if chars[lowerC] == 1 {
				return false
			}
			chars[lowerC]++
		}
	}
	return true
}
