Task Overview
=============

This is a rough outline of the high level tasks for silver-data.

1. [Lexer](#1-lexer)
2. [Parser](#2-parser)
3. [Use silver-data from client program](#3-use-silver-data-from-client-program)
4. [Testing](#4-testing)
5. [Versioning](#5-versioning)
6. [Publishing](#6-publishing)




1 Lexer
--------

Read in silver-data document and convert to token stream/token tree




2 Parser
---------

Take token tree and run through parsing phases to output usable data structure

* combine blocks together
* report syntax errors
* map AST nodes back to source document

3 Use silver-data from client program
-------------------------------------

* From a different module import the silver-data module
* Use silver-data module to read and use local config
* Construct/modify silver-data structures in memory
* Serialize silver-data structure back to document/file for storage


4 Testing
----------

Concurrent with 1,2,3 create test inputs and outputs

* document to token tree/stream
* token to AST
* AST to structure
* silver structure to document


5 Versioning
------------

Concurrent with 1,2,3,4 establish a versioning scheme



6 Publishing
------------

Licensing, packaging, etc