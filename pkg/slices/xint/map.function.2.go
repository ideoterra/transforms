// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xint

// ToBool maps a []Int to a []Bool.
func (MapInt) ToBool(aa []int, mapFn func(int) bool) []bool {
	bb := []bool{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToByte maps a []Int to a []Byte.
func (MapInt) ToByte(aa []int, mapFn func(int) byte) []byte {
	bb := []byte{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToComplex128 maps a []Int to a []Complex128.
func (MapInt) ToComplex128(aa []int, mapFn func(int) complex128) []complex128 {
	bb := []complex128{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToComplex64 maps a []Int to a []Complex64.
func (MapInt) ToComplex64(aa []int, mapFn func(int) complex64) []complex64 {
	bb := []complex64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToError maps a []Int to a []Error.
func (MapInt) ToError(aa []int, mapFn func(int) error) []error {
	bb := []error{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToFloat32 maps a []Int to a []Float32.
func (MapInt) ToFloat32(aa []int, mapFn func(int) float32) []float32 {
	bb := []float32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToFloat64 maps a []Int to a []Float64.
func (MapInt) ToFloat64(aa []int, mapFn func(int) float64) []float64 {
	bb := []float64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt maps a []Int to a []Int.
func (MapInt) ToInt(aa []int, mapFn func(int) int) []int {
	bb := []int{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt16 maps a []Int to a []Int16.
func (MapInt) ToInt16(aa []int, mapFn func(int) int16) []int16 {
	bb := []int16{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt32 maps a []Int to a []Int32.
func (MapInt) ToInt32(aa []int, mapFn func(int) int32) []int32 {
	bb := []int32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt64 maps a []Int to a []Int64.
func (MapInt) ToInt64(aa []int, mapFn func(int) int64) []int64 {
	bb := []int64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt8 maps a []Int to a []Int8.
func (MapInt) ToInt8(aa []int, mapFn func(int) int8) []int8 {
	bb := []int8{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToRune maps a []Int to a []Rune.
func (MapInt) ToRune(aa []int, mapFn func(int) rune) []rune {
	bb := []rune{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToString maps a []Int to a []String.
func (MapInt) ToString(aa []int, mapFn func(int) string) []string {
	bb := []string{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint maps a []Int to a []Uint.
func (MapInt) ToUint(aa []int, mapFn func(int) uint) []uint {
	bb := []uint{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint16 maps a []Int to a []Uint16.
func (MapInt) ToUint16(aa []int, mapFn func(int) uint16) []uint16 {
	bb := []uint16{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint32 maps a []Int to a []Uint32.
func (MapInt) ToUint32(aa []int, mapFn func(int) uint32) []uint32 {
	bb := []uint32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint64 maps a []Int to a []Uint64.
func (MapInt) ToUint64(aa []int, mapFn func(int) uint64) []uint64 {
	bb := []uint64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint8 maps a []Int to a []Uint8.
func (MapInt) ToUint8(aa []int, mapFn func(int) uint8) []uint8 {
	bb := []uint8{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUintptr maps a []Int to a []Uintptr.
func (MapInt) ToUintptr(aa []int, mapFn func(int) uintptr) []uintptr {
	bb := []uintptr{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
