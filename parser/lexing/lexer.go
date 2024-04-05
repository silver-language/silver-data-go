package lexing

import "experiment/parser/schema"

func Lex(source string, nodeType string) schema.AstNode {
	//

	result := new(schema.AstNode)

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
