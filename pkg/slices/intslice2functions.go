package slicexform

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

// IntSlice2All applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
func IntSlice2All(aa IntSlice2, test Test) bool {
	for _, s := range aa {
		if !test(s) {
			return false
		}
	}
	return true
}

// IntSlice2Any applies a test function to each element of the
// slice and returns true if the test function returns true for at least one
// item in the list.
//
// IntSlice2Any does not require that the source slice be sorted, and merely scans
// the slice, returning as soon as any element passes the supplied test. For
// a binary search, consider using sort.Search from the standard library.
func IntSlice2Any(aa IntSlice2, test Test) bool {
	for _, a := range aa {
		if test(a) {
			return true
		}
	}
	return false
}

//IntSlice2Append adds the supplied values to the end of the slice.
func IntSlice2Append(aa *IntSlice2, values ...IntSlice) {
	*aa = append(*aa, values...)
}

// IntSlice2Clear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func IntSlice2Clear(aa *IntSlice2) {
	*aa = nil
}

// IntSlice2Clone returns a copy of aa.
func IntSlice2Clone(aa IntSlice2) IntSlice2 {
	return append(IntSlice2{}, aa...)
}

// IntSlice2Collect applies a given function against each item in slice aa and
// each item of a slice bb, and returns the concatenation of each result.
//
//   Illustration:
//     aa:  		[A, B, C]
//     bb: 			[X, Y, Z]
//     collector:   func(a, b) { return a + b }
//     IntSlice2Collect(aa, bb, collector) -> [AX, AY, AZ, BX, BY, BZ, CX, XY, CZ]
func IntSlice2Collect(aa IntSlice2, bb IntSlice2, collector func(a, b IntSlice) IntSlice) IntSlice2 {
	cc := IntSlice2{}
	for _, a := range aa {
		for _, b := range bb {
			cc = append(cc, collector(a, b))
		}
	}
	return cc
}

// IntSlice2Count applies the supplied test function to each element of the slice,
// and returns the count of items for which the test returns true.
func IntSlice2Count(aa IntSlice2, test Test) int64 {
	matches := int64(0)
	for _, a := range aa {
		if test(a) {
			matches++
		}
	}
	return matches
}

// IntSlice2Dequeue returns a IntSlice2 containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func IntSlice2Dequeue(aa *IntSlice2) IntSlice2 {
	if len(*aa) == 0 {
		return IntSlice2{}
	}
	head := (*aa)[0]
	IntSlice2IntSlice2RemoveAt(aa, 0)
	return IntSlice2{head}
}

// IntSlice2Difference returns a new slice that contains items that are not common
// between aa and bb. The supplied equality function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. IntSlice2Distinct values from aa are listed ahead of those from bb in the
// resulting slice.
//
// Illustration:
//   aa: [1,2,3,3,1,4]
//   bb: [5,4,3,5]
//   equal: func(a, b) bool {return a == b}
//   IntSlice2Difference(aa, bb, equality) -> [1,2,1,5,5]
func IntSlice2Difference(aa, bb IntSlice2, equality Equality) IntSlice2 {
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

	cc := IntSlice2{}
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

// IntSlice2Distinct removes all duplicates from the slice, using the supplied equality
// function to determine equality.
func IntSlice2Distinct(aa *IntSlice2, equality Equality) {
	bb := IntSlice2{}
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
	IntSlice2Clear(aa)
	IntSlice2Append(aa, bb...)
}

// IntSlice2Empty returns true if the length of the slice is zero.
func IntSlice2Empty(aa IntSlice2) bool {
	return len(aa) == 0
}

// IntSlice2End returns the a IntSlice2 containing only the last element from aa.
func IntSlice2End(aa IntSlice2) IntSlice2 {
	if IntSlice2Empty(aa) {
		return IntSlice2{}
	}
	return IntSlice2{aa[len(aa)-1]}
}

// IntSlice2Enqueue places an item at the head of the slice.
func IntSlice2Enqueue(aa *IntSlice2, a IntSlice) {
	*aa = append(*aa, a)
	copy((*aa)[1:], (*aa)[:len(*aa)-1])
	(*aa)[0] = a
}

// IntSlice2Expand applies an expansion function to each element of aa, and flattens
// the results into a single IntSlice2.
//
//   Illustration (pseudocode):
//     aa: [AB, CD, EF]
//     expansion: func(a string) []string { return []string{a[0], a[1]}}
//     IntSlice2Expand(aa, expansion) -> [A, B, C, D, E, F]
func IntSlice2Expand(aa IntSlice2, expansion func(IntSlice) IntSlice2) IntSlice2 {
	bb := IntSlice2{}
	for _, a := range aa {
		IntSlice2Append(&bb, expansion(a)...)
	}
	return bb
}

// IntSlice2Filter removes all items from the slice for which the supplied test function
// returns true.
func IntSlice2Filter(aa *IntSlice2, test Test) {
	for i := len(*aa) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			IntSlice2IntSlice2RemoveAt(aa, int64(i))
		}
	}
}

