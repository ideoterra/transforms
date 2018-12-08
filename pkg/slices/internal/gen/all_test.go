package gen_test

import (
	"testing"

	"github.com/ideoterra/transforms/pkg/slices/internal/gen"
	"github.com/stretchr/testify/assert"
)

func init() {
	RegisterSpecs(
		Specification{
			FunctionName: "All",
			StandardPath: Behavior{
				Description: "Returns true if all elements pass test.",
				Expectation: func(t *testing.T) {
					s := []gen.TA{1, 2, 3, 4}
					test := func(p gen.TA) bool {
						return p.(int) < 5
					}
					assert.True(t, gen.All.Do(s, test))
				},
			},
			AlternativePath: Behavior{
				Description: "Returns false if not all elements pass test.",
				Expectation: func(t *testing.T) {
					s := []gen.TA{1, 2, 3, 4, 5}
					test := func(p gen.TA) bool {
						return p.(int) < 5
					}
					assert.False(t, gen.All.Do(s, test))
				},
			},
		},
	)
}
