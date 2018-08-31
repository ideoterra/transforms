package shared

// Continue instructs iterators about whether or not to keep iterating.
type Continue bool

const (
	// ContinueYes signals to an iterator that it should continue iterating.
	ContinueYes Continue = true

	// ContinueNo signals to an iterator that it should stop iterating.
	ContinueNo Continue = false
)
