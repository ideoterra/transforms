package slicexform

// Equality is a function that evalutes if two values share something in common.
type Equality func(a, b IntSlice) bool

// IntSlice2 is a one dimensional slice of IntSlice.
type IntSlice2 []IntSlice




// Test is a function that conditionally evaluates a IntSlice.
type Test func(IntSlice) bool
