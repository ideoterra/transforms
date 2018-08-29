// Package generic contains transform functions for SliceTypes.
//
// Function Naming Conventions:
// Often the same conceptual function can be implemented in more than one way.
// When multiple variants of a function are implemented in this package, each
// variant will start with the same base name, and the name will be suffixed
// with 0 or more flags that indicate details about the implementation. Such
// suffixes are as follows:
//
//   C: Functions with this suffix use concurrent operations internally, and
//		will typically require that a concurrency pool size be specified as an
//		argument. C functions guarantee that they will only use up to their
//		alotted number of concurrent goroutines for each invocation.
//
//   I: Functions with this suffix will have an index value threaded
// 		through each call to their applicable closure function.
//
//   R: Functions with this suffix operate against the underlaying slice
//		in reverse order (without incurring the penalty of actually
//		reversing the ordering of elements in the slice).
//
//   S: Functions with this suffix are optimized for an ordinal dataset
//		that has been presorted. Any time data is ordinal in nature, and
//		can be pre-sorted, there is typically a significant performance
//		advantage to using S variants.
//
// Parameter Naming Conventions:
// By convention, the source slice will be named `aa`. If multiple slices are
// to be supplied as arguments to a function, they are named `aa`, `bb`, `cc`,
// and so on.
//
// Mutability:
// The functions in this package will generally aim to mutate the underlaying
// slice unless doing so doesn't make sense given the nature of the operation
// or doing so would lead to confusion about which values are being mutated.
// Transforms that mutate the supplied sources will always require the slice
// as a pointer. Thus, if a function requires that a slice be passed
// as a pointer, it can be expected that the function is mutating the
// underlaying slice. Conversly, functions that require slices be passed as
// values, can be expected to be immutable.
//
// Null result handling:
// Transforms that reduce a result to a single value (such as Dequeue, or Fold)
// return a SliceType containing a single element rather than return a
// PrimitiveType. This is done to avoid edge cases associated with
// transformations that result in an null value. There are generally three
// options for how to handle an empty result. 1) Have the result be a pointer
// rather than a 'value', and set the result to nil if there is no result.
// 2) If there is no result, return a zero value. 3) Return the result as
// an empty slice if there is no result. Option 1 (returning nil) can be
// confusing if the underlaying slice contains pointers such as `[]*struct{}`.
// In such a case, it would be difficult to differentiate between a nil that is
// returned because the slice's head contained a nil pointer vs the slice being
// initially empty. A similar issue exists for option 2 (return a zero-value).
// Returning a zero value has a different implication from returning no value.
// As such, it makes sense to just return a slice in all circumstances. If the
// resulting slice is empty, we know there was no result returned, and confusion
// is avoided.
//
// Ordinality:
// Unless otherwise noted, the algorithms implemented for the transformations
// make no assumptions about the order or orderability of the dataset. For
// instance, the algorithms for `Any()` and `Difference()` are intentionally
// naive, so as to make as few assumptions about the nature of the data as
// possible. In cases where performance can be improved using a sorted dataset
// alternative functions are provided, such as in `AnyS()`, and
// `DifferenceS())`. By convention all functions that expect the inbound data
// to be sorted are suffixed with an 'S'. Such methods assume the data is sorted
// and may return unexpected results if the data is not, in fact, sorted.
//
// Equality (and Less) functions:
// Transforms that need to test the equality of slice elements are intentionaly
// left naive, and do not make any assumptions about how to test for equality.
// As a result, functions such as `Difference()` require an equality function
// to be supplied. For primitive types, typical equality (and `Less()`)
// functions are provided in the `eq` and `less` packages. It is encouraged to
// use the supplied equality (and less) functions for primitive types.
//
// Slice, Set, Stack, and other "non-native" Slice Functions:
// This package provides functions independent of the underlaying data
// structure. That is, even though the underlaying data type for all operations
// in this package is a slice, operations for other common data structures are
// provided for convenience. For example, several common set operations such as
// Intersection() are provided, even though this package operates on slices.
// While the implementations of these "non-native" functions are provided for
// convinience, that can come at the cost of performance because the underlaying
// datastructure simply may not be optimized for the operation being performed.
// With that said, efforts are made to provide reasonably performant
// algorythmic implementations for non-native operations. Other admissions are
// made to accomodate for the application of non-native functions to a slice.
// For instance, Sets, by definition, do not contain duplicates. However, there
// is no such requirement in a slice. As such, the set-type operations provided
// in this package are tollerant to duplicates.
package generic

import (
	"fmt"
	"math/big"
	"sync/atomic"

	"github.com/cheekybits/genny/generic"
)

// SliceType2 is a two dimensional slice of PrimitiveType
type SliceType2 SliceType

