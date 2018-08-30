package generic_test

import (
	"fmt"
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
)

func TestNullaryMethods(t *testing.T) {
	methodCalls := []func(generic.SliceType){
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

	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("Nullary test %v", i+1), test)
	}
}

func TestUnaryInt64Methods(t *testing.T) {
	methodCalls := []func(generic.SliceType, int64){
		func(aa generic.SliceType, i int64) { aa.Item(i) },
		func(aa generic.SliceType, i int64) { aa.ItemFuzzy(i) },
		func(aa generic.SliceType, i int64) { aa.RemoveAt(i) },
		func(aa generic.SliceType, i int64) { aa.Skip(i) },
		func(aa generic.SliceType, i int64) { aa.SplitAt(i) },
		func(aa generic.SliceType, i int64) { aa.Take(i) },
	}

	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			methodCall(generic.SliceType{}, int64(0))
		}
		t.Run(fmt.Sprintf("UnaryInt64 test %v", i+1), test)
	}
}
