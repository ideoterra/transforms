package generic_test

import (
	"fmt"
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
	"github.com/stretchr/testify/assert"
)

type Behavior struct {
	Description string
	Expectation func(t *testing.T)
}

type Specification struct {
	FunctionName    string
	StandardPath    Behavior
	AlternativePath Behavior
	EdgeCases       []Behavior
}

var Specifications = []Specification{
	Specification{
		FunctionName: "All",
		StandardPath: Behavior{
			Description: "Returns true if all elements pass test.",
			Expectation: func(t *testing.T) {
				s := generic.SliceType{1, 2, 3, 4}
				test := func(p generic.PrimitiveType) bool {
					return p.(int) < 5
				}
				assert.True(t, generic.All(s, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if not all elements pass test.",
			Expectation: func(t *testing.T) {
				s := generic.SliceType{1, 2, 3, 4, 5}
				test := func(p generic.PrimitiveType) bool {
					return p.(int) < 5
				}
				assert.False(t, generic.All(s, test))
			},
		},
	},
	Specification{
		FunctionName: "Any",
		StandardPath: Behavior{
			Description: "Returns true if any of the elements match.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 2
				}
				assert.True(t, generic.Any(aa, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if none of the elements match.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 5
				}
				assert.False(t, generic.Any(aa, test))
			},
		},
	},
	Specification{
		FunctionName: "Append",
		StandardPath: Behavior{
			Description: "Values are added to the end of the slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.SliceType{5, 6, 7, 8}
				generic.Append(&aa, bb...)
				assert.ElementsMatch(t, generic.SliceType{1, 2, 3, 4, 5, 6, 7, 8}, aa)
			},
		},
		AlternativePath: Behavior{
			Description: "No values supplied makes no change.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.SliceType{}
				generic.Append(&aa, bb...)
				assert.ElementsMatch(t, generic.SliceType{1, 2, 3, 4}, aa)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Nil passed as aa appends bb",
				Expectation: func(t *testing.T) {
					var aa generic.SliceType
					bb := generic.SliceType{5, 6, 7, 8}
					generic.Append(&aa, bb...)
					assert.ElementsMatch(t, generic.SliceType{5, 6, 7, 8}, aa)
				},
			},
		},
	},
	Specification{
		FunctionName: "Clear",
		StandardPath: Behavior{
			Description: "The slice is set to nil",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.Clear(&aa)
				assert.Nil(t, aa)
			},
		},
		AlternativePath: Behavior{
			Description: "An already nil slice can be cleared.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.Clear(&aa)
				generic.Append(&aa, 6, 7, 8)
				assert.ElementsMatch(t, generic.SliceType{6, 7, 8}, aa)
			},
		},
	},
	Specification{
		FunctionName: "Clone",
		StandardPath: Behavior{
			Description: "A new identical slice is allocated in memory.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.Clone(aa)
				if &aa == &bb {
					t.Error("Slices aa and bb should not have the same address")
				}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Slices are not deep cloned in this operation",
			Expectation: func(t *testing.T) {
				value := 1
				aa := generic.SliceType{&value}
				bb := generic.Clone(aa)
				a := aa[0].(*int)
				b := bb[0].(*int)
				if a != b {
					t.Error("Expected aa[0] and bb[0] to have the same address")
				}
			},
		},
	},
	Specification{
		FunctionName: "Collect",
		StandardPath: Behavior{
			Description: "Values are concatenated as expected.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B"}
				bb := generic.SliceType{"Y", "Z"}
				collector := func(a, b generic.PrimitiveType) generic.PrimitiveType {
					return a.(string) + b.(string)
				}
				cc := generic.Collect(aa, bb, collector)
				dd := generic.SliceType{"AY", "AZ", "BY", "BZ"}
				assert.ElementsMatch(t, cc, dd)
			},
		},
		AlternativePath: Behavior{
			Description: "",
			Expectation: func(t *testing.T) {
				t.Skip()
			},
		},
	},
	Specification{
		FunctionName: "Count",
		StandardPath: Behavior{
			Description: "Returns the correct count",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 == 0
				}
				assert.Equal(t, int64(2), generic.Count(aa, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns 0 if no matches",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 6
				}
				assert.Equal(t, int64(0), generic.Count(aa, test))
			},
		},
	},
	Specification{
		FunctionName: "Dequeue",
		StandardPath: Behavior{
			Description: "Removes and returns head.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				bb := generic.Dequeue(&aa)
				assert.Equal(t, 1, bb[0])
				cc := generic.SliceType{2, 3}
				assert.ElementsMatch(t, aa, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "If source slice is empty, empty slice is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				bb := generic.Dequeue(&aa)
				if len(bb) != 0 {
					t.Error("Expected bb to be empty.")
				}
			},
		},
	},
}

func TestTransforms(t *testing.T) {
	for _, specification := range Specifications {
		t.Run(specification.FunctionName+"StandardPath", specification.StandardPath.Expectation)
		t.Run(specification.FunctionName+"AlternativePath", specification.AlternativePath.Expectation)
		for i, edgeCase := range specification.EdgeCases {
			t.Run(fmt.Sprintf("%vEdgeCase%v", specification.FunctionName, i), edgeCase.Expectation)
		}
	}
}