// SliceType is a one dimensional slice of PrimitiveType.
type SliceType []PrimitiveType

// PrimitiveType is a placeholder for the type underpinning the generic SliceType.
type PrimitiveType generic.Type

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
func All(aa SliceType, test func(PrimitiveType) bool) bool {
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
func Any(aa SliceType, test func(PrimitiveType) bool) bool {
	for _, a := range aa {
		if test(a) {
			return true
		}
	}
	return false
}

//Append adds the supplied values to the end of the slice.
func Append(aa *SliceType, values ...PrimitiveType) {
	*aa = append(*aa, values...)
}

// Clear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func Clear(aa *SliceType) {
	*aa = nil
}

// Clone returns a copy of the slice.
func Clone(aa SliceType) SliceType {
	return append(SliceType{}, aa...)
}

// Collect applies a given function against each item in slice aa and
// each item of a slice bb, and returns the concatenation of each result.
//
//   Illustration:
//     aa:  		[A, B, C]
//     bb: 			[X, Y, Z]
//     collector:   func(a, b) { return a + b }
//     Collect(aa, bb, collector) -> [AX, AY, AZ, BX, BY, BZ, CX, XY, CZ]
func Collect(aa SliceType, bb SliceType, collector func(a, b PrimitiveType) PrimitiveType) SliceType {
	cc := SliceType{}
	for _, a := range aa {
		for _, b := range bb {
			cc = append(cc, collector(a, b))
		}
	}
	return cc
}

// Count applies the supplied test function to each element of the slice,
// and returns the count of items for which the test returns true.
func Count(aa SliceType, test func(PrimitiveType) bool) int64 {
	matches := int64(0)
	for _, a := range aa {
		if test(a) {
			matches++
		}
	}
	return matches
}

// Dequeue returns a SliceType containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func Dequeue(aa *SliceType) SliceType {
	if len(*aa) == 0 {
		return SliceType{}
	}
	head := (*aa)[0]
	RemoveAt(aa, 0)
	return SliceType{head}
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
func Difference(aa, bb SliceType, equal func(a, b PrimitiveType) bool) SliceType {
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

	cc := SliceType{}
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
func Distinct(aa *SliceType, equal func(a, b PrimitiveType) bool) {
	bb := SliceType{}
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
func Empty(aa SliceType) bool {
	return len(aa) == 0
}

// End returns the a SliceType containing only the last element from aa.
func End(aa SliceType) SliceType {
	if Empty(aa) {
		return SliceType{}
	}
	return SliceType{aa[len(aa)-1]}
}

// Enqueue places an item at the head of the slice.
func Enqueue(aa *SliceType, a PrimitiveType) {
	*aa = append(*aa, a)
	copy((*aa)[1:], (*aa)[:len(*aa)-1])
	(*aa)[0] = a
}

// Expand applies an expansion function to each element of aa, and flattens
// the results into a single SliceType.
//
//   Illustration (pseudocode):
//     aa: [AB, CD, EF]
//     expansion: func(a string) []string { return []string{a[0], a[1]}}
//     Expand(aa, expansion) -> [A, B, C, D, E, F]
func Expand(aa SliceType, expansion func(PrimitiveType) SliceType) SliceType {
	bb := SliceType{}
	for _, a := range aa {
		Append(&bb, expansion(a)...)
	}
	return bb
}

// Filter removes all items from the slice for which the supplied test function
// returns true.
func Filter(aa *SliceType, test func(PrimitiveType) bool) {
	for i := len(*aa) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, int64(i))
		}
	}
}

// FindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func FindIndex(aa SliceType, test func(PrimitiveType) bool) int64 {
	for i, a := range aa {
		if test(a) {
			return int64(i)
		}
	}
	return -1
}

// First returns a SliceType containing the first element in the slice for which
// the supplied test function returns true.
func First(aa SliceType, test func(PrimitiveType) bool) SliceType {
	bb := SliceType{}
	for _, a := range aa {
		if test(a) {
			Append(&bb, a)
			break
		}
	}
	return bb
}

// Fold applies a function to each item in slice aa, threading an accumulator
// through each iteration. The accumulated value is returned in a new SliceType
// once aa is fully scanned. Fold returns a SliceType rather than a
// PrimitiveType to be consistent with this package's Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func Fold(aa SliceType, acc PrimitiveType, folder func(a, acc PrimitiveType) PrimitiveType) SliceType {
	return FoldI(aa, acc, func(_ int64, a, acc PrimitiveType) PrimitiveType { return folder(a, acc) })
}

// FoldI applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned. Foldi returns a SliceType rather than a
// PrimitiveType to be consistent with this package's Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func FoldI(aa SliceType, acc PrimitiveType, folder func(i int64, a, acc PrimitiveType) PrimitiveType) SliceType {
	accumulation := acc
	for i, a := range aa {
		accumulation = folder(int64(i), a, accumulation)
	}
	return SliceType{accumulation}
}

