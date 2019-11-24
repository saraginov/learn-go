package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	sync.Mutex

		we've seen how channels are great for communication between goroutines.

		But what if we don't need communication? ie. make sure only one goroutine can access
		a variable at a time to avoid conflicts

			This concept is called mutual exclusion, and the conventional name for the data structure
			that provides it is mutex.

			Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

				Lock

				Unlock

			We can define a block of code to be executed in mutual exclusion by surrounding it with a
			call to Lock and Unlock as shown in the Inc method below.

			We can also use defer to ensure the mutex will be unlocked as in the Value method.

*/

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex // Mutex has a {1 0} or {0 0}  value, TODO:  figure out what this means exactly
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter from the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)

	fmt.Println(c.Value("somekey"))
}
