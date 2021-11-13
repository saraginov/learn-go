package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// send a value to notify that we're done
	done <- true
}

func main() {
	// done is used to notify antoher go routine that this function's work is done
	done := make(chan bool, 1)
	// start a worker goroutine, giving it the channel to notify on
	go worker(done)
	// both Println are
	fmt.Println("when will I print")
	fmt.Println("hey hey hey")

	// block until we receive a notification from the worker on the channel
	<-done
	fmt.Println("worker finished")
}
