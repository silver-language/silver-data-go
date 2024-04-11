package schema

var SilverLexer = Lexer{
	"document": NodeLexer{
		nodeType: "split",
		regex:    `\R`,
		subexp: []Submatch{
			{
				index:    0,
				nodeType: "line",
			},
		},
	},
	"line": NodeLexer{
		nodeType: "subex",
		regex:    `^(\s*)(.*)\n`,
		subexp: []Submatch{
			{
				index:    0,
				nodeType: "indent",
			},
			{
				index:    1,
				nodeType: "statement",
			},
		},
	},
	/*
		"statement": NodeLexer{
			nodeType: "statement",
			regex:    `([^:]*):(.*)`,
			subexp: []Subexp{
				{
					number:   1,
					nodeType: "name",
				},
				{
					number:   2,
					nodeType: "statement",
				},
			},
		},
	*/
}

type File struct {
	nodeType string
	regex    string
}

/*
I think I need individual node struct types for each node

I have a few options here
* have individual struct types for each node type
* reuse a single node type and extend (if possible)
* use a generic node type in a fairly generic way

I think go wants me to do this:
	https://go.dev/doc/effective_go#embedding



*/
