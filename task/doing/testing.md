Testing setup
=============


I'm rolling my own for now, will look at frameworks later.


- [ ] input a file, output the AST
- [ ] specify AST output format(s)
- [ ] specify wildcards for input file sets
- [ ] output pass/fails
- [ ] output as silver-data




Return after a while
--------------------

Had a long diversion here when i started looking into walkdir - went back to experiment with functions (anonymous etc), function typing and typing in general, some more OS/dir/file stuff.
Coming back to this I think I'll simplify a bit to try to get things moving.

* Instead of using walkdir I'll just use os.ReadDir
* will recombine the test folders (input/output) to make each folder a single topic test suite
* will just loop and filter to select the input files
* similarly filter for folders to create an 'all' list when test group unspecified
* that should pretty much allow for subfolders right away

What should the tests return?

	directory match yes/no
		inputfile	output file			match yes/no

Could keep a copy of these results in each directory also - not sure if needed?
