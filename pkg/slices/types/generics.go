package types

// SliceType2 is a two dimensional slice of PrimitiveType
type SliceType2 []SliceType

// SliceType is a one dimensional slice of PrimitiveType.
type SliceType []PrimitiveType

// PrimitiveType is a placeholder for the type underpinning the generic SliceType.
type PrimitiveType interface{}
