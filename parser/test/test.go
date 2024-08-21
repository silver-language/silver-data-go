package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"silver-data/parser-go/file"
	"silver-data/parser-go/lexer"
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
		generate TT,AST
		stringify
		record comparison against existing output
	return test results

regenerate output:
	calculate a set set based upon folders/filters/formats
	read each input file
		generate TT,AST
		stringify
		write result
	return output rewrites
*/

var inputFileSuffixes = [2]string{".agd", ".agl"}

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
		log.Printf("Error reading directory: %v \n %v \n", directoryPath, err)
	}

	for index, directoryItem := range directoryItems {
		log.Printf("%v: %v \n", index, directoryItem)

		//log.Printf("directory: %v \n", itemPath)

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
		log.Printf("Error reading directory: %v \n %v \n", directoryPath, err)
	}

	for index, directoryItem := range directoryItems {
		log.Printf("%v: %v \n", index, directoryItem)

		//log.Printf("directory: %v \n", itemPath)

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
	log.Printf("Test file: %v \n", fileName)

	/*
		//log.Printf("Test %v: %v \n", index, thisFile)
		filePath := fmt.Sprintf("%s%s", directory, fileName)

		input := ReadFile(filePath)

		abstractSyntaxTree := lexer.LexDocument(input)
		outputFilepath := fmt.Sprintf("%s.tt.json", outputFolder, filePath) //!!
		//log.Printf("outputFilepath: %v \n", outputFilepath)

		json, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
		if err != nil {
			log.Fatalf(err.Error())
		}
		WriteFile(outputFilepath, string(json))
	*/

}

func GenerateOutputFile(directory string, fileName string) {
	log.Printf("Generate output: %v \n", fileName)

	//log.Printf("Test %v: %v \n", index, thisFile)
	filePath := fmt.Sprintf("%s%s", directory, fileName)

	input := file.Read(filePath)

	abstractSyntaxTree := lexer.LexDocument(input)
	outputFilepath := fmt.Sprintf("%s.tt.json", filePath) //!!
	log.Printf("outputFilepath: %v \n", outputFilepath)

	json, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
	if err != nil {
		log.Fatalf(err.Error())
	}

	file.Write(outputFilepath, string(json))

}
