package generic

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/jecolasurdo/transforms/pkg/slices/types"
)

// Continue instructs iterators about whether or not to keep iterating.
type Continue bool

const (
	// ContinueYes signals to an iterator that it should continue iterating.
	ContinueYes Continue = true

	// ContinueNo signals to an iterator that it should stop iterating.
	ContinueNo Continue = false
)

// All applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
func All(aa types.SliceType, test func(types.PrimitiveType) bool) bool {
	for _, s := range aa {
		if !test(s) {
			return false
		}
	}
	return true
}

// Any applies a test function to each element of the
// slice and returns true if the test function returns true for at least one
// item in the list.
//
// Any does not require that the source slice be sorted, and merely scans
// the slice, returning as soon as any element passes the supplied test. For
// a binary search, consider using sort.Search from the standard library.
func Any(aa types.SliceType, test func(types.PrimitiveType) bool) bool {
	for _, a := range aa {
		if test(a) {
			return true
		}
	}
	return false
}

//Append adds the supplied values to the end of the slice.
func Append(aa *types.SliceType, values ...types.PrimitiveType) {
	*aa = append(*aa, values...)
}

// Clear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func Clear(aa *types.SliceType) {
	*aa = nil
}

// Clone returns a copy of the slice.
func Clone(aa types.SliceType) types.SliceType {
	return append(types.SliceType{}, aa...)
}

// Collect applies a given function against each item in slice aa and
// each item of a slice bb, and returns the concatenation of each result.
//
//   Illustration:
//     aa:  		[A, B, C]
//     bb: 			[X, Y, Z]
//     collector:   func(a, b) { return a + b }
//     Collect(aa, bb, collector) -> [AX, AY, AZ, BX, BY, BZ, CX, XY, CZ]
func Collect(aa types.SliceType, bb types.SliceType, collector func(a, b types.PrimitiveType) types.PrimitiveType) types.SliceType {
	cc := types.SliceType{}
	for _, a := range aa {
		for _, b := range bb {
			cc = append(cc, collector(a, b))
		}
	}
	return cc
}

// Count applies the supplied test function to each element of the slice,
// and returns the count of items for which the test returns true.
func Count(aa types.SliceType, test func(types.PrimitiveType) bool) int64 {
	matches := int64(0)
	for _, a := range aa {
		if test(a) {
			matches++
		}
	}
	return matches
}

// Dequeue returns a types.SliceType containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func Dequeue(aa *types.SliceType) types.SliceType {
	if len(*aa) == 0 {
		return types.SliceType{}
	}
	head := (*aa)[0]
	RemoveAt(aa, 0)
	return types.SliceType{head}
}

// Difference returns a new slice that contains items that are not common
// between aa and bb. The supplied equal function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. Distinct values from aa are listed ahead of those from bb in the
// resulting slice.
//
// Illustration:
//   aa: [1,2,3,3,1,4]
//   bb: [5,4,3,5]
//   equal: func(a, b) bool {return a == b}
//   Difference(aa, bb, equal) -> [1,2,1,5,5]
func Difference(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) types.SliceType {
	ii := make([]bool, len(aa))
	jj := make([]bool, len(bb))
	for i, a := range aa {
		for j, b := range bb {
			if equal(a, b) {
				ii[i] = true
				jj[j] = true
			}
		}
	}

	cc := types.SliceType{}
	for i, a := range aa {
		if !ii[i] {
			cc = append(cc, a)
		}
	}

	for j, b := range bb {
		if !jj[j] {
			cc = append(cc, b)
		}
	}

	return cc
}

// Distinct removes all duplicates from the slice, using the supplied equal
// function to determine equality.
func Distinct(aa *types.SliceType, equal func(a, b types.PrimitiveType) bool) {
	bb := types.SliceType{}
	dups := make([]bool, len(*aa))
	for i, a := range *aa {
		if !dups[i] {
			bb = append(bb, a)
		}
		for j := i + 1; j < len(*aa); j++ {
			if equal(a, (*aa)[j]) {
				dups[j] = true
			}
		}
	}
	Clear(aa)
	Append(aa, bb...)
}

