package generic

// SliceType2 is a two dimensional slice of PrimitiveType
type SliceType2 []SliceType

// SliceType is a one dimensional slice of PrimitiveType.
type SliceType []PrimitiveType

// PrimitiveType is a placeholder for the type underpinning the generic SliceType.
type PrimitiveType interface{}

type Test func(PrimitiveType) bool

type Equality func(a, b PrimitiveType) bool

func ptr(aa SliceType) *SliceType {
	return &aa
}

// All applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
func (aa *SliceType) All(test Test) bool {
	return All(*aa, test)
}

// Any applies a test function to each element of the
// slice and returns true if the test function returns true for at least one
// item in the list.
//
// Any does not require that the source slice be sorted, and merely scans
// the slice, returning as soon as any element passes the supplied test. For
// a binary search, consider using sort.Search from the standard library.
func (aa *SliceType) Any(test Test) bool {
	return Any(*aa, test)
}

//Append adds the supplied values to the end of the slice.
func (aa *SliceType) Append(values ...PrimitiveType) *SliceType {
	Append(aa, values...)
	return aa
}

// Clear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func (aa *SliceType) Clear() *SliceType {
	*aa = nil
	return aa
}

// Clone returns a copy of aa
func (aa *SliceType) Clone() *SliceType {
	return ptr(Clone(*aa))
}

// Collect applies a given function against each item in slice aa and
// each item of a slice bb, and returns the concatenation of each result.
func (aa *SliceType) Collect(bb SliceType, collector func(a, b PrimitiveType) PrimitiveType) *SliceType {
	return ptr(Collect(*aa, bb, collector))
}

// Count applies the supplied test function to each element of the slice,
// and returns the count of items for which the test returns true.
func (aa *SliceType) Count(test Test) int64 {
	return Count(*aa, test)
}

// Dequeue returns a SliceType containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func (aa *SliceType) Dequeue() *SliceType {
	return ptr(Dequeue(aa))
}

// Difference returns a new slice that contains items that are not common
// between aa and bb. The supplied equal function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. Distinct values from aa are listed ahead of those from bb in the
// resulting slice.
func (aa *SliceType) Difference(bb SliceType, equality Equality) *SliceType {
	return ptr(Difference(*aa, bb, equality))
}
