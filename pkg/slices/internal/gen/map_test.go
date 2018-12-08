package gen_test

import (
	"testing"

	"github.com/ideoterra/transforms/pkg/slices/internal/gen"
)

func init() {
	RegisterSpecs(
		Specification{
			FunctionName: "Map",
			StandardPath: Behavior{
				Description: "Maps the transform to each element",
				Expectation: func(t *testing.T) {
					aa := []gen.TA{1, 2, 3}
					mapFn := func(a gen.TA) gen.TB {
						return a.(int) * 2
					}
					bb := gen.Map.ToTB(aa, mapFn)
					cc := []gen.TB{2, 4, 6}
					assertSlicesEqual(t, bb, cc)
				},
			},
			AlternativePath: Behavior{
				Description: "",
				Expectation: func(t *testing.T) {
					t.Skip()
				},
			},
		},
	)
}
