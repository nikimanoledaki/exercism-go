// Package triangle determines if a triangle is equilateral, isoceles, or scalene.
package triangle

import (
	"math"
	"sort"
)

// Kind is a type of triangle.
type Kind string

const (
	// NaT is not a triangle
	NaT = "NaT"
	// Equ is an equilateral triangle
	Equ = "Equ"
	// Iso is an isoceles triangle
	Iso = "Iso"
	// Sca is a scalene triangle
	Sca = "Sca"
)

// KindFromSides identifies the Kind of triangle based on input.
func KindFromSides(a, b, c float64) Kind {

	sides := [3]float64{a, b, c}

	for _, f := range sides {
		if f <= 0 || math.IsNaN(f) || math.IsInf(f, 0) {
			return NaT
		}
	}

	sort.Float64s(sides[:])
	if sides[0]+sides[1] < sides[2] {
		return NaT
	} else if sides[0] == sides[1] && sides[0] == sides[2] {
		return Equ
	} else if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	} else {
		return Sca
	}

}
