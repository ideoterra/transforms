package gen

// Do executes the test against a []TA.
func (AllTA) Do(aa []TA, test func(TA) bool) bool {
	for _, a := range aa {
		if !test(a) {
			return false
		}
	}
	return true
}
