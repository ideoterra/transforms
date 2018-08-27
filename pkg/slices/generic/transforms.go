// Package generic contains transform functions for SliceTypes.
//
// By convention, the source slice will be named `aa`. If multiple slices are
// to be supplied as arguments to a function, they are named `aa`, `bb`, `cc`,
// and so on.
//
// Transforms that mutate the supplied sources will always require the slice
// as a pointer. If a function is operating on a slice value, it will not mutate
// the underlaying slice. However, if a function requires that a slice be passed
// as a pointer, it can be expected that the function is mutating the
// underlaying slice.
//
// Transforms that reduce a result to a single value (such as Dequeue, or Fold)
// return a SliceType containing a single element rather than a PrimitiveType.
// This is done to avoid edge cases associated with applying transformations
// on empty lists or that result in an empty value. There are generally three
// options for how to handle an empty result. 1) Have the result be a pointer
// rather than a 'value', and set the result to nil if there is no result.
// 2) If there is no result, just return a zero value. 3) Return the result as
// an empty slice if there is no result. Option 1 (returning nil) can be
// confusing if the underlaying slice contains pointers such as `[]*struct{}`
// in this case, it would be difficult to differentiate between a nil that is
// returned because the slice's head contained a nil pointer vs the slice being
// initially empty. A similar issue exists for option 2 (return a zero-value).
// Returning a zero value has a different implication from returning no value.
// As such, it seemed to make sense to just return a slice in all circumstances.
// If the slice is empty, we know there was no result returned, and confusion is
// avoided.
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

// Collect applies a given function against each item in the source list and
// each item of a second list, and returns the concatenation of each result.
//
//   Illustration:
//     source:     [A, B, C]
//     otherList:  [X, Y, Z]
//     collector:  sourceNode + listANode
//     source.Collect(otherList, collector) -> [AX, AY, AZ, BX, BY, BZ, CX, XY, CZ]
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

// // Difference returns a new slice that contains items that are not common
// // between slice aa and slice bb.
// // The optional equality function is used to compare each item in the slices.
// // If no equality function is supplied, the standard infix equality operator
// // `==` is applied to each element of the two slices.
// func Difference(aa, bb SliceType, equal ...Equality) SliceType {
// 	// seems like we should have two different types of generic transforms.
// 	// numeric transforms, where the type has a standard equality operator,
// 	// and non-numeric transforms, where the type does not have a standard equality
// 	// operator
// 	//
// 	// Another approach is to presume that no standard comparisons exist
// 	//
// 	// Yet another approach could be to supply a set of standard equality functions
// 	// That can be supplied if the user doesn't want to write their own. I kindof
// 	// like this option. As such, we will need to generate a set of equality
// 	// generic equality functions, one for each primitive type.
// }

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
