package generic

// AA is a slice of A.
type AA = []interface{}

// An A is a primitive of some type.
type A interface{}

// BB is a slice of B.
type BB = []interface{}

// B is a primitive of some type.
type B interface{}

// Map applies a transform to each element of the list, emitting a new list.
// Map is similar to Apply in that both project a transform across each element
// of a list. However, Apply mutates the source list, while Map does not
// mutate the source list. Thus, Map allows for the resulting list to be of a
// different type than the source list (at the cost of allocating a second
// list).
type Map struct{}

// ToB maps a slice of A to a slice of B.
func (Map) ToB(aa AA, mapFn func(A) B) BB {
	bb := BB{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