// IntSlice2FindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func IntSlice2FindIndex(aa IntSlice2, test Test) int64 {
	for i, a := range aa {
		if test(a) {
			return int64(i)
		}
	}
	return -1
}

// IntSlice2First returns a IntSlice2 containing the first element in the slice for which
// the supplied test function returns true.
func IntSlice2First(aa IntSlice2, test Test) IntSlice2 {
	bb := IntSlice2{}
	for _, a := range aa {
		if test(a) {
			IntSlice2Append(&bb, a)
			break
		}
	}
	return bb
}

// IntSlice2Fold applies a function to each item in slice aa, threading an accumulator
// through each iteration. The accumulated value is returned in a new IntSlice2
// once aa is fully scanned. IntSlice2Fold returns a IntSlice2 rather than a
// IntSlice to be consistent with this package's IntSlice2Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    IntSlice2Fold(aa, acc, folder) -> [11]
func IntSlice2Fold(aa IntSlice2, acc IntSlice, folder func(a, acc IntSlice) IntSlice) IntSlice2 {
	return IntSlice2IntSlice2FoldI(aa, acc, func(_ int64, a, acc IntSlice) IntSlice { return folder(a, acc) })
}

// IntSlice2IntSlice2FoldI applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned. IntSlice2Foldi returns a IntSlice2 rather than a
// IntSlice to be consistent with this package's IntSlice2Reduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    IntSlice2Fold(aa, acc, folder) -> [11]
func IntSlice2IntSlice2FoldI(aa IntSlice2, acc IntSlice, folder func(i int64, a, acc IntSlice) IntSlice) IntSlice2 {
	accumulation := acc
	for i, a := range aa {
		accumulation = folder(int64(i), a, accumulation)
	}
	return IntSlice2{accumulation}
}

// IntSlice2ForEach applies each element of the list to the given function.
// IntSlice2ForEach will stop iterating if fn return false.
func IntSlice2ForEach(aa IntSlice2, fn func(IntSlice) shared.Continue) {
	for _, a := range aa {
		if !fn(a) {
			return
		}
	}
}

