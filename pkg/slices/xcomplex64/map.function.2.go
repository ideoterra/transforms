// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xcomplex64

// ToError maps a []Complex64 to a []Error.
func (MapComplex64) ToError(aa []complex64, mapFn func(complex64) error) []error {
	bb := []error{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
