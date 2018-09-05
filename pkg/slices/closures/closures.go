package closures

// ConditionFn determines whether or not a value meets some condition.
type ConditionFn func(interface{}) bool

// EqualityFn determins whethre or not two values are equal.
type EqualityFn func(a, b interface{}) bool
