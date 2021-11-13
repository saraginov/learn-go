package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("make c1")
	// c1 := make(chan string, 1)
	c1 := make(chan string)
	// suppose we are executing an external call that returns its results on a
	// channel c1 after 2 seconds, note that the channel is buffered, so the
	// send in the goroutine is nonblocking. This is a common pattern to prevent
	// goroutine leaks in case the channel is never read
	fmt.Println("before first iife")
	go func() {
		fmt.Println("going to sleep for 2 seconds first iife")
		time.Sleep(2 * time.Second)
		fmt.Println("woke up from sleep in first iife")
		// "result 1" never prints...
		c1 <- "result 1"
	}()
	// select implements a timeout. res := <-c1 awaits the result and <-time
	// after awaits a value to be sent after the timeout of 1s.
	// since select proceeds with the first receive that's ready, we'll take the
	// timeout case if the operation takes more than the allowed 1s.
	// am I understanding this correctly in that
	fmt.Println("first select below")

	// the first case which executes is the only case which runs, time.After
	// takes place 1 second sooner than time.Sleep in IIFE above, I wonder what
	// will happen if both are the same? hypothesis is that first match will be
	// invoked... nope when both are at 2 * time.Second, timeout is still
	// triggered
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}
	fmt.Println("before c2")

	c2 := make(chan string, 1)
	fmt.Println("made c2")
	go func() {
		fmt.Println("second sleep now")
		time.Sleep(2 * time.Second)
		fmt.Println("woke up again")
		c2 <- "result 2"
	}()
	fmt.Println("second select")
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		// "timeout 2" didn't print either
		fmt.Println("timeout 2")
	}
	fmt.Println("app done")
}
