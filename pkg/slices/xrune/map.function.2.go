// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xrune

// ToBool maps a []Rune to a []Bool.
func (MapRune) ToBool(aa []rune, mapFn func(rune) bool) []bool {
	bb := []bool{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToByte maps a []Rune to a []Byte.
func (MapRune) ToByte(aa []rune, mapFn func(rune) byte) []byte {
	bb := []byte{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToComplex128 maps a []Rune to a []Complex128.
func (MapRune) ToComplex128(aa []rune, mapFn func(rune) complex128) []complex128 {
	bb := []complex128{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToComplex64 maps a []Rune to a []Complex64.
func (MapRune) ToComplex64(aa []rune, mapFn func(rune) complex64) []complex64 {
	bb := []complex64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToError maps a []Rune to a []Error.
func (MapRune) ToError(aa []rune, mapFn func(rune) error) []error {
	bb := []error{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToFloat32 maps a []Rune to a []Float32.
func (MapRune) ToFloat32(aa []rune, mapFn func(rune) float32) []float32 {
	bb := []float32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToFloat64 maps a []Rune to a []Float64.
func (MapRune) ToFloat64(aa []rune, mapFn func(rune) float64) []float64 {
	bb := []float64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt maps a []Rune to a []Int.
func (MapRune) ToInt(aa []rune, mapFn func(rune) int) []int {
	bb := []int{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt16 maps a []Rune to a []Int16.
func (MapRune) ToInt16(aa []rune, mapFn func(rune) int16) []int16 {
	bb := []int16{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt32 maps a []Rune to a []Int32.
func (MapRune) ToInt32(aa []rune, mapFn func(rune) int32) []int32 {
	bb := []int32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt64 maps a []Rune to a []Int64.
func (MapRune) ToInt64(aa []rune, mapFn func(rune) int64) []int64 {
	bb := []int64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt8 maps a []Rune to a []Int8.
func (MapRune) ToInt8(aa []rune, mapFn func(rune) int8) []int8 {
	bb := []int8{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToRune maps a []Rune to a []Rune.
func (MapRune) ToRune(aa []rune, mapFn func(rune) rune) []rune {
	bb := []rune{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToString maps a []Rune to a []String.
func (MapRune) ToString(aa []rune, mapFn func(rune) string) []string {
	bb := []string{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint maps a []Rune to a []Uint.
func (MapRune) ToUint(aa []rune, mapFn func(rune) uint) []uint {
	bb := []uint{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint16 maps a []Rune to a []Uint16.
func (MapRune) ToUint16(aa []rune, mapFn func(rune) uint16) []uint16 {
	bb := []uint16{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint32 maps a []Rune to a []Uint32.
func (MapRune) ToUint32(aa []rune, mapFn func(rune) uint32) []uint32 {
	bb := []uint32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint64 maps a []Rune to a []Uint64.
func (MapRune) ToUint64(aa []rune, mapFn func(rune) uint64) []uint64 {
	bb := []uint64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint8 maps a []Rune to a []Uint8.
func (MapRune) ToUint8(aa []rune, mapFn func(rune) uint8) []uint8 {
	bb := []uint8{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUintptr maps a []Rune to a []Uintptr.
func (MapRune) ToUintptr(aa []rune, mapFn func(rune) uintptr) []uintptr {
	bb := []uintptr{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
