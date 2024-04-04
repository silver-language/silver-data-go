package lexing

import "experiment/parser/schema"

func lex(source string, nodeType string) schema.AstNode {
	//

	astNode = &schema.AstNode{
		NodeType:   "",               // type of node
		Text:       "",               // the raw text of the node
		LineNumber: 1,                // the source line of this node
		CharStart:  1,                // start character of this node
		CharEnd:    1,                // end character of this node
		Child:      [], // zero or more child nodes
	}

	return astNode
}
