// Package twofer provides function for sharing with you and me.
package twofer

import "fmt"

// ShareWith prints "One for X, one for me." where x is the name provided.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
