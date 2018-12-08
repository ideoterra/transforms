package gen

// AllTA contains methods for applying an All trasnform to various types.
type AllTA struct{}

// All applies a test function to each element in the slice, and returns true if
// the test function returns true for all items in the slice.
var All AllTA
