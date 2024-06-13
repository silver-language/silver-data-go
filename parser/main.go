package main

import "silver-data/parser-go/test"

func main() {

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

	files := test.GetFiles("./test/simple/")
	test.GenerateOutput(files)

}
