package lexer

import (
	"fmt"
	"log"
	"regexp"
)

/* LexDocument
 */
func LexDocument(document string) TokenTree {

	tokenTree := TokenTree{
		NodeType:  "document",
		Text:      document,
		LineStart: 0,
		LineEnd:   0,
		CharStart: 0,
		CharEnd:   0,
	}

	documentLexer := SilverLexer["document"]
	LexTree(&tokenTree, &documentLexer, "document")
	return tokenTree
}

/* LexTree
 */
func LexTree(tokenTree *TokenTree, lexer *NodeLexer, nodeType string) {

	// first run any tests 	// leaving this out for the moment - will come back to

	// second, split this node

	switch lexer.Splitter {
	case "documentSplitter":
		{
			tokenTree.Child = linesplitNode(tokenTree, lexer)
			lastChild := tokenTree.Child[len(tokenTree.Child)-1]
			tokenTree.LineEnd = lastChild.LineEnd
			tokenTree.CharEnd = lastChild.CharEnd
		}
	case "subExpression":
		{
			tokenTree.Child = submatchNode(tokenTree, lexer)
		}
	default:
		{
			tokenTree.NodeType = fmt.Sprintf("%v: node type not found", tokenTree.NodeType)
		}
	}
	/*
		} else {
			// panic(fmt.Sprintf("Node type not found: %v", nodeType))
			// or return an error node type??
			tokenTree.NodeType = fmt.Sprintf("Node type not found: %v", tokenTree.NodeType)
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

func linesplitNode(tokenTree *TokenTree, lexer *NodeLexer) []TokenTree {
	re := regexp.MustCompile(lexer.Regex)
	split := re.Split(tokenTree.Text, -1)

	result := []TokenTree{}

	for i := 1; i <= len(split); i++ {
		result = append(result,
			TokenTree{
				NodeType:  "line",
				Text:      split[i-1],
				LineStart: i,
				LineEnd:   i,
				CharStart: 0,
				CharEnd:   len(split[i-1]),
				//Child:     []TokenTree,
			},
		)
	}

	return result
}

func submatchNode(tokenTree *TokenTree, lexer *NodeLexer) []TokenTree {
	result := new([]TokenTree)

	re := regexp.MustCompile(lexer.Regex)

	submatch := re.FindStringSubmatch(tokenTree.Text)
	log.Println(submatch)
	// there needs to be some distinction here

	return *result
}

/* I'm getting confused so here's an outline.


start with some text and a node type
send the text node type to the lexer which outputs some nodes
	for each node returned





*/
