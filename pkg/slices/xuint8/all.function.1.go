// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xuint8

// Do executes the test against a []Uint8.
func (AllUint8) Do(aa []uint8, test func(uint8) bool) bool {
	for _, a := range aa {
		if !test(a) {
			return false
		}
	}
	return true
}
