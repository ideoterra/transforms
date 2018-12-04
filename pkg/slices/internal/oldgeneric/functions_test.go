package oldgeneric_test

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ideoterra/transforms/pkg/slices/internal/oldgeneric"
	"github.com/ideoterra/transforms/pkg/slices/shared"
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
				s := []interface{}{1, 2, 3, 4}
				test := func(p interface{}) bool {
					return p.(int) < 5
				}
				assert.True(t, oldgeneric.All(s, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if not all elements pass test.",
			Expectation: func(t *testing.T) {
				s := []interface{}{1, 2, 3, 4, 5}
				test := func(p interface{}) bool {
					return p.(int) < 5
				}
				assert.False(t, oldgeneric.All(s, test))
			},
		},
	},
	Specification{
		FunctionName: "Any",
		StandardPath: Behavior{
			Description: "Returns true if any of the elements match.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) == 2
				}
				assert.True(t, oldgeneric.Any(aa, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if none of the elements match.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) == 5
				}
				assert.False(t, oldgeneric.Any(aa, test))
			},
		},
	},
	Specification{
		FunctionName: "Append",
		StandardPath: Behavior{
			Description: "Values are added to the end of the slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				bb := []interface{}{5, 6, 7, 8}
				oldgeneric.Append(&aa, bb...)
				assertSlicesEqual(t, []interface{}{1, 2, 3, 4, 5, 6, 7, 8}, aa)
			},
		},
		AlternativePath: Behavior{
			Description: "No values supplied makes no change.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				bb := []interface{}{}
				oldgeneric.Append(&aa, bb...)
				assertSlicesEqual(t, []interface{}{1, 2, 3, 4}, aa)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Nil passed as aa appends bb",
				Expectation: func(t *testing.T) {
					var aa []interface{}
					bb := []interface{}{5, 6, 7, 8}
					oldgeneric.Append(&aa, bb...)
					assertSlicesEqual(t, []interface{}{5, 6, 7, 8}, aa)
				},
			},
		},
	},

	Specification{
		FunctionName: "Clear",
		StandardPath: Behavior{
			Description: "The slice is set to nil",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.Clear(&aa)
				assert.Nil(t, aa)
			},
		},
		AlternativePath: Behavior{
			Description: "An already nil slice can be cleared.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.Clear(&aa)
				oldgeneric.Append(&aa, 6, 7, 8)
				assertSlicesEqual(t, []interface{}{6, 7, 8}, aa)
			},
		},
	},
	Specification{
		FunctionName: "Clone",
		StandardPath: Behavior{
			Description: "A new identical slice is allocated in memory.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				bb := oldgeneric.Clone(aa)
				if &aa == &bb {
					t.Error("Slices aa and bb should not have the same address")
				}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Slices are not deep cloned in this operation",
			Expectation: func(t *testing.T) {
				value := 1
				aa := []interface{}{&value}
				bb := oldgeneric.Clone(aa)
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
				aa := []interface{}{"A", "B"}
				bb := []interface{}{"Y", "Z"}
				collector := func(a, b interface{}) interface{} {
					return a.(string) + b.(string)
				}
				cc := oldgeneric.Collect(aa, bb, collector)
				dd := []interface{}{"AY", "AZ", "BY", "BZ"}
				assertSlicesEqual(t, cc, dd)
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
				aa := []interface{}{1, 2, 3, 4, 5}
				test := func(a interface{}) bool {
					return a.(int)%2 == 0
				}
				assert.Equal(t, int64(2), oldgeneric.Count(aa, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns 0 if no matches",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5}
				test := func(a interface{}) bool {
					return a.(int) == 6
				}
				assert.Equal(t, int64(0), oldgeneric.Count(aa, test))
			},
		},
	},
	Specification{
		FunctionName: "Dequeue",
		StandardPath: Behavior{
			Description: "Removes and returns head.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := oldgeneric.Dequeue(&aa)
				assert.Equal(t, 1, bb[0])
				cc := []interface{}{2, 3}
				assertSlicesEqual(t, aa, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "If source slice is empty, empty slice is returned.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				bb := oldgeneric.Dequeue(&aa)
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
				aa := []interface{}{1, 2, 3, 4}
				bb := []interface{}{5, 4, 3}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				cc := oldgeneric.Difference(aa, bb, equality)
				dd := []interface{}{1, 2, 5}
				assertSlicesEqual(t, cc, dd)
			},
		},
		AlternativePath: Behavior{
			Description: `Duplicates are handled.
						  Those from aa appear first in the result. bb appear
						  second. Order and duplicates should be maintained.`,
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 3, 1, 4}
				bb := []interface{}{5, 4, 3, 5}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				cc := oldgeneric.Difference(aa, bb, equality)
				dd := []interface{}{1, 2, 1, 5, 5}
				assertSlicesEqual(t, cc, dd)
			},
		},
	},
	Specification{
		FunctionName: "Distinct",
		StandardPath: Behavior{
			Description: "Duplicates are removed from the slice, mutating the original",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"Dani", "Riley", "Dani", "Tori", "Janice"}
				equality := func(a, b interface{}) bool {
					return a.(string) == b.(string)
				}
				oldgeneric.Distinct(&aa, equality)
				bb := []interface{}{"Dani", "Riley", "Tori", "Janice"}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{}
				assert.True(t, oldgeneric.Empty(aa))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if slice is not empty",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1}
				assert.False(t, oldgeneric.Empty(aa))
			},
		},
	},
	Specification{
		FunctionName: "End",
		StandardPath: Behavior{
			Description: "Returns a slice with the last element.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := oldgeneric.End(aa)
				assert.Equal(t, 3, bb[0])
			},
		},
		AlternativePath: Behavior{
			Description: "An empty slice returns an empty slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				bb := oldgeneric.End(aa)
				assert.True(t, oldgeneric.Empty(bb))
			},
		},
	},
	Specification{
		FunctionName: "Enqueue",
		StandardPath: Behavior{
			Description: "The value is added to the head of the slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				oldgeneric.Enqueue(&aa, 4)
				bb := []interface{}{4, 1, 2, 3}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{"AB", "CD", "EF"}
				expansion := func(a interface{}) []interface{} {
					b := string(a.(string)[0])
					c := string(a.(string)[1])
					return []interface{}{b, c}
				}
				bb := oldgeneric.Expand(aa, expansion)
				cc := []interface{}{"A", "B", "C", "D", "E", "F"}
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
	Specification{
		FunctionName: "Filter",
		StandardPath: Behavior{
			Description: "Items are removed for which the test function returns true.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int)%2 == 0
				}
				oldgeneric.Filter(&aa, test)
				bb := []interface{}{1, 3}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) >= 3
				}
				n := oldgeneric.FindIndex(aa, test)
				assert.Equal(t, int64(2), n)
			},
		},
		AlternativePath: Behavior{
			Description: "If no match is found, -1 is returned.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) >= 5
				}
				n := oldgeneric.FindIndex(aa, test)
				assert.Equal(t, int64(-1), n)
			},
		},
	},
	Specification{
		FunctionName: "First",
		StandardPath: Behavior{
			Description: "The first matching element is returned.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) >= 2
				}
				bb := oldgeneric.First(aa, test)
				cc := []interface{}{2}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "An empty slice is returned if there are no matches.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) >= 5
				}
				bb := oldgeneric.First(aa, test)
				cc := []interface{}{}
				assertSlicesEqual(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "Fold",
		StandardPath: Behavior{
			Description: "Fold accumulates values properly",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				folder := func(a, acc interface{}) interface{} {
					return a.(int) + acc.(int)
				}
				bb := oldgeneric.Fold(aa, 1, folder)
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
				aa := []interface{}{"A", "B", "C"}
				folder := func(i int64, a, acc interface{}) interface{} {
					return fmt.Sprintf("%v%v%v",
						acc.(string),
						strconv.Itoa(int(i)),
						a.(string))
				}
				bb := oldgeneric.FoldI(aa, "X", folder)
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
				aa := []interface{}{"A", "B", "C"}
				result := ""
				fn := func(a interface{}) shared.Continue {
					result = result + a.(string)
					return shared.ContinueYes
				}
				oldgeneric.ForEach(aa, fn)
				assert.Equal(t, "ABC", result)
			},
		},
		AlternativePath: Behavior{
			Description: "The iterator stops if false is returned.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C"}
				result := ""
				fn := func(a interface{}) shared.Continue {
					result = result + a.(string)
					return a.(string) != "B"
				}
				oldgeneric.ForEach(aa, fn)
				assert.Equal(t, "AB", result)
			},
		},
	},
	Specification{
		FunctionName: "ForEachC",
		StandardPath: Behavior{
			Description: "Each element of the list is applied to the function",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C"}
				mu := new(sync.Mutex)
				result := ""
				fn := func(a interface{}, cancelPending func() bool) shared.Continue {
					mu.Lock()
					defer mu.Unlock()
					result = result + a.(string)
					return shared.ContinueYes
				}
				oldgeneric.ForEachC(aa, 1, fn)
				bb := []interface{}{"ABC", "BAC", "CAB", "ACB", "BCA", "CBA"}
				if !oldgeneric.Any(bb, func(b interface{}) bool {
					return b.(string) == result
				}) {
					t.Errorf("Expected a variant of 'ABC', but got '%v'", result)
				}
			},
		},
		AlternativePath: Behavior{
			Description: "The function panic if a negative pool size is specified.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C"}
				mu := new(sync.Mutex)
				result := ""
				fn := func(a interface{}, cancelPending func() bool) shared.Continue {
					mu.Lock()
					defer mu.Unlock()
					result = result + a.(string)
					return shared.ContinueYes
				}
				assert.PanicsWithValue(t,
					"ForEachC: The concurrency pool size (c) must be non-negative.",
					func() { oldgeneric.ForEachC(aa, -1, fn) })
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
					aa := []interface{}{"A", "B", "C"}
					mu := new(sync.RWMutex)
					aIsRunning := false
					bIsRunning := false
					fn := func(a interface{}, cancelPending func() bool) shared.Continue {
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
							return shared.ContinueYes
						} else {
							halt := false
							for {
								mu.RLock()
								if aIsRunning && bIsRunning {
									halt = true
								}
								mu.RUnlock()
								if halt {
									return shared.ContinueNo
								}
							}
						}
					}
					oldgeneric.ForEachC(aa, 3, fn)
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
					aa := []interface{}{"A", "B", "C"}
					mu := new(sync.RWMutex)
					aIsRunning := false
					bIsRunning := false
					aExitedCleanly := false
					bExitedCleanly := false
					fn := func(a interface{}, cancelPending func() bool) shared.Continue {
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
							return shared.ContinueYes
						} else {
							halt := false
							for {
								mu.RLock()
								if aIsRunning && bIsRunning {
									halt = true
								}
								mu.RUnlock()
								if halt {
									return shared.ContinueNo
								}
							}
						}
					}
					oldgeneric.ForEachC(aa, 3, fn)
					assert.True(t, aExitedCleanly)
					assert.True(t, bExitedCleanly)
				},
			},
			Behavior{
				Description: `Upon cancellation, no previously backlogged work
							  will be scheduled.`,
				Expectation: func(t *testing.T) {
					// Elements "A" and "B" will spawn loops, which
					// check for pending cancellation on each iteration.
					//
					// Element "C" will spawn an operation that will wait until
					// "A" and "B" are both running, at which point "C" will
					// cancel further iterations. In response "A" and "B" will
					// then identify the cancellation, and will halt.
					//
					// Element "D" spawns an operation that will wait a couple
					// seconds, ensuring that the goroutine pool stays maxed
					// at 3 while A, B, and C wind down.
					//
					// "E" will write out a value upon initialization, but
					// because the pool size is only 3, A, B, C, and D will fill
					// the pool, causing E to be backlogged. As such, E should
					// never be marshalled to a goroutine if C requests a
					// cancellation.
					aa := []interface{}{"A", "B", "C", "D", "E"}
					mu := new(sync.RWMutex)
					aIsRunning := false
					bIsRunning := false
					eStarted := false
					fn := func(a interface{}, cancelPending func() bool) shared.Continue {
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
							return shared.ContinueYes
						} else if a.(string) == "C" {
							halt := false
							for {
								mu.RLock()
								if aIsRunning && bIsRunning {
									halt = true
								}
								mu.RUnlock()
								if halt {
									return shared.ContinueNo
								}
							}
						} else if a.(string) == "D" {
							time.Sleep(2 * time.Second)
							for cancelPending() == false {
							}
							return shared.ContinueNo
						} else if a.(string) == "E" {
							mu.Lock()
							eStarted = true
							mu.Unlock()
						}
						return shared.ContinueYes
					}
					oldgeneric.ForEachC(aa, 3, fn)
					assert.False(t, eStarted)
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
				aa := []interface{}{"A", "B", "C"}
				result := ""
				fn := func(a interface{}) shared.Continue {
					result = result + a.(string)
					return true
				}
				oldgeneric.ForEachR(aa, fn)
				assert.Equal(t, "CBA", result)
			},
		},
		AlternativePath: Behavior{
			Description: `The iterator stops when the function return true.`,
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C"}
				result := ""
				fn := func(a interface{}) shared.Continue {
					result = result + a.(string)
					return a.(string) != "B"
				}
				oldgeneric.ForEachR(aa, fn)
				assert.Equal(t, "CB", result)
			},
		},
	},
	Specification{
		FunctionName: "Group",
		StandardPath: Behavior{
			Description: "Elements are grouped as expected.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C", "D", "E", "F"}
				grouper := func(a interface{}) string {
					switch {
					case a.(string) == "A" || a.(string) == "B":
						return "1"
					case a.(string) == "C" || a.(string) == "D":
						return "2"
					default:
						return "3"
					}
				}
				bb := oldgeneric.Group(aa, grouper)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{"A", "B"},
					oldgeneric.SliceType{"C", "D"},
					oldgeneric.SliceType{"E", "F"},
				}
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
	Specification{
		FunctionName: "GroupByTrait",
		StandardPath: Behavior{
			Description: "Normally groups by trait.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"pigdog", "pigs", "dog", "pigdogs", "cat", "dogs", "pig"}
				trait := func(ai, an interface{}) bool {
					return strings.Index(an.(string), ai.(string)) == 0
				}
				equality := func(a, b interface{}) bool {
					return a.(string) == b.(string)
				}
				bb := oldgeneric.GroupByTrait(aa, trait, equality)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{"dog", "dogs"},
					oldgeneric.SliceType{"cat"},
					oldgeneric.SliceType{"pigdog", "pigs", "pigdogs", "pig"},
				}
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
	Specification{
		FunctionName: "GroupI",
		StandardPath: Behavior{
			Description: "Elements are grouped as expected.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C", "D", "E", "F"}
				grouper := func(i int64, a interface{}) string {
					switch {
					case i <= 1:
						return "1"
					case i <= 3:
						return "2"
					default:
						return "3"
					}
				}
				bb := oldgeneric.GroupI(aa, grouper)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{"A", "B"},
					oldgeneric.SliceType{"C", "D"},
					oldgeneric.SliceType{"E", "F"},
				}
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
	Specification{
		FunctionName: "Head",
		StandardPath: Behavior{
			Description: "Returns the first item from the slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := oldgeneric.Head(aa)
				assertSlicesEqual(t, bb, []interface{}{1})
			},
		},
		AlternativePath: Behavior{
			Description: "Returns an empty slice if the source slice is empty.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				bb := oldgeneric.Head(aa)
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
				aa := []interface{}{1, 2, 3}
				test := func(a interface{}) bool {
					return a.(int)%2 != 0
				}
				oldgeneric.InsertAfter(&aa, 9, test)
				bb := []interface{}{1, 9, 2, 3}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "No tests pass, inserts at end.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				test := func(a interface{}) bool {
					return a.(int) > 10
				}
				oldgeneric.InsertAfter(&aa, 9, test)
				bb := []interface{}{1, 2, 3, 9}
				assertSlicesEqual(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "InsertBefore",
		StandardPath: Behavior{
			Description: "Inserts before the first element passing the test.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int)%2 == 0
				}
				oldgeneric.InsertBefore(&aa, 9, test)
				bb := []interface{}{1, 9, 2, 3, 4}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Inserts at head if no tests pass",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) == 10
				}
				oldgeneric.InsertBefore(&aa, 9, test)
				bb := []interface{}{9, 1, 2, 3, 4}
				assertSlicesEqual(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "InsertAt",
		StandardPath: Behavior{
			Description: "Inserts properly in middle of list.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				oldgeneric.InsertAt(&aa, 9, 2)
				bb := []interface{}{1, 2, 9, 3}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Inserts into an empty list.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				oldgeneric.InsertAt(&aa, 9, 0)
				bb := []interface{}{9}
				assertSlicesEqual(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Negative index inserted at 0",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3}
					oldgeneric.InsertAt(&aa, 9, -2)
					bb := []interface{}{9, 1, 2, 3}
					assertSlicesEqual(t, aa, bb)
				},
			},
			Behavior{
				Description: "Index greater than length appended to end.",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3}
					oldgeneric.InsertAt(&aa, 9, 99)
					bb := []interface{}{1, 2, 3, 9}
					assertSlicesEqual(t, aa, bb)
				},
			},
		},
	},
	Specification{
		FunctionName: "Intersection",
		StandardPath: Behavior{
			Description: "Returns a slice of the commmon items.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 4, 2, 5}
				bb := []interface{}{4, 3, 7, 2}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				cc := oldgeneric.Intersection(aa, bb, equality)
				dd := []interface{}{4, 2}
				assertSlicesEqual(t, cc, dd)
			},
		},
		AlternativePath: Behavior{
			Description: "Duplicates are not retained",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 4, 2, 2, 2, 5, 4}
				bb := []interface{}{4, 3, 2, 7, 2}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				cc := oldgeneric.Intersection(aa, bb, equality)
				dd := []interface{}{4, 2}
				assertSlicesEqual(t, cc, dd)
			},
		},
	},
	Specification{
		FunctionName: "IsProperSubset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a proper subset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				bb := []interface{}{1, 2, 3, 4, 5}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsProperSubset(aa, bb, equality)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a proper subset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5}
				bb := []interface{}{1, 2, 3, 4, 5}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsProperSubset(aa, bb, equality)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "IsProperSuperset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a proper superset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5}
				bb := []interface{}{1, 2, 3, 4}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsProperSuperset(aa, bb, equality)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a proper superset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5}
				bb := []interface{}{1, 2, 3, 4, 5}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsProperSuperset(aa, bb, equality)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "IsSubset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a subset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				bb := []interface{}{1, 2, 3, 4}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsSubset(aa, bb, equality)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a subset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{6, 7, 8, 9, 0}
				bb := []interface{}{1, 2, 3, 4, 5}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsSubset(aa, bb, equality)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "IsSuperset",
		StandardPath: Behavior{
			Description: "Returns true if aa is a superset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				bb := []interface{}{1, 2, 3, 4}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsSuperset(aa, bb, equality)
				assert.True(t, result)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if aa is not a superset of bb",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5}
				bb := []interface{}{6, 7, 8, 9, 0}
				equality := func(a, b interface{}) bool {
					return a.(int) == b.(int)
				}
				result := oldgeneric.IsSuperset(aa, bb, equality)
				assert.False(t, result)
			},
		},
	},
	Specification{
		FunctionName: "Item",
		StandardPath: Behavior{
			Description: "Element at i is returned.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := oldgeneric.Item(aa, 1)
				cc := []interface{}{2}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "aa is empty, returns empty for any index",
			Expectation: func(t *testing.T) {
				for i := int64(-1); i <= 1; i++ {
					aa := []interface{}{}
					bb := oldgeneric.ItemFuzzy(aa, i)
					cc := []interface{}{}
					assertSlicesEqual(t, bb, cc)
				}
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "i < 0 returns empty",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3}
					bb := oldgeneric.Item(aa, -1)
					cc := []interface{}{}
					assertSlicesEqual(t, bb, cc)
				},
			},
			Behavior{
				Description: "i >= len(aa) returns empty",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3}
					bb := oldgeneric.Item(aa, 10)
					cc := []interface{}{}
					assertSlicesEqual(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "ItemFuzzy",
		StandardPath: Behavior{
			Description: "Element at i is returned.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := oldgeneric.ItemFuzzy(aa, 1)
				cc := []interface{}{2}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "aa is empty, returns empty for any index",
			Expectation: func(t *testing.T) {
				for i := int64(-1); i <= 1; i++ {
					aa := []interface{}{}
					bb := oldgeneric.ItemFuzzy(aa, i)
					cc := []interface{}{}
					assertSlicesEqual(t, bb, cc)
				}
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "i < 0 returns head",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3}
					bb := oldgeneric.ItemFuzzy(aa, -1)
					cc := []interface{}{1}
					assertSlicesEqual(t, bb, cc)
				},
			},
			Behavior{
				Description: "i >= len(aa) returns end",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3}
					bb := oldgeneric.ItemFuzzy(aa, 10)
					cc := []interface{}{3}
					assertSlicesEqual(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "Last",
		StandardPath: Behavior{
			Description: "Returns the last that matches the expectation.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
				test := func(a interface{}) bool {
					return a.(int)%2 != 0
				}
				bb := oldgeneric.Last(aa, test)
				cc := []interface{}{7}
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
	Specification{
		FunctionName: "Len",
		StandardPath: Behavior{
			Description: "Returns the length of the slice",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
				assert.Equal(t, 8, oldgeneric.Len(aa))
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
		FunctionName: "Apply",
		StandardPath: Behavior{
			Description: "Applies the transform to each element",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				mapFn := func(a interface{}) interface{} {
					return a.(int) * 2
				}
				oldgeneric.Apply(&aa, mapFn)
				bb := []interface{}{2, 4, 6}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{1, 2, 3}
				test := func(a interface{}) bool {
					return a.(int) == 4
				}
				assert.True(t, oldgeneric.None(aa, test))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns false if the test passes for any item",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				test := func(a interface{}) bool {
					return a.(int) == 2
				}
				assert.False(t, oldgeneric.None(aa, test))
			},
		},
	},
	Specification{
		FunctionName: "Pairwise",
		StandardPath: Behavior{
			Description: "Processes elements pairwise",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"W", "X", "Y", "Z"}
				xform := func(a, b interface{}) interface{} {
					return a.(string) + b.(string)
				}
				bb := oldgeneric.Pairwise(aa, "V", xform)
				cc := []interface{}{"VW", "WX", "XY", "YZ"}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Returns empty slice if aa is empty",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				xform := func(a, b interface{}) interface{} {
					return a.(string) + b.(string)
				}
				bb := oldgeneric.Pairwise(aa, "V", xform)
				cc := []interface{}{}
				assertSlicesEqual(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "Partition",
		StandardPath: Behavior{
			Description: "Parition splits the slice as expected",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5, 6}
				test := func(a interface{}) bool {
					return a.(int)%2 == 0
				}
				bb := oldgeneric.Partition(aa, test)
				cc := oldgeneric.SliceType{oldgeneric.SliceType{2, 4, 6}, []interface{}{1, 3, 5}}
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
	Specification{
		FunctionName: "Permutable",
		StandardPath: Behavior{
			Description: "Returns true if the slice has less than MaxInt64 permutations.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				for i := 0; i < 20; i++ {
					oldgeneric.Append(&aa, i)
				}
				assert.True(t, oldgeneric.Permutable(aa))
			},
		},
		AlternativePath: Behavior{
			Description: "Returns if the slice has more than MaxInt64 permutations.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				for i := 0; i < 21; i++ {
					oldgeneric.Append(&aa, i)
				}
				assert.False(t, oldgeneric.Permutable(aa))
			},
		},
	},
	Specification{
		FunctionName: "Permutations",
		StandardPath: Behavior{
			Description: "Returns the correct number of possible permutations.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5, 6}
				p := oldgeneric.Permutations(aa)
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
				aa := []interface{}{"A", "B", "C"}
				bb := oldgeneric.Permute(aa)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{"A", "B", "C"},
					oldgeneric.SliceType{"B", "A", "C"},
					oldgeneric.SliceType{"C", "A", "B"},
					oldgeneric.SliceType{"A", "C", "B"},
					oldgeneric.SliceType{"B", "C", "A"},
					oldgeneric.SliceType{"C", "B", "A"},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Original slice is unaltered.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{"A", "B", "C"}
				oldgeneric.Permute(aa)
				cc := []interface{}{"A", "B", "C"}
				assertSlicesEqual(t, aa, cc)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: `Permute should panic if the number of permutations
						  would exceed MaxInt64`,
				Expectation: func(t *testing.T) {
					aa := []interface{}{}
					for n := 0; n < 21; n++ {
						oldgeneric.Append(&aa, n)
					}
					assert.Panics(t, func() { oldgeneric.Permute(aa) })
				},
			},
			Behavior{
				Description: `If source is empty, empty slice should be returned.`,
				Expectation: func(t *testing.T) {
					aa := []interface{}{}
					bb := oldgeneric.Permute(aa)
					if len(bb) != 0 {
						t.Error("Expected bb to be empty, but it was not.")
					}
				},
			},
		},
	},
	Specification{
		FunctionName: "Pop",
		StandardPath: Behavior{
			Description: "Returns the head element from the slice, removing it from the slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := oldgeneric.Pop(&aa)
				cc := []interface{}{2, 3}
				dd := []interface{}{1}
				assertSlicesEqual(t, aa, cc)
				assertSlicesEqual(t, bb, dd)
			},
		},
		AlternativePath: Behavior{
			Description: "Slice is empty, returns an empty slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				bb := oldgeneric.Pop(&aa)
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
				aa := []interface{}{1, 2, 3}
				oldgeneric.Push(&aa, 9)
				bb := []interface{}{9, 1, 2, 3}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{1, 2, 3, 4}
				reducer := func(a, acc interface{}) interface{} {
					return a.(int) + acc.(int)
				}
				bb := oldgeneric.Reduce(aa, reducer)
				cc := []interface{}{10}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Reducing empty slice returns an empty slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				reducer := func(a, acc interface{}) interface{} {
					return a.(int) + acc.(int)
				}
				bb := oldgeneric.Reduce(aa, reducer)
				cc := []interface{}{}
				assertSlicesEqual(t, bb, cc)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Reducing a single element slice returns a single element slice.",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1}
					reducer := func(a, acc interface{}) interface{} {
						return a.(int) + acc.(int)
					}
					bb := oldgeneric.Reduce(aa, reducer)
					cc := []interface{}{1}
					assertSlicesEqual(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "Remove",
		StandardPath: Behavior{
			Description: "Removes the items that pass the test.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5, 6}
				test := func(a interface{}) bool {
					return a.(int)%2 == 0
				}
				oldgeneric.Remove(&aa, test)
				bb := []interface{}{1, 3, 5}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Does nothing if no items satisfy test.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5, 6}
				test := func(a interface{}) bool {
					return a.(int) == 10
				}
				oldgeneric.Remove(&aa, test)
				bb := []interface{}{1, 2, 3, 4, 5, 6}
				assertSlicesEqual(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "RemoveAt",
		StandardPath: Behavior{
			Description: "Removes the item at the specified index.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.RemoveAt(&aa, 2)
				bb := []interface{}{1, 2, 4}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Does nothing if slice is empty.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				oldgeneric.RemoveAt(&aa, 2)
				bb := []interface{}{}
				assertSlicesEqual(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Does nothing if slice is nil",
				Expectation: func(t *testing.T) {
					var aa []interface{}
					oldgeneric.RemoveAt(&aa, 2)
					assert.Nil(t, aa)
				},
			},
			Behavior{
				Description: "Does nothing if index is negative",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3, 4}
					oldgeneric.RemoveAt(&aa, -1)
					bb := []interface{}{1, 2, 3, 4}
					assertSlicesEqual(t, aa, bb)
				},
			},
			Behavior{
				Description: "Does nothing if index greater than max",
				Expectation: func(t *testing.T) {
					aa := []interface{}{1, 2, 3, 4}
					oldgeneric.RemoveAt(&aa, 10)
					bb := []interface{}{1, 2, 3, 4}
					assertSlicesEqual(t, aa, bb)
				},
			},
		},
	},
	Specification{
		FunctionName: "Reverse",
		StandardPath: Behavior{
			Description: "Reverses slice",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.Reverse(&aa)
				bb := []interface{}{4, 3, 2, 1}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Empty slice has no effect",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				oldgeneric.Reverse(&aa)
				bb := []interface{}{}
				assertSlicesEqual(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Nil slice has no effect",
				Expectation: func(t *testing.T) {
					var aa []interface{}
					oldgeneric.Reverse(&aa)
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
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.Skip(&aa, 2)
				bb := []interface{}{3, 4}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Empties the list if n >= len(aa)",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.Skip(&aa, 4)
				bb := []interface{}{}
				assertSlicesEqual(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Empty list does nothing",
				Expectation: func(t *testing.T) {
					aa := []interface{}{}
					oldgeneric.Skip(&aa, 4)
					bb := []interface{}{}
					assertSlicesEqual(t, aa, bb)
				},
			},
			Behavior{
				Description: "Nil list does nothing",
				Expectation: func(t *testing.T) {
					var aa []interface{}
					oldgeneric.Skip(&aa, 4)
					assert.Nil(t, aa)
				},
			},
			Behavior{
				Description: "n <= 0 does nothing",
				Expectation: func(t *testing.T) {
					nn := []int64{-1, 0}
					for _, n := range nn {
						aa := []interface{}{}
						oldgeneric.Skip(&aa, n)
						bb := []interface{}{}
						assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) < 3
				}
				oldgeneric.SkipWhile(&aa, test)
				bb := []interface{}{3, 4}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "Test never satisfied, does nothing",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) > 10
				}
				oldgeneric.SkipWhile(&aa, test)
				bb := []interface{}{1, 2, 3, 4}
				assertSlicesEqual(t, aa, bb)
			},
		},
	},
	Specification{
		FunctionName: "Sort",
		StandardPath: Behavior{
			Description: "Sorts",
			Expectation: func(t *testing.T) {
				aa := []interface{}{6, 3, 4, 2, 5}
				less := func(a, b interface{}) bool {
					return a.(int) < b.(int)
				}
				oldgeneric.Sort(&aa, less)
				bb := []interface{}{2, 3, 4, 5, 6}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{6, 7, 8, 9}
				test := func(a interface{}) bool {
					return a.(int) == 7
				}
				bb := oldgeneric.SplitAfter(aa, test)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{6, 7},
					oldgeneric.SliceType{8, 9},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "No match found, aa will be in SliceType[0]",
			Expectation: func(t *testing.T) {
				aa := []interface{}{6, 7, 8, 9}
				test := func(a interface{}) bool {
					return a.(int) == 10
				}
				bb := oldgeneric.SplitAfter(aa, test)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{6, 7, 8, 9},
					oldgeneric.SliceType{},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "SplitAt",
		StandardPath: Behavior{
			Description: "The slice is split as expected",
			Expectation: func(t *testing.T) {
				aa := []interface{}{6, 7, 8, 9}
				bb := oldgeneric.SplitAt(aa, 2)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{6, 7},
					oldgeneric.SliceType{8, 9},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "If the slice is empty, two empty slices are returned",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				bb := oldgeneric.SplitAt(aa, 2)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{},
					oldgeneric.SliceType{},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "If the slice is nil, two empty slices are returned",
				Expectation: func(t *testing.T) {
					var aa oldgeneric.SliceType
					bb := oldgeneric.SplitAt(aa, 2)
					cc := oldgeneric.SliceType{
						oldgeneric.SliceType{},
						oldgeneric.SliceType{},
					}
					assertSlicesEqual(t, bb, cc)
				},
			},
			Behavior{
				Description: "If i < 0, the full slice will be placed in SliceType[1]",
				Expectation: func(t *testing.T) {
					aa := []interface{}{6, 7, 8, 9}
					bb := oldgeneric.SplitAt(aa, -1)
					cc := oldgeneric.SliceType{
						oldgeneric.SliceType{},
						oldgeneric.SliceType{6, 7, 8, 9},
					}
					assertSlicesEqual(t, bb, cc)
				},
			},
			Behavior{
				Description: "If i >= len(aa), the full slice will be placed in SliceType[0]",
				Expectation: func(t *testing.T) {
					aa := []interface{}{6, 7, 8, 9}
					bb := oldgeneric.SplitAt(aa, 4)
					cc := oldgeneric.SliceType{
						oldgeneric.SliceType{6, 7, 8, 9},
						oldgeneric.SliceType{},
					}
					assertSlicesEqual(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "SplitBefore",
		StandardPath: Behavior{
			Description: "The slice is spit as expected",
			Expectation: func(t *testing.T) {
				aa := []interface{}{6, 7, 8, 9}
				test := func(a interface{}) bool {
					return a.(int) == 8
				}
				bb := oldgeneric.SplitBefore(aa, test)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{6, 7},
					oldgeneric.SliceType{8, 9},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "No match found, aa will be in SliceType[0]",
			Expectation: func(t *testing.T) {
				aa := []interface{}{6, 7, 8, 9}
				test := func(a interface{}) bool {
					return a.(int) == 10
				}
				bb := oldgeneric.SplitBefore(aa, test)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{6, 7, 8, 9},
					oldgeneric.SliceType{},
				}
				assertSlicesEqual(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "String",
		StandardPath: Behavior{
			Description: "Returns string representation of slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				s := oldgeneric.String(aa)
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
				aa := []interface{}{1, 2, 3, 4, 5}
				oldgeneric.SwapIndex(aa, 2, 4)
				bb := []interface{}{1, 2, 5, 4, 3}
				assertSlicesEqual(t, aa, bb)
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
					aa := []interface{}{1, 2, 3, 4, 5}
					oldgeneric.SwapIndex(aa, ii[0], ii[1])
					bb := []interface{}{1, 2, 3, 4, 5}
					assertSlicesEqual(t, aa, bb)
				}
			},
		},
	},
	Specification{
		FunctionName: "Tail",
		StandardPath: Behavior{
			Description: "Removes the head from the slice.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				oldgeneric.Tail(&aa)
				bb := []interface{}{2, 3}
				assertSlicesEqual(t, aa, bb)
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
				aa := []interface{}{1, 2, 3, 4}
				oldgeneric.Take(&aa, 2)
				bb := []interface{}{1, 2}
				assertSlicesEqual(t, aa, bb)
			},
		},
		AlternativePath: Behavior{
			Description: "If slice is empty, Take does nothing",
			Expectation: func(t *testing.T) {
				aa := []interface{}{}
				oldgeneric.Take(&aa, 2)
				bb := []interface{}{}
				assertSlicesEqual(t, aa, bb)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "If slice is nil, Take does nothing",
				Expectation: func(t *testing.T) {
					var aa []interface{}
					oldgeneric.Take(&aa, 2)
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
				aa := []interface{}{1, 2, 3, 4}
				test := func(a interface{}) bool {
					return a.(int) < 3
				}
				oldgeneric.TakeWhile(&aa, test)
				bb := []interface{}{1, 2}
				assertSlicesEqual(t, aa, bb)
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
		FunctionName: "Union",
		StandardPath: Behavior{
			Description: "Union appends bb to aa",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3}
				bb := []interface{}{4, 5, 6}
				oldgeneric.Union(&aa, bb)
				cc := []interface{}{1, 2, 3, 4, 5, 6}
				assertSlicesEqual(t, aa, cc)

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
		FunctionName: "Unzip",
		StandardPath: Behavior{
			Description: "Normally unzips.",
			Expectation: func(t *testing.T) {
				aa := []interface{}{1, 2, 3, 4, 5}
				bb := oldgeneric.Unzip(aa)
				cc := oldgeneric.SliceType{
					oldgeneric.SliceType{1, 3, 5},
					oldgeneric.SliceType{2, 4},
				}
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
	Specification{
		FunctionName: "WindowCentered",
		StandardPath: Behavior{
			Description: "Basic windowing works.",
			Expectation: func(t *testing.T) {
				windowFn := func(aa []interface{}) interface{} {
					sum := 0.0
					for _, a := range aa {
						sum += float64(a.(float64))
					}
					return sum / float64(len(aa))
				}
				aa := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
				bb := oldgeneric.WindowCentered(aa, 4, windowFn)
				cc := []interface{}{1.5, 2.0, 2.5, 3.5, 4.0}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Slice order is correct for odd window size",
			Expectation: func(t *testing.T) {
				windowFn := func(aa []interface{}) interface{} {
					result := ""
					for _, a := range aa {
						result = result + a.(string)
					}
					return result
				}
				aa := []interface{}{"1", "2", "3"}
				bb := oldgeneric.WindowCentered(aa, 3, windowFn)
				cc := []interface{}{"12", "123", "23"}
				assertSlicesEqual(t, bb, cc)
			},
		},
		EdgeCases: []Behavior{
			Behavior{
				Description: "Slice order is correct for even window size.",
				Expectation: func(t *testing.T) {
					windowFn := func(aa []interface{}) interface{} {
						result := ""
						for _, a := range aa {
							result = result + a.(string)
						}
						return result
					}
					aa := []interface{}{"1", "2", "3", "4", "5"}
					bb := oldgeneric.WindowCentered(aa, 4, windowFn)
					cc := []interface{}{"12", "123", "1234", "2345", "345"}
					assertSlicesEqual(t, bb, cc)
				},
			},
		},
	},
	Specification{
		FunctionName: "WindowLeft",
		StandardPath: Behavior{
			Description: "Basic windowing works.",
			Expectation: func(t *testing.T) {
				windowFn := func(aa []interface{}) interface{} {
					sum := 0.0
					for _, a := range aa {
						sum += float64(a.(float64))
					}
					return sum / float64(len(aa))
				}
				aa := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
				bb := oldgeneric.WindowLeft(aa, 4, windowFn)
				cc := []interface{}{2.5, 3.5, 4.0, 4.5, 5.0}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Slice order is is respected",
			Expectation: func(t *testing.T) {
				windowFn := func(aa []interface{}) interface{} {
					result := ""
					for _, a := range aa {
						result = result + a.(string)
					}
					return result
				}
				aa := []interface{}{"1", "2", "3"}
				bb := oldgeneric.WindowLeft(aa, 2, windowFn)
				cc := []interface{}{"12", "23", "3"}
				assertSlicesEqual(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "WindowRight",
		StandardPath: Behavior{
			Description: "Basic windowing works.",
			Expectation: func(t *testing.T) {
				windowFn := func(aa []interface{}) interface{} {
					sum := 0.0
					for _, a := range aa {
						sum += float64(a.(float64))
					}
					return sum / float64(len(aa))
				}
				aa := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
				bb := oldgeneric.WindowRight(aa, 4, windowFn)
				cc := []interface{}{1.0, 1.5, 2.0, 2.5, 3.5}
				assertSlicesEqual(t, bb, cc)
			},
		},
		AlternativePath: Behavior{
			Description: "Slice order is is respected",
			Expectation: func(t *testing.T) {
				windowFn := func(aa []interface{}) interface{} {
					result := ""
					for _, a := range aa {
						result = result + a.(string)
					}
					return result
				}
				aa := []interface{}{"1", "2", "3"}
				bb := oldgeneric.WindowRight(aa, 2, windowFn)
				cc := []interface{}{"1", "12", "23"}
				assertSlicesEqual(t, bb, cc)
			},
		},
	},
	Specification{
		FunctionName: "Zip",
		StandardPath: Behavior{
			Description: "Interleaves aa and bb",
			Expectation: func(t *testing.T) {
				type testCase struct {
					aa oldgeneric.SliceType
					bb oldgeneric.SliceType
					dd oldgeneric.SliceType
				}

				testCases := []testCase{
					testCase{
						aa: []interface{}{1, 2, 3},
						bb: []interface{}{7, 8, 9},
						dd: []interface{}{1, 7, 2, 8, 3, 9},
					},
					testCase{
						aa: []interface{}{1, 2},
						bb: []interface{}{7, 8, 9},
						dd: []interface{}{1, 7, 2, 8, 9},
					},
					testCase{
						aa: []interface{}{1, 2, 3},
						bb: []interface{}{7, 8},
						dd: []interface{}{1, 7, 2, 8, 3},
					},
					testCase{
						aa: []interface{}{},
						bb: []interface{}{7, 8, 9},
						dd: []interface{}{7, 8, 9},
					},
					testCase{
						aa: []interface{}{1, 2, 3},
						bb: []interface{}{},
						dd: []interface{}{1, 2, 3},
					},
					testCase{
						aa: []interface{}{},
						bb: []interface{}{},
						dd: []interface{}{},
					},
				}

				for _, tc := range testCases {
					cc := oldgeneric.Zip(tc.aa, tc.bb)
					assertSlicesEqual(t, cc, tc.dd)
				}
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
	t.Parallel()

	for _, specification := range Specifications {
		t.Run(specification.FunctionName+"StandardPath", specification.StandardPath.Expectation)
		t.Run(specification.FunctionName+"AlternativePath", specification.AlternativePath.Expectation)
		for i, edgeCase := range specification.EdgeCases {
			t.Run(fmt.Sprintf("%vEdgeCase%v", specification.FunctionName, i+1), edgeCase.Expectation)
		}
	}
}

func assertSlicesEqual(t *testing.T, xx, yy []interface{}) bool {
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
