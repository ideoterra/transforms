package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
	"github.com/jecolasurdo/transforms/pkg/slices/slicetypes"
)

func TestConversions(t *testing.T) {
	genericSlice := []generic.Generic{new(generic.Generic)}
	result := generic.AsUintSlice(genericSlice, func(value generic.Generic) uint {
		return uint(1)
	})
	assert.IsType(t, slicetypes.UintSlice{}, result)
	assert.Equal(t, 1, len(result))
}
