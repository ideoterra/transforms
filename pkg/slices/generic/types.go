package generic

// Equality is a function that evalutes if two values share something in common.
type Equality func(a, b PrimitiveType) bool

// Grouper is a function that returns a hash value for a supplied PrimitiveType.
type Grouper func(PrimitiveType) int64

// SliceType is a one dimensional slice of PrimitiveType.
type SliceType []PrimitiveType

// SliceType2 is a two dimensional slice of PrimitiveType
type SliceType2 []SliceType

// Test is a function that conditionally evaluates a PrimitiveType.
type Test func(PrimitiveType) bool
