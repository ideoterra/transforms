package genericiface

import (
	"math/big"

	"github.com/jecolasurdo/transforms/pkg/slices/shared"
)

type GenericSliceIface interface {
	Any(test func(interface{}) bool) bool
	Append(values ...interface{}) *GenericSliceIface
	Clear() *GenericSliceIface
	Clone() *GenericSliceIface
	Collect(bb *GenericSliceIface, collector func(a, b interface{}) interface{}) *GenericSliceIface
	Count(test func(interface{}) bool) int64
	Dequeue() *GenericSliceIface
	Difference(bb *GenericSliceIface, equality func(a, b interface{}) bool) *GenericSliceIface
	Distinct(equality func(a, b interface{}) bool) *GenericSliceIface
	Empty() bool
	End() *GenericSliceIface
	Enqueue(a interface{}) *GenericSliceIface
	Expand(expansion func(interface{}) GenericSliceIface) *GenericSliceIface
	Filter(test func(interface{}) bool) *GenericSliceIface
	FindIndex(test func(interface{}) bool) int64
	First(test func(interface{}) bool) *GenericSliceIface
	Fold(acc interface{}, folder func(a, acc interface{}) interface{}) *GenericSliceIface
	FoldI(acc interface{}, folder func(i int64, a, acc interface{}) interface{}) *GenericSliceIface
	ForEach(fn func(interface{}) shared.Continue) *GenericSliceIface
	ForEachC(c int, fn func(a interface{}, cancelPending func() bool) shared.Continue) *GenericSliceIface
	ForEachR(fn func(interface{}) shared.Continue) *GenericSliceIface
	Group(grouper func(interface{}) int64) *GenericSlice2Iface
	GroupI(grouper func(int64, interface{}) int64) *GenericSlice2Iface
	Head() *GenericSliceIface
	InsertAfter(b interface{}, test func(interface{}) bool) *GenericSliceIface
	InsertBefore(b interface{}, test func(interface{}) bool) *GenericSliceIface
	InsertAt(a interface{}, i int64) *GenericSliceIface
	Intersection(bb *GenericSliceIface, equality func(a, b interface{}) bool) *GenericSliceIface
	IsProperSubset(bb *GenericSliceIface, equality func(a, b interface{}) bool) bool
	IsProperSuperset(bb *GenericSliceIface, equality func(a, b interface{}) bool) bool
	IsSubset(bb *GenericSliceIface, equality func(a, b interface{}) bool) bool
	IsSuperset(bb *GenericSliceIface, equality func(a, b interface{}) bool) bool
	Item(i int64) *GenericSliceIface
	ItemFuzzy(i int64) *GenericSliceIface
	Last(test func(interface{}) bool) *GenericSliceIface
	Len() int
	Map(mapFn func(interface{}) interface{}) *GenericSliceIface
	None(test func(interface{}) bool) bool
	Pairwise(init interface{}, xform func(a, b interface{}) interface{}) *GenericSliceIface
	Partition(test func(interface{}) bool) *GenericSlice2Iface
	Permutable() bool
	Permutations() *big.Int
	Permute() *GenericSlice2Iface
	Pop() *GenericSliceIface
	Push(a interface{}) *GenericSliceIface
	Reduce(reducer func(a, acc interface{}) interface{}) *GenericSliceIface
	Remove(test func(interface{}) bool) *GenericSliceIface
	RemoveAt(i int64) *GenericSliceIface
	Reverse() *GenericSliceIface
	Skip(n int64) *GenericSliceIface
	SkipWhile(test func(interface{}) bool) *GenericSliceIface
	Sort(less func(a, b interface{}) bool) *GenericSliceIface
	SplitAfter(test func(interface{}) bool) *GenericSlice2Iface
	SplitAt(i int64) *GenericSlice2Iface
	SplitBefore(test func(interface{}) bool) *GenericSlice2Iface
	String() string
	SwapIndex(i, j int64) *GenericSliceIface
	Tail() *GenericSliceIface
	Take(n int64) *GenericSliceIface
	TakeWhile(test func(interface{}) bool) *GenericSliceIface
	Union(bb *GenericSliceIface) *GenericSliceIface
	Unzip() *GenericSlice2Iface
	WindowCentered(windowSize int64, windowFn func(window GenericSliceIface) interface{}) *GenericSliceIface
	WindowLeft(windowSize int64, windowFn func(window GenericSliceIface) interface{}) *GenericSliceIface
	WindowRight(windowSize int64, windowFn func(window GenericSliceIface) interface{}) *GenericSliceIface
	Zip(bb *GenericSliceIface) *GenericSliceIface
}

