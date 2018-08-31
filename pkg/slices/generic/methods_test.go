package generic_test

import (
	"fmt"
	"testing"

	"github.com/jecolasurdo/transforms/pkg/slices/generic"
	"github.com/jecolasurdo/transforms/pkg/slices/shared"
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
		func(aa generic.SliceType) { aa.Union(new(generic.SliceType)) },
		func(aa generic.SliceType) { aa.Zip(new(generic.SliceType)) },
	}

	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			methodCall(generic.SliceType{})
		}
		t.Run(fmt.Sprintf("UnaryValue test %v", i+1), test)
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
			testFn := func(_ generic.PrimitiveType) bool {
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
			aa.Distinct(func(a, b generic.PrimitiveType) bool { return true })
		},
		func(aa generic.SliceType) {
			aa.Expand(func(generic.PrimitiveType) generic.SliceType { return nil })
		},
		func(aa generic.SliceType) {
			aa.ForEach(func(generic.PrimitiveType) shared.Continue { return shared.ContinueNo })
		},
		func(aa generic.SliceType) {
			aa.ForEachR(func(generic.PrimitiveType) shared.Continue { return shared.ContinueNo })
		},
		func(aa generic.SliceType) {
			aa.Group(func(generic.PrimitiveType) int64 { return 0 })
		},
		func(aa generic.SliceType) {
			aa.GroupI(func(int64, generic.PrimitiveType) int64 { return 0 })
		},
		func(aa generic.SliceType) {
			aa.Map(func(generic.PrimitiveType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Reduce(func(a, b generic.PrimitiveType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Sort(func(a, b generic.PrimitiveType) bool { return false })
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
			aa.Difference(&aa, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.Intersection(&aa, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsProperSubset(&aa, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsProperSuperset(&aa, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsSubset(&aa, equality)
		},
		func(aa generic.SliceType, equality generic.Equality) {
			aa.IsSuperset(&aa, equality)
		}}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			equality := func(a, b generic.PrimitiveType) bool {
				return false
			}
			methodCall(generic.SliceType{}, equality)
		}
		t.Run(fmt.Sprintf("BinarySliceEquality test %v", i+1), test)
	}
}

func TestBinaryPrimitiveTestHappyPaths(t *testing.T) {
	methodCalls := []func(generic.SliceType, generic.PrimitiveType, generic.Test){
		func(aa generic.SliceType, b generic.PrimitiveType, test generic.Test) {
			aa.InsertAfter(b, test)
		},
		func(aa generic.SliceType, b generic.PrimitiveType, test generic.Test) {
			aa.InsertBefore(b, test)
		},
	}
	for i, methodCall := range methodCalls {
		test := func(t *testing.T) {
			testFn := func(generic.PrimitiveType) bool {
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
			aa.ForEachC(0, func(generic.PrimitiveType, func() bool) shared.Continue {
				return shared.ContinueNo
			})
		},
		func(aa generic.SliceType) {
			aa.WindowCentered(0, func(generic.SliceType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.WindowLeft(0, func(generic.SliceType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.WindowRight(0, func(generic.SliceType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Fold(primitiveZero, func(a, b generic.PrimitiveType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.FoldI(primitiveZero, func(i int64, a, b generic.PrimitiveType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Pairwise(primitiveZero, func(a, b generic.PrimitiveType) generic.PrimitiveType { return primitiveZero })
		},
		func(aa generic.SliceType) {
			aa.Collect(new(generic.SliceType), func(a, b generic.PrimitiveType) generic.PrimitiveType { return primitiveZero })
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
