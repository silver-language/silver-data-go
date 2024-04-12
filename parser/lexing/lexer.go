package lexing

import (
	"experiment/parser/schema"
	"fmt"
	"regexp"
)

type NodeArray []schema.AstNode

func Lex(astNode *schema.AstNode, lexer schema.Lexer) { //schema.AstNode

	// first run any tests
	// leaving this out for the moment - will come back to

	// second, try to split this node

	/*
		result := schema.AstNode{
			NodeType:  nodeType, // type of node
			Text:      source,   // the raw text of the node
			LineStart: 1,
			LineEnd:   1, // the source line of this node
			CharStart: 1, // start character of this node
			CharEnd:   1, // end character of this node
			//Child:     [], // zero or more child nodes
		}
	*/

	nodeLexer, ok := lexer[astNode.NodeType]
	if ok {
		switch nodeLexer.NodeType {
		case "linesplit":
			{
				astNode.Child = linesplitNode(*astNode, nodeLexer)
			}
		case "substring":
			{
				astNode.Child = submatchNode(*astNode, nodeLexer)
			}
		}
	} else {
		// panic(fmt.Sprintf("Node type not found: %v", nodeType))
		// or return an error node type??
		astNode.NodeType = fmt.Sprintf("Node type not found: %v", astNode.NodeType)
	}

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

	//return result
}

func linesplitNode(astNode schema.AstNode, lexer schema.NodeLexer) []schema.AstNode {
	re := regexp.MustCompile(lexer.Regex)
	split := re.Split(astNode.Text, -1)

	result := []schema.AstNode{}

	for i := 1; i <= len(split); i++ {
		result := append(result,
			schema.AstNode{
				NodeType:  "line",
				Text:      split[i-1],
				LineStart: i,
				LineEnd:   i,
				CharStart: 1,
				CharEnd:   len(split[i-1]),
				//Child:     []schema.AstNode,
			},
		)
	}

	return result
}

func submatchNode(astNode schema.AstNode, lexer schema.NodeLexer) []schema.AstNode {
	result := new([]schema.AstNode)
	return *result
}

/* I'm getting confused so here's an outline.


start with some text and a node type
send the text node type to the lexer which outputs some nodes
	for each node returned





*/
