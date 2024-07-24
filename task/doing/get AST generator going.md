Get AST generator going
=======================

Not working yet.
Get working.


Splitters
---------

Got the document-to-line splitter working last night.
I think anything that requires recombining lines can be done at later stages so this is no problem right now.
Not sure if I actually need to preserve the lineends in the AST - can do that if needed.

Indentation
-----------
Next I'll split lines into indentation and statements

All leading tabs are indentation, anything else becomes part of the statement.

Might need a way to indicate that a node terminates, ie that it has no children (a leaf).
The switch in the node lexer will fall through if it doesn't find a splitter, but might be nice to make it more explicit.

Setting a node type
-------------------

One possible option:
https://stackoverflow.com/questions/9993178/is-there-a-nice-way-to-simulate-a-maybe-or-option-type-in-go

This technique might be good for some cases, but it's probably not needed here yet.
See for example:
https://stackoverflow.com/questions/59964619/difference-using-pointer-in-struct-fields

There are a few caveats noted, so will just use another flag to indicate a terminal node for now.
Can change later if wanted.




AST node types - generic or specific?
-------------------------------------
At the moment I just have a single generic ASTNode that I was intending to return from each lexing phase, so I just end up with a tree of those.
Do I actually need different ast node structs for each node type, or is that getting ahead of things?
Is specific node types something i want to eject from parsing instead?



