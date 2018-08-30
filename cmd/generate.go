package main

import (
	"log"
	"strings"
)

type primitiveType struct {
	TypeName  string
	ZeroValue interface{}
}

var primitiveTypes = []primitiveType{
	primitiveType{"int", int(0)},
	primitiveType{"string", ""},
}

type typeNames struct {
	FileName      string
	PrimitiveType string
	SliceType     string
	SliceType2    string
}

func generateTypeNames() []typeNames {
	result := []typeNames{}
	for _, primitiveType := range primitiveTypes {
		oneDimensionalSliceType := typeNames{
			FileName:      primitiveType.TypeName + "slice.go",
			PrimitiveType: primitiveType.TypeName,
			SliceType:     strings.Title(primitiveType.TypeName) + "Slice",
			SliceType2:    strings.Title(primitiveType.TypeName) + "Slice2",
		}

		twoDimensionalSliceType := typeNames{
			FileName:      primitiveType.TypeName + "slice2.go",
			PrimitiveType: strings.Title(primitiveType.TypeName) + "Slice",
			SliceType:     strings.Title(primitiveType.TypeName) + "Slice2",
			SliceType2:    "[]" + strings.Title(primitiveType.TypeName) + "Slice2",
		}

		result = append(result, oneDimensionalSliceType, twoDimensionalSliceType)
		log.Println(oneDimensionalSliceType)
		log.Println(twoDimensionalSliceType)
	}
	return result
}

type conversionNames struct {
	FileName                string // uses SliceTypeA
	PrimitiveTypeA          string
	PrimitiveTypeB          string
	SliceTypeA              string
	SliceTypeB              string
	PrimitiveTypeBZeroValue interface{} // for unit test generation
}

func generateConversionNames() []conversionNames {
	result := []conversionNames{}
	for i, primitiveTypeA := range primitiveTypes {
		for j, primitiveTypeB := range primitiveTypes {
			if j == i {
				continue
			}
			oneDimensionalConversion := conversionNames{
				FileName:                primitiveTypeA.TypeName + "sliceconv.go",
				PrimitiveTypeA:          primitiveTypeA.TypeName,
				PrimitiveTypeB:          primitiveTypeB.TypeName,
				PrimitiveTypeBZeroValue: primitiveTypeB.ZeroValue,
				SliceTypeA:              strings.Title(primitiveTypeA.TypeName) + "Slice",
				SliceTypeB:              strings.Title(primitiveTypeB.TypeName) + "Slice",
			}
			twoDimensionalConversion := conversionNames{
				FileName:                primitiveTypeA.TypeName + "slice2conv.go",
				PrimitiveTypeA:          strings.Title(primitiveTypeA.TypeName) + "Slice",
				PrimitiveTypeB:          strings.Title(primitiveTypeB.TypeName) + "Slice",
				PrimitiveTypeBZeroValue: primitiveTypeB.ZeroValue,
				SliceTypeA:              strings.Title(primitiveTypeA.TypeName) + "Slice2",
				SliceTypeB:              strings.Title(primitiveTypeB.TypeName) + "Slice2",
			}
			result = append(result, oneDimensionalConversion, twoDimensionalConversion)
			log.Println(oneDimensionalConversion)
			log.Println(twoDimensionalConversion)
		}
	}
	return result
}

func main() {
	generateTypeNames()
	generateConversionNames()
}
