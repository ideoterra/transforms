package generic_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
	"github.com/jecolasurdo/transforms/pkg/slices/shared"
	"github.com/stretchr/testify/assert"
)

// As a rule, all methods in this package (methods.go) are just wrappers around
// a set of base functions (see functions.go). We want to keep a high level of
// test coverage while minimizing test-effort, so method sets are tested in bulk
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
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("Nullary test %v", i+1), test)
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
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("UnaryValue test %v", i+1), test)
	}
}

func TestUnaryPrimitiveMethodHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, interface{}){
		func(aa generic.SliceType, b interface{}) { aa.Append(b) },
		func(aa generic.SliceType, b interface{}) { aa.Enqueue(b) },
		func(aa generic.SliceType, b interface{}) { aa.Push(b) },
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			methodCall(generic.SliceType{}, primitiveZero)
		}
		t.Run(fmt.Sprintf("UnaryPrimitive test %v", i+1), test)
	}
}

func TestUnaryTestMethodHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, generic.Test){
		func(aa generic.SliceType, test generic.Test) { aa.All(test) },
		func(aa generic.SliceType, test generic.Test) { aa.Any(test) },
		func(aa generic.SliceType, test generic.Test) { aa.Count(test) },
		func(aa generic.SliceType, test generic.Test) { aa.Filter(test) },
		func(aa generic.SliceType, test generic.Test) { aa.FindIndex(test) },
		func(aa generic.SliceType, test generic.Test) { aa.First(test) },
		func(aa generic.SliceType, test generic.Test) { aa.Last(test) },
		func(aa generic.SliceType, test generic.Test) { aa.None(test) },
		func(aa generic.SliceType, test generic.Test) { aa.Partition(test) },
		func(aa generic.SliceType, test generic.Test) { aa.Remove(test) },
		func(aa generic.SliceType, test generic.Test) { aa.SkipWhile(test) },
		func(aa generic.SliceType, test generic.Test) { aa.SplitAfter(test) },
		func(aa generic.SliceType, test generic.Test) { aa.SplitBefore(test) },
		func(aa generic.SliceType, test generic.Test) { aa.TakeWhile(test) },
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			testFn := func(_ interface{}) bool {
				return true
			}
			methodCall(generic.SliceType{}, testFn)
		}
		t.Run(fmt.Sprintf("UnaryTest test %v", i+1), test)
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
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("UnaryClosure test %v", i+1), test)
	}
}

func TestBinarySliceEqualityHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, generic.Equality){
		func(aa generic.SliceType, equality generic.Equality) {
			aa.Difference(nil, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.Intersection(nil, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsProperSubset(nil, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsProperSuperset(nil, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsSubset(nil, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsSuperset(nil, equality)
		}}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			equality := func(a, b interface{}) bool {
				return false
			}
			methodCall(generic.SliceType{}, equality)
		}
		t.Run(fmt.Sprintf("BinarySliceEquality test %v", i+1), test)
	}
}

func TestBinaryPrimitiveTestHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, interface{}, generic.Test){
		func(aa generic.SliceType, b interface{}, test generic.Test) {
			aa.InsertAfter(b, test)
		},
		func(aa generic.SliceType, b interface{}, test generic.Test) {
			aa.InsertBefore(b, test)
		},
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			testFn := func(interface{}) bool {
				return false
			}
			methodCall(generic.SliceType{}, primitiveZero, testFn)
		}
		t.Run(fmt.Sprintf("BinaryPrimitiveTest test %v", i+1), test)
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
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("BinaryValueClosure test %v", i+1), test)
	}
}

func TestBinaryValueValueHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType){
		func(aa generic.SliceType) { aa.InsertAt(primitiveZero, 0) },
		func(aa generic.SliceType) { aa.SwapIndex(0, 0) },
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("BinaryValueValue test %v", i+1), test)
	}
}

func TestMutatingMethods(t *testing.T) {
	// The methods covered by this test are be expected to mutate their
	// receiver value. Core behavior of each method is covered in the function
	// tests (see function_test.go), so these tests are fairly cursory, in that
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
		func(aa *generic.SliceType) { aa.Clear() },
		func(aa *generic.SliceType) { aa.Dequeue() },
		func(aa *generic.SliceType) { aa.Distinct(equality) },
		func(aa *generic.SliceType) { aa.Enqueue(1) },
		func(aa *generic.SliceType) { aa.Filter(condition) },
		func(aa *generic.SliceType) { aa.InsertAfter(1, condition) },
		func(aa *generic.SliceType) { aa.InsertBefore(1, condition) },
		func(aa *generic.SliceType) { aa.InsertAt(1, 0) },
		func(aa *generic.SliceType) { aa.Map(func(a interface{}) interface{} { return a.(int) * 2 }) },
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
		test := func(t *testing.T) {
			aa := generic.SliceType{1, 1, 2, 1}
			methodCall(&aa)
			bb := generic.SliceType{1, 1, 2, 1}
			if reflect.DeepEqual(aa, bb) {
				t.Fail()
			}
		}
		t.Run(fmt.Sprintf("Mutation test %v", i+1), test)
	}
}

func TestNonMutatingMethods(t *testing.T) {
	// The methods covered by this test are be expected not to mutate their
	// receiver value. Core behavior of each method is covered in the function
	// tests (see function_test.go), so these tests are fairly cursory, in that
	// they only seek to verify that the receiver value has not changed after
	// each operation, but do not otherwise verify the outcome.

	// var equality = func(a, b interface{}) bool {
	// 	return a.(int) == b.(int)
	// }

	var condition = func(a interface{}) bool {
		return a.(int) > 0
	}

	methodCalls := []func(*generic.SliceType){
		func(aa *generic.SliceType) { aa.All(condition) },
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			aa := generic.SliceType{1, 2, 3}
			methodCall(&aa)
			bb := generic.SliceType{1, 2, 3}
			if !reflect.DeepEqual(aa, bb) {
				t.Fail()
			}
		}
		t.Run(fmt.Sprintf("Mutation test %v", i+1), test)
	}
}

func TestFoo(t *testing.T) {
	aa := &generic.SliceType{1, 2, 3}
	aa.
		Append(4).
		Map(func(a interface{}) interface{} { return a.(int) * 2 })
	bb := generic.SliceType{2, 4, 6, 8}
	assert.ElementsMatch(t, *aa, bb)
}
