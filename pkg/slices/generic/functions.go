package generic

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

// All applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
func All(aa SliceType, test Test) bool {
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
func Any(aa SliceType, test Test) bool {
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

// Clone returns a copy of aa.
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
func Count(aa SliceType, test Test) int64 {
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
// between aa and bb. The supplied equality function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. Distinct values from aa are listed ahead of those from bb in the
// resulting slice.
//
// Illustration:
//   aa: [1,2,3,3,1,4]
//   bb: [5,4,3,5]
//   equal: func(a, b) bool {return a == b}
//   Difference(aa, bb, equality) -> [1,2,1,5,5]
func Difference(aa, bb SliceType, equality Equality) SliceType {
	ii := make([]bool, len(aa))
	jj := make([]bool, len(bb))
	for i, a := range aa {
		for j, b := range bb {
			if equality(a, b) {
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

// Distinct removes all duplicates from the slice, using the supplied equality
// function to determine equality.
func Distinct(aa *SliceType, equality Equality) {
	bb := SliceType{}
	dups := make([]bool, len(*aa))
	for i, a := range *aa {
		if !dups[i] {
			bb = append(bb, a)
		}
		for j := i + 1; j < len(*aa); j++ {
			if equality(a, (*aa)[j]) {
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
func Filter(aa *SliceType, test Test) {
	for i := len(*aa) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, int64(i))
		}
	}
}

// FindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func FindIndex(aa SliceType, test Test) int64 {
	for i, a := range aa {
		if test(a) {
			return int64(i)
		}
	}
	return -1
}

// First returns a SliceType containing the first element in the slice for which
// the supplied test function returns true.
func First(aa SliceType, test Test) SliceType {
	bb := SliceType{}
	for _, a := range aa {
		if test(a) {
			Append(&bb, a)
			break
		}
	}
	return bb
}

// Flatten takes each slice of a SliceType2 and appends it to a new slice.
func Flatten(aa SliceType2) SliceType {
	bb := SliceType{}
	for _, a := range aa {
		Append(&bb, a)
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
func ForEach(aa SliceType, fn func(PrimitiveType) shared.Continue) {
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
// If any execution of fn returns shared.ContinueNo, ForEachC will cease marshalling
// any backlogged work, and will immediately set the cancellation flag to true.
// Any goroutines monitoring the cancelPending closure can wind down their
// activities as necessary. ForEachC will continue to block until all active
// goroutines exit cleanly.
func ForEachC(aa SliceType, c int, fn func(a PrimitiveType, cancelPending func() bool) shared.Continue) {
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
		go func(a PrimitiveType) {
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
func ForEachR(aa SliceType, fn func(PrimitiveType) shared.Continue) {
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

// GroupByTrait compares each item (a[i]) in the slice to every other item
// (a[n]) using the supplied trait function. Every item a[n] who shares a trait
// with a[i] is added to a slice that represents a group of items that express a
// potential trait. This potential trait is then compared to a slice of
// established traits using the supplied equality function. If the potential
// trait is a subset of any established trait, the potential trait is it is
// disregarded, othwewise, the potential trait is added as an established trait.
// If the potential trait is a superset of any established trait, each relevent
// established trait is disregarded.
//
//  Illustration (pseuodocode):
//    aa: [pigdog, pigs, dog, pigdogs, cat, dogs, pig]
//    trait: return strings.Index(a[i], a[n]) == 0
//    equal: return a[i] == a[j]
//    GroupByTrait(aa, trait, equality) ->
//			[
//			 [pigdogs, pigdog, pigs, pig],
//			 [cat],
//			 [dogs, dog],
//			]
func GroupByTrait(aa SliceType, trait func(ai, an PrimitiveType) bool, equality Equality) SliceType2 {
	establishedTraits := SliceType2{}
	for _, ai := range aa {
		potentialTrait := SliceType{}
		for _, an := range aa {
			if trait(ai, an) {
				Append(&potentialTrait, an)
			}
		}
		traitIsSubsetOfEstablished := false
		for i := len(establishedTraits) - 1; i >= 0; i-- {
			establishedTrait := establishedTraits[i]
			if IsSubset(potentialTrait, establishedTrait, equality) {
				traitIsSubsetOfEstablished = true
				break
			}
			if IsSuperset(potentialTrait, establishedTrait, equality) {
				establishedTraits = append(establishedTraits[:i], establishedTraits[i+1:]...)
			}
		}
		if !traitIsSubsetOfEstablished {
			establishedTraits = append(establishedTraits, potentialTrait)
		}
	}
	return establishedTraits
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
func InsertAfter(aa *SliceType, b PrimitiveType, test Test) {
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
func InsertBefore(aa *SliceType, b PrimitiveType, test Test) {
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
// to both aa and bb. Duplicates are removed in this operation.
func Intersection(aa, bb SliceType, equality Equality) SliceType {
	cc := SliceType{}
	ForEach(aa, func(a PrimitiveType) shared.Continue {
		ForEach(bb, func(b PrimitiveType) shared.Continue {
			if equality(a, b) && !Any(cc, func(c PrimitiveType) bool { return equality(a, c) }) {
				Append(&cc, a)
			}
			return shared.ContinueYes
		})
		return shared.ContinueYes
	})
	return cc
}

// IsProperSubset returns true if aa is a proper subset of bb.
// aa is considered a proper subset if all of its elements exist within bb, but
// bb also contains some elements that do not exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsProperSubset(aa, bb SliceType, equality Equality) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) > 0
}

// IsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsProperSuperset(aa, bb SliceType, equality Equality) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) > 0 && len(bb1) == 0
}

// IsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsSubset(aa, bb SliceType, equality Equality) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) >= 0
}

// IsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsSuperset(aa, bb SliceType, equality Equality) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) >= 0 && len(bb1) == 0
}

func removeIntersections(aa, bb SliceType, equality Equality) (SliceType, SliceType) {
	aa1 := Clone(aa)
	bb1 := Clone(bb)
	for ai := int64(len(aa1)) - 1; ai >= 0; ai-- {
		intersectionFound := false
		for bi := int64(len(bb1)) - 1; bi >= 0; bi-- {
			if equality((aa1)[ai], (bb1)[bi]) {
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
func Last(aa SliceType, test Test) SliceType {
	bb := SliceType{}
	ForEachR(aa, func(a PrimitiveType) shared.Continue {
		if test(a) {
			Append(&bb, a)
			return shared.ContinueNo
		}
		return shared.ContinueYes
	})
	return bb
}

// Len returns the length of aa.
func Len(aa SliceType) int {
	return len(aa)
}

// Map applies a tranform to each element of the list.
func Map(aa *SliceType, mapFn func(PrimitiveType) PrimitiveType) {
	for i, a := range *aa {
		(*aa)[i] = mapFn(a)
	}
}

// None applies a test function to each element in aa, and returns true if
// the test function returns false for all items.
func None(aa SliceType, test Test) bool {
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
	if Empty(aa) {
		return SliceType{}
	}
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
func Partition(aa SliceType, test Test) SliceType2 {
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
	if Empty(aa) {
		return SliceType2{}
	}

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
func Remove(aa *SliceType, test Test) {
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

// Skip removes the first n elements from aa.
//
// Note that Skip(aa, len(aa)) will remove all items from the list, but does not
// "clear" the slice, meaning that the list remains allocated in memory.
// To fully de-pointer the slice, and ensure it is available for garbage
// collection as soon as possible, consider using Clear().
func Skip(aa *SliceType, n int64) {
	if len(*aa) == 0 {
		return
	}
	*aa = (*aa)[n:]
}

// SkipWhile scans through aa starting at the head, and removes all
// elements from aa while the test function returns true.
// SkipWhile stops removing any further items from aa after the first test that
// returns false.
func SkipWhile(aa *SliceType, test Test) {
	// find the first index where the test would evaluate to false and skip
	// everything up to that index.
	findTest := func(a PrimitiveType) bool { return !test(a) }
	Skip(aa, FindIndex(*aa, findTest))
}

// Sort sorts aa, using the supplied less function to determine order.
// Sort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func Sort(aa *SliceType, less func(a, b PrimitiveType) bool) {
	lessI := func(i, j int) bool {
		return less((*aa)[i], (*aa)[j])
	}
	sort.SliceStable(*aa, lessI)
}

// SplitAfter finds the first element b for which a test function returns true,
// and returns a SliceType2 where SliceType2[0] contains the first half of aa
// and SliceType2[1] contains the second half of aa. Element b will be included
// in SliceType2[0]. If the no element can be found for which the test returns
// true, SliceType2[0] will contain aa, and SliceType2[1] will be empty.
func SplitAfter(aa SliceType, test Test) SliceType2 {
	return SplitAt(aa, FindIndex(aa, test)+1)
}

// SplitAt splits aa at index i, and returns a SliceType2 which contains the
// two split halves of aa. aa[i] will be included in SliceType2[1].
// If i < 0, all of aa will be placed in SliceType2[0] and SliceType2[1] will
// be empty. Conversly, if i >= len(aa), all of aa will be placed in
// SliceType2[1] and SliceType2[0] will be empty. If aa is nil or empty,
// SliceType2 will contain two empty slices.
func SplitAt(aa SliceType, i int64) SliceType2 {
	if len(aa) == 0 {
		return SliceType2{
			SliceType{},
			SliceType{},
		}
	}
	if i < 0 {
		i = 0
	}
	return SliceType2{
		aa[:i],
		aa[i:],
	}
}

// SplitBefore finds the first element b for which a test function returns true,
// and returns a SliceType2 where SliceType2[0] contains the first half of aa
// and SliceType2[1] contains the second half of aa. Element b will be included
// in SliceType2[1]
func SplitBefore(aa SliceType, test Test) SliceType2 {
	return SplitAt(aa, FindIndex(aa, test))
}

// String returns a string representation of aa, suitable for use
// with fmt.Print, or other similar functions. String should be regarded as
// informational, and should not be relied upon to formally serialize a
// SliceType.
func String(aa SliceType) string {
	jsonBytes, _ := json.Marshal(aa)
	return string(jsonBytes)
}

// SwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of aa, SwapIndex does nothing.
func SwapIndex(aa SliceType, i, j int64) {
	l := int64(len(aa))
	if i < 0 || j < 0 || i >= l || j >= l {
		return
	}
	aa[i], aa[j] = aa[j], aa[i]
}

// Tail removes the current head element from aa.
// This equivelant to RemoveAt(aa, 0)
func Tail(aa *SliceType) {
	RemoveAt(aa, 0)
}

// Take retains the first n elements of aa, and removes all remaining elements
// from the slice. If n < 0 or n >= len(aa), Take does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func Take(aa *SliceType, n int64) {
	if len(*aa) == 0 || n < 0 || n >= int64(len(*aa)) {
		return
	}
	*aa = (*aa)[:n]
}

// TakeWhile applies a test function to each element in aa, and retains all
// elements of aa so long as the test function returns true. As soon as the test
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func TakeWhile(aa *SliceType, test Test) {
	find := func(a PrimitiveType) bool {
		return !test(a)
	}
	Take(aa, FindIndex(*aa, find))
}

// Union appends slice bb to slice aa.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func Union(aa *SliceType, bb SliceType) {
	Append(aa, bb...)
}

// Unzip splits aa into a SliceType2, such that SliceType2[0] contains all odd
// indices from aa, and SliceType2[1] contains all even indices from aa.
func Unzip(aa SliceType) SliceType2 {
	odds := SliceType{}
	evens := SliceType{}
	for i, a := range aa {
		if i%2 != 0 {
			odds = append(odds, a)
		} else {
			evens = append(evens, a)
		}
	}
	return SliceType2{odds, evens}
}

// WindowCentered applies a windowing function across the aa, using a centered
// window of the specified size.
func WindowCentered(aa SliceType, windowSize int64, windowFn func(window SliceType) PrimitiveType) SliceType {
	cc := SliceType{}
	fullWindowReached := false
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := SliceType{}
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
	dd := SliceType(SplitAt(cc, frontTrim)[1])
	Reverse(&dd)
	ee := SliceType(SplitAt(dd, backTrim)[1])
	Reverse(&ee)
	return ee
}

// WindowLeft applies a windowing function across aa, using a left-sided window
// of the specified size.
func WindowLeft(aa SliceType, windowSize int64, windowFn func(window SliceType) PrimitiveType) SliceType {
	bb := SliceType{}
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := SliceType{}
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
func WindowRight(aa SliceType, windowSize int64, windowFn func(window SliceType) PrimitiveType) SliceType {
	aa1 := Clone(aa)
	defer Clear(&aa1)

	Reverse(&aa1)
	bb := SliceType{}
	for i := int64(0); i < int64(len(aa1)); i++ {
		currentWindow := SliceType{}
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
// new SliceType. aa[0] is evaluated first. Thus if aa and bb are the same
// length, slice aa will occupy the odd indices of the result slice, and bb
// will occupy the even indices of the result slice. If aa and bb are not
// the same length, Zip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func Zip(aa, bb SliceType) SliceType {
	if len(aa) == 0 {
		return bb
	}
	if len(bb) == 0 {
		return aa
	}

	cc := SliceType{}
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
