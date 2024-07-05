Add logger for debugging
========================

At the moment I've got debug prints all over the place, need to rationalise into something more organised.

For starters concatenating to a log would do, but something more structured for reporting would be ideal.

Turns out the inbuilt logger is a super-easy drop in replacement for fmt.Println so will go with that for a while.