// Empty returns true if the length of the slice is zero.
func Empty(aa types.SliceType) bool {
	return len(aa) == 0
}

// End returns the a types.SliceType containing only the last element from aa.
func End(aa types.SliceType) types.SliceType {
	if Empty(aa) {
		return types.SliceType{}
	}
	return types.SliceType{aa[len(aa)-1]}
}

// Enqueue places an item at the head of the slice.
func Enqueue(aa *types.SliceType, a types.PrimitiveType) {
	*aa = append(*aa, a)
	copy((*aa)[1:], (*aa)[:len(*aa)-1])
	(*aa)[0] = a
}

// Expand applies an expansion function to each element of aa, and flattens
// the results into a single types.SliceType.
//
//   Illustration (pseudocode):
//     aa: [AB, CD, EF]
//     expansion: func(a string) []string { return []string{a[0], a[1]}}
//     Expand(aa, expansion) -> [A, B, C, D, E, F]
func Expand(aa types.SliceType, expansion func(types.PrimitiveType) types.SliceType) types.SliceType {
	bb := types.SliceType{}
	for _, a := range aa {
		Append(&bb, expansion(a)...)
	}
	return bb
}

// Filter removes all items from the slice for which the supplied test function
// returns true.
func Filter(aa *types.SliceType, test func(types.PrimitiveType) bool) {
	for i := len(*aa) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, int64(i))
		}
	}
}

// FindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func FindIndex(aa types.SliceType, test func(types.PrimitiveType) bool) int64 {
	for i, a := range aa {
		if test(a) {
			return int64(i)
		}
	}
	return -1
}

// First returns a types.SliceType containing the first element in the slice for which
// the supplied test function returns true.
func First(aa types.SliceType, test func(types.PrimitiveType) bool) types.SliceType {
	bb := types.SliceType{}
	for _, a := range aa {
		if test(a) {
			Append(&bb, a)
			break
		}
	}
	return bb
}

// Fold applies a function to each item in slice aa, threading an accumulator
// through each iteration. The accumulated value is returned in a new types.SliceType
// once aa is fully scanned. Fold returns a types.SliceType rather than a
// types.PrimitiveType to be consistent with this package's Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func Fold(aa types.SliceType, acc types.PrimitiveType, folder func(a, acc types.PrimitiveType) types.PrimitiveType) types.SliceType {
	return FoldI(aa, acc, func(_ int64, a, acc types.PrimitiveType) types.PrimitiveType { return folder(a, acc) })
}

// FoldI applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned. Foldi returns a types.SliceType rather than a
// types.PrimitiveType to be consistent with this package's Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func FoldI(aa types.SliceType, acc types.PrimitiveType, folder func(i int64, a, acc types.PrimitiveType) types.PrimitiveType) types.SliceType {
	accumulation := acc
	for i, a := range aa {
		accumulation = folder(int64(i), a, accumulation)
	}
	return types.SliceType{accumulation}
}

