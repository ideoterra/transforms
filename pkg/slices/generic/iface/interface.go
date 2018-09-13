// Package genericiface establishes an interface that allows methods/functions
// to be decoupled from their concrete implementations.
package genericiface

import (
	"math/big"

	c "github.com/jecolasurdo/transforms/pkg/slices/generic/closures"
	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

// GenericSliceIface represents an object that knows how to do tranformations.
type GenericSliceIface interface {
	All(condition c.ConditionFn) bool
	Any(condition c.ConditionFn) bool
	Append(values ...interface{}) GenericSliceIface
	AsSlice() []interface{}
	Clear() GenericSliceIface
	Clone() GenericSliceIface
	Collect(bb []interface{}, collector func(a, b interface{}) interface{}) GenericSliceIface
	Count(condition c.ConditionFn) int64
	Dequeue() GenericSliceIface
	Difference(bb []interface{}, equality c.EqualityFn) GenericSliceIface
	Distinct(equality c.EqualityFn) GenericSliceIface
	Empty() bool
	End() GenericSliceIface
	Enqueue(a interface{}) GenericSliceIface
	Expand(expansion func(interface{}) []interface{}) GenericSliceIface
	Filter(condition c.ConditionFn) GenericSliceIface
	FindIndex(condition c.ConditionFn) int64
	First(condition c.ConditionFn) GenericSliceIface
	Fold(acc interface{}, folder func(a, acc interface{}) interface{}) GenericSliceIface
	FoldI(acc interface{}, folder func(i int64, a, acc interface{}) interface{}) GenericSliceIface
	ForEach(fn func(interface{}) shared.Continue) GenericSliceIface
	ForEachC(c int, fn func(a interface{}, cancelPending func() bool) shared.Continue) GenericSliceIface
	ForEachR(fn func(interface{}) shared.Continue) GenericSliceIface
	Group(grouper func(interface{}) int64) GenericSlice2Iface
	GroupI(grouper func(int64, interface{}) int64) GenericSlice2Iface
	Head() GenericSliceIface
	InsertAfter(b interface{}, condition c.ConditionFn) GenericSliceIface
	InsertBefore(b interface{}, condition c.ConditionFn) GenericSliceIface
	InsertAt(a interface{}, i int64) GenericSliceIface
	Intersection(bb []interface{}, equality c.EqualityFn) GenericSliceIface
	IsProperSubset(bb []interface{}, equality c.EqualityFn) bool
	IsProperSuperset(bb []interface{}, equality c.EqualityFn) bool
	IsSubset(bb []interface{}, equality c.EqualityFn) bool
	IsSuperset(bb []interface{}, equality c.EqualityFn) bool
	Item(i int64) GenericSliceIface
	ItemFuzzy(i int64) GenericSliceIface
	Last(condition c.ConditionFn) GenericSliceIface
	Len() int
	Map(mapFn func(interface{}) interface{}) GenericSliceIface
	None(condition c.ConditionFn) bool
	Pairwise(init interface{}, xform func(a, b interface{}) interface{}) GenericSliceIface
	Partition(condition c.ConditionFn) GenericSlice2Iface
	Permutable() bool
	Permutations() *big.Int
	Permute() GenericSlice2Iface
	Pop() GenericSliceIface
	Push(a interface{}) GenericSliceIface
	Reduce(reducer func(a, acc interface{}) interface{}) GenericSliceIface
	Remove(condition c.ConditionFn) GenericSliceIface
	RemoveAt(i int64) GenericSliceIface
	Reverse() GenericSliceIface
	Skip(n int64) GenericSliceIface
	SkipWhile(condition c.ConditionFn) GenericSliceIface
	Sort(less func(a, b interface{}) bool) GenericSliceIface
	SplitAfter(condition c.ConditionFn) GenericSlice2Iface
	SplitAt(i int64) GenericSlice2Iface
	SplitBefore(condition c.ConditionFn) GenericSlice2Iface
	String() string
	SwapIndex(i, j int64) GenericSliceIface
	Tail() GenericSliceIface
	Take(n int64) GenericSliceIface
	TakeWhile(condition c.ConditionFn) GenericSliceIface
	Union(bb *[]interface{}) GenericSliceIface
	Unzip() GenericSlice2Iface
	WindowCentered(windowSize int64, windowFn func(window []interface{}) interface{}) GenericSliceIface
	WindowLeft(windowSize int64, windowFn func(window []interface{}) interface{}) GenericSliceIface
	WindowRight(windowSize int64, windowFn func(window []interface{}) interface{}) GenericSliceIface
	Zip(bb *[]interface{}) GenericSliceIface
}

type GenericSlice2Iface interface{}