// IntSlice2IntSlice2ForEachC concurrently applies each element of the list to the given function.
// The elements of the list are marshalled to a pool of goroutines, where each
// element is passed to fn concurrently.
//
// The concurrency pool is limited to contain no more than c active goroutines
// at any time. Note that if a pool size of 0 is supplied, this method
// will block indefinitely. This function will panic if a negative value is
// supplied for c.
//
// If any execution of fn returns shared.ContinueNo, IntSlice2IntSlice2ForEachC will cease marshalling
// any backlogged work, and will immediately set the cancellation flag to true.
// IntSlice2Any goroutines monitoring the cancelPending closure can wind down their
// activities as necessary. IntSlice2IntSlice2ForEachC will continue to block until all active
// goroutines exit cleanly.
func IntSlice2IntSlice2ForEachC(aa IntSlice2, c int, fn func(a IntSlice, cancelPending func() bool) shared.Continue) {
	if c < 0 {
		panic("IntSlice2IntSlice2ForEachC: The concurrency pool size (c) must be non-negative.")
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
		go func(a IntSlice) {
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

// IntSlice2IntSlice2ForEachR applies each element of aa to a given function, scanning
// through the slice in reverse order, starting from the end and working towards
// the head.
func IntSlice2IntSlice2ForEachR(aa IntSlice2, fn func(IntSlice) shared.Continue) {
	for i := len(aa) - 1; i >= 0; i-- {
		if !fn(aa[i]) {
			return
		}
	}
}

// IntSlice2Group consolidates like-items into groups according to the supplied grouper
// function, and returns them as a []IntSlice2.
// The grouper function is expected to return a hash value which IntSlice2Group will use
// to determine into which bucket each element wil be placed.
func IntSlice2Group(aa IntSlice2, grouper func(IntSlice) int64) []IntSlice2 {
	return IntSlice2IntSlice2GroupI(aa, func(_ int64, a IntSlice) int64 { return grouper(a) })
}

// IntSlice2IntSlice2GroupI consolidates like-items into groups according to the supplied grouper
// function, and returns them as a []IntSlice2.
// The grouper function is expected to return a hash value which IntSlice2Group will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from aa is also passed into the grouper function.
func IntSlice2IntSlice2GroupI(aa IntSlice2, grouper func(int64, IntSlice) int64) []IntSlice2 {
	groupIntSlice2Map := map[int64]IntSlice2{}
	for i, a := range aa {
		hash := grouper(int64(i), a)
		if _, exists := groupIntSlice2Map[hash]; exists {
			groupIntSlice2Map[hash] = append(groupIntSlice2Map[hash], a)
		} else {
			groupIntSlice2Map[hash] = IntSlice2{a}
		}
	}
	group := []IntSlice2{}
	for _, bb := range groupIntSlice2Map {
		group = append(group, bb)
	}
	return group
}

// IntSlice2Head returns a IntSlice2 containing the first item from the aa. If aa is
// empty, the resulting IntSlice2 will be empty.
func IntSlice2Head(aa IntSlice2) IntSlice2 {
	if IntSlice2Empty(aa) {
		return IntSlice2{}
	}
	return IntSlice2{aa[0]}
}

// IntSlice2InsertAfter inserts an element in aa after the first element for which the
// supplied test function returns true. If none of the tests return true, the
// element is appended to the end of the aa.
func IntSlice2InsertAfter(aa *IntSlice2, b IntSlice, test Test) {
	var i int
	var a IntSlice
	for i, a = range *aa {
		if test(a) {
			break
		}
	}
	IntSlice2InsertAt(aa, b, int64(i+1))
}

// IntSlice2InsertBefore inserts an element in aa before the first element for which the
// supplied test function returns true. If none of the tests return true,
// the element is inserted at the head of aa.
func IntSlice2InsertBefore(aa *IntSlice2, b IntSlice, test Test) {
	var i int
	var a IntSlice
	for i, a = range *aa {
		if test(a) {
			break
		}
	}
	IntSlice2InsertAt(aa, b, int64(i-1))
}

// IntSlice2InsertAt inserts an element in aa at the specified index i, shifting the
// element originally at index i (and all subsequent elements) one position
// to the right. If i < 0, the element is inserted at index 0. If
// i >= len(aa), the value is appended to the end of aa.
func IntSlice2InsertAt(aa *IntSlice2, a IntSlice, i int64) {
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

// IntSlice2Intersection compares each element of aa to bb using the supplied equal
// function, and returns a IntSlice2 containing the elements which are common
// to both aa and bb. Duplicates are removed in this operation.
func IntSlice2Intersection(aa, bb IntSlice2, equality Equality) IntSlice2 {
	cc := IntSlice2{}
	IntSlice2ForEach(aa, func(a IntSlice) shared.Continue {
		IntSlice2ForEach(bb, func(b IntSlice) shared.Continue {
			if equality(a, b) && !IntSlice2Any(cc, func(c IntSlice) bool { return equality(a, c) }) {
				IntSlice2Append(&cc, a)
			}
			return shared.ContinueYes
		})
		return shared.ContinueYes
	})
	return cc
}

// IntSlice2IsProperSubset returns true if aa is a proper subset of bb.
// aa is considered a proper subset if all of its elements exist within bb, but
// bb also contains some elements that do not exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the IntSlice2Distinct
// operations to enforce uniqueness, if that is necessary.
func IntSlice2IsProperSubset(aa, bb IntSlice2, equality Equality) bool {
	aa1, bb1 := removeIntSlice2Intersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) > 0
}

// IntSlice2IsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the IntSlice2Distinct
// operations to enforce uniqueness, if that is necessary.
func IntSlice2IsProperSuperset(aa, bb IntSlice2, equality Equality) bool {
	aa1, bb1 := removeIntSlice2Intersections(aa, bb, equality)
	return len(aa1) > 0 && len(bb1) == 0
}

// IntSlice2IsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the IntSlice2Distinct
// operations to enforce uniqueness, if that is necessary.
func IntSlice2IsSubset(aa, bb IntSlice2, equality Equality) bool {
	aa1, bb1 := removeIntSlice2Intersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) >= 0
}

// IntSlice2IsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the IntSlice2Distinct
// operations to enforce uniqueness, if that is necessary.
func IntSlice2IsSuperset(aa, bb IntSlice2, equality Equality) bool {
	aa1, bb1 := removeIntSlice2Intersections(aa, bb, equality)
	return len(aa1) >= 0 && len(bb1) == 0
}

func removeIntSlice2Intersections(aa, bb IntSlice2, equality Equality) (IntSlice2, IntSlice2) {
	aa1 := IntSlice2Clone(aa)
	bb1 := IntSlice2Clone(bb)
	for ai := int64(len(aa1)) - 1; ai >= 0; ai-- {
		intersectionFound := false
		for bi := int64(len(bb1)) - 1; bi >= 0; bi-- {
			if equality((aa1)[ai], (bb1)[bi]) {
				intersectionFound = true
				IntSlice2IntSlice2RemoveAt(&bb1, bi)
			}
		}
		if intersectionFound {
			IntSlice2IntSlice2RemoveAt(&aa1, ai)
		}
	}
	return aa1, bb1
}

// IntSlice2Item returns a IntSlice2 containing the element at aa[i].
// If len(aa) == 0, i < 0, or, i >= len(aa), the resulting slice will be empty.
func IntSlice2Item(aa IntSlice2, i int64) IntSlice2 {
	if IntSlice2Empty(aa) || i < 0 || i >= int64(len(aa)) {
		return IntSlice2{}
	}
	return IntSlice2{aa[i]}
}

// IntSlice2IntSlice2ItemFuzzy returns a IntSlice2 containing the element at aa[i].
// If the supplied index is outside of the bounds of aa, IntSlice2IntSlice2ItemFuzzy will attempt
// to retrieve the head or end element of aa according to the following rules:
// If len(aa) == 0 an empty IntSlice2 is returned.
// If i < 0, the head of aa is returned.
// If i >= len(aa), the end of the aa is returned.
func IntSlice2IntSlice2ItemFuzzy(aa IntSlice2, i int64) IntSlice2 {
	if IntSlice2Empty(aa) {
		return IntSlice2{}
	}
	if i < 0 {
		return IntSlice2Head(aa)
	}
	if i >= int64(len(aa)) {
		return IntSlice2End(aa)
	}
	return IntSlice2{aa[i]}
}

// IntSlice2Last applies a test function to each element in aa, and returns a IntSlice2
// containing the last element for which the test returned true. If no elements
// pass the supplied test, the resulting IntSlice2 will be empty.
func IntSlice2Last(aa IntSlice2, test Test) IntSlice2 {
	bb := IntSlice2{}
	IntSlice2IntSlice2ForEachR(aa, func(a IntSlice) shared.Continue {
		if test(a) {
			IntSlice2Append(&bb, a)
			return shared.ContinueNo
		}
		return shared.ContinueYes
	})
	return bb
}

// IntSlice2Len returns the length of aa.
func IntSlice2Len(aa IntSlice2) int {
	return len(aa)
}

// IntSlice2Map applies a tranform to each element of the list.
func IntSlice2Map(aa *IntSlice2, mapFn func(IntSlice) IntSlice) {
	for i, a := range *aa {
		(*aa)[i] = mapFn(a)
	}
}

// IntSlice2None applies a test function to each element in aa, and returns true if
// the test function returns false for all items.
func IntSlice2None(aa IntSlice2, test Test) bool {
	return !IntSlice2Any(aa, test)
}

// IntSlice2Pairwise threads a transform function through aa, passing to the transform
// successive two-element pairs, aa[i-1] && aa[i]. For the first pairing
// the supplied init value is supplied as the initial element in the pair.
//
//   Illustration (pseudocode):
//     aa:  [W,X,Y,Z]
//     xform: func(a, b string) string { return a + b }
//     init: V
//     IntSlice2Pairwise(aa, init, xform) -> [VW, WX, XY, YZ]
func IntSlice2Pairwise(aa IntSlice2, init IntSlice, xform func(a, b IntSlice) IntSlice) IntSlice2 {
	if IntSlice2Empty(aa) {
		return IntSlice2{}
	}
	bb := IntSlice2{}
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

// IntSlice2Partition applies a test function to each element in aa, and returns
// a []IntSlice2 where []IntSlice2[0] contains a IntSlice2 with all elements for
// whom the test function returned true, and where []IntSlice2[1] contains a
// IntSlice2 with all elements for whom the test function returned false.
//
// IntSlice2Partition is a special case of the IntSlice2Group function.
func IntSlice2Partition(aa IntSlice2, test Test) []IntSlice2 {
	grouper := func(a IntSlice) int64 {
		if test(a) {
			return 1
		}
		return 0
	}
	return IntSlice2Group(aa, grouper)
}

// IntSlice2Permutable returns true if the number of permutations for aa exceeds
// MaxInt64.
func IntSlice2Permutable(aa IntSlice2) bool {
	return IntSlice2Permutations(aa).IsInt64()
}

// IntSlice2Permutations returns the number of permutations that exist given the current
// number of items in the aa.
func IntSlice2Permutations(aa IntSlice2) *big.Int {
	var f big.Int
	return f.MulRange(1, int64(len(aa)))
}

// IntSlice2Permute returns a []IntSlice2 which contains a IntSlice2 for each permutation
// of aa.
//
// This function will panic if it determines that the list is not permutable
// (see IntSlice2Permutable function).
//
// IntSlice2Permute makes no assumptions about whether or not the elements in aa are
// distinct. IntSlice2Permutations are created positionally, and do not involve any
// equality checks. As such, if it important that IntSlice2Permute operate on a set of
// distinct elements, pass aa through one of the IntSlice2Distinct transforms before
// passing it to IntSlice2Permute().
//
// IntSlice2Permute is implemented using Heap's algorithm.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func IntSlice2Permute(aa IntSlice2) []IntSlice2 {
	if IntSlice2Empty(aa) {
		return []IntSlice2{}
	}

	if !IntSlice2Permutable(aa) {
		panic(fmt.Sprintf("The number of permutations for this list (%v) exceeeds MaxInt64.", IntSlice2Permutations(aa)))
	}

	acc := []IntSlice2{}
	generate(int64(len(aa)), aa, &acc)
	return acc
}

func generate(n int64, aa IntSlice2, acc *[]IntSlice2) {
	if n == 1 {
		*acc = append(*acc, aa)
		return
	}

	for i := int64(0); i < n-1; i++ {
		generate(n-1, aa, acc)
		aa = IntSlice2Clone(aa)
		if n%2 != 0 {
			IntSlice2SwapIndex(aa, i, n-1)
		} else {
			IntSlice2SwapIndex(aa, 0, n-1)
		}
	}

	generate(n-1, aa, acc)
}

// IntSlice2Pop returns a IntSlice2 containing the head element from aa, and removes the
// element from aa. If aa is empty, the returned IntSlice2 will also be empty.
func IntSlice2Pop(aa *IntSlice2) IntSlice2 {
	bb := IntSlice2Head(*aa)
	IntSlice2IntSlice2RemoveAt(aa, 0)
	return bb
}

// IntSlice2Push places a prepends a new element at the head of aa.
func IntSlice2Push(aa *IntSlice2, a IntSlice) {
	IntSlice2InsertAt(aa, a, 0)
}

// IntSlice2Reduce applies a reducer function to each element in aa, threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new IntSlice2. If aa is empty, the resulting IntSlice2
// will also be empty.
//
//  Illustration:
//    aa: [1,2,3,4]
//    reducer: acc + sourceNode
//    IntSlice2Fold(aa, reducer) -> [10]
func IntSlice2Reduce(aa IntSlice2, reducer func(a, acc IntSlice) IntSlice) IntSlice2 {
	if len(aa) == 0 {
		return IntSlice2{}
	}
	accumulator := aa[0]
	if len(aa) > 1 {
		for i := 1; i < len(aa); i++ {
			accumulator = reducer(aa[i], accumulator)
		}
	}
	return IntSlice2{accumulator}
}

// IntSlice2Remove applies a test function to each item in the list, and removes any item
// for which the test returns true.
func IntSlice2Remove(aa *IntSlice2, test Test) {
	for i := int64(len(*aa)) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			IntSlice2IntSlice2RemoveAt(aa, i)
		}
	}
}

