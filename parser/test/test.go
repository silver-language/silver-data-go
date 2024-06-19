package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"silver-data/parser-go/lexing"
	"strings"
)

/*

defaults for:
	folder
	input filter
	ouptut formats

operations:
	test run
	regenerate output

test:
	calculate a set set based upon folders/filters/formats
	read each file
		generate AST
		stringify
		record comparison against existing output
	return test results

regenerate output:
	calculate a set set based upon folders/filters/formats
	read each input file
		generate AST
		stringify
		write result
	return output rewrites
*/

var inputFileSuffixes = [2]string{".agd", ".agl"}

func ReadFile(filepath string) string {
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v \n", err)
	}
	defer file.Close()

	fmt.Printf("File opened: %v \n", filepath)

	// Read file
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v \n", err)
	}
	fmt.Printf("File read: %v \n", filepath)

	return string(content)
}

func WriteFile(filepath string, content string) {
	//fmt.Printf("filepath: %v \n", filepath)
	err := os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v \n", err)
	}
	fmt.Printf("File written: %v \n", filepath)
}

/* Rewrite

readDir(directory string)
	read directory
	loop directoryItems
		match directory
			readDir(item)
		match file
			processFile(item)

processFile(file string)
	match inputtype
		generateOutput(file)
	match other
		ignore

*/

func TestDirectory(directoryPath string) {
	directoryItems, err := os.ReadDir(directoryPath)
	var itemPath string

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error reading directory: %v \n %v \n", directoryPath, err)
	}

	for index, directoryItem := range directoryItems {
		fmt.Printf("%v: %v \n", index, directoryItem)

		//fmt.Printf("directory: %v \n", itemPath)

		if directoryItem.IsDir() {
			itemPath = fmt.Sprintf("%s%s/", directoryPath, directoryItem.Name())
			TestDirectory(itemPath)
		} else {
			//itemPath = fmt.Sprintf("%s%s", directoryPath, directoryItem.Name())
			if isInputFile(directoryItem.Name()) {
				TestFile(directoryPath, directoryItem.Name())
			}
		}
	}
	//return testResults
} //TestDirectory

// GenerateDirectoryOutput
func GenerateDirectoryOutput(directoryPath string) {
	directoryItems, err := os.ReadDir(directoryPath)
	var itemPath string

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error reading directory: %v \n %v \n", directoryPath, err)
	}

	for index, directoryItem := range directoryItems {
		fmt.Printf("%v: %v \n", index, directoryItem)

		//fmt.Printf("directory: %v \n", itemPath)

		if directoryItem.IsDir() {
			itemPath = fmt.Sprintf("%s%s/", directoryPath, directoryItem.Name())
			GenerateDirectoryOutput(itemPath)
		} else {
			//itemPath = fmt.Sprintf("%s%s", directoryPath, directoryItem.Name())
			if isInputFile(directoryItem.Name()) {
				GenerateOutputFile(directoryPath, directoryItem.Name())
			}
		}
	}
	//return testResults
} //GenerateDirectoryOutput

func isInputFile(filePath string) bool {
	result := false
	for _, suffix := range inputFileSuffixes {
		result = result || strings.HasSuffix(filePath, suffix)
	}
	return result
}

func TestFile(directory string, fileName string) {
	fmt.Printf("Test file: %v \n", fileName)

	/*
		//fmt.Printf("Test %v: %v \n", index, thisFile)
		filePath := fmt.Sprintf("%s%s", directory, fileName)

		input := ReadFile(filePath)

		abstractSyntaxTree := lexing.LexDocument(input)
		outputFilepath := fmt.Sprintf("%s.ast.json", outputFolder, filePath) //!!
		//fmt.Printf("outputFilepath: %v \n", outputFilepath)

		json, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
		if err != nil {
			log.Fatalf(err.Error())
		}
		WriteFile(outputFilepath, string(json))
	*/

}

func GenerateOutputFile(directory string, fileName string) {
	fmt.Printf("Generate output: %v \n", fileName)

	//fmt.Printf("Test %v: %v \n", index, thisFile)
	filePath := fmt.Sprintf("%s%s", directory, fileName)

	input := ReadFile(filePath)

	abstractSyntaxTree := lexing.LexDocument(input)
	outputFilepath := fmt.Sprintf("%s.ast.json", filePath) //!!
	fmt.Printf("outputFilepath: %v \n", outputFilepath)

	json, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
	if err != nil {
		log.Fatalf(err.Error())
	}

	WriteFile(outputFilepath, string(json))

}
