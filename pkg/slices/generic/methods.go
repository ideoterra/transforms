package generic

import (
	"math/big"

	"github.com/jecolasurdo/transforms/pkg/slices/generic/closures"
	"github.com/jecolasurdo/transforms/pkg/slices/generic/iface"
	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

func unbox(tt1 []interface{}) *SliceType {
	tt2 := SliceType(tt1)
	return &tt2
}

func unbox2(tt1 [][]interface{}) *SliceType2 {
	tt2 := SliceType2(tt1)
	return &tt2
}

func box(tt1 SliceType) []interface{} {
	return ([]interface{})(tt1)
}

func boxP(tt1 *SliceType) *[]interface{} {
	return (*[]interface{})(tt1)
}

// All applies a condition function to each element in the slice, and returns true if
// the condition function returns true for all items in the slice.
func (tt1 *SliceType) All(condition closures.ConditionFn) bool {
	return All(*tt1, condition)
}

// Any applies a condition function to each element of the
// slice and returns true if the condition function returns true for at least one
// item in the list.
//
// Any does not require that the source slice be sorted, and merely scans
// the slice, returning as soon as any element passes the supplied condition. For
// a binary search, consider using sort.Search from the standard library.
func (tt1 *SliceType) Any(condition closures.ConditionFn) bool {
	return Any(box(*tt1), condition)
}

//Append adds the supplied values to the end of the slice.
func (tt1 *SliceType) Append(values ...interface{}) genericiface.GenericSliceIface {
	Append(boxP(tt1), values...)
	return tt1
}

//AsSlice returns the elements of the slice as a bare []interface{}.
func (tt1 *SliceType) AsSlice() []interface{} {
	return []interface{}(*tt1)
}

// Clear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func (tt1 *SliceType) Clear() genericiface.GenericSliceIface {
	*tt1 = nil
	return tt1
}

// Clone returns a copy of tt1
func (tt1 *SliceType) Clone() genericiface.GenericSliceIface {
	return unbox(Clone(box(*tt1)))
}

// Collect applies a given function against each item in slice tt1 and
// each item of a slice tt2, and returns the concatenation of each result.
func (tt1 *SliceType) Collect(tt2 []interface{}, collector func(t1 interface{}, t2 interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(Collect(box(*tt1), tt2, collector))
}

// Count applies the supplied condition function to each element of the slice,
// and returns the count of items for which the condition returns true.
func (tt1 *SliceType) Count(condition closures.ConditionFn) int64 {
	return Count(*tt1, condition)
}

// Dequeue returns a genericiface.GenericSliceIface containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func (tt1 *SliceType) Dequeue() genericiface.GenericSliceIface {
	return unbox(Dequeue(boxP(tt1)))
}

// Difference returns a new slice that contains items that are not common
// between tt1 and tt2. The supplied equality function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. Distinct values from tt1 are listed ahead of those from tt2 in the
// resulting slice.
func (tt1 *SliceType) Difference(tt2 []interface{}, equality closures.EqualityFn) genericiface.GenericSliceIface {
	return unbox(Difference(box(*tt1), tt2, equality))

}

// Distinct removes all duplicates from the slice, using the supplied equality
// function to determine equality.
func (tt1 *SliceType) Distinct(equality closures.EqualityFn) genericiface.GenericSliceIface {
	Distinct(boxP(tt1), equality)
	return tt1
}

// Empty returns true if the length of the slice is zero.
func (tt1 *SliceType) Empty() bool {
	return Empty(*tt1)
}

// End returns the a genericiface.GenericSliceIface containing only the last element from tt1.
func (tt1 *SliceType) End() genericiface.GenericSliceIface {
	return unbox(End(box(*tt1)))

}

// Enqueue places an item at the head of the slice.
func (tt1 *SliceType) Enqueue(t1 interface{}) genericiface.GenericSliceIface {
	Enqueue(boxP(tt1), t1)
	return tt1
}

// Expand applies an expansion function to each element of tt1, and flattens
// the results into a single genericiface.GenericSliceIface.
func (tt1 *SliceType) Expand(expansion func(interface{}) []interface{}) genericiface.GenericSliceIface {
	return unbox(Expand(box(*tt1), expansion))
}

// Filter removes all items from the slice for which the supplied condition function
// returns true.
func (tt1 *SliceType) Filter(condition closures.ConditionFn) genericiface.GenericSliceIface {
	Filter(boxP(tt1), condition)
	return tt1
}

