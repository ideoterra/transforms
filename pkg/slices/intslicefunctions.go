package slicexform

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

// IntSliceAll applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
func IntSliceAll(aa IntSlice, test Test) bool {
	for _, s := range aa {
		if !test(s) {
			return false
		}
	}
	return true
}

// IntSliceAny applies a test function to each element of the
// slice and returns true if the test function returns true for at least one
// item in the list.
//
// IntSliceAny does not require that the source slice be sorted, and merely scans
// the slice, returning as soon as any element passes the supplied test. For
// a binary search, consider using sort.Search from the standard library.
func IntSliceAny(aa IntSlice, test Test) bool {
	for _, a := range aa {
		if test(a) {
			return true
		}
	}
	return false
}

//IntSliceAppend adds the supplied values to the end of the slice.
func IntSliceAppend(aa *IntSlice, values ...int) {
	*aa = append(*aa, values...)
}

// IntSliceClear removes all of the items from the slice, setting the slice to nil
// such that any memory previously allocated to the slice can be garbage
// collected.
func IntSliceClear(aa *IntSlice) {
	*aa = nil
}

// IntSliceClone returns a copy of aa.
func IntSliceClone(aa IntSlice) IntSlice {
	return append(IntSlice{}, aa...)
}

// IntSliceCollect applies a given function against each item in slice aa and
// each item of a slice bb, and returns the concatenation of each result.
//
//   Illustration:
//     aa:  		[A, B, C]
//     bb: 			[X, Y, Z]
//     collector:   func(a, b) { return a + b }
//     IntSliceCollect(aa, bb, collector) -> [AX, AY, AZ, BX, BY, BZ, CX, XY, CZ]
func IntSliceCollect(aa IntSlice, bb IntSlice, collector func(a, b int) int) IntSlice {
	cc := IntSlice{}
	for _, a := range aa {
		for _, b := range bb {
			cc = append(cc, collector(a, b))
		}
	}
	return cc
}

// IntSliceCount applies the supplied test function to each element of the slice,
// and returns the count of items for which the test returns true.
func IntSliceCount(aa IntSlice, test Test) int64 {
	matches := int64(0)
	for _, a := range aa {
		if test(a) {
			matches++
		}
	}
	return matches
}

// IntSliceDequeue returns a IntSlice containing the head item from the source slice.
// The head item is removed from the source slice in this operation. If the
// source slice is initially empty, the resulting slice will also be empty.
func IntSliceDequeue(aa *IntSlice) IntSlice {
	if len(*aa) == 0 {
		return IntSlice{}
	}
	head := (*aa)[0]
	IntSliceIntSliceRemoveAt(aa, 0)
	return IntSlice{head}
}

// IntSliceDifference returns a new slice that contains items that are not common
// between aa and bb. The supplied equality function is used to compare values
// between each slice. Duplicates are retained through this process. As such,
// The elements in the slice that results from this transform may not be
// distinct. IntSliceDistinct values from aa are listed ahead of those from bb in the
// resulting slice.
//
// Illustration:
//   aa: [1,2,3,3,1,4]
//   bb: [5,4,3,5]
//   equal: func(a, b) bool {return a == b}
//   IntSliceDifference(aa, bb, equality) -> [1,2,1,5,5]
func IntSliceDifference(aa, bb IntSlice, equality Equality) IntSlice {
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

	cc := IntSlice{}
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

// IntSliceDistinct removes all duplicates from the slice, using the supplied equality
// function to determine equality.
func IntSliceDistinct(aa *IntSlice, equality Equality) {
	bb := IntSlice{}
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
	IntSliceClear(aa)
	IntSliceAppend(aa, bb...)
}

// IntSliceEmpty returns true if the length of the slice is zero.
func IntSliceEmpty(aa IntSlice) bool {
	return len(aa) == 0
}

// IntSliceEnd returns the a IntSlice containing only the last element from aa.
func IntSliceEnd(aa IntSlice) IntSlice {
	if IntSliceEmpty(aa) {
		return IntSlice{}
	}
	return IntSlice{aa[len(aa)-1]}
}

// IntSliceEnqueue places an item at the head of the slice.
func IntSliceEnqueue(aa *IntSlice, a int) {
	*aa = append(*aa, a)
	copy((*aa)[1:], (*aa)[:len(*aa)-1])
	(*aa)[0] = a
}

// IntSliceExpand applies an expansion function to each element of aa, and flattens
// the results into a single IntSlice.
//
//   Illustration (pseudocode):
//     aa: [AB, CD, EF]
//     expansion: func(a string) []string { return []string{a[0], a[1]}}
//     IntSliceExpand(aa, expansion) -> [A, B, C, D, E, F]
func IntSliceExpand(aa IntSlice, expansion func(int) IntSlice) IntSlice {
	bb := IntSlice{}
	for _, a := range aa {
		IntSliceAppend(&bb, expansion(a)...)
	}
	return bb
}

// IntSliceFilter removes all items from the slice for which the supplied test function
// returns true.
func IntSliceFilter(aa *IntSlice, test Test) {
	for i := len(*aa) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			IntSliceIntSliceRemoveAt(aa, int64(i))
		}
	}
}

