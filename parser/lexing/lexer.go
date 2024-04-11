package lexing

import (
	"experiment/parser/schema"
	"fmt"
	"regexp"
)

type NodeArray []schema.AstNode

func Lex(ast schema.AstNode, source string, nodeType string, lexer schema.Lexer) schema.AstNode {

	// fmt.Println(lexer[nodeType])

	// first run any tests
	// leaving this out for the moment - will come back to

	// second, try to split this node

	nodeLexer, ok := lexer[nodeType]
	if !ok {
		panic(fmt.Sprintf("Node type not found: %v", nodeType))
		// or return an error node type??
	}

	submatches := lexer[nodeType]

	result := new(schema.AstNode)

	/* given some text and a node type, attempt to split it into nodes
	there will need to be two levels to this
	*/

	//fmt.Println(abstractSyntaxTree)

	/*
		run the node through its tests
		if there is a match perform it's split
			then loop through the splits
			and lex those nodes
		else if there is no match eject or record error
	*/

	/*
		astNode = &schema.AstNode{
			NodeType:   "",               // type of node
			Text:       "",               // the raw text of the node
			LineNumber: 1,                // the source line of this node
			CharStart:  1,                // start character of this node
			CharEnd:    1,                // end character of this node
			Child:      [], // zero or more child nodes
		}
	*/

	return *result
}

func splitter(text string, regex string) []string {
	re := regexp.MustCompile(regex)
	result := re.Split(text, -1)
	return result
}

/* I'm getting confused so here's an outline.


start with some text and a node type
send the text node type to the lexer which outputs some nodes
	for each node returned





*/