// FindIndex returns the index of the first element in the slice for which the
// supplied condition function returns true. If no matches are found, -1 is returned.
func (tt1 *SliceType) FindIndex(condition closures.ConditionFn) int64 {
	return FindIndex(*tt1, condition)
}

// First returns a genericiface.GenericSliceIface containing the first element in the slice for which
// the supplied condition function returns true.
func (tt1 *SliceType) First(condition closures.ConditionFn) genericiface.GenericSliceIface {
	return unbox(First(box(*tt1), condition))

}

// Fold applies a function to each item in slice tt1, threading an accumulator
// through each iteration. The accumulated value is returned in a new genericiface.GenericSliceIface
// once tt1 is fully scanned. Fold returns a genericiface.GenericSliceIface rather than a
// interface{} to be consistent with this package's Reduce implementation.
func (tt1 *SliceType) Fold(acc interface{}, folder func(a, acc interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(Fold(box(*tt1), acc, folder))
}

// FoldI applies a function to each item in slice tt1, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once tt1 is fully scanned. Foldi returns a genericiface.GenericSliceIface rather than a
// interface{} to be consistent with this package's Reduce implementation.
func (tt1 *SliceType) FoldI(acc interface{}, folder func(i int64, a, acc interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(FoldI(box(*tt1), acc, folder))
}

// ForEach applies each element of the list to the given function.
// ForEach will stop iterating if fn return false.
func (tt1 *SliceType) ForEach(fn func(interface{}) shared.Continue) genericiface.GenericSliceIface {
	ForEach(*tt1, fn)
	return tt1
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
func (tt1 *SliceType) ForEachC(c int, fn func(t1 interface{}, cancelPending func() bool) shared.Continue) genericiface.GenericSliceIface {
	ForEachC(*tt1, c, fn)
	return tt1
}

// ForEachR applies each element of tt1 to a given function, scanning
// through the slice in reverse order, starting from the end and working towards
// the head.
func (tt1 *SliceType) ForEachR(fn func(interface{}) shared.Continue) genericiface.GenericSliceIface {
	ForEachR(*tt1, fn)
	return tt1
}

// Group consolidates like-items into groups according to the supplied grouper
// function, and returns them as a genericiface.GenericSlice2Iface.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed.
func (tt1 *SliceType) Group(grouper func(interface{}) int64) genericiface.GenericSlice2Iface {
	return unbox2(Group(box(*tt1), grouper))
}

// GroupI consolidates like-items into groups according to the supplied grouper
// function, and returns them as a genericiface.GenericSlice2Iface.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from tt1 is also passed into the grouper function.
func (tt1 *SliceType) GroupI(grouper func(int64, interface{}) int64) genericiface.GenericSlice2Iface {
	return unbox2(GroupI(box(*tt1), grouper))
}

// Head returns a genericiface.GenericSliceIface containing the first item from the tt1. If tt1 is
// empty, the resulting genericiface.GenericSliceIface will be empty.
func (tt1 *SliceType) Head() genericiface.GenericSliceIface {
	return unbox(Head(box(*tt1)))
}

// InsertAfter inserts an element in tt1 after the first element for which the
// supplied condition function returns true. If none of the tests return true, the
// element is appended to the end of the tt1.
func (tt1 *SliceType) InsertAfter(t2 interface{}, condition closures.ConditionFn) genericiface.GenericSliceIface {
	InsertAfter(boxP(tt1), t2, condition)
	return tt1
}

// InsertBefore inserts an element in tt1 before the first element for which the
// supplied condition function returns true. If none of the tests return true,
// the element is inserted at the head of tt1.
func (tt1 *SliceType) InsertBefore(t2 interface{}, condition closures.ConditionFn) genericiface.GenericSliceIface {
	InsertBefore(boxP(tt1), t2, condition)
	return tt1
}

// InsertAt inserts an element in tt1 at the specified index i, shifting the
// element originally at index i (and all subsequent elements) one position
// to the right. If i < 0, the element is inserted at index 0. If
// i >= len(tt1), the value is appended to the end of tt1.
func (tt1 *SliceType) InsertAt(t1 interface{}, i int64) genericiface.GenericSliceIface {
	InsertAt(boxP(tt1), t1, i)
	return tt1
}

// Intersection compares each element of tt1 to tt2 using the supplied equal
// function, and returns a genericiface.GenericSliceIface containing the elements which are common
// to both tt1 and tt2. Duplicates are removed in this operation.
func (tt1 *SliceType) Intersection(tt2 []interface{}, equality closures.EqualityFn) genericiface.GenericSliceIface {
	return unbox(Intersection(box(*tt1), tt2, equality))
}

// IsProperSubset returns true if tt1 is a proper subset of tt2.
// tt1 is considered a proper subset if all of its elements exist within tt2, but
// tt2 also contains some elements that do not exist within tt1.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (tt1 *SliceType) IsProperSubset(tt2 []interface{}, equality closures.EqualityFn) bool {
	return IsProperSubset(box(*tt1), tt2, equality)
}

// IsProperSuperset returns true if tt1 is a proper superset of tt2.
// tt1 is considered a proper superset if it contains all of tt2's elements, but
// tt1 also contains some elements that do not exist within tt2.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (tt1 *SliceType) IsProperSuperset(tt2 []interface{}, equality closures.EqualityFn) bool {
	return IsProperSuperset(box(*tt1), tt2, equality)
}

// IsSubset returns true if tt1 is a subset of tt2.
// tt1 is considered a subset if all of its elements exist within tt2.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (tt1 *SliceType) IsSubset(tt2 []interface{}, equality closures.EqualityFn) bool {
	return IsSubset(box(*tt1), tt2, equality)
}

