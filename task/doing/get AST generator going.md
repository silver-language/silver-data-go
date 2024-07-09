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

One possible option:
https://stackoverflow.com/questions/9993178/is-there-a-nice-way-to-simulate-a-maybe-or-option-type-in-go

