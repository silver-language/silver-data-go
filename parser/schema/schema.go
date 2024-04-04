package schema

/* need regexes and AST nodes


https://pkg.go.dev/math/big
*/

type AstNode struct {
	NodeType   string    // type of node
	Text       string    // the raw text of the node
	LineNumber uint      // the source line of this node
	CharStart  uint      // start character of this node
	CharEnd    uint      // end character of this node
	Child      []AstNode // zero or more child nodes
}

type Lexer struct {
	nodeType string
	regex    string
	test     []string
	subexp   []Subexp
}

/*
each lexer can either split the text immediately,

or depending on a test then pass onto another lexer


*/

type Subexp struct {
	number   int
	nodeType string
}

/*
can i eject more nodes from the lexer by using grouping patterns, or do i have to use a cursor?
*/
