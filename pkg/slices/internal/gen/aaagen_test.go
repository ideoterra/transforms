package gen_test

import (
	"fmt"
	"testing"

	"github.com/ideoterra/transforms/pkg/slices/internal/gen"
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

var Specifications = []Specification{}

func RegisterSpecs(spec ...Specification) {
	Specifications = append(Specifications, spec...)
}

func TestTransforms(t *testing.T) {
	for _, specification := range Specifications {
		t.Run(specification.FunctionName+"StandardPath", specification.StandardPath.Expectation)
		t.Run(specification.FunctionName+"AlternativePath", specification.AlternativePath.Expectation)
		for i, edgeCase := range specification.EdgeCases {
			t.Run(fmt.Sprintf("%vEdgeCase%v", specification.FunctionName, i+1), edgeCase.Expectation)
		}
	}
}

func assertSlicesEqual(t *testing.T, xx []gen.B, yy []gen.B) bool {
	// often dealing with using []interface{} as the key (hash) value in a map
	// which go doesn't like because []interface{} types are unhashable.
	// We convert the values to a string to get around this limitation.
	hash := func(z interface{}) string {
		return fmt.Sprintf("%v", z)
	}

	if len(xx) != len(yy) {
		t.Errorf("Expected lengths to match. Wanted %v, got %v", xx, yy)
		return false
	}
	diff := make(map[interface{}]int, len(xx))
	for _, x := range xx {
		diff[hash(x)]++
	}
	for _, y := range yy {
		hashy := hash(y)
		if _, ok := diff[hashy]; !ok {
			t.Errorf("Expected %v, but got %v", xx, yy)
			return false
		}
		diff[hashy] -= 1
		if diff[hashy] == 0 {
			delete(diff, hashy)
		}
	}
	if len(diff) == 0 {
		return true
	}

	t.Errorf("Expected %v, but got %v", xx, yy)
	return false
}
