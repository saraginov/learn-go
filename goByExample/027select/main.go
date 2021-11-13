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
		// use select to await both of these values simultaneously, printing each
		// one as it arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