// ForEach applies each element of the list to the given function.
// ForEach will stop iterating if fn return false.
func ForEach(aa SliceType, fn func(PrimitiveType) Continue) {
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
func ForEachC(aa SliceType, c int, fn func(a PrimitiveType, cancelPending func() bool) Continue) {
	if c < 0 {
		panic("ForEachC: The concurrency pool size (c) must be non-negative.")
	}
	halt := int64(0)
	cancelPending := func() bool {
		return atomic.LoadInt64(&halt) > 0
	}
	sem := make(chan struct{}, c)
	defer close(sem)
	for _, a := range aa {
		if halt > 0 {
			break
		}
		sem <- struct{}{}
		go func(a PrimitiveType) {
			defer func() { <-sem }()
			if !fn(a, cancelPending) {
				atomic.AddInt64(&halt, 1)
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
func ForEachR(aa SliceType, fn func(PrimitiveType) Continue) {
	for i := len(aa) - 1; i >= 0; i-- {
		if !fn(aa[i]) {
			return
		}
	}
}

// Group consolidates like-items into groups according to the supplied grouper
// function, and returns them as a SliceType2.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed.
func Group(aa SliceType, grouper func(PrimitiveType) int64) SliceType2 {
	return GroupI(aa, func(_ int64, a PrimitiveType) int64 { return grouper(a) })
}

// GroupI consolidates like-items into groups according to the supplied grouper
// function, and returns them as a SliceType2.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from aa is also passed into the grouper function.
func GroupI(aa SliceType, grouper func(int64, PrimitiveType) int64) SliceType2 {
	groupMap := map[int64]SliceType{}
	for i, a := range aa {
		hash := grouper(int64(i), a)
		if _, exists := groupMap[hash]; exists {
			groupMap[hash] = append(groupMap[hash], a)
		} else {
			groupMap[hash] = SliceType{a}
		}
	}
	group := SliceType2{}
	for _, bb := range groupMap {
		group = append(group, bb)
	}
	return group
}

// Head returns a SliceType containing the first item from the aa. If aa is
// empty, the resulting SliceType will be empty.
func Head(aa SliceType) SliceType {
	if Empty(aa) {
		return SliceType{}
	}
	return SliceType{aa[0]}
}

// InsertAfter inserts an element in aa after the first element for which the
// supplied test function returns true. If none of the tests return true, the
// element is appended to the end of the aa.
func InsertAfter(aa *SliceType, b PrimitiveType, test func(PrimitiveType) bool) {
	var i int
	var a PrimitiveType
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
func InsertBefore(aa *SliceType, b PrimitiveType, test func(PrimitiveType) bool) {
	var i int
	var a PrimitiveType
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
func InsertAt(aa *SliceType, a PrimitiveType, i int64) {
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
// function, and returns a SliceType containing the elements which are common
// to both aa and bb. Duplicates are removed in this process.
func Intersection(aa, bb SliceType, equal func(a, b PrimitiveType) bool) SliceType {
	cc := SliceType{}
	ForEach(aa, func(a PrimitiveType) Continue {
		ForEach(bb, func(b PrimitiveType) Continue {
			if equal(a, b) && !Any(cc, func(c PrimitiveType) bool { return equal(a, c) }) {
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
func IsProperSubset(aa, bb SliceType, equal func(a, b PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) == 0 && len(bb1) > 0
}

// IsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
func IsProperSuperset(aa, bb SliceType, equal func(a, b PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) > 0 && len(bb1) == 0
}

// IsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
func IsSubset(aa, bb SliceType, equal func(a, b PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) == 0 && len(bb1) >= 0
}

// IsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
func IsSuperset(aa, bb SliceType, equal func(a, b PrimitiveType) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equal)
	return len(aa1) >= 0 && len(bb1) == 0
}

func removeIntersections(aa, bb SliceType, equal func(a, b PrimitiveType) bool) (SliceType, SliceType) {
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

// Item returns a SliceType containing the element at aa[i].
// If len(aa) == 0, i < 0, or, i >= len(aa), the resulting slice will be empty.
func Item(aa SliceType, i int64) SliceType {
	if Empty(aa) || i < 0 || i >= int64(len(aa)) {
		return SliceType{}
	}
	return SliceType{aa[i]}
}

// ItemFuzzy returns a SliceType containing the element at aa[i].
// If the supplied index is outside of the bounds of aa, ItemFuzzy will attempt
// to retrieve the head or end element of aa according to the following rules:
// If len(aa) == 0 an empty SliceType is returned.
// If i < 0, the head of aa is returned.
// If i >= len(aa), the end of the aa is returned.
func ItemFuzzy(aa SliceType, i int64) SliceType {
	if Empty(aa) {
		return SliceType{}
	}
	if i < 0 {
		return Head(aa)
	}
	if i >= int64(len(aa)) {
		return End(aa)
	}
	return SliceType{aa[i]}
}

// Last applies a test function to each element in aa, and returns a SliceType
// containing the last element for which the test returned true. If no elements
// pass the supplied test, the resulting SliceType will be empty.
func Last(aa SliceType, test func(PrimitiveType) bool) SliceType {
	bb := SliceType{}
	ForEachR(aa, func(a PrimitiveType) Continue {
		if test(a) {
			Append(&bb, a)
			return ContinueNo
		}
		return ContinueYes
	})
	return bb
}

// Len returns the length of aa.
func Len(aa SliceType) int {
	return len(aa)
}

// Map applies a tranform to each element of the list.
func Map(aa SliceType, mapFn func(PrimitiveType) PrimitiveType) {
	for i, a := range aa {
		aa[i] = mapFn(a)
	}
}

// None applies a test function to each element in aa, and returns true if
// the test function returns false for all items.
func None(aa SliceType, test func(PrimitiveType) bool) bool {
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
func Pairwise(aa SliceType, init PrimitiveType, xform func(a, b PrimitiveType) PrimitiveType) SliceType {
	bb := SliceType{}
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
// a SliceType2 where SliceType2[0] contains a SliceType with all elements for
// whom the test function returned true, and where SliceType2[1] contains a
// SliceType with all elements for whom the test function returned false.
//
// Partition is a special case of the Group function.
func Partition(aa SliceType, test func(PrimitiveType) bool) SliceType2 {
	grouper := func(a PrimitiveType) int64 {
		if test(a) {
			return 1
		}
		return 0
	}
	return Group(aa, grouper)
}

// Permutable returns true if the number of permutations for aa exceeds
// MaxInt64.
func Permutable(aa SliceType) bool {
	return Permutations(aa).IsInt64()
}

// Permutations returns the number of permutations that exist given the current
// number of items in the aa.
func Permutations(aa SliceType) *big.Int {
	var f big.Int
	return f.MulRange(1, int64(len(aa)))
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
func Permute(aa SliceType) SliceType2 {
	if !Permutable(aa) {
		panic(fmt.Sprintf("The number of permutations for this list (%v) exceeeds MaxInt64.", Permutations(aa)))
	}

	acc := SliceType2{}
	generate(int64(len(aa)), aa, &acc)
	return acc
}

func generate(n int64, aa SliceType, acc *SliceType2) {
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

// Pop returns a SliceType containing the head element from aa, and removes the
// element from aa. If aa is empty, the returned SliceType will also be empty.
func Pop(aa *SliceType) SliceType {
	bb := Head(*aa)
	RemoveAt(aa, 0)
	return bb
}

// Push places a prepends a new element at the head of aa.
func Push(aa *SliceType, a PrimitiveType) {
	InsertAt(aa, a, 0)
}

// Reduce applies a reducer function to each element in aa, threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new SliceType. If aa is empty, the resulting SliceType
// will also be empty.
//
//  Illustration:
//    aa: [1,2,3,4]
//    reducer: acc + sourceNode
//    Fold(aa, reducer) -> [10]
func Reduce(aa SliceType, reducer func(a, acc PrimitiveType) PrimitiveType) SliceType {
	if len(aa) == 0 {
		return SliceType{}
	}
	accumulator := aa[0]
	if len(aa) > 1 {
		for i := 1; i < len(aa); i++ {
			accumulator = reducer(aa[i], accumulator)
		}
	}
	return SliceType{accumulator}
}

// Remove applies a test function to each item in the list, and removes any item
// for which the test returns true.
func Remove(aa *SliceType, test func(PrimitiveType) bool) {
	for i := int64(len(*aa)) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, i)
		}
	}
}

// RemoveAt removes the item at the specified index from the slice.
// If len(aa) == 0, aa == nil, i < 0, or i >= len(aa), this function will do
// nothing.
func RemoveAt(aa *SliceType, i int64) {
	if i < 0 || i >= int64(len(*aa)) {
		return
	}
	if len(*aa) > 0 {
		*aa = append((*aa)[:i], (*aa)[i+1:]...)
	}
}

// Reverse reverses the order of aa.
func Reverse(aa *SliceType) {
	for i := len(*aa)/2 - 1; i >= 0; i-- {
		j := len(*aa) - 1 - i
		(*aa)[i], (*aa)[j] = (*aa)[j], (*aa)[i]
	}
}

// SwapIndex swaps the elements at the specified indices.
func SwapIndex(aa SliceType, i, j int64) {
	aa[i], aa[j] = aa[j], aa[i]
}
