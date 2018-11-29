package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ideoterra/transforms/pkg/slices/generic"
	"github.com/ideoterra/transforms/pkg/slices/shared"
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
	basePath    = `/Users/joe/workspace/go/src/github.com/ideoterra/transforms/pkg/slices`
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
			if fileInfo.Name() == "generic" || fileInfo.Name() == "shared" {
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
			log.Println("Retrieving list of source files...")
			fileInfos, err := ioutil.ReadDir(genericPath)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Copying source files from generic to %v...\n", t.PackageName)
			newBaseName := filepath.Join(basePath, t.PackageName)
			for _, fileInfo := range fileInfos {
				if strings.Contains(fileInfo.Name(), "_test") || fileInfo.Name() == "doc.go" {
					continue
				}
				oldName := filepath.Join(genericPath, fileInfo.Name())
				newName := newBaseName + fileInfo.Name()
				err := copyFile(oldName, newName)
				if err != nil {
					log.Fatal(err)
				}
			}

			basicReplacementFiles := []string{
				"functions.go",
				"methods.go",
				"types.go",
			}
			for _, basicFile := range basicReplacementFiles {
				fileName := newBaseName + basicFile
				replaceTextInFile(fileName, "PrimitiveType", t.PrimitiveType)
				replaceTextInFile(fileName, "SliceType2", t.SliceType2)
				replaceTextInFile(fileName, "SliceType", t.SliceType)
				replaceTextInFile(fileName, "generic", "slicexform")
				if strings.Contains(fileName, "types.go") && t.IsLastGeneration {
					removeLinesContainingValue(fileName, "[]"+t.SliceType)
				}
				if strings.Contains(fileName, "functions.go") {
					functionNames := getFunctionNamesForFile(fileName)
					functionNames.Sort(func(a, b generic.PrimitiveType) bool {
						return len(a.(string)) < len(b.(string))
					}).Distinct(func(a, b generic.PrimitiveType) bool {
						name1 := a.(string)
						name2 := b.(string)
						if strings.Contains(name1, "Fold") && strings.Contains(name2, "Fold") {
							log.Println("break")
						}
						return strings.Contains(name1, name2)
					}).ForEach(func(a generic.PrimitiveType) shared.Continue {
						functionName := a.(string)
						replaceTextInFile(fileName, functionName, t.SliceType+functionName)
						return shared.ContinueYes
					})
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
		SliceType2:       strings.Title(p.TypeName) + "Slice2",
	}

	twoDimensionalSliceType := typeNames{
		IsLastGeneration: true,
		PackageName:      p.TypeName + "slice2",
		PrimitiveType:    strings.Title(p.TypeName) + "Slice",
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
func getFunctionNamesForFile(fileName string) generic.SliceType {
	functionNames := generic.SliceType{}
	const funcRegex = `func ([A-Z]\w*)\(`
	re := regexp.MustCompile(funcRegex)
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(input)
	for _, submatch := range re.FindAllStringSubmatch(contents, -1) {
		functionNames.Append(submatch[1])
	}
	return functionNames
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
