package loggeria

// IHook defines an interface to a log hook.
type IHook interface {
	// Run runs the hook with the event.
	Run(ie IMapBuffer, level Level, message string)
}

// ISampler defines an interface to a log sampler.
type ISampler interface {
	// Sample returns true if the event should be part of the sample, false if
	// the event should be dropped.
	Sample(lvl Level) bool
}
