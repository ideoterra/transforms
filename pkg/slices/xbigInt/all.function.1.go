// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xbigInt

import "math/big"

// Do executes the test against a []Big_Int.
func (AllBig_Int) Do(aa []big.Int, test func(big.Int) bool) bool {
	for _, a := range aa {
		if !test(a) {
			return false
		}
	}
	return true
}
