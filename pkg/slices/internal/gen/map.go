package gen

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
