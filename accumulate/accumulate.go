package accumulate

// Accumulate returns a new collection with the result of applying an operation on it.
func Accumulate(list []string, f func(string) string) []string {
	newList := make([]string, len(list))
	for i := range list {
		newList[i] = f(list[i])
	}
	return newList
}
