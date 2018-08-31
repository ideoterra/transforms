package generic

import (
	"math/big"

	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

func ptr(aa SliceType) *SliceType {
	return &aa
}

func ptr2(aa SliceType2) *SliceType2 {
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
func (aa *SliceType) Collect(bb *SliceType, collector func(a, b PrimitiveType) PrimitiveType) *SliceType {
	return ptr(Collect(*aa, *bb, collector))
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
// between aa and bb. The supplied equality function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. Distinct values from aa are listed ahead of those from bb in the
// resulting slice.
func (aa *SliceType) Difference(bb *SliceType, equality Equality) *SliceType {
	return ptr(Difference(*aa, *bb, equality))
}

// Distinct removes all duplicates from the slice, using the supplied equality
// function to determine equality.
func (aa *SliceType) Distinct(equality Equality) *SliceType {
	Distinct(aa, equality)
	return aa
}

// Empty returns true if the length of the slice is zero.
func (aa *SliceType) Empty() bool {
	return Empty(*aa)
}

// End returns the a SliceType containing only the last element from aa.
func (aa *SliceType) End() *SliceType {
	return ptr(End(*aa))
}

// Enqueue places an item at the head of the slice.
func (aa *SliceType) Enqueue(a PrimitiveType) *SliceType {
	Enqueue(aa, a)
	return aa
}

// Expand applies an expansion function to each element of aa, and flattens
// the results into a single SliceType.
func (aa *SliceType) Expand(expansion func(PrimitiveType) SliceType) *SliceType {
	return ptr(Expand(*aa, expansion))
}

// Filter removes all items from the slice for which the supplied test function
// returns true.
func (aa *SliceType) Filter(test Test) *SliceType {
	Filter(aa, test)
	return aa
}

// FindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func (aa *SliceType) FindIndex(test Test) int64 {
	return FindIndex(*aa, test)
}

// First returns a SliceType containing the first element in the slice for which
// the supplied test function returns true.
func (aa *SliceType) First(test Test) *SliceType {
	return ptr(First(*aa, test))
}

// Fold applies a function to each item in slice aa, threading an accumulator
// through each iteration. The accumulated value is returned in a new SliceType
// once aa is fully scanned. Fold returns a SliceType rather than a
// PrimitiveType to be consistent with this package's Reduce implementation.
func (aa *SliceType) Fold(acc PrimitiveType, folder func(a, acc PrimitiveType) PrimitiveType) *SliceType {
	return ptr(Fold(*aa, acc, folder))
}

// FoldI applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned. Foldi returns a SliceType rather than a
// PrimitiveType to be consistent with this package's Reduce implementation.
func (aa *SliceType) FoldI(acc PrimitiveType, folder func(i int64, a, acc PrimitiveType) PrimitiveType) *SliceType {
	return ptr(FoldI(*aa, acc, folder))
}

// ForEach applies each element of the list to the given function.
// ForEach will stop iterating if fn return false.
func (aa *SliceType) ForEach(fn func(PrimitiveType) shared.Continue) *SliceType {
	ForEach(*aa, fn)
	return aa
}

// ForEachC concurrently applies each element of the list to the given function.
// The elements of the list are marshalled to a pool of goroutines, where each
// element is passed to fn concurrently.
//
// The concurrency pool is limited to contain no more than c active goroutines
// at any time. Note that if a pool size of 0 is supplied, this method
// will block indefinitely. This function will panic if a negative value is
// supplied for c.
//
// If any execution of fn returns shared.ContinueNo, ForEachC will cease marshalling
// any backlogged work, and will immediately set the cancellation flag to true.
// Any goroutines monitoring the cancelPending closure can wind down their
// activities as necessary. ForEachC will continue to block until all active
// goroutines exit cleanly.
func (aa *SliceType) ForEachC(c int, fn func(a PrimitiveType, cancelPending func() bool) shared.Continue) *SliceType {
	ForEachC(*aa, c, fn)
	return aa
}

// ForEachR applies each element of aa to a given function, scanning
// through the slice in reverse order, starting from the end and working towards
// the head.
func (aa *SliceType) ForEachR(fn func(PrimitiveType) shared.Continue) *SliceType {
	ForEachR(*aa, fn)
	return aa
}

// Group consolidates like-items into groups according to the supplied grouper
// function, and returns them as a SliceType2.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed.
func (aa *SliceType) Group(grouper func(PrimitiveType) int64) *SliceType2 {
	return ptr2(Group(*aa, grouper))
}

// GroupI consolidates like-items into groups according to the supplied grouper
// function, and returns them as a SliceType2.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from aa is also passed into the grouper function.
func (aa *SliceType) GroupI(grouper func(int64, PrimitiveType) int64) *SliceType2 {
	return ptr2(GroupI(*aa, grouper))
}

// Head returns a SliceType containing the first item from the aa. If aa is
// empty, the resulting SliceType will be empty.
func (aa *SliceType) Head() *SliceType {
	return ptr(Head(*aa))
}

// InsertAfter inserts an element in aa after the first element for which the
// supplied test function returns true. If none of the tests return true, the
// element is appended to the end of the aa.
func (aa *SliceType) InsertAfter(b PrimitiveType, test Test) *SliceType {
	InsertAfter(aa, b, test)
	return aa
}

// InsertBefore inserts an element in aa before the first element for which the
// supplied test function returns true. If none of the tests return true,
// the element is inserted at the head of aa.
func (aa *SliceType) InsertBefore(b PrimitiveType, test Test) *SliceType {
	InsertBefore(aa, b, test)
	return aa
}

// InsertAt inserts an element in aa at the specified index i, shifting the
// element originally at index i (and all subsequent elements) one position
// to the right. If i < 0, the element is inserted at index 0. If
// i >= len(aa), the value is appended to the end of aa.
func (aa *SliceType) InsertAt(a PrimitiveType, i int64) *SliceType {
	InsertAt(aa, a, i)
	return aa
}

// Intersection compares each element of aa to bb using the supplied equal
// function, and returns a SliceType containing the elements which are common
// to both aa and bb. Duplicates are removed in this operation.
func (aa *SliceType) Intersection(bb *SliceType, equality Equality) *SliceType {
	return ptr(Intersection(*aa, *bb, equality))
}

// IsProperSubset returns true if aa is a proper subset of bb.
// aa is considered a proper subset if all of its elements exist within bb, but
// bb also contains some elements that do not exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (aa *SliceType) IsProperSubset(bb *SliceType, equality Equality) bool {
	return IsProperSubset(*aa, *bb, equality)
}

// IsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (aa *SliceType) IsProperSuperset(bb *SliceType, equality Equality) bool {
	return IsProperSuperset(*aa, *bb, equality)
}

// IsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (aa *SliceType) IsSubset(bb *SliceType, equality Equality) bool {
	return IsSubset(*aa, *bb, equality)
}

// IsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (aa *SliceType) IsSuperset(bb *SliceType, equality Equality) bool {
	return IsSuperset(*aa, *bb, equality)
}

// Item returns a SliceType containing the element at aa[i].
// If len(aa) == 0, i < 0, or, i >= len(aa), the resulting slice will be empty.
func (aa *SliceType) Item(i int64) *SliceType {
	return ptr(Item(*aa, i))
}

// ItemFuzzy returns a SliceType containing the element at aa[i].
// If the supplied index is outside of the bounds of ItemFuzzy will attempt
// to retrieve the head or end element of aa according to the following rules:
// If len(aa) == 0 an empty SliceType is returned.
// If i < 0, the head of aa is returned.
// If i >= len(aa), the end of the aa is returned.
func (aa *SliceType) ItemFuzzy(i int64) *SliceType {
	return ptr(ItemFuzzy(*aa, i))
}

// Last applies a test function to each element in and returns a SliceType
// containing the last element for which the test returned true. If no elements
// pass the supplied test, the resulting SliceType will be empty.
func (aa *SliceType) Last(test Test) *SliceType {
	return ptr(Last(*aa, test))
}

// Len returns the length of aa.
func (aa *SliceType) Len() int {
	return Len(*aa)
}

