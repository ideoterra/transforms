package intslice

// Equality is a function that evalutes if two values share something in common.
type Equality func(a, b int) bool

// IntSlice is a one dimensional slice of int.
type IntSlice []int

// IntSlice2 is a two dimensional slice of int
type IntSlice2 []IntSlice

// Test is a function that conditionally evaluates a int.
type Test func(int) bool