// IntSlice2IntSlice2RemoveAt removes the item at the specified index from the slice.
// If len(aa) == 0, aa == nil, i < 0, or i >= len(aa), this function will do
// nothing.
func IntSlice2IntSlice2RemoveAt(aa *IntSlice2, i int64) {
	if i < 0 || i >= int64(len(*aa)) {
		return
	}
	if len(*aa) > 0 {
		*aa = append((*aa)[:i], (*aa)[i+1:]...)
	}
}

// IntSlice2Reverse reverses the order of aa.
func IntSlice2Reverse(aa *IntSlice2) {
	for i := len(*aa)/2 - 1; i >= 0; i-- {
		j := len(*aa) - 1 - i
		(*aa)[i], (*aa)[j] = (*aa)[j], (*aa)[i]
	}
}

// IntSlice2Skip removes the first n elements from aa.
//
// Note that IntSlice2Skip(aa, len(aa)) will remove all items from the list, but does not
// "clear" the slice, meaning that the list remains allocated in memory.
// To fully de-pointer the slice, and ensure it is available for garbage
// collection as soon as possible, consider using IntSlice2Clear().
func IntSlice2Skip(aa *IntSlice2, n int64) {
	if len(*aa) == 0 {
		return
	}
	*aa = (*aa)[n:]
}

