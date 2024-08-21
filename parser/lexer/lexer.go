package lexer

import (
	"fmt"
	"log"
	"regexp"
	"silver-data/parser-go/schema"
)

// type NodeArray []schema.AstNode

/* LexDocument
 */
func LexDocument(document string) schema.AstNode {

	abstractSyntaxTree := schema.AstNode{
		NodeType:  "document",
		Text:      document,
		LineStart: 0,
		LineEnd:   0,
		CharStart: 0,
		CharEnd:   0,
	}

	documentLexer := schema.SilverLexer["document"]
	Lex(&abstractSyntaxTree, &documentLexer, "document")
	return abstractSyntaxTree
}

/* Lex
 */
func Lex(astNode *schema.AstNode, lexer *schema.NodeLexer, nodeType string) { //schema.AstNode

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

	switch lexer.Splitter {
	case "lineSplitter":
		{
			astNode.Child = linesplitNode(astNode, lexer)
			lastChild := astNode.Child[len(astNode.Child)-1]
			astNode.LineEnd = lastChild.LineEnd
			astNode.CharEnd = lastChild.CharEnd
		}
	case "subExpression":
		{
			astNode.Child = submatchNode(astNode, lexer)
		}
	default:
		{
			astNode.NodeType = fmt.Sprintf("%v: node type not found", astNode.NodeType)
		}
	}
	/*
		} else {
			// panic(fmt.Sprintf("Node type not found: %v", nodeType))
			// or return an error node type??
			astNode.NodeType = fmt.Sprintf("Node type not found: %v", astNode.NodeType)
		}
	*/

	/* given some text and a node type, attempt to split it into nodes
	there will need to be two levels to this
	*/

	//log.Println(abstractSyntaxTree)

	/*
		run the node through its tests
		if there is a match perform it's split
			then loop through the splits
			and lex those nodes
		else if there is no match eject or record error
	*/

	//return result
}

func linesplitNode(astNode *schema.AstNode, lexer *schema.NodeLexer) []schema.AstNode {
	re := regexp.MustCompile(lexer.Regex)
	split := re.Split(astNode.Text, -1)

	result := []schema.AstNode{}

	for i := 1; i <= len(split); i++ {
		result = append(result,
			schema.AstNode{
				NodeType:  "line",
				Text:      split[i-1],
				LineStart: i,
				LineEnd:   i,
				CharStart: 0,
				CharEnd:   len(split[i-1]),
				//Child:     []schema.AstNode,
			},
		)
	}

	return result
}

func submatchNode(astNode *schema.AstNode, lexer *schema.NodeLexer) []schema.AstNode {
	result := new([]schema.AstNode)

	re := regexp.MustCompile(lexer.Regex)

	submatch := re.FindStringSubmatch(astNode.Text)
	log.Println(submatch)
	// there needs to be some distinction here

	return *result
}

/* I'm getting confused so here's an outline.


start with some text and a node type
send the text node type to the lexer which outputs some nodes
	for each node returned





*/