// IntSliceFindIndex returns the index of the first element in the slice for which the
// supplied test function returns true. If no matches are found, -1 is returned.
func IntSliceFindIndex(aa IntSlice, test Test) int64 {
	for i, a := range aa {
		if test(a) {
			return int64(i)
		}
	}
	return -1
}

// IntSliceFirst returns a IntSlice containing the first element in the slice for which
// the supplied test function returns true.
func IntSliceFirst(aa IntSlice, test Test) IntSlice {
	bb := IntSlice{}
	for _, a := range aa {
		if test(a) {
			IntSliceAppend(&bb, a)
			break
		}
	}
	return bb
}

// IntSliceFold applies a function to each item in slice aa, threading an accumulator
// through each iteration. The accumulated value is returned in a new IntSlice
// once aa is fully scanned. IntSliceFold returns a IntSlice rather than a
// int to be consistent with this package's IntSliceReduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    IntSliceFold(aa, acc, folder) -> [11]
func IntSliceFold(aa IntSlice, acc int, folder func(a, acc int) int) IntSlice {
	return IntSliceIntSliceFoldI(aa, acc, func(_ int64, a, acc int) int { return folder(a, acc) })
}

// IntSliceIntSliceFoldI applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned. IntSliceFoldi returns a IntSlice rather than a
// int to be consistent with this package's IntSliceReduce implementation.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    IntSliceFold(aa, acc, folder) -> [11]
func IntSliceIntSliceFoldI(aa IntSlice, acc int, folder func(i int64, a, acc int) int) IntSlice {
	accumulation := acc
	for i, a := range aa {
		accumulation = folder(int64(i), a, accumulation)
	}
	return IntSlice{accumulation}
}

// IntSliceForEach applies each element of the list to the given function.
// IntSliceForEach will stop iterating if fn return false.
func IntSliceForEach(aa IntSlice, fn func(int) shared.Continue) {
	for _, a := range aa {
		if !fn(a) {
			return
		}
	}
}

