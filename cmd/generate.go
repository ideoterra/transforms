package main

import (
	"log"
	"strings"
)

var primitiveTypes = []string{
	"int",
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
		result = append(result, typeNames{
			FileName:      primitiveType + "slice",
			PrimitiveType: primitiveType,
			SliceType:     strings.Title(primitiveType) + "Slice",
			SliceType2:    strings.Title(primitiveType) + "Slice2",
		})
		result = append(result, typeNames{
			FileName:      primitiveType + "slice2",
			PrimitiveType: strings.Title(primitiveType) + "Slice",
			SliceType:     strings.Title(primitiveType) + "Slice2",
			SliceType2:    "[]" + strings.Title(primitiveType) + "Slice2",
		})
	}
	return result
}

func main() {
	log.Println(generateTypeNames())
}
