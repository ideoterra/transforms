package generic

// AsSliceTypeB applies a convert function to each element in aa and returns a
// new SliceTypeB.
func (aa *SliceTypeA) AsSliceTypeB(convert func(PrimitiveTypeA) PrimitiveTypeB) *SliceTypeB {
	return ptrB(AsSliceTypeB(*aa, convert))
}

func ptrB(aa SliceTypeB) *SliceTypeB {
	return &aa
}