// Map applies a tranform to each element of the list.
func (aa *SliceType) Map(mapFn func(PrimitiveType) PrimitiveType) *SliceType {
	Map(aa, mapFn)
	return aa
}

// None applies a test function to each element in and returns true if
// the test function returns false for all items.
func (aa *SliceType) None(test Test) bool {
	return None(*aa, test)
}

// Pairwise threads a transform function through passing to the transform
// successive two-element pairs, aa[i-1] && aa[i]. For the first pairing
// the supplied init value is supplied as the initial element in the pair.
func (aa *SliceType) Pairwise(init PrimitiveType, xform func(a, b PrimitiveType) PrimitiveType) *SliceType {
	return ptr(Pairwise(*aa, init, xform))
}

// Partition applies a test function to each element in and returns
// a SliceType2 where SliceType2[0] contains a SliceType with all elements for
// whom the test function returned true, and where SliceType2[1] contains a
// SliceType with all elements for whom the test function returned false.
//
// Partition is a special case of the Group function.
func (aa *SliceType) Partition(test Test) *SliceType2 {
	return ptr2(Partition(*aa, test))
}

// Permutable returns true if the number of permutations for aa exceeds
// MaxInt64.
func (aa *SliceType) Permutable() bool {
	return Permutable(*aa)
}

// Permutations returns the number of permutations that exist given the current
// number of items in the aa.
func (aa *SliceType) Permutations() *big.Int {
	return Permutations(*aa)
}

// Permute returns a SliceType2 which contains a SliceType for each permutation
// of aa.
//
// This function will panic if it determines that the list is not permutable
// (see Permutable function).
//
// Permute makes no assumptions about whether or not the elements in aa are
// distinct. Permutations are created positionally, and do not involve any
// equality checks. As such, if it important that Permute operate on a set of
// distinct elements, pass aa through one of the Distinct transforms before
// passing it to Permute().
//
// Permute is implemented using Heap's algorithm.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func (aa *SliceType) Permute() *SliceType2 {
	return ptr2(Permute(*aa))
}

// Pop returns a SliceType containing the head element from and removes the
// element from aa. If aa is empty, the returned SliceType will also be empty.
func (aa *SliceType) Pop() *SliceType {
	Pop(aa)
	return aa
}

// Push places a prepends a new element at the head of aa.
func (aa *SliceType) Push(a PrimitiveType) *SliceType {
	Push(aa, a)
	return aa
}

// Reduce applies a reducer function to each element in threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new SliceType. If aa is empty, the resulting SliceType
// will also be empty.
func (aa *SliceType) Reduce(reducer func(a, acc PrimitiveType) PrimitiveType) *SliceType {
	return ptr(Reduce(*aa, reducer))

}

// Remove applies a test function to each item in the list, and removes any item
// for which the test returns true.
func (aa *SliceType) Remove(test Test) *SliceType {
	Remove(aa, test)
	return aa
}

// RemoveAt removes the item at the specified index from the slice.
// If len(aa) == 0, aa == nil, i < 0, or i >= len(aa), this function will do
// nothing.
func (aa *SliceType) RemoveAt(i int64) *SliceType {
	RemoveAt(aa, i)
	return aa
}

// Reverse reverses the order of aa.
func (aa *SliceType) Reverse() *SliceType {
	Reverse(aa)
	return aa
}

// Skip removes the first n elements from aa.
//
// Note that Skip(len(aa)) will remove all items from the list, but does not
// "clear" the slice, meaning that the list remains allocated in memory.
// To fully de-pointer the slice, and ensure it is available for garbage
// collection as soon as possible, consider using Clear().
func (aa *SliceType) Skip(n int64) *SliceType {
	Skip(aa, n)
	return aa
}

// SkipWhile scans through aa starting at the head, and removes all
// elements from aa while the test function returns true.
// SkipWhile stops removing any further items from aa after the first test that
// returns false.
func (aa *SliceType) SkipWhile(test Test) *SliceType {
	SkipWhile(aa, test)
	return aa
}

// Sort sorts using the supplied less function to determine order.
// Sort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func (aa *SliceType) Sort(less func(a, b PrimitiveType) bool) *SliceType {
	Sort(aa, less)
	return aa
}