// ForEach applies each element of the list to the given function.
// ForEach will stop iterating if fn return false.
func ForEach(aa types.SliceType, fn func(types.PrimitiveType) Continue) {
	for _, a := range aa {
		if !fn(a) {
			return
		}
	}
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
// If any execution of fn returns ContinueNo, ForEachC will cease marshalling
// any backlogged work, and will immediately set the cancellation flag to true.
// Any goroutines monitoring the cancelPending closure can wind down their
// activities as necessary. ForEachC will continue to block until all active
// goroutines exit cleanly.
func ForEachC(aa types.SliceType, c int, fn func(a types.PrimitiveType, cancelPending func() bool) Continue) {
	if c < 0 {
		panic("ForEachC: The concurrency pool size (c) must be non-negative.")
	}
	mu := new(sync.RWMutex)
	halt := int64(0)
	cancelPending := func() bool {
		mu.RLock()
		defer mu.RUnlock()
		return halt > 0
	}
	sem := make(chan struct{}, c)
	defer close(sem)
	for _, a := range aa {
		mu.RLock()
		stop := halt > 0
		mu.RUnlock()
		if stop {
			break
		}
		sem <- struct{}{}
		go func(a types.PrimitiveType) {
			defer func() { <-sem }()
			if !fn(a, cancelPending) {
				mu.Lock()
				halt++
				mu.Unlock()
			}
		}(a)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}

// ForEachR applies each element of aa to a given function, scanning
// through the slice in reverse order, starting from the end and working towards
// the head.
func ForEachR(aa types.SliceType, fn func(types.PrimitiveType) Continue) {
	for i := len(aa) - 1; i >= 0; i-- {
		if !fn(aa[i]) {
			return
		}
	}
}

// Group consolidates like-items into groups according to the supplied grouper
// function, and returns them as a types.SliceType2.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed.
func Group(aa types.SliceType, grouper func(types.PrimitiveType) int64) types.SliceType2 {
	return GroupI(aa, func(_ int64, a types.PrimitiveType) int64 { return grouper(a) })
}

// GroupI consolidates like-items into groups according to the supplied grouper
// function, and returns them as a types.SliceType2.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from aa is also passed into the grouper function.
func GroupI(aa types.SliceType, grouper func(int64, types.PrimitiveType) int64) types.SliceType2 {
	groupMap := map[int64]types.SliceType{}
	for i, a := range aa {
		hash := grouper(int64(i), a)
		if _, exists := groupMap[hash]; exists {
			groupMap[hash] = append(groupMap[hash], a)
		} else {
			groupMap[hash] = types.SliceType{a}
		}
	}
	group := types.SliceType2{}
	for _, bb := range groupMap {
		group = append(group, bb)
	}
	return group
}

// Head returns a types.SliceType containing the first item from the aa. If aa is
// empty, the resulting types.SliceType will be empty.
func Head(aa types.SliceType) types.SliceType {
	if Empty(aa) {
		return types.SliceType{}
	}
	return types.SliceType{aa[0]}
}

// InsertAfter inserts an element in aa after the first element for which the
// supplied test function returns true. If none of the tests return true, the
// element is appended to the end of the aa.
func InsertAfter(aa *types.SliceType, b types.PrimitiveType, test func(types.PrimitiveType) bool) {
	var i int
	var a types.PrimitiveType
	for i, a = range *aa {
		if test(a) {
			break
		}
	}
	InsertAt(aa, b, int64(i+1))
}

// InsertBefore inserts an element in aa before the first element for which the
// supplied test function returns true. If none of the tests return true,
// the element is inserted at the head of aa.
func InsertBefore(aa *types.SliceType, b types.PrimitiveType, test func(types.PrimitiveType) bool) {
	var i int
	var a types.PrimitiveType
	for i, a = range *aa {
		if test(a) {
			break
		}
	}
	InsertAt(aa, b, int64(i-1))
}

// InsertAt inserts an element in aa at the specified index i, shifting the
// element originally at index i (and all subsequent elements) one position
// to the right. If i < 0, the element is inserted at index 0. If
// i >= len(aa), the value is appended to the end of aa.
func InsertAt(aa *types.SliceType, a types.PrimitiveType, i int64) {
	*aa = append(*aa, a)
	if i >= int64(len(*aa)) {
		return
	}
	if i < 0 {
		i = 0
	}
	copy((*aa)[i+1:], (*aa)[i:])
	(*aa)[i] = a
}

// Intersection compares each element of aa to bb using the supplied equal
// function, and returns a types.SliceType containing the elements which are common
// to both aa and bb. Duplicates are removed in this operation.
func Intersection(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) types.SliceType {
	cc := types.SliceType{}
	ForEach(aa, func(a types.PrimitiveType) Continue {
		ForEach(bb, func(b types.PrimitiveType) Continue {
			if equal(a, b) && !Any(cc, func(c types.PrimitiveType) bool { return equal(a, c) }) {
				Append(&cc, a)
			}
			return ContinueYes
		})
		return ContinueYes
	})
	return cc
}

// IsProperSubset returns true if aa is a proper subset of bb.
// aa is considered a proper subset if all of its elements exist within bb, but
// bb also contains some elements that do not exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsProperSubset(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) == 0 && len(bb1) > 0
}

// IsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsProperSuperset(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) > 0 && len(bb1) == 0
}

// IsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsSubset(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) == 0 && len(bb1) >= 0
}

// IsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsSuperset(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) >= 0 && len(bb1) == 0
}

func removeIntersections(aa, bb types.SliceType, equal func(a, b types.PrimitiveType) bool) (types.SliceType, types.SliceType) {
	aa1 := Clone(aa)
	bb1 := Clone(bb)
	for ai := int64(len(aa1)) - 1; ai >= 0; ai-- {
		intersectionFound := false
		for bi := int64(len(bb1)) - 1; bi >= 0; bi-- {
			if equal((aa1)[ai], (bb1)[bi]) {
				intersectionFound = true
				RemoveAt(&bb1, bi)
			}
		}
		if intersectionFound {
			RemoveAt(&aa1, ai)
		}
	}
	return aa1, bb1
}

// Item returns a types.SliceType containing the element at aa[i].
// If len(aa) == 0, i < 0, or, i >= len(aa), the resulting slice will be empty.
func Item(aa types.SliceType, i int64) types.SliceType {
	if Empty(aa) || i < 0 || i >= int64(len(aa)) {
		return types.SliceType{}
	}
	return types.SliceType{aa[i]}
}

// ItemFuzzy returns a types.SliceType containing the element at aa[i].
// If the supplied index is outside of the bounds of aa, ItemFuzzy will attempt
// to retrieve the head or end element of aa according to the following rules:
// If len(aa) == 0 an empty types.SliceType is returned.
// If i < 0, the head of aa is returned.
// If i >= len(aa), the end of the aa is returned.
func ItemFuzzy(aa types.SliceType, i int64) types.SliceType {
	if Empty(aa) {
		return types.SliceType{}
	}
	if i < 0 {
		return Head(aa)
	}
	if i >= int64(len(aa)) {
		return End(aa)
	}
	return types.SliceType{aa[i]}
}

// Last applies a test function to each element in aa, and returns a types.SliceType
// containing the last element for which the test returned true. If no elements
// pass the supplied test, the resulting types.SliceType will be empty.
func Last(aa types.SliceType, test func(types.PrimitiveType) bool) types.SliceType {
	bb := types.SliceType{}
	ForEachR(aa, func(a types.PrimitiveType) Continue {
		if test(a) {
			Append(&bb, a)
			return ContinueNo
		}
		return ContinueYes
	})
	return bb
}

// Len returns the length of aa.
func Len(aa types.SliceType) int {
	return len(aa)
}

// Map applies a tranform to each element of the list.
func Map(aa *types.SliceType, mapFn func(types.PrimitiveType) types.PrimitiveType) {
	for i, a := range *aa {
		(*aa)[i] = mapFn(a)
	}
}

// None applies a test function to each element in aa, and returns true if
// the test function returns false for all items.
func None(aa types.SliceType, test func(types.PrimitiveType) bool) bool {
	return !Any(aa, test)
}

// Pairwise threads a transform function through aa, passing to the transform
// successive two-element pairs, aa[i-1] && aa[i]. For the first pairing
// the supplied init value is supplied as the initial element in the pair.
//
//   Illustration (pseudocode):
//     aa:  [W,X,Y,Z]
//     xform: func(a, b string) string { return a + b }
//     init: V
//     Pairwise(aa, init, xform) -> [VW, WX, XY, YZ]
func Pairwise(aa types.SliceType, init types.PrimitiveType, xform func(a, b types.PrimitiveType) types.PrimitiveType) types.SliceType {
	bb := types.SliceType{}
	i := 0
	a1, a2 := init, aa[i]
	for {
		bb = append(bb, xform(a1, a2))
		i++
		if i >= len(aa) {
			break
		}
		a1, a2 = aa[i-1], aa[i]
	}
	return bb
}

// Partition applies a test function to each element in aa, and returns
// a types.SliceType2 where types.SliceType2[0] contains a types.SliceType with all elements for
// whom the test function returned true, and where types.SliceType2[1] contains a
// types.SliceType with all elements for whom the test function returned false.
//
// Partition is a special case of the Group function.
func Partition(aa types.SliceType, test func(types.PrimitiveType) bool) types.SliceType2 {
	grouper := func(a types.PrimitiveType) int64 {
		if test(a) {
			return 1
		}
		return 0
	}
	return Group(aa, grouper)
}

