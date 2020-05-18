// Package proverb returns the "for want of a nail" proverb.
package proverb

import (
	"fmt"
)

// Proverb returns the proverb given a slice of strings.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return rhyme
	}

	result := []string{}

	for i, piece := range rhyme {
		if i < len(rhyme)-1 {
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", piece, rhyme[i+1]))
		} else {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		}
	}

	return result

}
