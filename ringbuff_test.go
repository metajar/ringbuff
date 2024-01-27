package ringbuff

import (
	"io"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	capacity := 4
	ringBuffer := NewRingBuffer[int](capacity)

	// Test pushing and popping
	for i := 1; i <= capacity; i++ {
		ringBuffer.Push(i)
		if ringBuffer.Len() != i {
			t.Errorf("Expected buffer length %d, got %d", i, ringBuffer.Len())
		}
	}

	// Test overwriting and popping
	extraValues := []int{5, 6}
	for _, v := range extraValues {
		ringBuffer.Push(v)
	}

	expectedValues := []int{3, 4, 5, 6} // Values 1 and 2 should be overwritten
	for _, expected := range expectedValues {
		val, err := ringBuffer.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if val != expected {
			t.Errorf("Expected value %d, got %d", expected, val)
		}
	}

	// Test popping from empty buffer
	_, err := ringBuffer.Pop()
	if err != io.EOF {
		t.Errorf("Expected io.EOF, got %v", err)
	}
}

func TestRingBufferEmpty(t *testing.T) {
	ringBuffer := NewRingBuffer[int](1)

	// Test empty buffer
	if ringBuffer.Len() != 0 {
		t.Errorf("Expected buffer length 0, got %d", ringBuffer.Len())
	}

	_, err := ringBuffer.Pop()
	if err != io.EOF {
		t.Errorf("Expected io.EOF for empty buffer, got %v", err)
	}
}
