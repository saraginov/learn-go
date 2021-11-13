package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// When we run this program, we see the output of the blocking call first,
// then the output of the 2 goroutines
func main() {
	// invoking function f(s) synchronously.
	f("direct")
	// To invoke this function in a goroutine, use `go f(s)`. This new goroutine
	// will execute concurrently with the calling one.
	go f("goroutine")
	// can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// the 2 go routines are running asynchronously in separte goroutines
	// put the program to sleep in order to wait for them to finish
	// for a more robust approach, use WaitGroup
	time.Sleep(time.Second)
	fmt.Println("done")
}
