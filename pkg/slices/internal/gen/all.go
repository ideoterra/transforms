package gen

// All applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
type All struct{}

// Do executes the test against a []TA.
func (All) Do(aa []TA, test func(TA) bool) bool {
	for _, a := range aa {
		if !test(a) {
			return false
		}
	}
	return true
}
