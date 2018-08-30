package generic_test

import (
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
)

func TestAsSliceTypeBMethod(t *testing.T) {
	aa := generic.SliceTypeA{}
	converter := func(a generic.PrimitiveTypeA) generic.PrimitiveTypeB {
		return primitiveBZero
	}
	aa.AsSliceTypeB(converter)
}