// SplitAfter finds the first element b for which a test function returns true,
// and returns a SliceType2 where SliceType2[0] contains the first half of aa
// and SliceType2[1] contains the second half of aa. Element b will be included
// in SliceType2[0]. If the no element can be found for which the test returns
// true, SliceType2[0] will contain and SliceType2[1] will be empty.
func (aa *SliceType) SplitAfter(test Test) *SliceType2 {
	return ptr2(SplitAfter(*aa, test))
}

// SplitAt splits aa at index i, and returns a SliceType2 which contains the
// two split halves of aa. aa[i] will be included in SliceType2[1].
// If i < 0, all of aa will be placed in SliceType2[0] and SliceType2[1] will
// be empty. Conversly, if i >= len(aa), all of aa will be placed in
// SliceType2[1] and SliceType2[0] will be empty. If aa is nil or empty,
// SliceType2 will contain two empty slices.
func (aa *SliceType) SplitAt(i int64) *SliceType2 {
	return ptr2(SplitAt(*aa, i))
}

// SplitBefore finds the first element b for which a test function returns true,
// and returns a SliceType2 where SliceType2[0] contains the first half of aa
// and SliceType2[1] contains the second half of aa. Element b will be included
// in SliceType2[1]
func (aa *SliceType) SplitBefore(test Test) *SliceType2 {
	return ptr2(SplitBefore(*aa, test))
}

// String returns a string representation of suitable for use
// with fmt.Print, or other similar functions. String should be regarded as
// informational, and should not be relied upon to formally serialize a
// SliceType.
func (aa *SliceType) String() string {
	return String(*aa)
}

// SwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of SwapIndex does nothing.
func (aa *SliceType) SwapIndex(i, j int64) *SliceType {
	SwapIndex(*aa, i, j)
	return aa
}

// Tail removes the current head element from aa.
// This equivelant to RemoveAt(0)
func (aa *SliceType) Tail() *SliceType {
	Tail(aa)
	return aa
}

// Take retains the first n elements of and removes all remaining elements
// from the slice. If n < 0 or n >= len(aa), Take does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func (aa *SliceType) Take(n int64) *SliceType {
	Take(aa, n)
	return aa
}

// TakeWhile applies a test function to each element in and retains all
// elements of aa so long as the test function returns true. As soon as the test
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func (aa *SliceType) TakeWhile(test Test) *SliceType {
	TakeWhile(aa, test)
	return aa
}

// Union appends slice bb to slice aa.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func (aa *SliceType) Union(bb *SliceType) *SliceType {
	Union(aa, *bb)
	return aa
}

// Unzip splits aa into a SliceType2, such that SliceType2[0] contains all odd
// indices from and SliceType2[1] contains all even indices from aa.
func (aa *SliceType) Unzip() *SliceType2 {
	return ptr2(Unzip(*aa))
}

// WindowCentered applies a windowing function across the using a centered
// window of the specified size.
func (aa *SliceType) WindowCentered(windowSize int64, windowFn func(window SliceType) PrimitiveType) *SliceType {
	return ptr(WindowCentered(*aa, windowSize, windowFn))
}

// WindowLeft applies a windowing function across using a left-sided window
// of the specified size.
func (aa *SliceType) WindowLeft(windowSize int64, windowFn func(window SliceType) PrimitiveType) *SliceType {
	return ptr(WindowLeft(*aa, windowSize, windowFn))
}

// WindowRight applies a windowing function across using a right-sided
// window of the specified size.
func (aa *SliceType) WindowRight(windowSize int64, windowFn func(window SliceType) PrimitiveType) *SliceType {
	return ptr(WindowRight(*aa, windowSize, windowFn))
}

// Zip interleaves the contents of aa with bb, and returns the result as a
// new SliceType. aa[0] is evaluated first. Thus if aa and bb are the same
// length, slice aa will occupy the odd indices of the result slice, and bb
// will occupy the even indices of the result slice. If aa and bb are not
// the same length, Zip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func (aa *SliceType) Zip(bb *SliceType) *SliceType {
	return ptr(Zip(*aa, *bb))
}
