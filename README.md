# RingBuffer - A Generic Circular Buffer in Go

## Introduction
`RingBuffer` is a generic, type-safe circular buffer implementation in Go (Golang). It allows efficient, FIFO (first-in-first-out) operations in a fixed-size buffer. Ideal for scenarios where you need a fixed-size, rolling storage of elements.

## Features
- Generic: Can store any data type.
- Automatic Overwriting: When full, new elements replace the oldest.
- Concurrent Access: Suitable for producer-consumer scenarios. (Note: Does not include synchronization primitives, you'll need to handle concurrency as per your use case.)
- Simple API: Easy to use methods for common operations.

## Installation
To install the package, use the following `go get` command:

```sh
go get github.com/metajar/ringbuff
```

## Usage
Here's a simple example of how to use the `RingBuffer`:

```go
package main

import (
    "fmt"
    "github.com/metajar/ringbuff" // Replace with the correct import path
)

func main() {
    ringBuffer := ringbuff.NewRingBuffer[int](4)

    // Adding elements
    ringBuffer.Push(1)
    ringBuffer.Push(2)
    ringBuffer.Push(3)
    ringBuffer.Push(4)

    // Popping and printing elements
    for i := 0; i < 4; i++ {
        value, err := ringBuffer.Pop()
        if err != nil {
            fmt.Println("Error:", err)
            break
        }
        fmt.Println(value)
    }
}
```

## Running Tests
To run the tests for the `RingBuffer`, navigate to the package directory and use the `go test` command:

```sh
go test
```


