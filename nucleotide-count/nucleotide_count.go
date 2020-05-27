package dna

import "errors"

// DNA is a list of nucleotides.
type DNA string

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, n := range d {
		if n != 'A' && n != 'C' && n != 'G' && n != 'T' {
			return h, errors.New("Not valid")
		}
		h[n]++
	}
	return h, nil
}
