package strand

var complements = map[rune]string{
	'C': "G",
	'G': "C",
	'T': "A",
	'A': "U",
}

// ToRNA returns a RNA complement (per RNA transcription) given a DNA strand.
func ToRNA(dna string) string {
	var rna string
	for _, n := range dna {
		rna += complements[n]
	}

	return rna
}
