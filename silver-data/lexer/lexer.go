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

	token := Token{
		Type: "document",
		Text: document,
		//Line:      0,
		CharStart: 0,
		CharEnd:   0,
		Split:     true,
		//Terminal:  false,
		Child: []Token{},
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
			linesplitDocument(token, nodeLexer)
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

	for index, child := range token.Child {
		log.Printf("Lex child: %#v \n", &child)
		if child.Split {
			itemLexer := lexer[child.Type]
			LexTree(&token.Child[index], &itemLexer, child.Type)
		}

	}

}

/* linesplitDocument
 */
func linesplitDocument(documentToken *Token, lexer *NodeLexer) {

	// pre-clean newlines
	documentToken.Text = strings.ReplaceAll(documentToken.Text, "\r\n", "\n")
	documentToken.Text = strings.ReplaceAll(documentToken.Text, "\r", "\n")

	//result := []Token{}
	var newToken Token

	//re := regexp.MustCompile(lexer.Regex)
	//split := re.Split(token.Text, -1)

	split := strings.Split(documentToken.Text, "\n")

	for i, line := range split {

		log.Println(len(strings.TrimSpace(line)))

		if len(strings.TrimSpace(line)) == 0 {
			log.Println("if line len == 0")
			newToken = Token{
				Type:      "line-empty",
				Text:      line,
				Line:      i + 1,
				CharStart: 0,
				CharEnd:   len(line),
				Split:     false,
				//Terminal:  true,
				//Child:     []Token{},
			}
		} else {
			log.Println("else")
			newToken = Token{
				Type:      "line",
				Text:      line,
				Line:      i + 1,
				CharStart: 0,
				CharEnd:   len(line),
				Split:     true,
				//Terminal:  false,
				Child: []Token{},
			}
		}
		documentToken.Child = append(documentToken.Child, newToken)
	}
	lastChild := documentToken.Child[len(documentToken.Child)-1]
	//documentToken.Line = lastChild.Line
	documentToken.CharEnd = lastChild.CharEnd
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
				Line:      token.Line,
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
