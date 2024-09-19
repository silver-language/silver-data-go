/*
Schema
*/
package lexer

/*
 */

type Token struct {
	Type      string  // type of token
	Text      string  // raw text of the token
	Line      int     // end line of this token
	CharStart int     // start character of this token
	CharEnd   int     // end character of this token
	Terminal  bool    // whether this is a terminal node
	Child     []Token // zero or more child tree nodes
}

/*
type Token struct {
	NodeType  string // type of token
	Text      string // raw text of the token
	LineStart int    // start line of this token
	LineEnd   int    // end line of this token
	CharStart int    // start character of this token
	CharEnd   int    // end character of this token
}
*/

type Lexer map[string]NodeLexer

type NodeLexer struct {
	Terminal bool
	Splitter string
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
