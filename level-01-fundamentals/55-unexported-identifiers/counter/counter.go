// Package counter demonstrates ENCAPSULATION using unexported
// identifiers: the internal count is hidden from outside packages,
// reachable only through this package's deliberately exposed,
// EXPORTED API.
package counter

// value is UNEXPORTED — lowercase first letter. No package other than
// counter itself can read or write this field directly, from anywhere,
// under any circumstances. This is Go's entire encapsulation mechanism.
type Counter struct {
	value int
	step  int // also unexported — an internal implementation detail
}

// New is the EXPORTED constructor — the only way another package can
// obtain a working Counter, since Counter's own fields are unexported
// and can't be set directly with a struct literal from outside.
func New(step int) *Counter {
	return &Counter{value: 0, step: step}
}

// Increment is EXPORTED — it's the deliberately exposed way to modify
// the unexported `value` field from outside this package.
func (c *Counter) Increment() {
	c.value += c.step
}

// Value is EXPORTED — a deliberately exposed READ-ONLY view of the
// unexported `value` field.
func (c *Counter) Value() int {
	return c.value
}
