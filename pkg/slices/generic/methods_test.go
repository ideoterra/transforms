package generic_test

import (
	"fmt"
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
)

func TestNullaryMethods(t *testing.T) {
	nullaryMethodCalls := []func(generic.SliceType){
		func(aa generic.SliceType) { aa.Clear() },
		func(aa generic.SliceType) { aa.Clone() },
		func(aa generic.SliceType) { aa.Dequeue() },
		func(aa generic.SliceType) { aa.Empty() },
		func(aa generic.SliceType) { aa.End() },
		func(aa generic.SliceType) { aa.Head() },
		func(aa generic.SliceType) { aa.Len() },
		func(aa generic.SliceType) { aa.Permutable() },
		func(aa generic.SliceType) { aa.Permutations() },
		func(aa generic.SliceType) { aa.Permute() },
		func(aa generic.SliceType) { aa.Pop() },
		func(aa generic.SliceType) { aa.Reverse() },
		func(aa generic.SliceType) { _ = aa.String() },
		func(aa generic.SliceType) { aa.Tail() },
		func(aa generic.SliceType) { aa.Unzip() },
	}

	for i, nullaryMethodCall := range nullaryMethodCalls {
		test := func(t *testing.T) {
			nullaryMethodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("Nullary test %v", i+1), test)
	}
}