type GenericSlice2Iface interface {
	Any(test func(GenericSliceIface) bool) bool
	Append(values ...GenericSliceIface) *GenericSlice2Iface
	Clear() *GenericSlice2Iface
	Clone() *GenericSlice2Iface
	Collect(bb *GenericSlice2Iface, collector func(a, b GenericSliceIface) GenericSliceIface) *GenericSlice2Iface
	Count(test func(GenericSliceIface) bool) int64
	Dequeue() *GenericSlice2Iface
	Difference(bb *GenericSlice2Iface, equality func(a, b GenericSliceIface) bool) *GenericSlice2Iface
	Distinct(equality func(a, b GenericSliceIface) bool) *GenericSlice2Iface
	Empty() bool
	End() *GenericSlice2Iface
	Enqueue(a GenericSliceIface) *GenericSlice2Iface
	Expand(expansion func(GenericSliceIface) GenericSlice2Iface) *GenericSlice2Iface
	Filter(test func(GenericSliceIface) bool) *GenericSlice2Iface
	FindIndex(test func(GenericSliceIface) bool) int64
	First(test func(GenericSliceIface) bool) *GenericSlice2Iface
	Fold(acc GenericSliceIface, folder func(a, acc GenericSliceIface) GenericSliceIface) *GenericSlice2Iface
	FoldI(acc GenericSliceIface, folder func(i int64, a, acc GenericSliceIface) GenericSliceIface) *GenericSlice2Iface
	ForEach(fn func(GenericSliceIface) shared.Continue) *GenericSlice2Iface
	ForEachC(c int, fn func(a GenericSliceIface, cancelPending func() bool) shared.Continue) *GenericSlice2Iface
	ForEachR(fn func(GenericSliceIface) shared.Continue) *GenericSlice2Iface
	Group(grouper func(GenericSliceIface) int64) *[]GenericSlice2Iface
	GroupI(grouper func(int64, GenericSliceIface) int64) *[]GenericSlice2Iface
	Head() *GenericSlice2Iface
	InsertAfter(b GenericSliceIface, test func(GenericSliceIface) bool) *GenericSlice2Iface
	InsertBefore(b GenericSliceIface, test func(GenericSliceIface) bool) *GenericSlice2Iface
	InsertAt(a GenericSliceIface, i int64) *GenericSlice2Iface
	Intersection(bb *GenericSlice2Iface, equality func(a, b GenericSliceIface) bool) *GenericSlice2Iface
	IsProperSubset(bb *GenericSlice2Iface, equality func(a, b GenericSliceIface) bool) bool
	IsProperSuperset(bb *GenericSlice2Iface, equality func(a, b GenericSliceIface) bool) bool
	IsSubset(bb *GenericSlice2Iface, equality func(a, b GenericSliceIface) bool) bool
	IsSuperset(bb *GenericSlice2Iface, equality func(a, b GenericSliceIface) bool) bool
	Item(i int64) *GenericSlice2Iface
	ItemFuzzy(i int64) *GenericSlice2Iface
	Last(test func(GenericSliceIface) bool) *GenericSlice2Iface
	Len() int
	Map(mapFn func(GenericSliceIface) GenericSliceIface) *GenericSlice2Iface
	None(test func(GenericSliceIface) bool) bool
	Pairwise(init GenericSliceIface, xform func(a, b GenericSliceIface) GenericSliceIface) *GenericSlice2Iface
	Partition(test func(GenericSliceIface) bool) *[]GenericSlice2Iface
	Permutable() bool
	Permutations() *big.Int
	Permute() *[]GenericSlice2Iface
	Pop() *GenericSlice2Iface
	Push(a GenericSliceIface) *GenericSlice2Iface
	Reduce(reducer func(a, acc GenericSliceIface) GenericSliceIface) *GenericSlice2Iface
	Remove(test func(GenericSliceIface) bool) *GenericSlice2Iface
	RemoveAt(i int64) *GenericSlice2Iface
	Reverse() *GenericSlice2Iface
	Skip(n int64) *GenericSlice2Iface
	SkipWhile(test func(GenericSliceIface) bool) *GenericSlice2Iface
	Sort(less func(a, b GenericSliceIface) bool) *GenericSlice2Iface
	SplitAfter(test func(GenericSliceIface) bool) *[]GenericSlice2Iface
	SplitAt(i int64) *[]GenericSlice2Iface
	SplitBefore(test func(GenericSliceIface) bool) *[]GenericSlice2Iface
	String() string
	SwapIndex(i, j int64) *GenericSlice2Iface
	Tail() *GenericSlice2Iface
	Take(n int64) *GenericSlice2Iface
	TakeWhile(test func(GenericSliceIface) bool) *GenericSlice2Iface
	Union(bb *GenericSlice2Iface) *GenericSlice2Iface
	Unzip() *[]GenericSlice2Iface
	WindowCentered(windowSize int64, windowFn func(window GenericSlice2Iface) GenericSliceIface) *GenericSlice2Iface
	WindowLeft(windowSize int64, windowFn func(window GenericSlice2Iface) GenericSliceIface) *GenericSlice2Iface
	WindowRight(windowSize int64, windowFn func(window GenericSlice2Iface) GenericSliceIface) *GenericSlice2Iface
	Zip(bb *GenericSlice2Iface) *GenericSlice2Iface
}
