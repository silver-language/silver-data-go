package main

import (
	lexing "experiment/parser/lexing"
	schema "experiment/parser/schema"
	"fmt"
	"os"
)

func main() {

	// read file

	filename := "./test/factorial.agl"

	// show file stats
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(file.Stat())

	// read the file
	content, err := os.ReadFile(filename)

	// fmt.Println(reflect.TypeOf(content))	// content is a slice

	if err != nil {
		fmt.Println("Err")
	}
	//fmt.Println(string(content))

	contentString := string(content)

	abstractSyntaxTree := lexing.Lex(contentString, "document", schema.SilverLexer)

	fmt.Println(abstractSyntaxTree)

	defer file.Close()
}
