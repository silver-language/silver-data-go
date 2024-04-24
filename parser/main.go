package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	lexing "silver-data/parser-go/lexing"
	schema "silver-data/parser-go/schema"
)

func main() {

	testFolder := "./test/file/"
	testFilename := "factorial.agl"
	//outputFilename := fmt.Sprintf("%s.ast", inputFilename)
	testFilePath := fmt.Sprintf("%v%v", testFolder, testFilename)
	outputFilePath := fmt.Sprintf("%s.ast.json", testFilePath)

	// Open file
	file, err := os.Open(testFilePath)
	if err != nil {
		fmt.Println("Error opening file: %v", err)
	}
	defer file.Close()

	// Read file
	content, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Println("Error reading file: %v", err)
	}

	// Parse file content
	contentString := string(content)

	abstractSyntaxTree := schema.AstNode{
		NodeType:  "document",
		Text:      contentString,
		LineStart: 1,
		LineEnd:   1,
		CharStart: 1,
		CharEnd:   1,
	}

	tmp := schema.SilverLexer["document"]

	lexing.Lex(&abstractSyntaxTree, &tmp, "document")

	// Write output

	output, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
	if err != nil {
		log.Fatalf(err.Error())
	}

	//output := fmt.Sprintf("%+v", abstractSyntaxTree)
	err = os.WriteFile(outputFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
