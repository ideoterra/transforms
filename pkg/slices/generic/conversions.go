package generic

// PrimitiveTypeA is a base type.
type PrimitiveTypeA interface{}

// PrimitiveTypeB is a base type.
type PrimitiveTypeB interface{}

// SliceTypeA is a one dimensional slice of PrimitiveTypeA
type SliceTypeA []PrimitiveTypeA

// SliceTypeB is a two dimensional slice of PrimitiveTypeB
type SliceTypeB []PrimitiveTypeB

// AsSliceTypeB applies a converst function to each element in aa and returns a
// new SliceTypeB.
func AsSliceTypeB(aa SliceTypeA, convert func(PrimitiveTypeA) PrimitiveTypeB) SliceTypeB {
	bb := SliceTypeB{}
	for _, a := range aa {
		bb = append(bb, convert(a))
	}
	return bb
}
