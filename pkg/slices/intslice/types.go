package intslice

// Equality is a function that evalutes if two values share something in common.
type Equality func(a, b int) bool

// IntSlice is a one dimensional slice of int.
type IntSlice []int

// Test is a function that conditionally evaluates a int.
type Test func(int) bool
