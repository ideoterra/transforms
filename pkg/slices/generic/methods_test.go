package generic_test

import (
	"fmt"
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
)

// As a rule, all methods in this package are just wrappers around a set of
// base functions. We want to keep a high level of test coverage, so method
// sets are tested in bulk where possible for simple happy-path operation.
//
// Additional tests are added as necessary, but the bulk of the deeper testing
// is handled in functions_test.go.

var primitiveZero = interface{}(nil)

func TestNullaryMethodHappyPaths(t *testing.T) {
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

func TestUnaryInt64MethodHappyPaths(t *testing.T) {
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

func TestUnaryPrimitiveMethodHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, generic.PrimitiveType){
		func(aa generic.SliceType, b generic.PrimitiveType) { aa.Append(b) },
		func(aa generic.SliceType, b generic.PrimitiveType) { aa.Enqueue(b) },
		func(aa generic.SliceType, b generic.PrimitiveType) { aa.Push(b) },
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			methodCall(generic.SliceType{}, primitiveZero)
		}
		t.Run(fmt.Sprintf("UnaryPrimitive test %v", i+1), test)
	}
}