// IsSuperset returns true if tt1 is a superset of tt2.
// tt1 is considered a superset if all of tt2's elements exist within tt1.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func (tt1 *SliceType) IsSuperset(tt2 []interface{}, equality closures.EqualityFn) bool {
	return IsSuperset(box(*tt1), tt2, equality)
}

// Item returns a genericiface.GenericSliceIface containing the element at tt1[i].
// If len(tt1) == 0, i < 0, or, i >= len(tt1), the resulting slice will be empty.
func (tt1 *SliceType) Item(i int64) genericiface.GenericSliceIface {
	return unbox(Item(box(*tt1), i))
}

// ItemFuzzy returns a genericiface.GenericSliceIface containing the element at tt1[i].
// If the supplied index is outside of the bounds of ItemFuzzy will attempt
// to retrieve the head or end element of tt1 according to the following rules:
// If len(tt1) == 0 an empty genericiface.GenericSliceIface is returned.
// If i < 0, the head of tt1 is returned.
// If i >= len(tt1), the end of the tt1 is returned.
func (tt1 *SliceType) ItemFuzzy(i int64) genericiface.GenericSliceIface {
	return unbox(ItemFuzzy(box(*tt1), i))
}

// Last applies a condition function to each element in and returns a genericiface.GenericSliceIface
// containing the last element for which the condition returned true. If no elements
// pass the supplied condition, the resulting genericiface.GenericSliceIface will be empty.
func (tt1 *SliceType) Last(condition closures.ConditionFn) genericiface.GenericSliceIface {
	return unbox(Last(box(*tt1), condition))
}

// Len returns the length of tt1.
func (tt1 *SliceType) Len() int {
	return Len(box(*tt1))
}

// Map applies a tranform to each element of the list.
func (tt1 *SliceType) Map(mapFn func(interface{}) interface{}) genericiface.GenericSliceIface {
	Map(boxP(tt1), mapFn)
	return tt1
}

// None applies a condition function to each element in and returns true if
// the condition function returns false for all items.
func (tt1 *SliceType) None(condition closures.ConditionFn) bool {
	return None(box(*tt1), condition)
}

// Pairwise threads a transform function through passing to the transform
// successive two-element pairs, tt1[i-1] && tt1[i]. For the first pairing
// the supplied init value is supplied as the initial element in the pair.
func (tt1 *SliceType) Pairwise(init interface{}, xform func(t1 interface{}, t2 interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(Pairwise(box(*tt1), init, xform))
}

// Partition applies a condition function to each element in and returns
// a genericiface.GenericSlice2Iface where genericiface.GenericSlice2Iface[0] contains a genericiface.GenericSliceIface with all elements for
// whom the condition function returned true, and where genericiface.GenericSlice2Iface[1] contains a
// genericiface.GenericSliceIface with all elements for whom the condition function returned false.
//
// Partition is a special case of the Group function.
func (tt1 *SliceType) Partition(condition closures.ConditionFn) genericiface.GenericSlice2Iface {
	return unbox2(Partition(box(*tt1), condition))
}

// Permutable returns true if the number of permutations for tt1 exceeds
// MaxInt64.
func (tt1 *SliceType) Permutable() bool {
	return Permutable(*tt1)
}

// Permutations returns the number of permutations that exist given the current
// number of items in the tt1.
func (tt1 *SliceType) Permutations() *big.Int {
	return Permutations(*tt1)
}

// Permute returns a genericiface.GenericSlice2Iface which contains a genericiface.GenericSliceIface for each permutation
// of tt1.
//
// This function will panic if it determines that the list is not permutable
// (see Permutable function).
//
// Permute makes no assumptions about whether or not the elements in tt1 are
// distinct. Permutations are created positionally, and do not involve any
// equality checks. As such, if it important that Permute operate on a set of
// distinct elements, pass tt1 through one of the Distinct transforms before
// passing it to Permute().
//
// Permute is implemented using Heap's algorithm.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func (tt1 *SliceType) Permute() genericiface.GenericSlice2Iface {
	return (Permute(*tt1))
}