// IntSlice2IntSlice2SkipWhile scans through aa starting at the head, and removes all
// elements from aa while the test function returns true.
// IntSlice2IntSlice2SkipWhile stops removing any further items from aa after the first test that
// returns false.
func IntSlice2IntSlice2SkipWhile(aa *IntSlice2, test Test) {
	// find the first index where the test would evaluate to false and skip
	// everything up to that index.
	findTest := func(a IntSlice) bool { return !test(a) }
	IntSlice2Skip(aa, IntSlice2FindIndex(*aa, findTest))
}

// IntSlice2Sort sorts aa, using the supplied less function to determine order.
// IntSlice2Sort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func IntSlice2Sort(aa *IntSlice2, less func(a, b IntSlice) bool) {
	lessI := func(i, j int) bool {
		return less((*aa)[i], (*aa)[j])
	}
	sort.SliceStable(*aa, lessI)
}

// IntSlice2SplitAfter finds the first element b for which a test function returns true,
// and returns a []IntSlice2 where []IntSlice2[0] contains the first half of aa
// and []IntSlice2[1] contains the second half of aa. Element b will be included
// in []IntSlice2[0]. If the no element can be found for which the test returns
// true, []IntSlice2[0] will contain aa, and []IntSlice2[1] will be empty.
func IntSlice2SplitAfter(aa IntSlice2, test Test) []IntSlice2 {
	return IntSlice2SplitAt(aa, IntSlice2FindIndex(aa, test)+1)
}

