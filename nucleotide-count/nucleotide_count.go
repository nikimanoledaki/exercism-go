package dna

// DNA is a list of nucleotides.
type DNA string

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	return h, nil
}
