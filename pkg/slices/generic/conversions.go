package generic

import (
	"math/big"

	"github.com/cheekybits/genny/generic"
	"github.com/jecolasurdo/transforms/pkg/slices/slicetypes"
)

//Generic is a placeholder for an interface{}
type Generic generic.Type

//AsUintSlice converts a generic slice to a uint slice
func AsUintSlice(genericSlice []Generic, conversion func(value Generic) uint) slicetypes.UintSlice {
	r := make(slicetypes.UintSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsUint8Slice converts a generic slice to a uint8 slice
func AsUint8Slice(genericSlice []Generic, conversion func(value Generic) uint8) slicetypes.Uint8Slice {
	r := make(slicetypes.Uint8Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsUint16Slice converts a generic slice to a uint16 slice
func AsUint16Slice(genericSlice []Generic, conversion func(value Generic) uint16) slicetypes.Uint16Slice {
	r := make(slicetypes.Uint16Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsUint32Slice converts a generic slice to a uint32 slice
func AsUint32Slice(genericSlice []Generic, conversion func(value Generic) uint32) slicetypes.Uint32Slice {
	r := make(slicetypes.Uint32Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsUint64Slice converts a generic slice to a uint64 slice
func AsUint64Slice(genericSlice []Generic, conversion func(value Generic) uint64) slicetypes.Uint64Slice {
	r := make(slicetypes.Uint64Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsIntSlice converts a generic slice to an int slice
func AsIntSlice(genericSlice []Generic, conversion func(value Generic) int) slicetypes.IntSlice {
	r := make(slicetypes.IntSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsInt8Slice converts a generic slice to an int8 slice
func AsInt8Slice(genericSlice []Generic, conversion func(value Generic) int8) slicetypes.Int8Slice {
	r := make(slicetypes.Int8Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsInt16Slice converts a generic slice to an int64 slice
func AsInt16Slice(genericSlice []Generic, conversion func(value Generic) int16) slicetypes.Int16Slice {
	r := make(slicetypes.Int16Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsInt32Slice converts a generic slice to an Int32Slice
func AsInt32Slice(genericSlice []Generic, conversion func(value Generic) int32) slicetypes.Int32Slice {
	r := make(slicetypes.Int32Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsInt64Slice converts a generic slice to an Int64Slice
func AsInt64Slice(genericSlice []Generic, conversion func(value Generic) int64) slicetypes.Int64Slice {
	r := make(slicetypes.Int64Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsFloat32Slice converts a generic slice to a Float32Slice
func AsFloat32Slice(genericSlice []Generic, conversion func(value Generic) float32) slicetypes.Float32Slice {
	r := make(slicetypes.Float32Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsFloat64Slice converts a generic slice to a Float64Slice
func AsFloat64Slice(genericSlice []Generic, conversion func(value Generic) float64) slicetypes.Float64Slice {
	r := make(slicetypes.Float64Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsComplex64Slice converts a generic slice to a Complex64Slice
func AsComplex64Slice(genericSlice []Generic, conversion func(value Generic) complex64) slicetypes.Complex64Slice {
	r := make(slicetypes.Complex64Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsComplex128Slice converts a generic slice to a Complex128Slice
func AsComplex128Slice(genericSlice []Generic, conversion func(value Generic) complex128) slicetypes.Complex128Slice {
	r := make(slicetypes.Complex128Slice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsByteSlice converts a generic slice to a ByteSlice
func AsByteSlice(genericSlice []Generic, conversion func(value Generic) byte) slicetypes.ByteSlice {
	r := make(slicetypes.ByteSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsRuneSlice converts a generic slice to a RuneSlice
func AsRuneSlice(genericSlice []Generic, conversion func(value Generic) rune) slicetypes.RuneSlice {
	r := make(slicetypes.RuneSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsUintptrSlice converts a generic slice to a UintptrSlice
func AsUintptrSlice(genericSlice []Generic, conversion func(value Generic) uintptr) slicetypes.UintptrSlice {
	r := make(slicetypes.UintptrSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsStringSlice converts a generic slice to a StringSlice
func AsStringSlice(genericSlice []Generic, conversion func(value Generic) string) slicetypes.StringSlice {
	r := make(slicetypes.StringSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsInterfaceSlice converts a generic slice to an InterfaceSlice
func AsInterfaceSlice(genericSlice []Generic, conversion func(value Generic) interface{}) slicetypes.InterfaceSlice {
	r := make(slicetypes.InterfaceSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsBigIntSlice converts a generic slice to a BigIntSlice
func AsBigIntSlice(genericSlice []Generic, conversion func(value Generic) *big.Int) slicetypes.BigIntSlice {
	r := make(slicetypes.BigIntSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}

//AsBigFloatSlice converts a generic slice to a BigFloatSlice
func AsBigFloatSlice(genericSlice []Generic, conversion func(value Generic) *big.Float) slicetypes.BigFloatSlice {
	r := make(slicetypes.BigFloatSlice, len(genericSlice))
	for i, v := range genericSlice {
		r[i] = conversion(v)
	}
	return r
}
