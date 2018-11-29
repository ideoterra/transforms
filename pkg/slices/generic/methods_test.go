package generic_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ideoterra/transforms/pkg/slices/generic"
	"github.com/ideoterra/transforms/pkg/slices/generic/closures"
	"github.com/ideoterra/transforms/pkg/slices/shared"
)

// As a rule, all methods in this package (methods.go) are just wrappers around
// a set of base functions (see functions.go). We want to keep a high level of
// condition coverage while minimizing condition-effort, so method sets are tested in bulk
// where possible for simple happy-path operation.
//
// Additional tests are added as necessary, but the bulk of the deeper testing
// is handled in functions_test.go.

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
		condition := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("Nullary condition %v", i+1), condition)
	}
}

func TestUnaryValueMethodHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType){
		func(aa generic.SliceType) { aa.Item(0) },
		func(aa generic.SliceType) { aa.ItemFuzzy(0) },
		func(aa generic.SliceType) { aa.RemoveAt(0) },
		func(aa generic.SliceType) { aa.Skip(0) },
		func(aa generic.SliceType) { aa.SplitAt(0) },
		func(aa generic.SliceType) { aa.Take(0) },
		func(aa generic.SliceType) { aa.Union(new([]interface{})) },
		func(aa generic.SliceType) { aa.Zip(new([]interface{})) },
	}

	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("UnaryValue condition %v", i+1), condition)
	}
}

func TestUnaryPrimitiveMethodHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, interface{}){
		func(aa generic.SliceType, b interface{}) { aa.Append(b) },
		func(aa generic.SliceType, b interface{}) { aa.Enqueue(b) },
		func(aa generic.SliceType, b interface{}) { aa.Push(b) },
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			methodCall(generic.SliceType{}, primitiveZero)
		}
		t.Run(fmt.Sprintf("UnaryPrimitive condition %v", i+1), condition)
	}
}

func TestUnaryTestMethodHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, closures.ConditionFn){
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.All(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.Any(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.Count(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.Filter(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.FindIndex(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.First(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.Last(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.None(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.Partition(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.Remove(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.SkipWhile(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.SplitAfter(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.SplitBefore(condition) },
		func(aa generic.SliceType, condition closures.ConditionFn) { aa.TakeWhile(condition) },
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			testFn := func(_ interface{}) bool {
				return true
			}
			methodCall(generic.SliceType{}, testFn)
		}
		t.Run(fmt.Sprintf("UnaryTest condition %v", i+1), condition)
	}
}

func TestUnaryClosureHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType){
		func(aa generic.SliceType) {
			aa.Distinct(func(a, b interface{}) bool { return true })
		},
		func(aa generic.SliceType) {
			aa.Expand(func(interface{}) []interface{} { return nil })
		},
		func(aa generic.SliceType) {
			aa.ForEach(func(interface{}) shared.Continue { return shared.ContinueNo })
		},
		func(aa generic.SliceType) {
			aa.ForEachR(func(interface{}) shared.Continue { return shared.ContinueNo })
		},
		func(aa generic.SliceType) {
			aa.Group(func(interface{}) int64 { return 0 })
		},
		func(aa generic.SliceType) {
			aa.GroupI(func(int64, interface{}) int64 { return 0 })
		},
		func(aa generic.SliceType) {
			aa.Map(func(interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Reduce(func(a, b interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Sort(func(a, b interface{}) bool { return false })
		},
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("UnaryClosure condition %v", i+1), condition)
	}
}

func TestBinarySliceEqualityHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, closures.EqualityFn){
		func(aa generic.SliceType, equality closures.EqualityFn) {
			aa.Difference(nil, equality)
		},
		func(aa generic.SliceType, equality closures.EqualityFn) {
			aa.Intersection(nil, equality)
		},
		func(aa generic.SliceType, equality closures.EqualityFn) {
			aa.IsProperSubset(nil, equality)
		},
		func(aa generic.SliceType, equality closures.EqualityFn) {
			aa.IsProperSuperset(nil, equality)
		},
		func(aa generic.SliceType, equality closures.EqualityFn) {
			aa.IsSubset(nil, equality)
		},
		func(aa generic.SliceType, equality closures.EqualityFn) {
			aa.IsSuperset(nil, equality)
		}}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			equality := func(a, b interface{}) bool {
				return false
			}
			methodCall(generic.SliceType{}, equality)
		}
		t.Run(fmt.Sprintf("BinarySliceEquality condition %v", i+1), condition)
	}
}

func TestBinaryPrimitiveTestHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, interface{}, closures.ConditionFn){
		func(aa generic.SliceType, b interface{}, condition closures.ConditionFn) {
			aa.InsertAfter(b, condition)
		},
		func(aa generic.SliceType, b interface{}, condition closures.ConditionFn) {
			aa.InsertBefore(b, condition)
		},
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			testFn := func(interface{}) bool {
				return false
			}
			methodCall(generic.SliceType{}, primitiveZero, testFn)
		}
		t.Run(fmt.Sprintf("BinaryPrimitiveTest condition %v", i+1), condition)
	}
}

func TestBinaryValueClosureHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType){
		func(aa generic.SliceType) {
			aa.ForEachC(0, func(interface{}, func() bool) shared.Continue {
				return shared.ContinueNo
			})
		},
		func(aa generic.SliceType) {
			aa.WindowCentered(0, func([]interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.WindowLeft(0, func([]interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.WindowRight(0, func([]interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Fold(primitiveZero, func(a, b interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.FoldI(primitiveZero, func(i int64, a, b interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Pairwise(primitiveZero, func(a, b interface{}) interface{} { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Collect([]interface{}{}, func(a, b interface{}) interface{} { return primitiveZero })
		},
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("BinaryValueClosure condition %v", i+1), condition)
	}
}

func TestBinaryValueValueHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType){
		func(aa generic.SliceType) { aa.InsertAt(primitiveZero, 0) },
		func(aa generic.SliceType) { aa.SwapIndex(0, 0) },
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("BinaryValueValue condition %v", i+1), condition)
	}
}

func TestMutatingMethods(t *testing.T) {
	// The methods covered by this condition are expected to mutate their receiver
	// value. Core behavior of each method is covered in the function
	// tests (see functions_test.go), so these tests are fairly cursory in that
	// they only seek to verify that the value has changed after an operation,
	// but don't go so far as to check how the value changed.

	var equality = func(a, b interface{}) bool {
		return a.(int) == b.(int)
	}

	var condition = func(a interface{}) bool {
		return a.(int) == 1
	}

	var sliceForUnionTest = []interface{}{1}

	methodCalls := []func(*generic.SliceType){
		func(aa *generic.SliceType) { aa.Append(1) },
		func(aa *generic.SliceType) { aa.Apply(func(a interface{}) interface{} { return a.(int) * 2 }) },
		func(aa *generic.SliceType) { aa.Clear() },
		func(aa *generic.SliceType) { aa.Dequeue() },
		func(aa *generic.SliceType) { aa.Distinct(equality) },
		func(aa *generic.SliceType) { aa.Enqueue(1) },
		func(aa *generic.SliceType) { aa.Filter(condition) },
		func(aa *generic.SliceType) { aa.InsertAfter(1, condition) },
		func(aa *generic.SliceType) { aa.InsertBefore(1, condition) },
		func(aa *generic.SliceType) { aa.InsertAt(1, 0) },
		func(aa *generic.SliceType) { aa.Pop() },
		func(aa *generic.SliceType) { aa.Push(1) },
		func(aa *generic.SliceType) { aa.Remove(condition) },
		func(aa *generic.SliceType) { aa.RemoveAt(1) },
		func(aa *generic.SliceType) { aa.Reverse() },
		func(aa *generic.SliceType) { aa.Skip(1) },
		func(aa *generic.SliceType) { aa.SkipWhile(condition) },
		func(aa *generic.SliceType) { aa.Sort(func(a, b interface{}) bool { return a.(int) < b.(int) }) },
		func(aa *generic.SliceType) { aa.SwapIndex(0, 2) },
		func(aa *generic.SliceType) { aa.Tail() },
		func(aa *generic.SliceType) { aa.Take(1) },
		func(aa *generic.SliceType) { aa.TakeWhile(condition) },
		func(aa *generic.SliceType) { aa.Union(&sliceForUnionTest) },
	}
	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			aa := generic.SliceType{1, 1, 2, 1}
			methodCall(&aa)
			bb := generic.SliceType{1, 1, 2, 1}
			if reflect.DeepEqual(aa, bb) {
				t.Fail()
			}
		}
		t.Run(fmt.Sprintf("Mutation condition %v", i+1), condition)
	}
}

func TestNonMutatingMethods(t *testing.T) {
	// The methods covered by this condition are be expected not to mutate their
	// receiver value. Core behavior of each method is covered in the function
	// tests (see functions_test.go), so these tests are fairly cursory in that
	// they only seek to verify that the receiver value has not changed after
	// each operation, but do not otherwise verify the outcomes.

	var equality = func(a, b interface{}) bool {
		return a.(int) == b.(int)
	}

	var condition = func(a interface{}) bool {
		return a.(int) > 0
	}

	var window = func(window []interface{}) interface{} {
		return window[0]
	}

	sliceForZipTest := []interface{}{4, 5, 6}

	methodCalls := []func(*generic.SliceType){
		func(aa *generic.SliceType) { aa.All(condition) },
		func(aa *generic.SliceType) { aa.Any(condition) },
		func(aa *generic.SliceType) { aa.Clone() },
		func(aa *generic.SliceType) {
			aa.Collect([]interface{}{1, 2}, func(a, b interface{}) interface{} {
				return a.(int) * b.(int)
			})
		},
		func(aa *generic.SliceType) { aa.Count(condition) },
		func(aa *generic.SliceType) { aa.Difference([]interface{}{2, 3}, equality) },
		func(aa *generic.SliceType) { aa.Empty() },
		func(aa *generic.SliceType) { aa.End() },
		func(aa *generic.SliceType) {
			aa.Expand(func(a interface{}) []interface{} { return []interface{}{1, 2} })
		},
		func(aa *generic.SliceType) { aa.FindIndex(condition) },
		func(aa *generic.SliceType) { aa.First(condition) },
		func(aa *generic.SliceType) {
			aa.Fold(2, func(a, acc interface{}) interface{} {
				return acc.(int) * a.(int)
			})
		},
		func(aa *generic.SliceType) {
			aa.FoldI(2, func(_ int64, a, acc interface{}) interface{} {
				return acc.(int) * a.(int)
			})
		},
		func(aa *generic.SliceType) {
			aa.ForEach(func(_ interface{}) shared.Continue { return shared.ContinueYes })
		},
		func(aa *generic.SliceType) {
			aa.ForEachC(1, func(_ interface{}, _ func() bool) shared.Continue { return shared.ContinueYes })
		},
		func(aa *generic.SliceType) {
			aa.ForEachR(func(_ interface{}) shared.Continue { return shared.ContinueYes })
		},
		func(aa *generic.SliceType) { aa.Group(func(_ interface{}) int64 { return 0 }) },
		func(aa *generic.SliceType) { aa.GroupI(func(i int64, _ interface{}) int64 { return i }) },
		func(aa *generic.SliceType) { aa.Head() },
		func(aa *generic.SliceType) { aa.Intersection([]interface{}{1, 2}, equality) },
		func(aa *generic.SliceType) { aa.IsProperSubset([]interface{}{1, 2}, equality) },
		func(aa *generic.SliceType) { aa.IsProperSuperset([]interface{}{1, 2}, equality) },
		func(aa *generic.SliceType) { aa.IsSubset([]interface{}{1, 2}, equality) },
		func(aa *generic.SliceType) { aa.IsSuperset([]interface{}{1, 2}, equality) },
		func(aa *generic.SliceType) { aa.Item(0) },
		func(aa *generic.SliceType) { aa.ItemFuzzy(0) },
		func(aa *generic.SliceType) { aa.Last(condition) },
		func(aa *generic.SliceType) { aa.Len() },
		func(aa *generic.SliceType) { aa.None(condition) },
		func(aa *generic.SliceType) {
			aa.Pairwise(1, func(a, b interface{}) interface{} {
				return a.(int) * b.(int)
			})
		},
		func(aa *generic.SliceType) { aa.Partition(condition) },
		func(aa *generic.SliceType) { aa.Permutable() },
		func(aa *generic.SliceType) { aa.Permutations() },
		func(aa *generic.SliceType) { aa.Permute() },
		func(aa *generic.SliceType) {
			aa.Reduce(func(a, acc interface{}) interface{} {
				return a.(int) + acc.(int)
			})
		},
		func(aa *generic.SliceType) { aa.SplitAfter(condition) },
		func(aa *generic.SliceType) { aa.SplitAt(1) },
		func(aa *generic.SliceType) { aa.SplitBefore(condition) },
		func(aa *generic.SliceType) { _ = aa.String() },
		func(aa *generic.SliceType) { aa.Unzip() },
		func(aa *generic.SliceType) { aa.WindowCentered(2, window) },
		func(aa *generic.SliceType) { aa.WindowLeft(2, window) },
		func(aa *generic.SliceType) { aa.WindowRight(2, window) },
		func(aa *generic.SliceType) { aa.Zip(&sliceForZipTest) },
	}

	for i, methodCall := range methodCalls {
		condition := func(t *testing.T) {
			aa := generic.SliceType{1, 2, 3}
			methodCall(&aa)
			bb := generic.SliceType{1, 2, 3}
			if !reflect.DeepEqual(aa, bb) {
				t.Fail()
			}
		}
		t.Run(fmt.Sprintf("Mutation condition %v", i+1), condition)
	}
}
