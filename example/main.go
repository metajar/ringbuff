package main

import (
	"fmt"
	ringbuff "github.com/metajar/ringbuff" // Replace with your package path
	"io"
)

func main() {
	ringBuffer := ringbuff.NewRingBuffer[int](3)

	// Adding elements
	ringBuffer.Push(1)
	ringBuffer.Push(2)
	ringBuffer.Push(3)
	ringBuffer.Push(4)
	// Popping and printing elements
	for i := 0; i < 4; i++ {
		value, err := ringBuffer.Pop()
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(value)
	}

	// Output should be: 1 2 3
}
