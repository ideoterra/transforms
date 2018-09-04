package genericiface

import (
	"math/big"

	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

type GenericSliceIface interface {
	All(test func(interface{}) bool) bool
	Any(test func(interface{}) bool) bool
	Append(values ...interface{}) GenericSliceIface
	Clear() GenericSliceIface
	Clone() GenericSliceIface
	Collect(bb []interface{}, collector func(a, b interface{}) interface{}) GenericSliceIface
	Count(test func(interface{}) bool) int64
	Dequeue() GenericSliceIface
	Difference(bb []interface{}, equality func(a, b interface{}) bool) GenericSliceIface
	Distinct(equality func(a, b interface{}) bool) GenericSliceIface
	Empty() bool
	End() GenericSliceIface
	Enqueue(a interface{}) GenericSliceIface
	Expand(expansion func(interface{}) []interface{}) GenericSliceIface
	Filter(test func(interface{}) bool) GenericSliceIface
	FindIndex(test func(interface{}) bool) int64
	First(test func(interface{}) bool) GenericSliceIface
	Fold(acc interface{}, folder func(a, acc interface{}) interface{}) GenericSliceIface
	FoldI(acc interface{}, folder func(i int64, a, acc interface{}) interface{}) GenericSliceIface
	ForEach(fn func(interface{}) shared.Continue) GenericSliceIface
	ForEachC(c int, fn func(a interface{}, cancelPending func() bool) shared.Continue) GenericSliceIface
	ForEachR(fn func(interface{}) shared.Continue) GenericSliceIface
	Group(grouper func(interface{}) int64) GenericSlice2Iface
	GroupI(grouper func(int64, interface{}) int64) GenericSlice2Iface
	Head() GenericSliceIface
	InsertAfter(b interface{}, test func(interface{}) bool) GenericSliceIface
	InsertBefore(b interface{}, test func(interface{}) bool) GenericSliceIface
	InsertAt(a interface{}, i int64) GenericSliceIface
	Intersection(bb []interface{}, equality func(a, b interface{}) bool) GenericSliceIface
	IsProperSubset(bb []interface{}, equality func(a, b interface{}) bool) bool
	IsProperSuperset(bb []interface{}, equality func(a, b interface{}) bool) bool
	IsSubset(bb []interface{}, equality func(a, b interface{}) bool) bool
	IsSuperset(bb []interface{}, equality func(a, b interface{}) bool) bool
	Item(i int64) GenericSliceIface
	ItemFuzzy(i int64) GenericSliceIface
	Last(test func(interface{}) bool) GenericSliceIface
	Len() int
	Map(mapFn func(interface{}) interface{}) GenericSliceIface
	None(test func(interface{}) bool) bool
	Pairwise(init interface{}, xform func(a, b interface{}) interface{}) GenericSliceIface
	Partition(test func(interface{}) bool) GenericSlice2Iface
	Permutable() bool
	Permutations() *big.Int
	Permute() GenericSlice2Iface
	Pop() GenericSliceIface
	Push(a interface{}) GenericSliceIface
	Reduce(reducer func(a, acc interface{}) interface{}) GenericSliceIface
	Remove(test func(interface{}) bool) GenericSliceIface
	RemoveAt(i int64) GenericSliceIface
	Reverse() GenericSliceIface
	Skip(n int64) GenericSliceIface
	SkipWhile(test func(interface{}) bool) GenericSliceIface
	Sort(less func(a, b interface{}) bool) GenericSliceIface
	SplitAfter(test func(interface{}) bool) GenericSlice2Iface
	SplitAt(i int64) GenericSlice2Iface
	SplitBefore(test func(interface{}) bool) GenericSlice2Iface
	String() string
	SwapIndex(i, j int64) GenericSliceIface
	Tail() GenericSliceIface
	Take(n int64) GenericSliceIface
	TakeWhile(test func(interface{}) bool) GenericSliceIface
	Union(bb *[]interface{}) GenericSliceIface
	Unzip() GenericSlice2Iface
	WindowCentered(windowSize int64, windowFn func(window []interface{}) interface{}) GenericSliceIface
	WindowLeft(windowSize int64, windowFn func(window []interface{}) interface{}) GenericSliceIface
	WindowRight(windowSize int64, windowFn func(window []interface{}) interface{}) GenericSliceIface
	Zip(bb *[]interface{}) GenericSliceIface
}

type GenericSlice2Iface interface{}