// IntSliceIntSliceForEachC concurrently applies each element of the list to the given function.
// The elements of the list are marshalled to a pool of goroutines, where each
// element is passed to fn concurrently.
//
// The concurrency pool is limited to contain no more than c active goroutines
// at any time. Note that if a pool size of 0 is supplied, this method
// will block indefinitely. This function will panic if a negative value is
// supplied for c.
//
// If any execution of fn returns shared.ContinueNo, IntSliceIntSliceForEachC will cease marshalling
// any backlogged work, and will immediately set the cancellation flag to true.
// IntSliceAny goroutines monitoring the cancelPending closure can wind down their
// activities as necessary. IntSliceIntSliceForEachC will continue to block until all active
// goroutines exit cleanly.
func IntSliceIntSliceForEachC(aa IntSlice, c int, fn func(a int, cancelPending func() bool) shared.Continue) {
	if c < 0 {
		panic("IntSliceIntSliceForEachC: The concurrency pool size (c) must be non-negative.")
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
		go func(a int) {
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

// IntSliceIntSliceForEachR applies each element of aa to a given function, scanning
// through the slice in reverse order, starting from the end and working towards
// the head.
func IntSliceIntSliceForEachR(aa IntSlice, fn func(int) shared.Continue) {
	for i := len(aa) - 1; i >= 0; i-- {
		if !fn(aa[i]) {
			return
		}
	}
}

// IntSliceGroup consolidates like-items into groups according to the supplied grouper
// function, and returns them as a IntSlice2.
// The grouper function is expected to return a hash value which IntSliceGroup will use
// to determine into which bucket each element wil be placed.
func IntSliceGroup(aa IntSlice, grouper func(int) int64) IntSlice2 {
	return IntSliceIntSliceGroupI(aa, func(_ int64, a int) int64 { return grouper(a) })
}

// IntSliceIntSliceGroupI consolidates like-items into groups according to the supplied grouper
// function, and returns them as a IntSlice2.
// The grouper function is expected to return a hash value which IntSliceGroup will use
// to determine into which bucket each element wil be placed. For convenience
// the index value from aa is also passed into the grouper function.
func IntSliceIntSliceGroupI(aa IntSlice, grouper func(int64, int) int64) IntSlice2 {
	groupIntSliceMap := map[int64]IntSlice{}
	for i, a := range aa {
		hash := grouper(int64(i), a)
		if _, exists := groupIntSliceMap[hash]; exists {
			groupIntSliceMap[hash] = append(groupIntSliceMap[hash], a)
		} else {
			groupIntSliceMap[hash] = IntSlice{a}
		}
	}
	group := IntSlice2{}
	for _, bb := range groupIntSliceMap {
		group = append(group, bb)
	}
	return group
}

// IntSliceHead returns a IntSlice containing the first item from the aa. If aa is
// empty, the resulting IntSlice will be empty.
func IntSliceHead(aa IntSlice) IntSlice {
	if IntSliceEmpty(aa) {
		return IntSlice{}
	}
	return IntSlice{aa[0]}
}

// IntSliceInsertAfter inserts an element in aa after the first element for which the
// supplied test function returns true. If none of the tests return true, the
// element is appended to the end of the aa.
func IntSliceInsertAfter(aa *IntSlice, b int, test Test) {
	var i int
	var a int
	for i, a = range *aa {
		if test(a) {
			break
		}
	}
	IntSliceInsertAt(aa, b, int64(i+1))
}

// IntSliceInsertBefore inserts an element in aa before the first element for which the
// supplied test function returns true. If none of the tests return true,
// the element is inserted at the head of aa.
func IntSliceInsertBefore(aa *IntSlice, b int, test Test) {
	var i int
	var a int
	for i, a = range *aa {
		if test(a) {
			break
		}
	}
	IntSliceInsertAt(aa, b, int64(i-1))
}

// IntSliceInsertAt inserts an element in aa at the specified index i, shifting the
// element originally at index i (and all subsequent elements) one position
// to the right. If i < 0, the element is inserted at index 0. If
// i >= len(aa), the value is appended to the end of aa.
func IntSliceInsertAt(aa *IntSlice, a int, i int64) {
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

// IntSliceIntersection compares each element of aa to bb using the supplied equal
// function, and returns a IntSlice containing the elements which are common
// to both aa and bb. Duplicates are removed in this operation.
func IntSliceIntersection(aa, bb IntSlice, equality Equality) IntSlice {
	cc := IntSlice{}
	IntSliceForEach(aa, func(a int) shared.Continue {
		IntSliceForEach(bb, func(b int) shared.Continue {
			if equality(a, b) && !IntSliceAny(cc, func(c int) bool { return equality(a, c) }) {
				IntSliceAppend(&cc, a)
			}
			return shared.ContinueYes
		})
		return shared.ContinueYes
	})
	return cc
}

// IntSliceIsProperSubset returns true if aa is a proper subset of bb.
// aa is considered a proper subset if all of its elements exist within bb, but
// bb also contains some elements that do not exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the IntSliceDistinct
// operations to enforce uniqueness, if that is necessary.
func IntSliceIsProperSubset(aa, bb IntSlice, equality Equality) bool {
	aa1, bb1 := removeIntSliceIntersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) > 0
}

// IntSliceIsProperSuperset returns true if aa is a proper superset of bb.
// aa is considered a proper superset if it contains all of bb's elements, but
// aa also contains some elements that do not exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the IntSliceDistinct
// operations to enforce uniqueness, if that is necessary.
func IntSliceIsProperSuperset(aa, bb IntSlice, equality Equality) bool {
	aa1, bb1 := removeIntSliceIntersections(aa, bb, equality)
	return len(aa1) > 0 && len(bb1) == 0
}

// IntSliceIsSubset returns true if aa is a subset of bb.
// aa is considered a subset if all of its elements exist within bb.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a subset to be larger than its superset. Use the IntSliceDistinct
// operations to enforce uniqueness, if that is necessary.
func IntSliceIsSubset(aa, bb IntSlice, equality Equality) bool {
	aa1, bb1 := removeIntSliceIntersections(aa, bb, equality)
	return len(aa1) == 0 && len(bb1) >= 0
}

// IntSliceIsSuperset returns true if aa is a superset of bb.
// aa is considered a superset if all of bb's elements exist within aa.
// Note: This operation does not enforce that each element be unique, thus, it
// is possible for a superset to be smaller than its subset. Use the IntSliceDistinct
// operations to enforce uniqueness, if that is necessary.
func IntSliceIsSuperset(aa, bb IntSlice, equality Equality) bool {
	aa1, bb1 := removeIntSliceIntersections(aa, bb, equality)
	return len(aa1) >= 0 && len(bb1) == 0
}

func removeIntSliceIntersections(aa, bb IntSlice, equality Equality) (IntSlice, IntSlice) {
	aa1 := IntSliceClone(aa)
	bb1 := IntSliceClone(bb)
	for ai := int64(len(aa1)) - 1; ai >= 0; ai-- {
		intersectionFound := false
		for bi := int64(len(bb1)) - 1; bi >= 0; bi-- {
			if equality((aa1)[ai], (bb1)[bi]) {
				intersectionFound = true
				IntSliceIntSliceRemoveAt(&bb1, bi)
			}
		}
		if intersectionFound {
			IntSliceIntSliceRemoveAt(&aa1, ai)
		}
	}
	return aa1, bb1
}

// IntSliceItem returns a IntSlice containing the element at aa[i].
// If len(aa) == 0, i < 0, or, i >= len(aa), the resulting slice will be empty.
func IntSliceItem(aa IntSlice, i int64) IntSlice {
	if IntSliceEmpty(aa) || i < 0 || i >= int64(len(aa)) {
		return IntSlice{}
	}
	return IntSlice{aa[i]}
}

// IntSliceIntSliceItemFuzzy returns a IntSlice containing the element at aa[i].
// If the supplied index is outside of the bounds of aa, IntSliceIntSliceItemFuzzy will attempt
// to retrieve the head or end element of aa according to the following rules:
// If len(aa) == 0 an empty IntSlice is returned.
// If i < 0, the head of aa is returned.
// If i >= len(aa), the end of the aa is returned.
func IntSliceIntSliceItemFuzzy(aa IntSlice, i int64) IntSlice {
	if IntSliceEmpty(aa) {
		return IntSlice{}
	}
	if i < 0 {
		return IntSliceHead(aa)
	}
	if i >= int64(len(aa)) {
		return IntSliceEnd(aa)
	}
	return IntSlice{aa[i]}
}

// IntSliceLast applies a test function to each element in aa, and returns a IntSlice
// containing the last element for which the test returned true. If no elements
// pass the supplied test, the resulting IntSlice will be empty.
func IntSliceLast(aa IntSlice, test Test) IntSlice {
	bb := IntSlice{}
	IntSliceIntSliceForEachR(aa, func(a int) shared.Continue {
		if test(a) {
			IntSliceAppend(&bb, a)
			return shared.ContinueNo
		}
		return shared.ContinueYes
	})
	return bb
}

// IntSliceLen returns the length of aa.
func IntSliceLen(aa IntSlice) int {
	return len(aa)
}

// IntSliceMap applies a tranform to each element of the list.
func IntSliceMap(aa *IntSlice, mapFn func(int) int) {
	for i, a := range *aa {
		(*aa)[i] = mapFn(a)
	}
}

// IntSliceNone applies a test function to each element in aa, and returns true if
// the test function returns false for all items.
func IntSliceNone(aa IntSlice, test Test) bool {
	return !IntSliceAny(aa, test)
}

// IntSlicePairwise threads a transform function through aa, passing to the transform
// successive two-element pairs, aa[i-1] && aa[i]. For the first pairing
// the supplied init value is supplied as the initial element in the pair.
//
//   Illustration (pseudocode):
//     aa:  [W,X,Y,Z]
//     xform: func(a, b string) string { return a + b }
//     init: V
//     IntSlicePairwise(aa, init, xform) -> [VW, WX, XY, YZ]
func IntSlicePairwise(aa IntSlice, init int, xform func(a, b int) int) IntSlice {
	if IntSliceEmpty(aa) {
		return IntSlice{}
	}
	bb := IntSlice{}
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

// IntSlicePartition applies a test function to each element in aa, and returns
// a IntSlice2 where IntSlice2[0] contains a IntSlice with all elements for
// whom the test function returned true, and where IntSlice2[1] contains a
// IntSlice with all elements for whom the test function returned false.
//
// IntSlicePartition is a special case of the IntSliceGroup function.
func IntSlicePartition(aa IntSlice, test Test) IntSlice2 {
	grouper := func(a int) int64 {
		if test(a) {
			return 1
		}
		return 0
	}
	return IntSliceGroup(aa, grouper)
}

// IntSlicePermutable returns true if the number of permutations for aa exceeds
// MaxInt64.
func IntSlicePermutable(aa IntSlice) bool {
	return IntSlicePermutations(aa).IsInt64()
}

// IntSlicePermutations returns the number of permutations that exist given the current
// number of items in the aa.
func IntSlicePermutations(aa IntSlice) *big.Int {
	var f big.Int
	return f.MulRange(1, int64(len(aa)))
}

// IntSlicePermute returns a IntSlice2 which contains a IntSlice for each permutation
// of aa.
//
// This function will panic if it determines that the list is not permutable
// (see IntSlicePermutable function).
//
// IntSlicePermute makes no assumptions about whether or not the elements in aa are
// distinct. IntSlicePermutations are created positionally, and do not involve any
// equality checks. As such, if it important that IntSlicePermute operate on a set of
// distinct elements, pass aa through one of the IntSliceDistinct transforms before
// passing it to IntSlicePermute().
//
// IntSlicePermute is implemented using Heap's algorithm.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func IntSlicePermute(aa IntSlice) IntSlice2 {
	if IntSliceEmpty(aa) {
		return IntSlice2{}
	}

	if !IntSlicePermutable(aa) {
		panic(fmt.Sprintf("The number of permutations for this list (%v) exceeeds MaxInt64.", IntSlicePermutations(aa)))
	}

	acc := IntSlice2{}
	generate(int64(len(aa)), aa, &acc)
	return acc
}

func generate(n int64, aa IntSlice, acc *IntSlice2) {
	if n == 1 {
		*acc = append(*acc, aa)
		return
	}

	for i := int64(0); i < n-1; i++ {
		generate(n-1, aa, acc)
		aa = IntSliceClone(aa)
		if n%2 != 0 {
			IntSliceSwapIndex(aa, i, n-1)
		} else {
			IntSliceSwapIndex(aa, 0, n-1)
		}
	}

	generate(n-1, aa, acc)
}

// IntSlicePop returns a IntSlice containing the head element from aa, and removes the
// element from aa. If aa is empty, the returned IntSlice will also be empty.
func IntSlicePop(aa *IntSlice) IntSlice {
	bb := IntSliceHead(*aa)
	IntSliceIntSliceRemoveAt(aa, 0)
	return bb
}

// IntSlicePush places a prepends a new element at the head of aa.
func IntSlicePush(aa *IntSlice, a int) {
	IntSliceInsertAt(aa, a, 0)
}

// IntSliceReduce applies a reducer function to each element in aa, threading an
// accumulator through each iteration. The resulting accumulation is returned
// as an element of a new IntSlice. If aa is empty, the resulting IntSlice
// will also be empty.
//
//  Illustration:
//    aa: [1,2,3,4]
//    reducer: acc + sourceNode
//    IntSliceFold(aa, reducer) -> [10]
func IntSliceReduce(aa IntSlice, reducer func(a, acc int) int) IntSlice {
	if len(aa) == 0 {
		return IntSlice{}
	}
	accumulator := aa[0]
	if len(aa) > 1 {
		for i := 1; i < len(aa); i++ {
			accumulator = reducer(aa[i], accumulator)
		}
	}
	return IntSlice{accumulator}
}

// IntSliceRemove applies a test function to each item in the list, and removes any item
// for which the test returns true.
func IntSliceRemove(aa *IntSlice, test Test) {
	for i := int64(len(*aa)) - 1; i >= 0; i-- {
		if test((*aa)[i]) {
			IntSliceIntSliceRemoveAt(aa, i)
		}
	}
}

// IntSliceIntSliceRemoveAt removes the item at the specified index from the slice.
// If len(aa) == 0, aa == nil, i < 0, or i >= len(aa), this function will do
// nothing.
func IntSliceIntSliceRemoveAt(aa *IntSlice, i int64) {
	if i < 0 || i >= int64(len(*aa)) {
		return
	}
	if len(*aa) > 0 {
		*aa = append((*aa)[:i], (*aa)[i+1:]...)
	}
}

// IntSliceReverse reverses the order of aa.
func IntSliceReverse(aa *IntSlice) {
	for i := len(*aa)/2 - 1; i >= 0; i-- {
		j := len(*aa) - 1 - i
		(*aa)[i], (*aa)[j] = (*aa)[j], (*aa)[i]
	}
}

// IntSliceSkip removes the first n elements from aa.
//
// Note that IntSliceSkip(aa, len(aa)) will remove all items from the list, but does not
// "clear" the slice, meaning that the list remains allocated in memory.
// To fully de-pointer the slice, and ensure it is available for garbage
// collection as soon as possible, consider using IntSliceClear().
func IntSliceSkip(aa *IntSlice, n int64) {
	if len(*aa) == 0 {
		return
	}
	*aa = (*aa)[n:]
}

// IntSliceIntSliceSkipWhile scans through aa starting at the head, and removes all
// elements from aa while the test function returns true.
// IntSliceIntSliceSkipWhile stops removing any further items from aa after the first test that
// returns false.
func IntSliceIntSliceSkipWhile(aa *IntSlice, test Test) {
	// find the first index where the test would evaluate to false and skip
	// everything up to that index.
	findTest := func(a int) bool { return !test(a) }
	IntSliceSkip(aa, IntSliceFindIndex(*aa, findTest))
}

// IntSliceSort sorts aa, using the supplied less function to determine order.
// IntSliceSort is a convenience wrapper around the stdlib sort.SliceStable
// function.
func IntSliceSort(aa *IntSlice, less func(a, b int) bool) {
	lessI := func(i, j int) bool {
		return less((*aa)[i], (*aa)[j])
	}
	sort.SliceStable(*aa, lessI)
}

// IntSliceSplitAfter finds the first element b for which a test function returns true,
// and returns a IntSlice2 where IntSlice2[0] contains the first half of aa
// and IntSlice2[1] contains the second half of aa. Element b will be included
// in IntSlice2[0]. If the no element can be found for which the test returns
// true, IntSlice2[0] will contain aa, and IntSlice2[1] will be empty.
func IntSliceSplitAfter(aa IntSlice, test Test) IntSlice2 {
	return IntSliceSplitAt(aa, IntSliceFindIndex(aa, test)+1)
}

// IntSliceSplitAt splits aa at index i, and returns a IntSlice2 which contains the
// two split halves of aa. aa[i] will be included in IntSlice2[1].
// If i < 0, all of aa will be placed in IntSlice2[0] and IntSlice2[1] will
// be empty. Conversly, if i >= len(aa), all of aa will be placed in
// IntSlice2[1] and IntSlice2[0] will be empty. If aa is nil or empty,
// IntSlice2 will contain two empty slices.
func IntSliceSplitAt(aa IntSlice, i int64) IntSlice2 {
	if len(aa) == 0 {
		return IntSlice2{
			IntSlice{},
			IntSlice{},
		}
	}
	if i < 0 {
		i = 0
	}
	return IntSlice2{
		aa[:i],
		aa[i:],
	}
}

// IntSliceSplitBefore finds the first element b for which a test function returns true,
// and returns a IntSlice2 where IntSlice2[0] contains the first half of aa
// and IntSlice2[1] contains the second half of aa. Element b will be included
// in IntSlice2[1]
func IntSliceSplitBefore(aa IntSlice, test Test) IntSlice2 {
	return IntSliceSplitAt(aa, IntSliceFindIndex(aa, test))
}

// IntSliceString returns a string representation of aa, suitable for use
// with fmt.Print, or other similar functions. IntSliceString should be regarded as
// informational, and should not be relied upon to formally serialize a
// IntSlice.
func IntSliceString(aa IntSlice) string {
	jsonBytes, _ := json.Marshal(aa)
	return string(jsonBytes)
}

// IntSliceSwapIndex swaps the elements at the specified indices. If either i or j is
// out of the bounds of aa, IntSliceSwapIndex does nothing.
func IntSliceSwapIndex(aa IntSlice, i, j int64) {
	l := int64(len(aa))
	if i < 0 || j < 0 || i >= l || j >= l {
		return
	}
	aa[i], aa[j] = aa[j], aa[i]
}

// IntSliceTail removes the current head element from aa.
// This equivelant to IntSliceIntSliceRemoveAt(aa, 0)
func IntSliceTail(aa *IntSlice) {
	IntSliceIntSliceRemoveAt(aa, 0)
}

// IntSliceTake retains the first n elements of aa, and removes all remaining elements
// from the slice. If n < 0 or n >= len(aa), IntSliceTake does nothing. If n == 0, all
// elements are removed from the slice (but the slice is not de-pointered).
func IntSliceTake(aa *IntSlice, n int64) {
	if len(*aa) == 0 || n < 0 || n >= int64(len(*aa)) {
		return
	}
	*aa = (*aa)[:n]
}

// IntSliceIntSliceTakeWhile applies a test function to each element in aa, and retains all
// elements of aa so long as the test function returns true. As soon as the test
// function returns false, take stops evaluating any further, and abandons the
// rest of the slice.
func IntSliceIntSliceTakeWhile(aa *IntSlice, test Test) {
	find := func(a int) bool {
		return !test(a)
	}
	IntSliceTake(aa, IntSliceFindIndex(*aa, find))
}

// IntSliceUnion appends slice bb to slice aa.
// Note: This operation does not remove any duplicates from the slice, as a
// similar operation would when operating on a formal Set.
func IntSliceUnion(aa *IntSlice, bb IntSlice) {
	IntSliceAppend(aa, bb...)
}

// IntSliceUnzip splits aa into a IntSlice2, such that IntSlice2[0] contains all odd
// indices from aa, and IntSlice2[1] contains all even indices from aa.
func IntSliceUnzip(aa IntSlice) IntSlice2 {
	odds := IntSlice{}
	evens := IntSlice{}
	for i, a := range aa {
		if i%2 != 0 {
			odds = append(odds, a)
		} else {
			evens = append(evens, a)
		}
	}
	return IntSlice2{odds, evens}
}

// IntSliceWindowCentered applies a windowing function across the aa, using a centered
// window of the specified size.
func IntSliceWindowCentered(aa IntSlice, windowSize int64, windowFn func(window IntSlice) int) IntSlice {
	cc := IntSlice{}
	fullWindowReached := false
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := IntSlice{}
		a := aa[i]
		for n := int64(1); n <= windowSize; n++ {
			IntSliceAppend(&currentWindow, a)
			if !fullWindowReached && n >= windowSize {
				fullWindowReached = true
			}
			if !fullWindowReached {
				IntSliceAppend(&cc, windowFn(currentWindow))
			}
			if i+n >= int64(len(aa)) {
				break
			}
			a = aa[i+n]
		}
		IntSliceAppend(&cc, windowFn(currentWindow))
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
	dd := IntSlice(IntSliceSplitAt(cc, frontTrim)[1])
	IntSliceReverse(&dd)
	ee := IntSlice(IntSliceSplitAt(dd, backTrim)[1])
	IntSliceReverse(&ee)
	return ee
}

// IntSliceWindowLeft applies a windowing function across aa, using a left-sided window
// of the specified size.
func IntSliceWindowLeft(aa IntSlice, windowSize int64, windowFn func(window IntSlice) int) IntSlice {
	bb := IntSlice{}
	for i := int64(0); i < int64(len(aa)); i++ {
		currentWindow := IntSlice{}
		for n := int64(0); n < windowSize; n++ {
			if i+n >= int64(len(aa)) {
				break
			}
			IntSliceAppend(&currentWindow, aa[i+n])
		}
		IntSliceAppend(&bb, windowFn(currentWindow))
	}
	return bb
}

// IntSliceWindowRight applies a windowing function across aa, using a right-sided
// window of the specified size.
func IntSliceWindowRight(aa IntSlice, windowSize int64, windowFn func(window IntSlice) int) IntSlice {
	aa1 := IntSliceClone(aa)
	defer IntSliceClear(&aa1)

	IntSliceReverse(&aa1)
	bb := IntSlice{}
	for i := int64(0); i < int64(len(aa1)); i++ {
		currentWindow := IntSlice{}
		for n := int64(0); n < windowSize; n++ {
			if i+n >= int64(len(aa1)) {
				break
			}
			IntSliceAppend(&currentWindow, aa1[i+n])
		}
		IntSliceReverse(&currentWindow)
		IntSliceAppend(&bb, windowFn(currentWindow))
	}
	IntSliceReverse(&bb)
	return bb
}

// IntSliceZip interleaves the contents of aa with bb, and returns the result as a
// new IntSlice. aa[0] is evaluated first. Thus if aa and bb are the same
// length, slice aa will occupy the odd indices of the result slice, and bb
// will occupy the even indices of the result slice. If aa and bb are not
// the same length, IntSliceZip will interleave as many values as possible, and will
// simply append the remaining values for the longer of the two slices to the
// end of the result slice.
func IntSliceZip(aa, bb IntSlice) IntSlice {
	if len(aa) == 0 {
		return bb
	}
	if len(bb) == 0 {
		return aa
	}

	cc := IntSlice{}
	aaIntSliceEndReached, bbIntSliceEndReached := false, false
	for i := 0; aaIntSliceEndReached == false && bbIntSliceEndReached == false; i++ {
		if i >= len(aa) {
			aaIntSliceEndReached = true
		}
		if i >= len(bb) {
			bbIntSliceEndReached = true
		}
		if i%2 != 0 {
			if !aaIntSliceEndReached {
				IntSliceAppend(&cc, aa[i])
			}
			if !bbIntSliceEndReached {
				IntSliceAppend(&cc, bb[i])
			}
		} else {
			if !bbIntSliceEndReached {
				IntSliceAppend(&cc, bb[i])
			}
			if !aaIntSliceEndReached {
				IntSliceAppend(&cc, aa[i])
			}
		}
	}
	return cc
}
