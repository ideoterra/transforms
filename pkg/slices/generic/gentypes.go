package generic

// This file contains types that are used by the generic package only, and are
// removed by generate.go for generated packages.

// PrimitiveType is the type for SliceType elements.
type PrimitiveType interface{}

// PrimitiveTypeA is a PrimitiveType used for testing.
type PrimitiveTypeA interface{}

// PrimitiveTypeB is a PrimitiveType used for testing.
type PrimitiveTypeB interface{}

// SliceTypeA is a one dimensional slice of PrimitiveTypeA used for testing.
type SliceTypeA []PrimitiveTypeA

// SliceTypeB is a one dimensional slice of PrimitiveTypeB used for testing.
type SliceTypeB []PrimitiveTypeB
