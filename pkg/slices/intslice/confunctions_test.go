package generic_test

import (
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
	"github.com/stretchr/testify/assert"
)

func TestAsSliceTypeB(t *testing.T) {
	aa := generic.SliceTypeA{primitiveAValue}
	bb := generic.SliceTypeB{primitiveBValue}
	convert := func(a generic.PrimitiveTypeA) generic.PrimitiveTypeB {
		return primitiveBValue
	}
	cc := generic.AsSliceTypeB(aa, convert)
	assert.ElementsMatch(t, bb, cc)
}
