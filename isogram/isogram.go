package isogram

import "unicode"

// IsIsogram returns true if inputted word is an isogram.
func IsIsogram(word string) bool {
	chars := Histogram(word)
	for k := range chars {
		if chars[k] > 1 {
			return false
		}
	}
	return true
}

// Histogram returns a frequency count of characters in word.
func Histogram(word string) map[rune]int {
	chars := make(map[rune]int)
	for _, c := range word {
		if unicode.IsLetter(c) {
			chars[unicode.ToLower(c)]++
		}
	}
	return chars
}