// Permutable returns true if the number of permutations for aa exceeds
// MaxInt64.
func Permutable(aa types.SliceType) bool {
	return Permutations(aa).IsInt64()
}

// Permutations returns the number of permutations that exist given the current
// number of items in the aa.
func Permutations(aa types.SliceType) *big.Int {
	var f big.Int
	return f.MulRange(1, int64(len(aa)))
}

// Permute returns a types.SliceType2 which contains a types.SliceType for each permutation
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
func Permute(aa types.SliceType) types.SliceType2 {
	if !Permutable(aa) {
		panic(fmt.Sprintf("The number of permutations for this list (%v) exceeeds MaxInt64.", Permutations(aa)))
	}

	acc := types.SliceType2{}
	generate(int64(len(aa)), aa, &acc)
	return acc
}

func generate(n int64, aa types.SliceType, acc *types.SliceType2) {
	if n == 1 {
		*acc = append(*acc, aa)
		return
	}

	for i := int64(0); i < n-1; i++ {
		generate(n-1, aa, acc)
		aa = Clone(aa)
		if n%2 != 0 {
			SwapIndex(aa, i, n-1)
		} else {
			SwapIndex(aa, 0, n-1)
		}
	}

	generate(n-1, aa, acc)
}

// Pop returns a types.SliceType containing the head element from aa, and removes the
// element from aa. If aa is empty, the returned types.SliceType will also be empty.
func Pop(aa *types.SliceType) types.SliceType {
	bb := Head(*aa)
	RemoveAt(aa, 0)
	return bb
}

// Push places a prepends a new element at the head of aa.
func Push(aa *types.SliceType, a types.PrimitiveType) {
	InsertAt(aa, a, 0)
}

// Reduce applies a reducer function to each element in aa, threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new types.SliceType. If aa is empty, the resulting types.SliceType
// will also be empty.
//
//  Illustration:
//    aa: [1,2,3,4]
//    reducer: acc + sourceNode
//    Fold(aa, reducer) -> [10]
func Reduce(aa types.SliceType, reducer func(a, acc types.PrimitiveType) types.PrimitiveType) types.SliceType {
	if len(aa) == 0 {
		return types.SliceType{}
	}
	accumulator := aa[0]
	if len(aa) > 1 {
		for i := 1; i < len(aa); i++ {
			accumulator = reducer(aa[i], accumulator)
		}
	}
	return types.SliceType{accumulator}
}

// Remove applies a test function to each item in the list, and removes any item
// for which the test returns true.
func Remove(aa *types.SliceType, test func(types.PrimitiveType) bool) {
	for i := int64(len(*aa)) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, i)
		}
	}
}

// RemoveAt removes the item at the specified index from the slice.
// If len(aa) == 0, aa == nil, i < 0, or i >= len(aa), this function will do
// nothing.
func RemoveAt(aa *types.SliceType, i int64) {
	if i < 0 || i >= int64(len(*aa)) {
		return
	}
	if len(*aa) > 0 {
		*aa = append((*aa)[:i], (*aa)[i+1:]...)
	}
}

// Reverse reverses the order of aa.
func Reverse(aa *types.SliceType) {
	for i := len(*aa)/2 - 1; i >= 0; i-- {
		j := len(*aa) - 1 - i
		(*aa)[i], (*aa)[j] = (*aa)[j], (*aa)[i]
	}
}

// Skip removes the first n elements from aa.
//
// Note that Skip(aa, len(aa)) will remove all items from the list, but does not
// "clear" the slice, meaning that the list remains allocated in memory.
// To fully de-pointer the slice, and ensure it is available for garbage
// collection as soon as possible, consider using Clear().
func Skip(aa *types.SliceType, n int64) {
	if len(*aa) == 0 {
		return
	}
	*aa = (*aa)[n:]
}

