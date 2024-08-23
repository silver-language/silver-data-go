package lexer

import (
	"fmt"
	"log"
	"regexp"
	"silver-data/parser-go/util"
)

var lexer = SilverLexer

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

	documentLexer := lexer["document"]
	LexTree(&tokenTree, &documentLexer, "document")
	return tokenTree
}

/* LexTree
 */
func LexTree(tokenTree *TokenTree, nodeLexer *NodeLexer, nodeType string) {

	log.Printf("LexTree: %v, %v \n", nodeType, nodeLexer.Splitter)

	// first run any tests 	// leaving this out for the moment - will come back to

	// second, split this node

	switch nodeLexer.Splitter {
	case "documentSplitter":
		{
			linesplitNode(tokenTree, nodeLexer)
		}
	case "subExpression":
		{
			submatchNode(tokenTree, nodeLexer)
		}
	default:
		{
			tokenTree.NodeType = fmt.Sprintf("%v: node type not found", tokenTree.NodeType)
		}
	}

	/*
		run the node through its tests
		if there is a match perform it's split
			then loop through the splits
			and lex those nodes
		else if there is no match eject or record error
	*/

	for index, item := range tokenTree.Child {
		log.Printf("Lex child: %v, %v \n", item.NodeType, &item)
		itemLexer := lexer[item.NodeType]
		LexTree(&tokenTree.Child[index], &itemLexer, item.NodeType)
	}

}

/* linesplitNode
 */
func linesplitNode(tokenTree *TokenTree, lexer *NodeLexer) {
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

	tokenTree.Child = result
	lastChild := tokenTree.Child[len(tokenTree.Child)-1]
	tokenTree.LineEnd = lastChild.LineEnd
	tokenTree.CharEnd = lastChild.CharEnd
}

/* submatchNode
 */
func submatchNode(tokenTree *TokenTree, lexer *NodeLexer) { //[]TokenTree
	//result := new([]TokenTree)

	re := regexp.MustCompile(lexer.Regex)

	submatch := util.FindSubmatches(tokenTree.Text, *re)

	log.Println(submatch)
	// there needs to be some distinction here

	for i, value := range submatch[1:] {
		tokenTree.Child = append(tokenTree.Child,
			TokenTree{
				NodeType:  lexer.Subexp[i].nodeType,
				Text:      value.String,
				LineStart: tokenTree.LineStart,
				LineEnd:   tokenTree.LineEnd,
				CharStart: value.Start,
				CharEnd:   value.End,
				//Child:     []TokenTree,
			},
		)

	}

	//return *result
}

/* I'm getting confused so here's an outline.


start with some text and a node type
send the text node type to the lexer which outputs some nodes
	for each node returned





*/
