package generic

import (
	"math/big"

	"github.com/cheekybits/genny/generic"
)

//Generic is a placeholder for an interface{}
type Generic generic.Type

//AsUintSlice converts a []Generic to a UintSlice
func AsUintSlice(s []Generic, conversion func(value Generic) uint) UintSlice {
	r := make(UintSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint8Slice converts a []Generic to a Uint8Slice
func AsUint8Slice(s []Generic, conversion func(value Generic) uint8) Uint8Slice {
	r := make(Uint8Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint16Slice converts a []Generic to a Uint16Slice
func AsUint16Slice(s []Generic, conversion func(value Generic) uint16) Uint16Slice {
	r := make(Uint16Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint32Slice converts a []Generic to a Uint32Slice
func AsUint32Slice(s []Generic, conversion func(value Generic) uint32) Uint32Slice {
	r := make(Uint32Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUint64Slice converts a []Generic to a Uint64Slice
func AsUint64Slice(s []Generic, conversion func(value Generic) uint64) Uint64Slice {
	r := make(Uint64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsIntSlice converts a []Generic to an IntSlice
func AsIntSlice(s []Generic, conversion func(value Generic) int) IntSlice {
	r := make(IntSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt8Slice converts a []Generic to an Int8Slice
func AsInt8Slice(s []Generic, conversion func(value Generic) int8) Int8Slice {
	r := make(Int8Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt16Slice converts a []Generic to an Int16Slice
func AsInt16Slice(s []Generic, conversion func(value Generic) int16) Int16Slice {
	r := make(Int16Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt32Slice converts a []Generic to an Int32Slice
func AsInt32Slice(s []Generic, conversion func(value Generic) int32) Int32Slice {
	r := make(Int32Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInt64Slice converts a []Generic to an Int64Slice
func AsInt64Slice(s []Generic, conversion func(value Generic) int64) Int64Slice {
	r := make(Int64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsFloat32Slice converts a []Generic to a Float32Slice
func AsFloat32Slice(s []Generic, conversion func(value Generic) float32) Float32Slice {
	r := make(Float32Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsFloat64Slice converts a []Generic to a Float64Slice
func AsFloat64Slice(s []Generic, conversion func(value Generic) float64) Float64Slice {
	r := make(Float64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsComplex64Slice converts a []Generic to a Complex64Slice
func AsComplex64Slice(s []Generic, conversion func(value Generic) complex64) Complex64Slice {
	r := make(Complex64Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsComplex128Slice converts a []Generic to a Complex128Slice
func AsComplex128Slice(s []Generic, conversion func(value Generic) complex128) Complex128Slice {
	r := make(Complex128Slice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsByteSlice converts a []Generic to a ByteSlice
func AsByteSlice(s []Generic, conversion func(value Generic) byte) ByteSlice {
	r := make(ByteSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsRuneSlice converts a []Generic to a RuneSlice
func AsRuneSlice(s []Generic, conversion func(value Generic) rune) RuneSlice {
	r := make(RuneSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsUintptrSlice converts a []Generic to a UintptrSlice
func AsUintptrSlice(s []Generic, conversion func(value Generic) uintptr) UintptrSlice {
	r := make(UintptrSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsStringSlice converts a []Generic to a StringSlice
func AsStringSlice(s []Generic, conversion func(value Generic) string) StringSlice {
	r := make(StringSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsInterfaceSlice converts a []Generic to an InterfaceSlice
func AsInterfaceSlice(s []Generic, conversion func(value Generic) interface{}) InterfaceSlice {
	r := make(InterfaceSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsBigIntSlice converts a []Generic to a BigIntSlice
func AsBigIntSlice(s []Generic, conversion func(value Generic) *big.Int) BigIntSlice {
	r := make(BigIntSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}

//AsBigFloatSlice converts a []Generic to a BigFloatSlice
func AsBigFloatSlice(s []Generic, conversion func(value Generic) *big.Float) BigFloatSlice {
	r := make(BigFloatSlice, len(s))
	for i, v := range s {
		r[i] = conversion(v)
	}
	return r
}
