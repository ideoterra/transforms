package generic_test

import (
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
	"github.com/stretchr/testify/assert"
)

var primitiveAZero = 1
var primitiveBZero = "1"

func TestAsSliceTypeB(t *testing.T) {
	aa := generic.SliceTypeA{primitiveAZero}
	bb := generic.SliceTypeB{primitiveBZero}
	convert := func(a generic.PrimitiveTypeA) generic.PrimitiveTypeB {
		return primitiveBZero
	}
	cc := generic.AsSliceTypeB(aa, convert)
	assert.ElementsMatch(t, bb, cc)
}
