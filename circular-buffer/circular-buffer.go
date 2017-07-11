package store

// Implements a circular buffer. This is not a ring buffer,
// as it doesn't provide separate read/write pointers.  Instead,
// it behaves more like a MongoDB capped collection.  It accepts
// new values until it has run out of it's fixed size, then it
// wraps around and begins overwriting existing values.  Reads
// are always in order however, no matter what the underlying buffer
// looks like at the time.
// See Also: https://docs.mongodb.com/manual/core/capped-collections/
type CircularBuffer struct {
	buf     []interface{}
	size    int
	tail    int
	wrapped bool
}

func New(size int) *CircularBuffer {
	return &CircularBuffer{make([]interface{}, size), size, 0, false}
}

func (cb *CircularBuffer) Append(v interface{}) {
	cb.buf[cb.tail] = v
	cb.tail = cb.tail + 1
	if cb.tail >= cb.size {
		cb.tail = 0
		cb.wrapped = true
	}
}

func (cb *CircularBuffer) Slice() []interface{} {
	if cb.wrapped {
		return append(cb.buf[cb.tail:], cb.buf[:cb.tail]...)
	} else {
		return cb.buf[:cb.tail]
	}
}
