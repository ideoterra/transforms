// Package generic contains transform functions for SliceTypes.
//
// Parameter Naming:
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
package generic

import (
	"github.com/cheekybits/genny/generic"
)

// SliceType is a placeholder for a generic slice.
type SliceType []PrimitiveType

// PrimitiveType is a placeholder for the type underpinning the generic SliceType.
type PrimitiveType generic.Type

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

// First returns a SliceType contaiing the first element in the slice for which
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
// through each iteration. The accumulated value is returned once aa is fully
// scanned.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func Fold(aa SliceType, acc PrimitiveType, folder func(a, acc PrimitiveType) PrimitiveType) PrimitiveType {
	return Foldi(aa, acc, func(_ int64, a, acc PrimitiveType) PrimitiveType { return folder(a, acc) })
}

// Foldi applies a function to each item in slice aa, threading an accumulator
// and an index value through each iteration. The accumulated value is returned
// once aa is fully scanned.
//
//  Illustration:
//    aa: [1,2,3,4]
//    acc:    1
//    folder: acc + sourceNode
//    Fold(aa, acc, folder) -> [11]
func Foldi(aa SliceType, acc PrimitiveType, folder func(i int64, a, acc PrimitiveType) PrimitiveType) PrimitiveType {
	accumulation := acc
	for i, a := range aa {
		accumulation = folder(int64(i), a, accumulation)
	}
	return accumulation
}

// ForEach applies each element of the list to the given function.
func ForEach(aa SliceType, fn func(PrimitiveType)) {
	for _, a := range aa {
		fn(a)
	}
}

//Remove applies a test function to each item in the list, and removes all items
//for which the test returns true.
// func Remove(aa *SliceType, test func(PrimitiveType) bool) {
// 	originalListLength := int64(len(*aa))
// 	for i := originalListLength - 1; i >= 0; i-- {
// 		if test((*aa)[i]) {
// 			RemoveAt(aa, i)
// 		}
// 	}
// }

//RemoveAt removes the item at the specified index from the slice.
func RemoveAt(aa *SliceType, i int64) {
	*aa = append((*aa)[:i], (*aa)[i+1:]...)
}
