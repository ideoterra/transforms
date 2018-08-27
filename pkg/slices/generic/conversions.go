package generic

import (
	"math/big"

	"github.com/cheekybits/genny/generic"
	"github.com/jecolasurdo/transforms/pkg/slices/slicetypes"
)

//Generic is a placeholder for an interface{}
type Generic generic.Type

//AsUintSlice converts a []Generic to a UintSlice
func AsUintSlice(s []Generic, conversion func(value Generic) uint) slicetypes.UintSlice {
	r := make(slicetypes.UintSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint8Slice converts a []Generic to a Uint8Slice
func AsUint8Slice(s []Generic, conversion func(value Generic) uint8) slicetypes.Uint8Slice {
	r := make(slicetypes.Uint8Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint16Slice converts a []Generic to a Uint16Slice
func AsUint16Slice(s []Generic, conversion func(value Generic) uint16) slicetypes.Uint16Slice {
	r := make(slicetypes.Uint16Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint32Slice converts a []Generic to a Uint32Slice
func AsUint32Slice(s []Generic, conversion func(value Generic) uint32) slicetypes.Uint32Slice {
	r := make(slicetypes.Uint32Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint64Slice converts a []Generic to a Uint64Slice
func AsUint64Slice(s []Generic, conversion func(value Generic) uint64) slicetypes.Uint64Slice {
	r := make(slicetypes.Uint64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsIntSlice converts a []Generic to an IntSlice
func AsIntSlice(s []Generic, conversion func(value Generic) int) slicetypes.IntSlice {
	r := make(slicetypes.IntSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt8Slice converts a []Generic to an Int8Slice
func AsInt8Slice(s []Generic, conversion func(value Generic) int8) slicetypes.Int8Slice {
	r := make(slicetypes.Int8Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt16Slice converts a []Generic to an Int16Slice
func AsInt16Slice(s []Generic, conversion func(value Generic) int16) slicetypes.Int16Slice {
	r := make(slicetypes.Int16Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt32Slice converts a []Generic to an Int32Slice
func AsInt32Slice(s []Generic, conversion func(value Generic) int32) slicetypes.Int32Slice {
	r := make(slicetypes.Int32Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt64Slice converts a []Generic to an Int64Slice
func AsInt64Slice(s []Generic, conversion func(value Generic) int64) slicetypes.Int64Slice {
	r := make(slicetypes.Int64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsFloat32Slice converts a []Generic to a Float32Slice
func AsFloat32Slice(s []Generic, conversion func(value Generic) float32) slicetypes.Float32Slice {
	r := make(slicetypes.Float32Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsFloat64Slice converts a []Generic to a Float64Slice
func AsFloat64Slice(s []Generic, conversion func(value Generic) float64) slicetypes.Float64Slice {
	r := make(slicetypes.Float64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsComplex64Slice converts a []Generic to a Complex64Slice
func AsComplex64Slice(s []Generic, conversion func(value Generic) complex64) slicetypes.Complex64Slice {
	r := make(slicetypes.Complex64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsComplex128Slice converts a []Generic to a Complex128Slice
func AsComplex128Slice(s []Generic, conversion func(value Generic) complex128) slicetypes.Complex128Slice {
	r := make(slicetypes.Complex128Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsByteSlice converts a []Generic to a ByteSlice
func AsByteSlice(s []Generic, conversion func(value Generic) byte) slicetypes.ByteSlice {
	r := make(slicetypes.ByteSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsRuneSlice converts a []Generic to a RuneSlice
func AsRuneSlice(s []Generic, conversion func(value Generic) rune) slicetypes.RuneSlice {
	r := make(slicetypes.RuneSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUintptrSlice converts a []Generic to a UintptrSlice
func AsUintptrSlice(s []Generic, conversion func(value Generic) uintptr) slicetypes.UintptrSlice {
	r := make(slicetypes.UintptrSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsStringSlice converts a []Generic to a StringSlice
func AsStringSlice(s []Generic, conversion func(value Generic) string) slicetypes.StringSlice {
	r := make(slicetypes.StringSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInterfaceSlice converts a []Generic to an InterfaceSlice
func AsInterfaceSlice(s []Generic, conversion func(value Generic) interface{}) slicetypes.InterfaceSlice {
	r := make(slicetypes.InterfaceSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsBigIntSlice converts a []Generic to a BigIntSlice
func AsBigIntSlice(s []Generic, conversion func(value Generic) *big.Int) slicetypes.BigIntSlice {
	r := make(slicetypes.BigIntSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsBigFloatSlice converts a []Generic to a BigFloatSlice
func AsBigFloatSlice(s []Generic, conversion func(value Generic) *big.Float) slicetypes.BigFloatSlice {
	r := make(slicetypes.BigFloatSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}
