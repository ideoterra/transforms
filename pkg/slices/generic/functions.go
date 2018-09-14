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
func All(aa []interface{}, test func(interface{}) bool) bool {
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
func Any(aa []interface{}, test func(interface{}) bool) bool {
	for _, a := range aa {
		if test(a) {
			return true
		}
	}
	return false
}

//Append adds the supplied values to the end of the slice.
func Append(aa *[]interface{}, values ...interface{}) {
	*aa = append(*aa, values...)
}

// Clear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func Clear(aa *[]interface{}) {
	*aa = nil
}

// Clone returns a copy of aa.
func Clone(aa []interface{}) []interface{} {
	return append([]interface{}{}, aa...)
}

// Collect applies a given function against each item in slice aa and
// each item of a slice bb, and returns the concatenation of each result.
//
//   Illustration:
//     aa:  		[A, B, C]
//     bb: 			[X, Y, Z]
//     collector:   func(a, b) { return a + b }
//     Collect(aa, bb, collector) -> [AX, AY, AZ, BX, BY, BZ, CX, XY, CZ]
func Collect(aa []interface{}, bb []interface{}, collector func(a, b interface{}) interface{}) []interface{} {
	cc := []interface{}{}
	for _, a := range aa {
		for _, b := range bb {
			cc = append(cc, collector(a, b))
		}
	}
	return cc
}

// Count applies the supplied test function to each element of the slice,
// and returns the count of items for which the test returns true.
func Count(aa []interface{}, test func(interface{}) bool) int64 {
	matches := int64(0)
	for _, a := range aa {
		if test(a) {
			matches++
		}
	}
	return matches
}

// Dequeue returns a []interface{} containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func Dequeue(aa *[]interface{}) []interface{} {
	if len(*aa) == 0 {
		return []interface{}{}
	}
	head := (*aa)[0]
	RemoveAt(aa, 0)
	return []interface{}{head}
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
func Difference(aa, bb []interface{}, equality func(a, b interface{}) bool) []interface{} {
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

	cc := []interface{}{}
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
func Distinct(aa *[]interface{}, equality func(a, b interface{}) bool) {
	bb := []interface{}{}
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
func Empty(aa []interface{}) bool {
	return len(aa) == 0
}

// End returns the a []interface{} containing only the last element from aa.
func End(aa []interface{}) []interface{} {
	if Empty(aa) {
		return []interface{}{}
	}
	return []interface{}{aa[len(aa)-1]}
}

// Enqueue places an item at the head of the slice.
func Enqueue(aa *[]interface{}, a interface{}) {
	*aa = append(*aa, a)
	copy((*aa)[1:], (*aa)[:len(*aa)-1])
	(*aa)[0] = a
}

// Expand applies an expansion function to each element of aa, and flattens
// the results into a single []interface{}.
//
//   Illustration (pseudocode):
//     aa: [AB, CD, EF]
//     expansion: func(a string) []string { return []string{a[0], a[1]}}
//     Expand(aa, expansion) -> [A, B, C, D, E, F]
func Expand(aa []interface{}, expansion func(interface{}) []interface{}) []interface{} {
	bb := []interface{}{}
	for _, a := range aa {
		Append(&bb, expansion(a)...)
	}
	return bb
}

// Filter removes all items from the slice for which the supplied test function
// returns true.
func Filter(aa *[]interface{}, test func(interface{}) bool) {
	for i := len(*aa) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, int64(i))
		}
	}
}

// FindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func FindIndex(aa []interface{}, test func(interface{}) bool) int64 {
	for i, a := range aa {
		if test(a) {
			return int64(i)
		}
	}
	return -1
}

// First returns a []interface{} containing the first element in the slice for which
// the supplied test function returns true.
func First(aa []interface{}, test func(interface{}) bool) []interface{} {
	bb := []interface{}{}
	for _, a := range aa {
		if test(a) {
			Append(&bb, a)
			break
		}
	}
	return bb
}

// Flatten takes each slice of a [][]interface{} and appends it to a new slice.
func Flatten(aa [][]interface{}) []interface{} {
	bb := []interface{}{}
	for _, a := range aa {
		Append(&bb, a)
	}
	return bb
}

