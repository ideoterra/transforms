package gen

// ToTB maps a []TA to a []TB.
func (MapTA) ToTB(aa []TA, mapFn func(TA) TB) []TB {
	bb := []TB{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
