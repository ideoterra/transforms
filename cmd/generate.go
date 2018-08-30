package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type primitiveType struct {
	TypeName  string
	ZeroValue interface{}
}

var primitiveTypes = []primitiveType{
	primitiveType{"int", int(0)},
}

type typeNames struct {
	DirName       string
	PrimitiveType string
	SliceType     string
	SliceType2    string
}

const basePath = `/Users/joe/workspace/go/src/github.com/jecolasurdo/transforms/pkg/slices`

func main() {
	for _, p := range primitiveTypes {
		typeNames := generateTypeNames(p)
		for _, t := range typeNames {
			newDirName := filepath.Join(basePath, t.DirName)
			err := os.MkdirAll(newDirName, os.ModeDir)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func generateTypeNames(p primitiveType) []typeNames {
	result := []typeNames{}
	oneDimensionalSliceType := typeNames{
		DirName:       p.TypeName + "slice",
		PrimitiveType: p.TypeName,
		SliceType:     strings.Title(p.TypeName) + "Slice",
		SliceType2:    strings.Title(p.TypeName) + "Slice2",
	}

	twoDimensionalSliceType := typeNames{
		DirName:       p.TypeName + "slice2",
		PrimitiveType: strings.Title(p.TypeName) + "Slice",
		SliceType:     strings.Title(p.TypeName) + "Slice2",
		SliceType2:    "[]" + strings.Title(p.TypeName) + "Slice2",
	}

	result = append(result, oneDimensionalSliceType, twoDimensionalSliceType)
	log.Println(oneDimensionalSliceType)
	log.Println(twoDimensionalSliceType)
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