// IntSlice2SplitAt splits aa at index i, and returns a []IntSlice2 which contains the
// two split halves of aa. aa[i] will be included in []IntSlice2[1].
// If i < 0, all of aa will be placed in []IntSlice2[0] and []IntSlice2[1] will
// be empty. Conversly, if i >= len(aa), all of aa will be placed in
// []IntSlice2[1] and []IntSlice2[0] will be empty. If aa is nil or empty,
// []IntSlice2 will contain two empty slices.
func IntSlice2SplitAt(aa IntSlice2, i int64) []IntSlice2 {
	if len(aa) == 0 {
		return []IntSlice2{
			IntSlice2{},
			IntSlice2{},
		}
	}
	if i < 0 {
		i = 0
	}
	return []IntSlice2{
		aa[:i],
		aa[i:],
	}
}

// IntSlice2SplitBefore finds the first element b for which a test function returns true,
// and returns a []IntSlice2 where []IntSlice2[0] contains the first half of aa
// and []IntSlice2[1] contains the second half of aa. Element b will be included
// in []IntSlice2[1]
func IntSlice2SplitBefore(aa IntSlice2, test Test) []IntSlice2 {
	return IntSlice2SplitAt(aa, IntSlice2FindIndex(aa, test))
}

// IntSlice2String returns a string representation of aa, suitable for use
// with fmt.Print, or other similar functions. IntSlice2String should be regarded as
// informational, and should not be relied upon to formally serialize a
// IntSlice2.
func IntSlice2String(aa IntSlice2) string {
	jsonBytes, _ := json.Marshal(aa)
	return string(jsonBytes)
}

