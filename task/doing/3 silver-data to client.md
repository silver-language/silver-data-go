Silver data to client
=====================

```yaml
type: epic
status: not started
blocks: publishing
priority: low
```

Thoughts
--------
This has been going through my head lately.
There are a few layers to this and how a parser for another language might work will probably depend on what features that language has.

Ideally:
* create named types dynamically as specified
* translate silver-data specified types to something locally applicable

For  a language with not-so-good typing features it might be better for the client to receive something like this (or just omit the types altogether):

```
[name]:
	type: Typename as specified by silver-data
	value: single or multiple
```