package ringbuff

import "io"

// RingBuffer is a circular buffer.
type RingBuffer[T any] struct {
	buffer         []T
	head, tail, sz int
	capacity       int
}

// NewRingBuffer creates a new RingBuffer with the given capacity.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{buffer: make([]T, capacity), capacity: capacity}
}

// Push adds the specified element to the ring buffer.
// If the buffer is full, the oldest element will be overwritten.
// The head pointer is incremented to point to the next available slot.
// If the buffer is already full, the tail pointer is also incremented to maintain the buffer's size.
// The size of the buffer is incremented if it is not already at the maximum capacity.
func (r *RingBuffer[T]) Push(elem T) {
	r.buffer[r.head] = elem
	r.head = (r.head + 1) % r.capacity
	if r.sz == r.capacity {
		r.tail = (r.tail + 1) % r.capacity
	} else {
		r.sz++
	}
}

// Pop retrieves and removes the oldest element from the buffer.
// It returns the retrieved element and nil if successful,
// or a zero value of type T and io.EOF if the buffer is empty.
func (r *RingBuffer[T]) Pop() (T, error) {
	if r.sz == 0 {
		return *new(T), io.EOF
	}
	elem := r.buffer[r.tail]
	r.tail = (r.tail + 1) % r.capacity
	r.sz--
	return elem, nil
}

// Cap returns the capacity of the ring buffer.
func (r *RingBuffer[T]) Cap() int { return r.capacity }

// Len returns the number of elements in the RingBuffer.
func (r *RingBuffer[T]) Len() int { return r.sz }