// IntSlice2SwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of aa, IntSlice2SwapIndex does nothing.
func IntSlice2SwapIndex(aa IntSlice2, i, j int64) {
	l := int64(len(aa))
	if i < 0 || j < 0 || i >= l || j >= l {
		return
	}
	aa[i], aa[j] = aa[j], aa[i]
}

// IntSlice2Tail removes the current head element from aa.
// This equivelant to IntSlice2IntSlice2RemoveAt(aa, 0)
func IntSlice2Tail(aa *IntSlice2) {
	IntSlice2IntSlice2RemoveAt(aa, 0)
}

// IntSlice2Take retains the first n elements of aa, and removes all remaining elements
// from the slice. If n < 0 or n >= len(aa), IntSlice2Take does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func IntSlice2Take(aa *IntSlice2, n int64) {
	if len(*aa) == 0 || n < 0 || n >= int64(len(*aa)) {
		return
	}
	*aa = (*aa)[:n]
}

// IntSlice2IntSlice2TakeWhile applies a test function to each element in aa, and retains all
// elements of aa so long as the test function returns true. As soon as the test
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func IntSlice2IntSlice2TakeWhile(aa *IntSlice2, test Test) {
	find := func(a IntSlice) bool {
		return !test(a)
	}
	IntSlice2Take(aa, IntSlice2FindIndex(*aa, find))
}

// IntSlice2Union appends slice bb to slice aa.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func IntSlice2Union(aa *IntSlice2, bb IntSlice2) {
	IntSlice2Append(aa, bb...)
}

// IntSlice2Unzip splits aa into a []IntSlice2, such that []IntSlice2[0] contains all odd
// indices from aa, and []IntSlice2[1] contains all even indices from aa.
func IntSlice2Unzip(aa IntSlice2) []IntSlice2 {
	odds := IntSlice2{}
	evens := IntSlice2{}
	for i, a := range aa {
		if i%2 != 0 {
			odds = append(odds, a)
		} else {
			evens = append(evens, a)
		}
	}
	return []IntSlice2{odds, evens}
}

