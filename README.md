loggeria

# Logger Interface Abstraction

It defines a logger, a map buffer with 18 common types and 18 homogeneous lists and specials and shortcuts,
a heterogeneous list buffer with 18 common types, ability to nest one into another. Also hook and sampler
interfaces. Careful and thoughtful work was donw, to make this as rational as possible.

```go

	// recommended form
	logger.WRN(func(e loggeria.E) {
		e.Str("k", v).Str("k2", v2).Msg("hello")
	}).Send()

	// one-liner
	logger.WRN(func(e loggeria.E) { e.Str("k", v).Str("k2", v2).Msg("hello") }).Send()

```

The benefits of this approach, compared to zerolog:

1) Simpler and not duplicated data model
2) Buffers can be used on its own, outside of logger
3) Simpler and more explicit to understand
4) Add to p.3, only one explicit finalizer Send(), no ambiguity

5) No need for high-perf Disable/IfDisabled functionality; if the logger is disabled, or the level must not be logged, or the filter decided it shouldn't be logged, then the data will not be generated at all.


### Adding a subsys path:

```go

	logger.Sub("a/s/d").WRN(func(e loggeria.E) {
		e.Str("k", v).Str("k2", v2).Msg("hello")
	}).Send()

```

Note, that it's added __on the logger__, not on the buffer; this lets the buffer stay isolated, and do filtering
before starting to construct the data.


