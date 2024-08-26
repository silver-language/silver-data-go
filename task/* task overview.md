Task Overview
=============

This is a rough outline of the high level tasks for the silver-data parser.

1. [Lexer](#1-lexer)
2. [Parser](#2-parser)
3. [Testing](#3-testing)
4. [Use silver-data from client program](#3-use-silver-data-from-client-program)
5. [Versioning](#5-versioning)
6. [Publishing](#6-publishing)




1 Lexer
--------

Read in silver-data document and convert to token stream/token tree




2 Parser
---------

Take token tree and run through parsing phases to output usable data structure



3 Use silver-data from client program
-------------------------------------

* Read in silver-data from another program
* Construct/modify silver-data structures in memory
* Serialize silver-data structure back to document for storage


4 Testing
----------

Concurrent with 1,2,3 create test inputs and outputs



5 Versioning
------------

Concurrent with 1,2,3,4 establish a versioning scheme



6 Publishing
------------

Licensing, packaging, etc