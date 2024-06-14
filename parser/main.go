package main

import (
	"fmt"
	"os"
	"silver-data/parser-go/test"
)

func main() {

	fmt.Println("silver-data-parser-go")

	//testFolder := "./test/file/"
	//testFilename := "factorial.agl"
	//outputFilename := fmt.Sprintf("%s.ast", inputFilename)
	//testFilepath := fmt.Sprintf("%v%v", testFolder, testFilename)
	//outputFilepath := fmt.Sprintf("%s.ast.json", testFilepath)

	/*
		// Parse file content
		content := test.ReadFile(testFilepath)

		abstractSyntaxTree := schema.AstNode{
			NodeType:  "document",
			Text:      content,
			LineStart: 1,
			LineEnd:   1,
			CharStart: 1,
			CharEnd:   1,
		}

		tmp := schema.SilverLexer["document"]

		lexing.Lex(&abstractSyntaxTree, &tmp, "document")

		// Write output

		json, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
		if err != nil {
			log.Fatalf(err.Error())
		}

		test.WriteFile(outputFilepath, string(json))
	*/

	// files := test.GetFiles("./test/simple/")
	//test.GenerateOutput(files)

	if len(os.Args) > 1 {
		task := os.Args[1]
		switch task {
		case "test":
			test.TestDirectory("./test/")
		case "generate":
			test.GenerateDirectoryOutput("./test/")
		default:
			fmt.Println("No matching task")
		}
	} else {
		fmt.Println("Try: test")
	}
}