// SkipWhile scans through aa starting at the head, and removes all
// elements from aa while the test function returns true.
// SkipWhile stops removing any further items from aa after the first test that
// returns false.
func SkipWhile(aa *types.SliceType, test func(types.PrimitiveType) bool) {
	// find the first index where the test would evaluate to false and skip
	// everything up to that index.
	findTest := func(a types.PrimitiveType) bool { return !test(a) }
	Skip(aa, FindIndex(*aa, findTest))
}

// Sort sorts aa, using the supplied less function to determine order.
// Sort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func Sort(aa *types.SliceType, less func(a, b types.PrimitiveType) bool) {
	lessI := func(i, j int) bool {
		return less((*aa)[i], (*aa)[j])
	}
	sort.SliceStable(*aa, lessI)
}

// SplitAfter finds the first element b for which a test function returns true,
// and returns a types.SliceType2 where types.SliceType2[0] contains the first half of aa
// and types.SliceType2[1] contains the second half of aa. Element b will be included
// in types.SliceType2[0]. If the no element can be found for which the test returns
// true, types.SliceType2[0] will contain aa, and types.SliceType2[1] will be empty.
func SplitAfter(aa types.SliceType, test func(types.PrimitiveType) bool) types.SliceType2 {
	return SplitAt(aa, FindIndex(aa, test)+1)
}

// SplitAt splits aa at index i, and returns a types.SliceType2 which contains the
// two split halves of aa. aa[i] will be included in types.SliceType2[1].
// If i < 0, all of aa will be placed in types.SliceType2[0] and types.SliceType2[1] will
// be empty. Conversly, if i >= len(aa), all of aa will be placed in
// types.SliceType2[1] and types.SliceType2[0] will be empty. If aa is nil or empty,
// types.SliceType2 will contain two empty slices.
func SplitAt(aa types.SliceType, i int64) types.SliceType2 {
	if len(aa) == 0 {
		return types.SliceType2{
			types.SliceType{},
			types.SliceType{},
		}
	}
	if i < 0 {
		i = 0
	}
	return types.SliceType2{
		aa[:i],
		aa[i:],
	}
}

// SplitBefore finds the first element b for which a test function returns true,
// and returns a types.SliceType2 where types.SliceType2[0] contains the first half of aa
// and types.SliceType2[1] contains the second half of aa. Element b will be included
// in types.SliceType2[1]
func SplitBefore(aa types.SliceType, test func(types.PrimitiveType) bool) types.SliceType2 {
	return SplitAt(aa, FindIndex(aa, test))
}

// String returns a string representation of aa, suitable for use
// with fmt.Print, or other similar functions. String should be regarded as
// informational, and should not be relied upon to formally serialize a
// types.SliceType.
func String(aa types.SliceType) string {
	jsonBytes, _ := json.Marshal(aa)
	return string(jsonBytes)
}

// SwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of aa, SwapIndex does nothing.
func SwapIndex(aa types.SliceType, i, j int64) {
	l := int64(len(aa))
	if i < 0 || j < 0 || i >= l || j >= l {
		return
	}
	aa[i], aa[j] = aa[j], aa[i]
}

// Tail removes the current head element from aa.
// This equivelant to RemoveAt(aa, 0)
func Tail(aa *types.SliceType) {
	RemoveAt(aa, 0)
}

// Take retains the first n elements of aa, and removes all remaining elements
// from the slice. If n < 0 or n >= len(aa), Take does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func Take(aa *types.SliceType, n int64) {
	if len(*aa) == 0 || n < 0 || n >= int64(len(*aa)) {
		return
	}
	*aa = (*aa)[:n]
}

// TakeWhile applies a test function to each element in aa, and retains all
// elements of aa so long as the test function returns true. As soon as the test
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func TakeWhile(aa *types.SliceType, test func(types.PrimitiveType) bool) {
	find := func(a types.PrimitiveType) bool {
		return !test(a)
	}
	Take(aa, FindIndex(*aa, find))
}

// Union appends slice bb to slice aa.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func Union(aa *types.SliceType, bb types.SliceType) {
	Append(aa, bb...)
}