// Pop returns a genericiface.GenericSliceIface containing the head element from and removes the
// element from tt1. If tt1 is empty, the returned genericiface.GenericSliceIface will also be empty.
func (tt1 *SliceType) Pop() genericiface.GenericSliceIface {
	Pop(boxP(tt1))
	return tt1
}

// Push places a prepends a new element at the head of tt1.
func (tt1 *SliceType) Push(t1 interface{}) genericiface.GenericSliceIface {
	Push(boxP(tt1), t1)
	return tt1
}

// Reduce applies a reducer function to each element in threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new genericiface.GenericSliceIface. If tt1 is empty, the resulting genericiface.GenericSliceIface
// will also be empty.
func (tt1 *SliceType) Reduce(reducer func(a, acc interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(Reduce(box(*tt1), reducer))
}

// Remove applies a condition function to each item in the list, and removes any item
// for which the condition returns true.
func (tt1 *SliceType) Remove(condition closures.ConditionFn) genericiface.GenericSliceIface {
	Remove(boxP(tt1), condition)
	return tt1
}

// RemoveAt removes the item at the specified index from the slice.
// If len(tt1) == 0, tt1 == nil, i < 0, or i >= len(tt1), this function will do
// nothing.
func (tt1 *SliceType) RemoveAt(i int64) genericiface.GenericSliceIface {
	RemoveAt(boxP(tt1), i)
	return tt1
}

// Reverse reverses the order of tt1.
func (tt1 *SliceType) Reverse() genericiface.GenericSliceIface {
	Reverse(boxP(tt1))
	return tt1
}

// Skip removes the first n elements from tt1.
//
// Note that Skip(len(tt1)) will remove all items from the list, but does not
// "clear" the slice, meaning that the list remains allocated in memory.
// To fully de-pointer the slice, and ensure it is available for garbage
// collection as soon as possible, consider using Clear().
func (tt1 *SliceType) Skip(n int64) genericiface.GenericSliceIface {
	Skip(boxP(tt1), n)
	return tt1
}

// SkipWhile scans through tt1 starting at the head, and removes all
// elements from tt1 while the condition function returns true.
// SkipWhile stops removing any further items from tt1 after the first condition that
// returns false.
func (tt1 *SliceType) SkipWhile(condition closures.ConditionFn) genericiface.GenericSliceIface {
	SkipWhile(boxP(tt1), condition)
	return tt1
}

// Sort sorts using the supplied less function to determine order.
// Sort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func (tt1 *SliceType) Sort(less func(t1 interface{}, t2 interface{}) bool) genericiface.GenericSliceIface {
	Sort(boxP(tt1), less)
	return tt1
}

// SplitAfter finds the first element b for which a condition function returns true,
// and returns a genericiface.GenericSlice2Iface where genericiface.GenericSlice2Iface[0] contains the first half of tt1
// and genericiface.GenericSlice2Iface[1] contains the second half of tt1. Element b will be included
// in genericiface.GenericSlice2Iface[0]. If the no element can be found for which the condition returns
// true, genericiface.GenericSlice2Iface[0] will contain and genericiface.GenericSlice2Iface[1] will be empty.
func (tt1 *SliceType) SplitAfter(condition closures.ConditionFn) genericiface.GenericSlice2Iface {
	return unbox2(SplitAfter(box(*tt1), condition))
}

