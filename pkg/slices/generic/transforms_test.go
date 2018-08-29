package generic_test

import (
	"fmt"
	"strconv"
	"sync"
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
	Specification{
		FunctionName: "Difference",
		StandardPath: Behavior{
			Description: "Returns the difference between two slices",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.SliceType{5, 4, 3}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				cc := generic.Difference(aa, bb, equal)
				dd := generic.SliceType{1, 2, 5}
				assert.ElementsMatch(t, cc, dd)
			},
		},
		AlternativePath: Behavior{
			Description: `Duplicates are handled.
						  Those from aa appear first in the result. bb appear
						  second. Order and duplicates should be maintained.`,
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 3, 1, 4}
				bb := generic.SliceType{5, 4, 3, 5}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				cc := generic.Difference(aa, bb, equal)
				dd := generic.SliceType{1, 2, 1, 5, 5}
				assert.ElementsMatch(t, cc, dd)
			},
		},
	},
	Specification{
		FunctionName: "Distinct",
		StandardPath: Behavior{
			Description: "Duplicates are removed from the slice, mutating the original",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"Dani", "Riley", "Dani", "Tori", "Janice"}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(string) == b.(string)
				}
				generic.Distinct(&aa, equal)
				bb := generic.SliceType{"Dani", "Riley", "Tori", "Janice"}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "Empty",
		StandardPath: Behavior{
			Description: "Returns true if slice is empty",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				assert.True(t, generic.Empty(aa))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if slice is not empty",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1}
				assert.False(t, generic.Empty(aa))
			},
		},
	},
	Specification{
		FunctionName: "End",
		StandardPath: Behavior{
			Description: "Returns a slice with the last element.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				bb := generic.End(aa)
				assert.Equal(t, 3, bb[0])
			},
		},
		AlternativePath: Behavior{
			Description: "An empty slice returns an empty slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				bb := generic.End(aa)
				assert.True(t, generic.Empty(bb))
			},
		},
	},
	Specification{
		FunctionName: "Enqueue",
		StandardPath: Behavior{
			Description: "The value is added to the head of the slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				generic.Enqueue(&aa, 4)
				bb := generic.SliceType{4, 1, 2, 3}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "Expand",
		StandardPath: Behavior{
			Description: "Expands the supplied list according to the expansion",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"AB", "CD", "EF"}
				expansion := func(a generic.PrimitiveType) generic.SliceType {
					b := string(a.(string)[0])
					c := string(a.(string)[1])
					return generic.SliceType{b, c}
				}
				bb := generic.Expand(aa, expansion)
				cc := generic.SliceType{"A", "B", "C", "D", "E", "F"}
				assert.ElementsMatch(t, bb, cc)
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
		FunctionName: "Filter",
		StandardPath: Behavior{
			Description: "Items are removed for which the test function returns true.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 == 0
				}
				generic.Filter(&aa, test)
				bb := generic.SliceType{1, 3}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "FindIndex",
		StandardPath: Behavior{
			Description: "The first matching element is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) >= 3
				}
				n := generic.FindIndex(aa, test)
				assert.Equal(t, int64(2), n)
			},
		},
		AlternativePath: Behavior{
			Description: "If no match is found, -1 is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) >= 5
				}
				n := generic.FindIndex(aa, test)
				assert.Equal(t, int64(-1), n)
			},
		},
	},
	Specification{
		FunctionName: "First",
		StandardPath: Behavior{
			Description: "The first matching element is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) >= 2
				}
				bb := generic.First(aa, test)
				cc := generic.SliceType{2}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "An empty slice is returned if there are no matches.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) >= 5
				}
				bb := generic.First(aa, test)
				cc := generic.SliceType{}
				assert.ElementsMatch(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "Fold",
		StandardPath: Behavior{
			Description: "Fold accumulates values properly",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				folder := func(a, acc generic.PrimitiveType) generic.PrimitiveType {
					return a.(int) + acc.(int)
				}
				bb := generic.Fold(aa, 1, folder)
				assert.Equal(t, 11, bb[0].(int))
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
		FunctionName: "FoldI",
		StandardPath: Behavior{
			Description: "FoldI accumulates values properly",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				folder := func(i int64, a, acc generic.PrimitiveType) generic.PrimitiveType {
					return fmt.Sprintf("%v%v%v",
						acc.(string),
						strconv.Itoa(int(i)),
						a.(string))
				}
				bb := generic.FoldI(aa, "X", folder)
				assert.Equal(t, "X0A1B2C", bb[0].(string))
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
		FunctionName: "ForEach",
		StandardPath: Behavior{
			Description: "Each element of the list is applied to the function",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				result := ""
				fn := func(a generic.PrimitiveType) generic.Continue {
					result = result + a.(string)
					return generic.ContinueYes
				}
				generic.ForEach(aa, fn)
				assert.Equal(t, "ABC", result)
			},
		},
		AlternativePath: Behavior{
			Description: "The iterator stops if false is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				result := ""
				fn := func(a generic.PrimitiveType) generic.Continue {
					result = result + a.(string)
					return a.(string) != "B"
				}
				generic.ForEach(aa, fn)
				assert.Equal(t, "AB", result)
			},
		},
	},
	Specification{
		FunctionName: "ForEachC",
		StandardPath: Behavior{
			Description: "Each element of the list is applied to the function",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				mu := new(sync.Mutex)
				result := ""
				fn := func(a generic.PrimitiveType, cancelPending func() bool) generic.Continue {
					mu.Lock()
					defer mu.Unlock()
					result = result + a.(string)
					return generic.ContinueYes
				}
				generic.ForEachC(aa, 1, fn)
				bb := generic.SliceType{"ABC", "BAC", "CAB", "ACB", "BCA", "CBA"}
				if !generic.Any(bb, func(b generic.PrimitiveType) bool {
					return b.(string) == result
				}) {
					t.Errorf("Expected a variant of 'ABC', but got '%v'", result)
				}
			},
		},
		AlternativePath: Behavior{
			Description: "The function panic if a negative pool size is specified.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				mu := new(sync.Mutex)
				result := ""
				fn := func(a generic.PrimitiveType, cancelPending func() bool) generic.Continue {
					mu.Lock()
					defer mu.Unlock()
					result = result + a.(string)
					return generic.ContinueYes
				}
				assert.PanicsWithValue(t,
					"ForEachC: The concurrency pool size (c) must be non-negative.",
					func() { generic.ForEachC(aa, -1, fn) })
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: `Long running operations wind down when a 
							  cancellation is broadcast`,
				Expectation: func(t *testing.T) {
					// Elements "A" and "B" will spawn infinitely loops, which
					// check for pending cancellation on each iteration.
					// Element "C" will spawn an operation that will wait until
					// "A" and "B" are both running, at which point "C" will
					// cancel further iterations.
					// "A" and "B" will then identify the cancellation, and
					// will halt.
					//
					// If this test is failing, it should time out. This can
					// be verified by removing the "halt = true" assignment and
					// manually confirming that the test would time out in that
					// case.
					aa := generic.SliceType{"A", "B", "C"}
					mu := new(sync.RWMutex)
					aIsRunning := false
					bIsRunning := false
					fn := func(a generic.PrimitiveType, cancelPending func() bool) generic.Continue {
						if a.(string) == "A" || a.(string) == "B" {
							mu.Lock()
							if a.(string) == "A" {
								aIsRunning = true
							} else {
								bIsRunning = true
							}
							mu.Unlock()
							for cancelPending() == false {
							}
							return generic.ContinueYes
						} else {
							halt := false
							for {
								mu.RLock()
								if aIsRunning && bIsRunning {
									halt = true
								}
								mu.RUnlock()
								if halt {
									return generic.ContinueNo
								}
							}
						}
					}
					generic.ForEachC(aa, 3, fn)
				},
			},
			Behavior{
				Description: `Upon cancellation, active goroutines are allowed
							  to wind down before the function returns.`,
				Expectation: func(t *testing.T) {
					// Elements "A" and "B" will spawn infinitely loops, which
					// check for pending cancellation on each iteration.
					// Element "C" will spawn an operation that will wait until
					// "A" and "B" are both running, at which point "C" will
					// cancel further iterations.
					// "A" and "B" will then identify the cancellation, and
					// will halt.
					//
					// "A", and "B" will each write out a value upon
					// cancellation, which is used to verify that the function
					// blocked until all goroutines exited cleanly.
					aa := generic.SliceType{"A", "B", "C"}
					mu := new(sync.RWMutex)
					aIsRunning := false
					bIsRunning := false
					aExitedCleanly := false
					bExitedCleanly := false
					fn := func(a generic.PrimitiveType, cancelPending func() bool) generic.Continue {
						if a.(string) == "A" || a.(string) == "B" {
							mu.Lock()
							if a.(string) == "A" {
								aIsRunning = true
							} else {
								bIsRunning = true
							}
							mu.Unlock()
							for cancelPending() == false {
							}
							mu.Lock()
							if a.(string) == "A" {
								aExitedCleanly = true
							} else {
								bExitedCleanly = true
							}
							mu.Unlock()
							return generic.ContinueYes
						} else {
							halt := false
							for {
								mu.RLock()
								if aIsRunning && bIsRunning {
									halt = true
								}
								mu.RUnlock()
								if halt {
									return generic.ContinueNo
								}
							}
						}
					}
					generic.ForEachC(aa, 3, fn)
					assert.True(t, aExitedCleanly)
					assert.True(t, bExitedCleanly)
				},
			},
		},
	},
	Specification{
		FunctionName: "ForEachR",
		StandardPath: Behavior{
			Description: `Each element of the list is applied to the function
						  in reverse order`,
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				result := ""
				fn := func(a generic.PrimitiveType) generic.Continue {
					result = result + a.(string)
					return true
				}
				generic.ForEachR(aa, fn)
				assert.Equal(t, "CBA", result)
			},
		},
		AlternativePath: Behavior{
			Description: `The iterator stops when the function return true.`,
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				result := ""
				fn := func(a generic.PrimitiveType) generic.Continue {
					result = result + a.(string)
					return a.(string) != "B"
				}
				generic.ForEachR(aa, fn)
				assert.Equal(t, "CB", result)
			},
		},
	},
	Specification{
		FunctionName: "Group",
		StandardPath: Behavior{
			Description: "Elements are grouped as expected.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C", "D", "E", "F"}
				grouper := func(a generic.PrimitiveType) int64 {
					switch {
					case a.(string) == "A" || a.(string) == "B":
						return 1
					case a.(string) == "C" || a.(string) == "D":
						return 2
					default:
						return 3
					}
				}
				bb := generic.Group(aa, grouper)
				cc := generic.SliceType2{
					generic.SliceType{"A", "B"},
					generic.SliceType{"C", "D"},
					generic.SliceType{"E", "F"},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "",
			Expectation: func(t *testing.T) {
				t.Skip()
			},
		},
	}, Specification{
		FunctionName: "GroupI",
		StandardPath: Behavior{
			Description: "Elements are grouped as expected.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C", "D", "E", "F"}
				grouper := func(i int64, a generic.PrimitiveType) int64 {
					switch {
					case i <= 1:
						return 1
					case i <= 3:
						return 2
					default:
						return 3
					}
				}
				bb := generic.GroupI(aa, grouper)
				cc := generic.SliceType2{
					generic.SliceType{"A", "B"},
					generic.SliceType{"C", "D"},
					generic.SliceType{"E", "F"},
				}
				assert.ElementsMatch(t, bb, cc)
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
		FunctionName: "Head",
		StandardPath: Behavior{
			Description: "Returns the first item from the slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				bb := generic.Head(aa)
				assert.ElementsMatch(t, bb, generic.SliceType{1})
			},
		},
		AlternativePath: Behavior{
			Description: "Returns an empty slice if the source slice is empty.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				bb := generic.Head(aa)
				if len(bb) != 0 {
					t.Error("Expected bb to be empty.")
				}
			},
		},
	},
	Specification{
		FunctionName: "InsertAfter",
		StandardPath: Behavior{
			Description: "Inserts after the first element passing the test",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 != 0
				}
				generic.InsertAfter(&aa, 9, test)
				bb := generic.SliceType{1, 9, 2, 3}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "No tests pass, inserts at end.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) > 10
				}
				generic.InsertAfter(&aa, 9, test)
				bb := generic.SliceType{1, 2, 3, 9}
				assert.ElementsMatch(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "InsertBefore",
		StandardPath: Behavior{
			Description: "Inserts before the first element passing the test.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 == 0
				}
				generic.InsertBefore(&aa, 9, test)
				bb := generic.SliceType{1, 9, 2, 3, 4}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Inserts at head if no tests pass",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 10
				}
				generic.InsertBefore(&aa, 9, test)
				bb := generic.SliceType{9, 1, 2, 3, 4}
				assert.ElementsMatch(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "InsertAt",
		StandardPath: Behavior{
			Description: "Inserts properly in middle of list.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				generic.InsertAt(&aa, 9, 2)
				bb := generic.SliceType{1, 2, 9, 3}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Inserts into an empty list.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				generic.InsertAt(&aa, 9, 0)
				bb := generic.SliceType{9}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Negative index inserted at 0",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3}
					generic.InsertAt(&aa, 9, -2)
					bb := generic.SliceType{9, 1, 2, 3}
					assert.ElementsMatch(t, aa, bb)
				},
			},
			Behavior{
				Description: "Index greater than length appended to end.",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3}
					generic.InsertAt(&aa, 9, 99)
					bb := generic.SliceType{1, 2, 3, 9}
					assert.ElementsMatch(t, aa, bb)
				},
			},
		},
	},
	Specification{
		FunctionName: "Intersection",
		StandardPath: Behavior{
			Description: "Returns a slice of the commmon items.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 4, 2, 5}
				bb := generic.SliceType{4, 3, 7, 2}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				cc := generic.Intersection(aa, bb, equal)
				dd := generic.SliceType{4, 2}
				assert.ElementsMatch(t, cc, dd)
			},
		},
		AlternativePath: Behavior{
			Description: "Duplicates are not retained",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 4, 2, 2, 2, 5, 4}
				bb := generic.SliceType{4, 3, 2, 7, 2}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				cc := generic.Intersection(aa, bb, equal)
				dd := generic.SliceType{4, 2}
				assert.ElementsMatch(t, cc, dd)
			},
		},
	},
	Specification{
		FunctionName: "IsProperSubset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a proper subset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.SliceType{1, 2, 3, 4, 5}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsProperSubset(aa, bb, equal)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a proper subset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				bb := generic.SliceType{1, 2, 3, 4, 5}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsProperSubset(aa, bb, equal)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "IsProperSuperset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a proper superset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				bb := generic.SliceType{1, 2, 3, 4}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsProperSuperset(aa, bb, equal)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a proper superset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				bb := generic.SliceType{1, 2, 3, 4, 5}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsProperSuperset(aa, bb, equal)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "IsSubset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a subset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.SliceType{1, 2, 3, 4}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsSubset(aa, bb, equal)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a subset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 7, 8, 9, 0}
				bb := generic.SliceType{1, 2, 3, 4, 5}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsSubset(aa, bb, equal)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "IsSuperset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a superset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				bb := generic.SliceType{1, 2, 3, 4}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsSuperset(aa, bb, equal)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a superset of bb",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				bb := generic.SliceType{6, 7, 8, 9, 0}
				equal := func(a, b generic.PrimitiveType) bool {
					return a.(int) == b.(int)
				}
				result := generic.IsSuperset(aa, bb, equal)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "Item",
		StandardPath: Behavior{
			Description: "Element at i is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				bb := generic.Item(aa, 1)
				cc := generic.SliceType{2}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "aa is empty, returns empty for any index",
			Expectation: func(t *testing.T) {
				for i := int64(-1); i <= 1; i++ {
					aa := generic.SliceType{}
					bb := generic.ItemFuzzy(aa, i)
					cc := generic.SliceType{}
					assert.ElementsMatch(t, bb, cc)
				}
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "i < 0 returns empty",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3}
					bb := generic.Item(aa, -1)
					cc := generic.SliceType{}
					assert.ElementsMatch(t, bb, cc)
				},
			},
			Behavior{
				Description: "i >= len(aa) returns empty",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3}
					bb := generic.Item(aa, 10)
					cc := generic.SliceType{}
					assert.ElementsMatch(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "ItemFuzzy",
		StandardPath: Behavior{
			Description: "Element at i is returned.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				bb := generic.ItemFuzzy(aa, 1)
				cc := generic.SliceType{2}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "aa is empty, returns empty for any index",
			Expectation: func(t *testing.T) {
				for i := int64(-1); i <= 1; i++ {
					aa := generic.SliceType{}
					bb := generic.ItemFuzzy(aa, i)
					cc := generic.SliceType{}
					assert.ElementsMatch(t, bb, cc)
				}
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "i < 0 returns head",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3}
					bb := generic.ItemFuzzy(aa, -1)
					cc := generic.SliceType{1}
					assert.ElementsMatch(t, bb, cc)
				},
			},
			Behavior{
				Description: "i >= len(aa) returns end",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3}
					bb := generic.ItemFuzzy(aa, 10)
					cc := generic.SliceType{3}
					assert.ElementsMatch(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "Last",
		StandardPath: Behavior{
			Description: "Returns the last that matches the expectation.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5, 6, 7, 8}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 != 0
				}
				bb := generic.Last(aa, test)
				cc := generic.SliceType{7}
				assert.ElementsMatch(t, bb, cc)
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
		FunctionName: "Len",
		StandardPath: Behavior{
			Description: "Returns the length of the slice",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5, 6, 7, 8}
				assert.Equal(t, 8, generic.Len(aa))
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
		FunctionName: "Map",
		StandardPath: Behavior{
			Description: "Applies the transform to each element",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				mapFn := func(a generic.PrimitiveType) generic.PrimitiveType {
					return a.(int) * 2
				}
				generic.Map(aa, mapFn)
				bb := generic.SliceType{2, 4, 6}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "None",
		StandardPath: Behavior{
			Description: "Returns true if the test fails for all items",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 4
				}
				assert.True(t, generic.None(aa, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if the test passes for any item",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 2
				}
				assert.False(t, generic.None(aa, test))
			},
		},
	},
	Specification{
		FunctionName: "Pairwise",
		StandardPath: Behavior{
			Description: "Processes elements pairwise",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"W", "X", "Y", "Z"}
				xform := func(a, b generic.PrimitiveType) generic.PrimitiveType {
					return a.(string) + b.(string)
				}
				bb := generic.Pairwise(aa, "V", xform)
				cc := generic.SliceType{"VW", "WX", "XY", "YZ"}
				assert.ElementsMatch(t, bb, cc)
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
		FunctionName: "Partition",
		StandardPath: Behavior{
			Description: "Parition splits the slice as expected",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5, 6}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 == 0
				}
				bb := generic.Partition(aa, test)
				cc := generic.SliceType2{generic.SliceType{2, 4, 6}, generic.SliceType{1, 3, 5}}
				assert.ElementsMatch(t, bb, cc)
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
		FunctionName: "Permutable",
		StandardPath: Behavior{
			Description: "Returns true if the slice has less than MaxInt64 permutations.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				for i := 0; i < 20; i++ {
					generic.Append(&aa, i)
				}
				assert.True(t, generic.Permutable(aa))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns if the slice has more than MaxInt64 permutations.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				for i := 0; i < 21; i++ {
					generic.Append(&aa, i)
				}
				assert.False(t, generic.Permutable(aa))
			},
		},
	},
	Specification{
		FunctionName: "Permutations",
		StandardPath: Behavior{
			Description: "Returns the correct number of possible permutations.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5, 6}
				p := generic.Permutations(aa)
				assert.Equal(t, int64(720), p.Int64())
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
		FunctionName: "Permute",
		StandardPath: Behavior{
			Description: "Creates permutations.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				bb := generic.Permute(aa)
				cc := generic.SliceType2{
					generic.SliceType{"A", "B", "C"},
					generic.SliceType{"B", "A", "C"},
					generic.SliceType{"C", "A", "B"},
					generic.SliceType{"A", "C", "B"},
					generic.SliceType{"B", "C", "A"},
					generic.SliceType{"C", "B", "A"},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Original slice is unaltered.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{"A", "B", "C"}
				generic.Permute(aa)
				cc := generic.SliceType{"A", "B", "C"}
				assert.ElementsMatch(t, aa, cc)
			},
		},
	},
	Specification{
		FunctionName: "Pop",
		StandardPath: Behavior{
			Description: "Returns the head element from the slice, removing it from the slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				bb := generic.Pop(&aa)
				cc := generic.SliceType{2, 3}
				dd := generic.SliceType{1}
				assert.ElementsMatch(t, aa, cc)
				assert.ElementsMatch(t, bb, dd)
			},
		},
		AlternativePath: Behavior{
			Description: "Slice is empty, returns an empty slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				bb := generic.Pop(&aa)
				if len(bb) > 0 {
					t.Error("Expected bb to be empty.")
				}
			},
		},
	},
	Specification{
		FunctionName: "Push",
		StandardPath: Behavior{
			Description: "Inserts a new element at the head of the slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				generic.Push(&aa, 9)
				bb := generic.SliceType{9, 1, 2, 3}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "Reduce",
		StandardPath: Behavior{
			Description: "Slice is reduced as expected.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				reducer := func(a, acc generic.PrimitiveType) generic.PrimitiveType {
					return a.(int) + acc.(int)
				}
				bb := generic.Reduce(aa, reducer)
				cc := generic.SliceType{10}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Reducing empty slice returns an empty slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				reducer := func(a, acc generic.PrimitiveType) generic.PrimitiveType {
					return a.(int) + acc.(int)
				}
				bb := generic.Reduce(aa, reducer)
				cc := generic.SliceType{}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Reducing a single element slice returns a single element slice.",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1}
					reducer := func(a, acc generic.PrimitiveType) generic.PrimitiveType {
						return a.(int) + acc.(int)
					}
					bb := generic.Reduce(aa, reducer)
					cc := generic.SliceType{1}
					assert.ElementsMatch(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "Remove",
		StandardPath: Behavior{
			Description: "Removes the items that pass the test.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5, 6}
				test := func(a generic.PrimitiveType) bool {
					return a.(int)%2 == 0
				}
				generic.Remove(&aa, test)
				bb := generic.SliceType{1, 3, 5}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Does nothing if no items satisfy test.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5, 6}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 10
				}
				generic.Remove(&aa, test)
				bb := generic.SliceType{1, 2, 3, 4, 5, 6}
				assert.ElementsMatch(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "RemoveAt",
		StandardPath: Behavior{
			Description: "Removes the item at the specified index.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.RemoveAt(&aa, 2)
				bb := generic.SliceType{1, 2, 4}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Does nothing if slice is empty.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				generic.RemoveAt(&aa, 2)
				bb := generic.SliceType{}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Does nothing if slice is nil",
				Expectation: func(t *testing.T) {
					var aa generic.SliceType
					generic.RemoveAt(&aa, 2)
					assert.Nil(t, aa)
				},
			},
			Behavior{
				Description: "Does nothing if index is negative",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3, 4}
					generic.RemoveAt(&aa, -1)
					bb := generic.SliceType{1, 2, 3, 4}
					assert.ElementsMatch(t, aa, bb)
				},
			},
			Behavior{
				Description: "Does nothing if index greater than max",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{1, 2, 3, 4}
					generic.RemoveAt(&aa, 10)
					bb := generic.SliceType{1, 2, 3, 4}
					assert.ElementsMatch(t, aa, bb)
				},
			},
		},
	},
	Specification{
		FunctionName: "Reverse",
		StandardPath: Behavior{
			Description: "Reverses slice",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.Reverse(&aa)
				bb := generic.SliceType{4, 3, 2, 1}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Empty slice has no effect",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				generic.Reverse(&aa)
				bb := generic.SliceType{}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Nil slice has no effect",
				Expectation: func(t *testing.T) {
					var aa generic.SliceType
					generic.Reverse(&aa)
					assert.Nil(t, aa)
				},
			},
		},
	},
	Specification{
		FunctionName: "Skip",
		StandardPath: Behavior{
			Description: "Skips the first n elements",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.Skip(&aa, 2)
				bb := generic.SliceType{3, 4}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Empties the list if n >= len(aa)",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.Skip(&aa, 4)
				bb := generic.SliceType{}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Empty list does nothing",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{}
					generic.Skip(&aa, 4)
					bb := generic.SliceType{}
					assert.ElementsMatch(t, aa, bb)
				},
			},
			Behavior{
				Description: "Nil list does nothing",
				Expectation: func(t *testing.T) {
					var aa generic.SliceType
					generic.Skip(&aa, 4)
					assert.Nil(t, aa)
				},
			},
			Behavior{
				Description: "n <= 0 does nothing",
				Expectation: func(t *testing.T) {
					nn := []int64{-1, 0}
					for _, n := range nn {
						aa := generic.SliceType{}
						generic.Skip(&aa, n)
						bb := generic.SliceType{}
						assert.ElementsMatch(t, aa, bb)
					}
				},
			},
		},
	},
	Specification{
		FunctionName: "SkipWhile",
		StandardPath: Behavior{
			Description: "",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) < 3
				}
				generic.SkipWhile(&aa, test)
				bb := generic.SliceType{3, 4}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Test never satisfied, does nothing",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) > 10
				}
				generic.SkipWhile(&aa, test)
				bb := generic.SliceType{1, 2, 3, 4}
				assert.ElementsMatch(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "Sort",
		StandardPath: Behavior{
			Description: "Sorts",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 3, 4, 2, 5}
				less := func(a, b generic.PrimitiveType) bool {
					return a.(int) < b.(int)
				}
				generic.Sort(&aa, less)
				bb := generic.SliceType{2, 3, 4, 5, 6}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "SplitAfter",
		StandardPath: Behavior{
			Description: "The slice is spit as expected",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 7, 8, 9}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 7
				}
				bb := generic.SplitAfter(aa, test)
				cc := generic.SliceType2{
					generic.SliceType{6, 7},
					generic.SliceType{8, 9},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "No match found, aa will be in SliceType2[0]",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 7, 8, 9}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 10
				}
				bb := generic.SplitAfter(aa, test)
				cc := generic.SliceType2{
					generic.SliceType{6, 7, 8, 9},
					generic.SliceType{},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "SplitAt",
		StandardPath: Behavior{
			Description: "The slice is split as expected",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 7, 8, 9}
				bb := generic.SplitAt(aa, 2)
				cc := generic.SliceType2{
					generic.SliceType{6, 7},
					generic.SliceType{8, 9},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "If the slice is empty, two empty slices are returned",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				bb := generic.SplitAt(aa, 2)
				cc := generic.SliceType2{
					generic.SliceType{},
					generic.SliceType{},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "If the slice is nil, two empty slices are returned",
				Expectation: func(t *testing.T) {
					var aa generic.SliceType
					bb := generic.SplitAt(aa, 2)
					cc := generic.SliceType2{
						generic.SliceType{},
						generic.SliceType{},
					}
					assert.ElementsMatch(t, bb, cc)
				},
			},
			Behavior{
				Description: "If i < 0, the full slice will be placed in SliceType2[1]",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{6, 7, 8, 9}
					bb := generic.SplitAt(aa, -1)
					cc := generic.SliceType2{
						generic.SliceType{},
						generic.SliceType{6, 7, 8, 9},
					}
					assert.ElementsMatch(t, bb, cc)
				},
			},
			Behavior{
				Description: "If i >= len(aa), the full slice will be placed in SliceType2[0]",
				Expectation: func(t *testing.T) {
					aa := generic.SliceType{6, 7, 8, 9}
					bb := generic.SplitAt(aa, 4)
					cc := generic.SliceType2{
						generic.SliceType{6, 7, 8, 9},
						generic.SliceType{},
					}
					assert.ElementsMatch(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "SplitBefore",
		StandardPath: Behavior{
			Description: "The slice is spit as expected",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 7, 8, 9}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 8
				}
				bb := generic.SplitBefore(aa, test)
				cc := generic.SliceType2{
					generic.SliceType{6, 7},
					generic.SliceType{8, 9},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "No match found, aa will be in SliceType2[0]",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{6, 7, 8, 9}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) == 10
				}
				bb := generic.SplitBefore(aa, test)
				cc := generic.SliceType2{
					generic.SliceType{6, 7, 8, 9},
					generic.SliceType{},
				}
				assert.ElementsMatch(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "String",
		StandardPath: Behavior{
			Description: "Returns string representation of slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				s := generic.String(aa)
				assert.Equal(t, "[1,2,3]", s)
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
		FunctionName: "SwapIndex",
		StandardPath: Behavior{
			Description: "Swaps the specified indices.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4, 5}
				generic.SwapIndex(aa, 2, 4)
				bb := generic.SliceType{1, 2, 5, 4, 3}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "If either index is out of range, swap does nothing.",
			Expectation: func(t *testing.T) {
				indices := [][]int64{
					{-10, -9},
					{-10, 3},
					{-10, 10},
					{3, -10},
					{3, 10},
					{10, -10},
					{10, 3},
					{10, 9},
				}
				for _, ii := range indices {
					aa := generic.SliceType{1, 2, 3, 4, 5}
					generic.SwapIndex(aa, ii[0], ii[1])
					bb := generic.SliceType{1, 2, 3, 4, 5}
					assert.ElementsMatch(t, aa, bb)
				}
			},
		},
	},
	Specification{
		FunctionName: "Tail",
		StandardPath: Behavior{
			Description: "Removes the head from the slice.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3}
				generic.Tail(&aa)
				bb := generic.SliceType{2, 3}
				assert.ElementsMatch(t, aa, bb)
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
		FunctionName: "Take",
		StandardPath: Behavior{
			Description: "Normally retains first n elements",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				generic.Take(&aa, 2)
				bb := generic.SliceType{1, 2}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "If slice is empty, Take does nothing",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{}
				generic.Take(&aa, 2)
				bb := generic.SliceType{}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "If slice is nil, Take does nothing",
				Expectation: func(t *testing.T) {
					var aa generic.SliceType
					generic.Take(&aa, 2)
					assert.Nil(t, aa)
				},
			},
		},
	},
	Specification{
		FunctionName: "TakeWhile",
		StandardPath: Behavior{
			Description: "Takes while the test is true.",
			Expectation: func(t *testing.T) {
				aa := generic.SliceType{1, 2, 3, 4}
				test := func(a generic.PrimitiveType) bool {
					return a.(int) < 3
				}
				generic.TakeWhile(&aa, test)
				bb := generic.SliceType{1, 2}
				assert.ElementsMatch(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "",
			Expectation: func(t *testing.T) {
				t.Skip()
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
