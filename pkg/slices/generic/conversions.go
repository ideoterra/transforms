package generic

// import (
// 	"math/big"

// 	"github.com/cheekybits/genny/generic"
// 	"github.com/jecolasurdo/transforms/pkg/slices/types"
// )

// //Generic is a placeholder for an interface{}
// type Generic generic.Type

// //AsUintSlice converts a []Generic to a UintSlice
// func AsUintSlice(s []Generic, conversion func(value Generic) uint) types.UintSlice {
// 	r := make(types.UintSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsUint8Slice converts a []Generic to a Uint8Slice
// func AsUint8Slice(s []Generic, conversion func(value Generic) uint8) types.Uint8Slice {
// 	r := make(types.Uint8Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsUint16Slice converts a []Generic to a Uint16Slice
// func AsUint16Slice(s []Generic, conversion func(value Generic) uint16) types.Uint16Slice {
// 	r := make(types.Uint16Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsUint32Slice converts a []Generic to a Uint32Slice
// func AsUint32Slice(s []Generic, conversion func(value Generic) uint32) types.Uint32Slice {
// 	r := make(types.Uint32Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsUint64Slice converts a []Generic to a Uint64Slice
// func AsUint64Slice(s []Generic, conversion func(value Generic) uint64) types.Uint64Slice {
// 	r := make(types.Uint64Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsIntSlice converts a []Generic to an IntSlice
// func AsIntSlice(s []Generic, conversion func(value Generic) int) types.IntSlice {
// 	r := make(types.IntSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsInt8Slice converts a []Generic to an Int8Slice
// func AsInt8Slice(s []Generic, conversion func(value Generic) int8) types.Int8Slice {
// 	r := make(types.Int8Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsInt16Slice converts a []Generic to an Int16Slice
// func AsInt16Slice(s []Generic, conversion func(value Generic) int16) types.Int16Slice {
// 	r := make(types.Int16Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsInt32Slice converts a []Generic to an Int32Slice
// func AsInt32Slice(s []Generic, conversion func(value Generic) int32) types.Int32Slice {
// 	r := make(types.Int32Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsInt64Slice converts a []Generic to an Int64Slice
// func AsInt64Slice(s []Generic, conversion func(value Generic) int64) types.Int64Slice {
// 	r := make(types.Int64Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsFloat32Slice converts a []Generic to a Float32Slice
// func AsFloat32Slice(s []Generic, conversion func(value Generic) float32) types.Float32Slice {
// 	r := make(types.Float32Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsFloat64Slice converts a []Generic to a Float64Slice
// func AsFloat64Slice(s []Generic, conversion func(value Generic) float64) types.Float64Slice {
// 	r := make(types.Float64Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsComplex64Slice converts a []Generic to a Complex64Slice
// func AsComplex64Slice(s []Generic, conversion func(value Generic) complex64) types.Complex64Slice {
// 	r := make(types.Complex64Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsComplex128Slice converts a []Generic to a Complex128Slice
// func AsComplex128Slice(s []Generic, conversion func(value Generic) complex128) types.Complex128Slice {
// 	r := make(types.Complex128Slice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsByteSlice converts a []Generic to a ByteSlice
// func AsByteSlice(s []Generic, conversion func(value Generic) byte) types.ByteSlice {
// 	r := make(types.ByteSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsRuneSlice converts a []Generic to a RuneSlice
// func AsRuneSlice(s []Generic, conversion func(value Generic) rune) types.RuneSlice {
// 	r := make(types.RuneSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsUintptrSlice converts a []Generic to a UintptrSlice
// func AsUintptrSlice(s []Generic, conversion func(value Generic) uintptr) types.UintptrSlice {
// 	r := make(types.UintptrSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsStringSlice converts a []Generic to a StringSlice
// func AsStringSlice(s []Generic, conversion func(value Generic) string) types.StringSlice {
// 	r := make(types.StringSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsInterfaceSlice converts a []Generic to an InterfaceSlice
// func AsInterfaceSlice(s []Generic, conversion func(value Generic) interface{}) types.InterfaceSlice {
// 	r := make(types.InterfaceSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsBigIntSlice converts a []Generic to a BigIntSlice
// func AsBigIntSlice(s []Generic, conversion func(value Generic) *big.Int) types.BigIntSlice {
// 	r := make(types.BigIntSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }

// //AsBigFloatSlice converts a []Generic to a BigFloatSlice
// func AsBigFloatSlice(s []Generic, conversion func(value Generic) *big.Float) types.BigFloatSlice {
// 	r := make(types.BigFloatSlice, len(s))
// 	for i, v := range s {
// 		r[i] = conversion(v)
// 	}
// 	return r
// }
