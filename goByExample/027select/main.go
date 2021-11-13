package main

import (
	"fmt"
	"time"
)

// total execution time is approx 2 seconds since both 1 and 2 second Sleeps
// execute concurrently!!!
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	// each channel receives a value after some amount of time, to simulate e.g.
	// blocking RPC operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("i", i)
		// use select to await both of these values simultaneously, printing each
		// one as it arrives
		// during first for loop iteration, until c1, or c2 is received we are
		// blocked, nothing happens, in our case c1 arrives after 1 second, then
		// second loop iteration is immediately invoked and we are blocked again
		// 1 second later the 2nd iife is invoked and c2 arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
