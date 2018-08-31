package intslice2

import (
	"github.com/jecolasurdo/transforms/pkg/slices/intslice"
)

// Equality is a function that evalutes if two values share something in common.
type Equality func(a, b intslice.IntSlice) bool

// IntSlice2 is a one dimensional slice of intslice.IntSlice.
type IntSlice2 []intslice.IntSlice

// Test is a function that conditionally evaluates a intslice.IntSlice.
type Test func(intslice.IntSlice) bool
