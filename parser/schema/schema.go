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
}

/*
can i eject more nodes from the lexer by using grouping patterns, or do i have to use a cursor?


*/
