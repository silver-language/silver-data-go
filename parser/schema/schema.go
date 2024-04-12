package schema

/* need regexes and AST nodes
https://pkg.go.dev/math/big
*/

type Lexer map[string]NodeLexer

type NodeLexer struct {
	NodeType string
	Regex    string
	Test     []string
	Subexp   []Submatch
}

type Submatch struct {
	index    int
	nodeType string
	match    string
}

/*
type SimpleSplitter struct {
	nodeType string
	regexp   string
	output   string
}

type SubmatchSplitter struct {
	regex string
}
*/

/*
each lexer can either split the text immediately,
or depending on a test then pass onto another lexer

as I don't know yet how to do an 'or' type in go (enum, sum, union, choice, etc)
I'll just create separate node fields for branches and leaves that can zero or more elements.

I also need to figure out the best way to compose structs in Go.

*/

/*
can i eject more nodes from the lexer by using grouping patterns, or do i have to use a cursor?
*/

type AstNode struct {
	NodeType  string    // type of node
	Text      string    // raw text of the node
	LineStart int       // start line of this node
	LineEnd   int       // end line of this node
	CharStart int       // start character of this node
	CharEnd   int       // end character of this node
	Child     []AstNode // zero or more child nodes
}
