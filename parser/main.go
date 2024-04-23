package main

import (
	"encoding/json"
	lexing "experiment/parser/lexing"
	schema "experiment/parser/schema"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFilename := "./test/factorial.agl"
	//outputFilename := fmt.Sprintf("%s.ast", inputFilename)
	outputFilename := fmt.Sprintf("%s.ast.json", inputFilename)

	// Open and read file
	filename := "./test/factorial.agl"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file")
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
	err = os.WriteFile(outputFilename, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
