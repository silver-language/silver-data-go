package lexer

import (
	"fmt"
	"log"
	"regexp"
	"silver-language/silver-data/util"
	"strings"
)

var lexer = SilverLexer

/* LexDocument
 */
func LexDocument(document string) Token {

	// pre-clean newlines
	document = strings.ReplaceAll(document, "\r\n", "\n")
	document = strings.ReplaceAll(document, "\r", "\n")

	token := Token{
		Type:      "document",
		Text:      document,
		LineStart: 0,
		LineEnd:   0,
		CharStart: 0,
		CharEnd:   0,
		Split:     true,
		Child:     []Token{},
	}

	documentLexer := lexer["document"]
	LexTree(&token, &documentLexer, "document")
	return token
}

/* LexTree
 */
func LexTree(token *Token, nodeLexer *NodeLexer, nodeType string) {

	log.Printf("LexTree: %v, %v \n", nodeType, nodeLexer.Splitter)

	// first run any tests 	// leaving this out for the moment - will come back to

	// second, split this node

	switch nodeLexer.Splitter {
	case "documentSplitter":
		{
			linesplitNode(token, nodeLexer)
		}
	case "subExpression":
		{
			submatchNode(token, nodeLexer)
		}
	case "none":
		{
			// createTerminalNode()
		}
	default:
		{
			token.Type = fmt.Sprintf("%v: node type not found", token.Type)
		}
	}

	/*
		run the node through its tests
		if there is a match perform it's split
			then loop through the splits
			and lex those nodes
		else if there is no match eject or record error
	*/

	for index, item := range token.Child {
		//log.Printf("Lex child: %v, %v \n", item.Type, &item)
		itemLexer := lexer[item.Type]
		LexTree(&token.Child[index], &itemLexer, item.Type)
	}

}

/* linesplitNode
 */
func linesplitNode(token *Token, lexer *NodeLexer) {

	re := regexp.MustCompile(lexer.Regex)
	split := re.Split(token.Text, -1)

	result := []Token{}

	for i := 1; i <= len(split); i++ {
		result = append(result,
			Token{
				Type:      "line",
				Text:      split[i-1],
				LineStart: i,
				LineEnd:   i,
				CharStart: 0,
				CharEnd:   len(split[i-1]),
				Split:     true,
				Child:     []Token{},
			},
		)
	}

	token.Child = result
	lastChild := token.Child[len(token.Child)-1]
	token.LineEnd = lastChild.LineEnd
	token.CharEnd = lastChild.CharEnd
}

/* submatchNode
 */
func submatchNode(token *Token, lexer *NodeLexer) { //[]Token
	//result := new([]Token)

	re := regexp.MustCompile(lexer.Regex)

	submatch := util.FindSubmatches(token.Text, *re)

	log.Printf("%#v", submatch)
	// there needs to be some distinction here

	for i, value := range submatch[1:] {
		token.Child = append(token.Child,
			Token{
				Type:      lexer.Subexp[i].nodeType,
				Text:      value.String,
				LineStart: token.LineStart,
				LineEnd:   token.LineEnd,
				CharStart: value.Start,
				CharEnd:   value.End,
				//Child:     []Token,
			},
		)

	}

	//return *result
}

func createTerminalNode(token *Token, lexer *NodeLexer) {

}

/* I'm getting confused so here's an outline.


start with some text and a node type
send the text node type to the lexer which outputs some nodes
	for each node returned





*/
