package twofer

func ShareWith(name string) string {
	// ShareWith returns "One for X, one for me." where X is the name,
	// and returns the string if the name is missing.
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