// IntSlice2WindowCentered applies a windowing function across the aa, using a centered
// window of the specified size.
func IntSlice2WindowCentered(aa IntSlice2, windowSize int64, windowFn func(window IntSlice2) IntSlice) IntSlice2 {
	cc := IntSlice2{}
	fullWindowReached := false
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := IntSlice2{}
		a := aa[i]
		for n := int64(1); n <= windowSize; n++ {
			IntSlice2Append(&currentWindow, a)
			if !fullWindowReached && n >= windowSize {
				fullWindowReached = true
			}
			if !fullWindowReached {
				IntSlice2Append(&cc, windowFn(currentWindow))
			}
			if i+n >= int64(len(aa)) {
				break
			}
			a = aa[i+n]
		}
		IntSlice2Append(&cc, windowFn(currentWindow))
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
	dd := IntSlice2(IntSlice2SplitAt(cc, frontTrim)[1])
	IntSlice2Reverse(&dd)
	ee := IntSlice2(IntSlice2SplitAt(dd, backTrim)[1])
	IntSlice2Reverse(&ee)
	return ee
}

// IntSlice2WindowLeft applies a windowing function across aa, using a left-sided window
// of the specified size.
func IntSlice2WindowLeft(aa IntSlice2, windowSize int64, windowFn func(window IntSlice2) IntSlice) IntSlice2 {
	bb := IntSlice2{}
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := IntSlice2{}
		for n := int64(0); n < windowSize; n++ {
			if i+n >= int64(len(aa)) {
				break
			}
			IntSlice2Append(&currentWindow, aa[i+n])
		}
		IntSlice2Append(&bb, windowFn(currentWindow))
	}
	return bb
}

// IntSlice2WindowRight applies a windowing function across aa, using a right-sided
// window of the specified size.
func IntSlice2WindowRight(aa IntSlice2, windowSize int64, windowFn func(window IntSlice2) IntSlice) IntSlice2 {
	aa1 := IntSlice2Clone(aa)
	defer IntSlice2Clear(&aa1)

	IntSlice2Reverse(&aa1)
	bb := IntSlice2{}
	for i := int64(0); i < int64(len(aa1)); i++ {
		currentWindow := IntSlice2{}
		for n := int64(0); n < windowSize; n++ {
			if i+n >= int64(len(aa1)) {
				break
			}
			IntSlice2Append(&currentWindow, aa1[i+n])
		}
		IntSlice2Reverse(&currentWindow)
		IntSlice2Append(&bb, windowFn(currentWindow))
	}
	IntSlice2Reverse(&bb)
	return bb
}

// IntSlice2Zip interleaves the contents of aa with bb, and returns the result as a
// new IntSlice2. aa[0] is evaluated first. Thus if aa and bb are the same
// length, slice aa will occupy the odd indices of the result slice, and bb
// will occupy the even indices of the result slice. If aa and bb are not
// the same length, IntSlice2Zip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func IntSlice2Zip(aa, bb IntSlice2) IntSlice2 {
	if len(aa) == 0 {
		return bb
	}
	if len(bb) == 0 {
		return aa
	}

	cc := IntSlice2{}
	aaIntSlice2EndReached, bbIntSlice2EndReached := false, false
	for i := 0; aaIntSlice2EndReached == false && bbIntSlice2EndReached == false; i++ {
		if i >= len(aa) {
			aaIntSlice2EndReached = true
		}
		if i >= len(bb) {
			bbIntSlice2EndReached = true
		}
		if i%2 != 0 {
			if !aaIntSlice2EndReached {
				IntSlice2Append(&cc, aa[i])
			}
			if !bbIntSlice2EndReached {
				IntSlice2Append(&cc, bb[i])
			}
		} else {
			if !bbIntSlice2EndReached {
				IntSlice2Append(&cc, bb[i])
			}
			if !aaIntSlice2EndReached {
				IntSlice2Append(&cc, aa[i])
			}
		}
	}
	return cc
}
