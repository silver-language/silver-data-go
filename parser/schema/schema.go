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

type Lexer map[string]NodeLexer

type NodeLexer struct {
	nodeType string
	regex    string
	test     []string
	subexp   []Subexp
}

type SimpleSplitter struct {
	nodeType string
	regexp   string
	output   string
}

/*
each lexer can either split the text immediately,
or depending on a test then pass onto another lexer

as I don't know yet how to do an 'or' type in go (enum, sum, union, choice, etc)
I'll just create separate node fields for branches and leaves

*/

type Subexp struct {
	number   int
	nodeType string
}

/*
can i eject more nodes from the lexer by using grouping patterns, or do i have to use a cursor?
*/