// SplitAt splits tt1 at index i, and returns a genericiface.GenericSlice2Iface which contains the
// two split halves of tt1. tt1[i] will be included in genericiface.GenericSlice2Iface[1].
// If i < 0, all of tt1 will be placed in genericiface.GenericSlice2Iface[0] and genericiface.GenericSlice2Iface[1] will
// be empty. Conversly, if i >= len(tt1), all of tt1 will be placed in
// genericiface.GenericSlice2Iface[1] and genericiface.GenericSlice2Iface[0] will be empty. If tt1 is nil or empty,
// genericiface.GenericSlice2Iface will contain two empty slices.
func (tt1 *SliceType) SplitAt(i int64) genericiface.GenericSlice2Iface {
	return unbox2(SplitAt(box(*tt1), i))
}

// SplitBefore finds the first element b for which a condition function returns true,
// and returns a genericiface.GenericSlice2Iface where genericiface.GenericSlice2Iface[0] contains the first half of tt1
// and genericiface.GenericSlice2Iface[1] contains the second half of tt1. Element b will be included
// in genericiface.GenericSlice2Iface[1]
func (tt1 *SliceType) SplitBefore(condition closures.ConditionFn) genericiface.GenericSlice2Iface {
	return unbox2(SplitBefore(box(*tt1), condition))
}

// String returns a string representation of suitable for use
// with fmt.Print, or other similar functions. String should be regarded as
// informational, and should not be relied upon to formally serialize a
// genericiface.GenericSliceIface.
func (tt1 *SliceType) String() string {
	return String(box(*tt1))
}

// SwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of SwapIndex does nothing.
func (tt1 *SliceType) SwapIndex(i, j int64) genericiface.GenericSliceIface {
	SwapIndex(box(*tt1), i, j)
	return tt1
}

// Tail removes the current head element from tt1.
// This equivelant to RemoveAt(0)
func (tt1 *SliceType) Tail() genericiface.GenericSliceIface {
	Tail(boxP(tt1))
	return tt1
}

// Take retains the first n elements of and removes all remaining elements
// from the slice. If n < 0 or n >= len(tt1), Take does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func (tt1 *SliceType) Take(n int64) genericiface.GenericSliceIface {
	Take(boxP(tt1), n)
	return tt1
}

// TakeWhile applies a condition function to each element in and retains all
// elements of tt1 so long as the condition function returns true. As soon as the condition
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func (tt1 *SliceType) TakeWhile(condition closures.ConditionFn) genericiface.GenericSliceIface {
	TakeWhile(boxP(tt1), condition)
	return tt1
}

// Union appends slice tt2 to slice tt1.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func (tt1 *SliceType) Union(tt2 *[]interface{}) genericiface.GenericSliceIface {
	Union(boxP(tt1), *tt2)
	return tt1
}

// Unzip splits tt1 into a genericiface.GenericSlice2Iface, such that genericiface.GenericSlice2Iface[0] contains all odd
// indices from and genericiface.GenericSlice2Iface[1] contains all even indices from tt1.
func (tt1 *SliceType) Unzip() genericiface.GenericSlice2Iface {
	return unbox2(Unzip(box(*tt1)))
}

// WindowCentered applies a windowing function across the using a centered
// window of the specified size.
func (tt1 *SliceType) WindowCentered(windowSize int64, windowFn func(window []interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(WindowCentered(box(*tt1), windowSize, windowFn))
}

// WindowLeft applies a windowing function across using a left-sided window
// of the specified size.
func (tt1 *SliceType) WindowLeft(windowSize int64, windowFn func(window []interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(WindowLeft(box(*tt1), windowSize, windowFn))
}

// WindowRight applies a windowing function across using a right-sided
// window of the specified size.
func (tt1 *SliceType) WindowRight(windowSize int64, windowFn func(window []interface{}) interface{}) genericiface.GenericSliceIface {
	return unbox(WindowRight(box(*tt1), windowSize, windowFn))
}

// Zip interleaves the contents of tt1 with tt2, and returns the result as a
// new genericiface.GenericSliceIface. tt1[0] is evaluated first. Thus if tt1 and tt2 are the same
// length, slice tt1 will occupy the odd indices of the result slice, and tt2
// will occupy the even indices of the result slice. If tt1 and tt2 are not
// the same length, Zip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func (tt1 *SliceType) Zip(tt2 *[]interface{}) genericiface.GenericSliceIface {
	return unbox(Zip(box(*tt1), *tt2))
}

var _ genericiface.GenericSliceIface = (genericiface.GenericSliceIface)(nil)
