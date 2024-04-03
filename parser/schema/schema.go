package schema

/* need regexes and AST nodes
 */

type AstNode struct {
	nodeType string
	value    string
}

type Lexer struct {
	nodeType string
	regex    string
	test     []string
	subexp   []Subexp
}

/* each lexer can either split the text immediately,
or depending on a test then pass onto another lexer
*/

type Subexp struct {
	number   int
	nodeType string
}

/*
can i eject more nodes from the lexer by using grouping patterns, or do i have to use a cursor?


*/
