// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xstring

// Do executes the test against a []String.
func (AllString) Do(aa []string, test func(string) bool) bool {
	for _, a := range aa {
		if !test(a) {
			return false
		}
	}
	return true
}
