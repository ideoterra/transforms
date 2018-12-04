package gen

import (
	"github.com/cheekybits/genny/generic"
)

// An A is a primitive type.
type A generic.Type

// B is a primitive type.
type B generic.Type

// Map contains map functions.
type Map struct{}

// ToB maps a []A to a []B.
func (Map) ToB(aa []A, mapFn func(A) B) []B {
	bb := []B{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
