package main

import (
	"io"
	"io/ioutil"
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
	IsLastGeneration bool // tells the generator to fix up the []SliceType2 values in types.go
	PackageName      string
	PrimitiveType    string
	SliceType        string
	SliceType2       string
}

const (
	basePath    = `/Users/joe/workspace/go/src/github.com/jecolasurdo/transforms/pkg/slices`
	genericPath = basePath + "/generic"
)

func main() {
	for _, p := range primitiveTypes {
		typeNames := generateTypeNames(p)

		log.Println("Purging generated type directories...")
		fileInfos, err := ioutil.ReadDir(basePath)
		if err != nil {
			log.Fatal(err)
		}
		for _, fileInfo := range fileInfos {
			if fileInfo.Name() == "generic" {
				continue
			}

			fileToRemove := filepath.Join(basePath, fileInfo.Name())
			log.Printf("Purging %v...", fileToRemove)
			err := os.RemoveAll(fileToRemove)
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, t := range typeNames {
			newDirName := filepath.Join(basePath, t.PackageName)
			log.Printf("Creating new path %v\n", newDirName)
			err := os.MkdirAll(newDirName, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Retrieving list of source files...")
			fileInfos, err := ioutil.ReadDir(genericPath)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Copying source files from generic to %v...\n", t.PackageName)
			for _, fileInfo := range fileInfos {
				oldName := filepath.Join(genericPath, fileInfo.Name())
				newName := filepath.Join(newDirName, fileInfo.Name())
				err := copyFile(oldName, newName)
				if err != nil {
					log.Fatal(err)
				}
			}

			basicReplacementFiles := []string{
				"doc.go",
				"functions.go",
				"methods.go",
				"types.go",
			}
			for _, basicFile := range basicReplacementFiles {
				fileName := filepath.Join(newDirName, basicFile)
				replaceTextInFile(fileName, "PrimitiveType", t.PrimitiveType)
				replaceTextInFile(fileName, "SliceType2", t.SliceType2)
				replaceTextInFile(fileName, "SliceType", t.SliceType)
				replaceTextInFile(fileName, "generic", t.PackageName)
				if basicFile == "types.go" && t.IsLastGeneration {
					removeLinesContainingValue(fileName, "[]"+t.SliceType)
				}
			}
		}
	}
}

func generateTypeNames(p primitiveType) []typeNames {
	result := []typeNames{}
	oneDimensionalSliceType := typeNames{
		IsLastGeneration: false,
		PackageName:      p.TypeName + "slice",
		PrimitiveType:    p.TypeName,
		SliceType:        strings.Title(p.TypeName) + "Slice",
		SliceType2:       p.TypeName + "slice2." + strings.Title(p.TypeName) + "Slice2",
	}

	twoDimensionalSliceType := typeNames{
		IsLastGeneration: true,
		PackageName:      p.TypeName + "slice2",
		PrimitiveType:    oneDimensionalSliceType.PackageName + "." + strings.Title(p.TypeName) + "Slice",
		SliceType:        strings.Title(p.TypeName) + "Slice2",
		SliceType2:       "[]" + strings.Title(p.TypeName) + "Slice2",
	}

	result = append(result, oneDimensionalSliceType, twoDimensionalSliceType)
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

func copyFile(src, dst string) error {
	// https://stackoverflow.com/a/21061062/3434541
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func replaceTextInFile(fileName, old, new string) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	oldContent := string(input)
	newContent := strings.Replace(oldContent, old, new, -1)
	err = ioutil.WriteFile(fileName, []byte(newContent), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func removeLinesContainingValue(fileName, value string) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, value) {
			lines[i] = ""
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
