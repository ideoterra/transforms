// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xuint32

// ToError maps a []Uint32 to a []Error.
func (MapUint32) ToError(aa []uint32, mapFn func(uint32) error) []error {
	bb := []error{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