// Fold applies a function to each item in slice aa, threading an accumulator
// through each iteration. The accumulated value is returned in a new []interface{}
// once aa is fully scanned. Fold returns a []interface{} rather than a
// interface{} to be consistent with this package's Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func Fold(aa []interface{}, acc interface{}, folder func(a, acc interface{}) interface{}) []interface{} {
	return FoldI(aa, acc, func(_ int64, a, acc interface{}) interface{} { return folder(a, acc) })
}

// FoldI applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned. Foldi returns a []interface{} rather than a
// interface{} to be consistent with this package's Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func FoldI(aa []interface{}, acc interface{}, folder func(i int64, a, acc interface{}) interface{}) []interface{} {
	accumulation := acc
	for i, a := range aa {
		accumulation = folder(int64(i), a, accumulation)
	}
	return []interface{}{accumulation}
}

// ForEach applies each element of the list to the given function.
// ForEach will stop iterating if fn return false.
func ForEach(aa []interface{}, fn func(interface{}) shared.Continue) {
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
func ForEachC(aa []interface{}, c int, fn func(a interface{}, cancelPending func() bool) shared.Continue) {
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
		go func(a interface{}) {
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
func ForEachR(aa []interface{}, fn func(interface{}) shared.Continue) {
	for i := len(aa) - 1; i >= 0; i-- {
		if !fn(aa[i]) {
			return
		}
	}
}

// Group consolidates like-items into groups according to the supplied grouper
// function, and returns them as a [][]interface{}.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed.
func Group(aa []interface{}, grouper func(interface{}) int64) []interface{} {
	return GroupI(aa, func(_ int64, a interface{}) int64 { return grouper(a) })
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
func GroupByTrait(aa []interface{}, trait func(ai, an interface{}) bool, equality func(a, b interface{}) bool) []interface{} {
	establishedTraits := []interface{}{}
	for _, ai := range aa {
		potentialTrait := []interface{}{}
		for _, an := range aa {
			if trait(ai, an) {
				Append(&potentialTrait, an)
			}
		}
		traitIsSubsetOfEstablished := false
		for i := len(establishedTraits) - 1; i >= 0; i-- {
			establishedTrait := establishedTraits[i]
			if IsSubset(potentialTrait, establishedTrait.([]interface{}), equality) {
				traitIsSubsetOfEstablished = true
				break
			}
			if IsSuperset(potentialTrait, establishedTrait.([]interface{}), equality) {
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
// function, and returns them as a [][]interface{}.
// The grouper function is expected to return a hash value which Group will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from aa is also passed into the grouper function.
func GroupI(aa []interface{}, grouper func(int64, interface{}) int64) []interface{} {
	groupMap := map[int64][]interface{}{}
	for i, a := range aa {
		hash := grouper(int64(i), a)
		if _, exists := groupMap[hash]; exists {
			groupMap[hash] = append(groupMap[hash], a)
		} else {
			groupMap[hash] = []interface{}{a}
		}
	}
	group := []interface{}{}
	for _, bb := range groupMap {
		group = append(group, bb)
	}
	return group
}

// Head returns a []interface{} containing the first item from the aa. If aa is
// empty, the resulting []interface{} will be empty.
func Head(aa []interface{}) []interface{} {
	if Empty(aa) {
		return []interface{}{}
	}
	return []interface{}{aa[0]}
}

// InsertAfter inserts an element in aa after the first element for which the
// supplied test function returns true. If none of the tests return true, the
// element is appended to the end of the aa.
func InsertAfter(aa *[]interface{}, b interface{}, test func(interface{}) bool) {
	var i int
	var a interface{}
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
func InsertBefore(aa *[]interface{}, b interface{}, test func(interface{}) bool) {
	var i int
	var a interface{}
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
func InsertAt(aa *[]interface{}, a interface{}, i int64) {
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
// function, and returns a []interface{} containing the elements which are common
// to both aa and bb. Duplicates are removed in this operation.
func Intersection(aa, bb []interface{}, equality func(a, b interface{}) bool) []interface{} {
	cc := []interface{}{}
	ForEach(aa, func(a interface{}) shared.Continue {
		ForEach(bb, func(b interface{}) shared.Continue {
			if equality(a, b) && !Any(cc, func(c interface{}) bool { return equality(a, c) }) {
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
func IsProperSubset(aa, bb []interface{}, equality func(a, b interface{}) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) > 0
}

// IsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsProperSuperset(aa, bb []interface{}, equality func(a, b interface{}) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) > 0 && len(bb1) == 0
}

// IsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsSubset(aa, bb []interface{}, equality func(a, b interface{}) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) >= 0
}

// IsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the Distinct
// operations to enforce uniqueness, if that is necessary.
func IsSuperset(aa, bb []interface{}, equality func(a, b interface{}) bool) bool {
	aa1, bb1 := removeIntersections(aa, bb, equality)
	return len(aa1) >= 0 && len(bb1) == 0
}

func removeIntersections(aa, bb []interface{}, equality func(a, b interface{}) bool) ([]interface{}, []interface{}) {
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

// Item returns a []interface{} containing the element at aa[i].
// If len(aa) == 0, i < 0, or, i >= len(aa), the resulting slice will be empty.
func Item(aa []interface{}, i int64) []interface{} {
	if Empty(aa) || i < 0 || i >= int64(len(aa)) {
		return []interface{}{}
	}
	return []interface{}{aa[i]}
}

// ItemFuzzy returns a []interface{} containing the element at aa[i].
// If the supplied index is outside of the bounds of aa, ItemFuzzy will attempt
// to retrieve the head or end element of aa according to the following rules:
// If len(aa) == 0 an empty []interface{} is returned.
// If i < 0, the head of aa is returned.
// If i >= len(aa), the end of the aa is returned.
func ItemFuzzy(aa []interface{}, i int64) []interface{} {
	if Empty(aa) {
		return []interface{}{}
	}
	if i < 0 {
		return Head(aa)
	}
	if i >= int64(len(aa)) {
		return End(aa)
	}
	return []interface{}{aa[i]}
}

// Last applies a test function to each element in aa, and returns a []interface{}
// containing the last element for which the test returned true. If no elements
// pass the supplied test, the resulting []interface{} will be empty.
func Last(aa []interface{}, test func(interface{}) bool) []interface{} {
	bb := []interface{}{}
	ForEachR(aa, func(a interface{}) shared.Continue {
		if test(a) {
			Append(&bb, a)
			return shared.ContinueNo
		}
		return shared.ContinueYes
	})
	return bb
}

// Len returns the length of aa.
func Len(aa []interface{}) int {
	return len(aa)
}

// Map applies a tranform to each element of the list.
func Map(aa *[]interface{}, mapFn func(interface{}) interface{}) {
	for i, a := range *aa {
		(*aa)[i] = mapFn(a)
	}
}

// None applies a test function to each element in aa, and returns true if
// the test function returns false for all items.
func None(aa []interface{}, test func(interface{}) bool) bool {
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
func Pairwise(aa []interface{}, init interface{}, xform func(a, b interface{}) interface{}) []interface{} {
	if Empty(aa) {
		return []interface{}{}
	}
	bb := []interface{}{}
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
// a [][]interface{} where [][]interface{}[0] contains a []interface{} with all elements for
// whom the test function returned true, and where [][]interface{}[1] contains a
// []interface{} with all elements for whom the test function returned false.
//
// Partition is a special case of the Group function.
func Partition(aa []interface{}, test func(interface{}) bool) []interface{} {
	grouper := func(a interface{}) int64 {
		if test(a) {
			return 1
		}
		return 0
	}
	return Group(aa, grouper)
}

// Permutable returns true if the number of permutations for aa exceeds
// MaxInt64.
func Permutable(aa []interface{}) bool {
	return Permutations(aa).IsInt64()
}

// Permutations returns the number of permutations that exist given the current
// number of items in the aa.
func Permutations(aa []interface{}) *big.Int {
	var f big.Int
	return f.MulRange(1, int64(len(aa)))
}

// Permute returns a [][]interface{} which contains a []interface{} for each permutation
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
func Permute(aa []interface{}) []interface{} {
	if Empty(aa) {
		return []interface{}{}
	}

	if !Permutable(aa) {
		panic(fmt.Sprintf("The number of permutations for this list (%v) exceeeds MaxInt64.", Permutations(aa)))
	}

	acc := []interface{}{}
	generate(int64(len(aa)), aa, &acc)
	return acc
}

func generate(n int64, aa []interface{}, acc *[]interface{}) {
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

// Pop returns a []interface{} containing the head element from aa, and removes the
// element from aa. If aa is empty, the returned []interface{} will also be empty.
func Pop(aa *[]interface{}) []interface{} {
	bb := Head(*aa)
	RemoveAt(aa, 0)
	return bb
}

// Push places a prepends a new element at the head of aa.
func Push(aa *[]interface{}, a interface{}) {
	InsertAt(aa, a, 0)
}

// Reduce applies a reducer function to each element in aa, threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new []interface{}. If aa is empty, the resulting []interface{}
// will also be empty.
//
//  Illustration:
//    aa: [1,2,3,4]
//    reducer: acc + sourceNode
//    Fold(aa, reducer) -> [10]
func Reduce(aa []interface{}, reducer func(a, acc interface{}) interface{}) []interface{} {
	if len(aa) == 0 {
		return []interface{}{}
	}
	accumulator := aa[0]
	if len(aa) > 1 {
		for i := 1; i < len(aa); i++ {
			accumulator = reducer(aa[i], accumulator)
		}
	}
	return []interface{}{accumulator}
}

// Remove applies a test function to each item in the list, and removes any item
// for which the test returns true.
func Remove(aa *[]interface{}, test func(interface{}) bool) {
	for i := int64(len(*aa)) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			RemoveAt(aa, i)
		}
	}
}

// RemoveAt removes the item at the specified index from the slice.
// If len(aa) == 0, aa == nil, i < 0, or i >= len(aa), this function will do
// nothing.
func RemoveAt(aa *[]interface{}, i int64) {
	if i < 0 || i >= int64(len(*aa)) {
		return
	}
	if len(*aa) > 0 {
		*aa = append((*aa)[:i], (*aa)[i+1:]...)
	}
}

// Reverse reverses the order of aa.
func Reverse(aa *[]interface{}) {
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
func Skip(aa *[]interface{}, n int64) {
	if len(*aa) == 0 {
		return
	}
	*aa = (*aa)[n:]
}

// SkipWhile scans through aa starting at the head, and removes all
// elements from aa while the test function returns true.
// SkipWhile stops removing any further items from aa after the first test that
// returns false.
func SkipWhile(aa *[]interface{}, test func(interface{}) bool) {
	// find the first index where the test would evaluate to false and skip
	// everything up to that index.
	findTest := func(a interface{}) bool { return !test(a) }
	Skip(aa, FindIndex(*aa, findTest))
}

// Sort sorts aa, using the supplied less function to determine order.
// Sort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func Sort(aa *[]interface{}, less func(a, b interface{}) bool) {
	lessI := func(i, j int) bool {
		return less((*aa)[i], (*aa)[j])
	}
	sort.SliceStable(*aa, lessI)
}

// SplitAfter finds the first element b for which a test function returns true,
// and returns a [][]interface{} where [][]interface{}[0] contains the first half of aa
// and [][]interface{}[1] contains the second half of aa. Element b will be included
// in [][]interface{}[0]. If the no element can be found for which the test returns
// true, [][]interface{}[0] will contain aa, and [][]interface{}[1] will be empty.
func SplitAfter(aa []interface{}, test func(interface{}) bool) [][]interface{} {
	return SplitAt(aa, FindIndex(aa, test)+1)
}

// SplitAt splits aa at index i, and returns a [][]interface{} which contains the
// two split halves of aa. aa[i] will be included in [][]interface{}[1].
// If i < 0, all of aa will be placed in [][]interface{}[0] and [][]interface{}[1] will
// be empty. Conversly, if i >= len(aa), all of aa will be placed in
// [][]interface{}[1] and [][]interface{}[0] will be empty. If aa is nil or empty,
// [][]interface{} will contain two empty slices.
func SplitAt(aa []interface{}, i int64) [][]interface{} {
	if len(aa) == 0 {
		return [][]interface{}{
			[]interface{}{},
			[]interface{}{},
		}
	}
	if i < 0 {
		i = 0
	}
	return [][]interface{}{
		aa[:i],
		aa[i:],
	}
}

// SplitBefore finds the first element b for which a test function returns true,
// and returns a [][]interface{} where [][]interface{}[0] contains the first half of aa
// and [][]interface{}[1] contains the second half of aa. Element b will be included
// in [][]interface{}[1]
func SplitBefore(aa []interface{}, test func(interface{}) bool) [][]interface{} {
	return SplitAt(aa, FindIndex(aa, test))
}

// String returns a string representation of aa, suitable for use
// with fmt.Print, or other similar functions. String should be regarded as
// informational, and should not be relied upon to formally serialize a
// []interface{}.
func String(aa []interface{}) string {
	jsonBytes, _ := json.Marshal(aa)
	return string(jsonBytes)
}

// SwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of aa, SwapIndex does nothing.
func SwapIndex(aa []interface{}, i, j int64) {
	l := int64(len(aa))
	if i < 0 || j < 0 || i >= l || j >= l {
		return
	}
	aa[i], aa[j] = aa[j], aa[i]
}

// Tail removes the current head element from aa.
// This equivelant to RemoveAt(aa, 0)
func Tail(aa *[]interface{}) {
	RemoveAt(aa, 0)
}

// Take retains the first n elements of aa, and removes all remaining elements
// from the slice. If n < 0 or n >= len(aa), Take does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func Take(aa *[]interface{}, n int64) {
	if len(*aa) == 0 || n < 0 || n >= int64(len(*aa)) {
		return
	}
	*aa = (*aa)[:n]
}

// TakeWhile applies a test function to each element in aa, and retains all
// elements of aa so long as the test function returns true. As soon as the test
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func TakeWhile(aa *[]interface{}, test func(interface{}) bool) {
	find := func(a interface{}) bool {
		return !test(a)
	}
	Take(aa, FindIndex(*aa, find))
}

// Union appends slice bb to slice aa.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func Union(aa *[]interface{}, bb []interface{}) {
	Append(aa, bb...)
}

// Unzip splits aa into a [][]interface{}, such that [][]interface{}[0] contains all odd
// indices from aa, and [][]interface{}[1] contains all even indices from aa.
func Unzip(aa []interface{}) [][]interface{} {
	odds := []interface{}{}
	evens := []interface{}{}
	for i, a := range aa {
		if i%2 != 0 {
			odds = append(odds, a)
		} else {
			evens = append(evens, a)
		}
	}
	return [][]interface{}{odds, evens}
}

// WindowCentered applies a windowing function across the aa, using a centered
// window of the specified size.
func WindowCentered(aa []interface{}, windowSize int64, windowFn func(window []interface{}) interface{}) []interface{} {
	cc := []interface{}{}
	fullWindowReached := false
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := []interface{}{}
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
	dd := []interface{}(SplitAt(cc, frontTrim)[1])
	Reverse(&dd)
	ee := []interface{}(SplitAt(dd, backTrim)[1])
	Reverse(&ee)
	return ee
}

// WindowLeft applies a windowing function across aa, using a left-sided window
// of the specified size.
func WindowLeft(aa []interface{}, windowSize int64, windowFn func(window []interface{}) interface{}) []interface{} {
	bb := []interface{}{}
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := []interface{}{}
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
func WindowRight(aa []interface{}, windowSize int64, windowFn func(window []interface{}) interface{}) []interface{} {
	aa1 := Clone(aa)
	defer Clear(&aa1)

	Reverse(&aa1)
	bb := []interface{}{}
	for i := int64(0); i < int64(len(aa1)); i++ {
		currentWindow := []interface{}{}
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
// new []interface{}. aa[0] is evaluated first. Thus if aa and bb are the same
// length, slice aa will occupy the odd indices of the result slice, and bb
// will occupy the even indices of the result slice. If aa and bb are not
// the same length, Zip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func Zip(aa, bb []interface{}) []interface{} {
	if len(aa) == 0 {
		return bb
	}
	if len(bb) == 0 {
		return aa
	}

	cc := []interface{}{}
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
