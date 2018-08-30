package generic

// AsSliceTypeB applies a convert function to each element in aa and returns a
// new SliceTypeB.
func AsSliceTypeB(aa SliceTypeA, convert func(PrimitiveTypeA) PrimitiveTypeB) SliceTypeB {
	bb := SliceTypeB{}
	for _, a := range aa {
		bb = append(bb, convert(a))
	}
	return bb
}