// Unzip splits aa into a types.SliceType2, such that types.SliceType2[0] contains all odd
// indices from aa, and types.SliceType2[1] contains all even indices from aa.
func Unzip(aa types.SliceType) types.SliceType2 {
	odds := types.SliceType{}
	evens := types.SliceType{}
	for i, a := range aa {
		if i%2 != 0 {
			odds = append(odds, a)
		} else {
			evens = append(evens, a)
		}
	}
	return types.SliceType2{odds, evens}
}

// WindowCentered applies a windowing function across the aa, using a centered
// window of the specified size.
func WindowCentered(aa types.SliceType, windowSize int64, windowFn func(window types.SliceType) types.PrimitiveType) types.SliceType {
	cc := types.SliceType{}
	fullWindowReached := false
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := types.SliceType{}
		a := aa[i]
		for n := int64(1); n <= windowSize; n++ {
			Append(&currentWindow, a)
			if !fullWindowReached && n >= windowSize {
				fullWindowReached = true
			}
			if !fullWindowReached {
				Append(&cc, windowFn(currentWindow))
			}
			if i+n >= int64(len(aa)) {
				break
			}
			a = aa[i+n]
		}
		Append(&cc, windowFn(currentWindow))
	}
	trimSize := windowSize - 1
	var frontTrim, backTrim int64
	if trimSize%2 == 0 {
		frontTrim = trimSize / 2
		backTrim = frontTrim
	} else {
		frontTrim = trimSize / 2
		backTrim = frontTrim + 1
	}
	dd := types.SliceType(SplitAt(cc, frontTrim)[1])
	Reverse(&dd)
	ee := types.SliceType(SplitAt(dd, backTrim)[1])
	Reverse(&ee)
	return ee
}

// WindowLeft applies a windowing function across aa, using a left-sided window
// of the specified size.
func WindowLeft(aa types.SliceType, windowSize int64, windowFn func(window types.SliceType) types.PrimitiveType) types.SliceType {
	bb := types.SliceType{}
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := types.SliceType{}
		for n := int64(0); n < windowSize; n++ {
			if i+n >= int64(len(aa)) {
				break
			}
			Append(&currentWindow, aa[i+n])
		}
		Append(&bb, windowFn(currentWindow))
	}
	return bb
}

// WindowRight applies a windowing function across aa, using a right-sided
// window of the specified size.
func WindowRight(aa types.SliceType, windowSize int64, windowFn func(window types.SliceType) types.PrimitiveType) types.SliceType {
	aa1 := Clone(aa)
	defer Clear(&aa1)

	Reverse(&aa1)
	bb := types.SliceType{}
	for i := int64(0); i < int64(len(aa1)); i++ {
		currentWindow := types.SliceType{}
		for n := int64(0); n < windowSize; n++ {
			if i+n >= int64(len(aa1)) {
				break
			}
			Append(&currentWindow, aa1[i+n])
		}
		Reverse(&currentWindow)
		Append(&bb, windowFn(currentWindow))
	}
	Reverse(&bb)
	return bb
}

// Zip interleaves the contents of aa with bb, and returns the result as a
// new types.SliceType. aa[0] is evaluated first. Thus if aa and bb are the same
// length, slice aa will occupy the odd indices of the result slice, and bb
// will occupy the even indices of the result slice. If aa and bb are not
// the same length, Zip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func Zip(aa, bb types.SliceType) types.SliceType {
	if len(aa) == 0 {
		return bb
	}
	if len(bb) == 0 {
		return aa
	}

	cc := types.SliceType{}
	aaEndReached, bbEndReached := false, false
	for i := 0; aaEndReached == false && bbEndReached == false; i++ {
		if i >= len(aa) {
			aaEndReached = true
		}
		if i >= len(bb) {
			bbEndReached = true
		}
		if i%2 != 0 {
			if !aaEndReached {
				Append(&cc, aa[i])
			}
			if !bbEndReached {
				Append(&cc, bb[i])
			}
		} else {
			if !bbEndReached {
				Append(&cc, bb[i])
			}
			if !aaEndReached {
				Append(&cc, aa[i])
			}
		}
	}
	return cc
}
