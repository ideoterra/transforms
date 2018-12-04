// Package generic contains transform functions for SliceTypes.
//
// Function Naming Conventions:
// Often the same conceptual function can be implemented in more than one way.
// When multiple variants of a function are implemented in this package, each
// variant will start with the same base name, and the name will be suffixed
// with 0 or more characters that indicate details about the implementation.
// Such suffixes are as follows:
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
// values can be expected to be immutable.
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
// possible. In cases where performance can be improved using a sorted dataset,
// alternate function variants are provided, such as in `AnyS()`, and
// `DifferenceS())`. Bear in mind that passing unsorted data to a function
// variant that expects sorted data, will likely result in an incorrect result.
//
// Equality functions:
// Transforms that need to test the equality of slice elements are intentionaly
// left naive, and do not make any assumptions about how to test for equality.
// As a result, functions such as `Difference()` require an equality function
// to be supplied. For primitive types, typical equality functions are provided
// in the `eq` and `less` packages. It is encouraged to use the supplied
// equality functions for primitive types.
//
// Inclusion of non-native slice operations:
// This package provides functions independent of the underlaying data
// structure. That is, even though the base data structure for all operations in
// this package is the slice, operations for other common data structures are
// provided for convenience. For example, several common set operations such as
// Intersection() are provided. Stack and queue operations such as Push() and
// Dequeue() are also provided, as well as many others.
// While the implementations of these "non-native" operations are provided for
// convenience, that can come at the cost of performance because the underlaying
// datastructure simply may not be optimized for the operation being performed.
// With that said, efforts are made to provide reasonably performant
// implementations for all operations. Similarly, because a slice structure is
// fundamentally different from other data structures, some behavior for
// non-native operations (such as the set operations) are implemented more
// loosely than they would be for their native types. For instance, Sets, by
// definition, do not contain duplicates. However, there is no such requirement
// in a slice. As such, the set-type operations provided in this package are
// allowed to be tollerant to duplicates. These differences are noted in the
// description for each method, as warranted.package generic
package generic
