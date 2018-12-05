package gen

// Map contains map functions.
type Map struct{}

// ToTB maps a []TA to a []TB.
func (Map) ToTB(aa []TA, mapFn func(TA) TB) []TB {
	bb := []TB{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
