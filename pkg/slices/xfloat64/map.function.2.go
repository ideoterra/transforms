// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package xfloat64

// ToBool maps a []Float64 to a []Bool.
func (MapFloat64) ToBool(aa []float64, mapFn func(float64) bool) []bool {
	bb := []bool{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToByte maps a []Float64 to a []Byte.
func (MapFloat64) ToByte(aa []float64, mapFn func(float64) byte) []byte {
	bb := []byte{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToComplex128 maps a []Float64 to a []Complex128.
func (MapFloat64) ToComplex128(aa []float64, mapFn func(float64) complex128) []complex128 {
	bb := []complex128{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToComplex64 maps a []Float64 to a []Complex64.
func (MapFloat64) ToComplex64(aa []float64, mapFn func(float64) complex64) []complex64 {
	bb := []complex64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToError maps a []Float64 to a []Error.
func (MapFloat64) ToError(aa []float64, mapFn func(float64) error) []error {
	bb := []error{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToFloat32 maps a []Float64 to a []Float32.
func (MapFloat64) ToFloat32(aa []float64, mapFn func(float64) float32) []float32 {
	bb := []float32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToFloat64 maps a []Float64 to a []Float64.
func (MapFloat64) ToFloat64(aa []float64, mapFn func(float64) float64) []float64 {
	bb := []float64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt maps a []Float64 to a []Int.
func (MapFloat64) ToInt(aa []float64, mapFn func(float64) int) []int {
	bb := []int{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt16 maps a []Float64 to a []Int16.
func (MapFloat64) ToInt16(aa []float64, mapFn func(float64) int16) []int16 {
	bb := []int16{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt32 maps a []Float64 to a []Int32.
func (MapFloat64) ToInt32(aa []float64, mapFn func(float64) int32) []int32 {
	bb := []int32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt64 maps a []Float64 to a []Int64.
func (MapFloat64) ToInt64(aa []float64, mapFn func(float64) int64) []int64 {
	bb := []int64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToInt8 maps a []Float64 to a []Int8.
func (MapFloat64) ToInt8(aa []float64, mapFn func(float64) int8) []int8 {
	bb := []int8{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToRune maps a []Float64 to a []Rune.
func (MapFloat64) ToRune(aa []float64, mapFn func(float64) rune) []rune {
	bb := []rune{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToString maps a []Float64 to a []String.
func (MapFloat64) ToString(aa []float64, mapFn func(float64) string) []string {
	bb := []string{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint maps a []Float64 to a []Uint.
func (MapFloat64) ToUint(aa []float64, mapFn func(float64) uint) []uint {
	bb := []uint{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint16 maps a []Float64 to a []Uint16.
func (MapFloat64) ToUint16(aa []float64, mapFn func(float64) uint16) []uint16 {
	bb := []uint16{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint32 maps a []Float64 to a []Uint32.
func (MapFloat64) ToUint32(aa []float64, mapFn func(float64) uint32) []uint32 {
	bb := []uint32{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint64 maps a []Float64 to a []Uint64.
func (MapFloat64) ToUint64(aa []float64, mapFn func(float64) uint64) []uint64 {
	bb := []uint64{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUint8 maps a []Float64 to a []Uint8.
func (MapFloat64) ToUint8(aa []float64, mapFn func(float64) uint8) []uint8 {
	bb := []uint8{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}

// ToUintptr maps a []Float64 to a []Uintptr.
func (MapFloat64) ToUintptr(aa []float64, mapFn func(float64) uintptr) []uintptr {
	bb := []uintptr{}
	for _, a := range aa {
		bb = append(bb, mapFn(a))
	}
	return bb
}
