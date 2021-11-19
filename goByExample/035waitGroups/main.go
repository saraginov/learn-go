package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// WaitGroup is used to wait for all goroutines launched within main to finish
	// If a WaitGroup is explicitly passed into functions, it should be done by
	// pointer
	var wg sync.WaitGroup

	// launch goroutines, increment WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		// wg.Add increments WaitGroup by 1, each wg.Done() in closure below
		// decrements, wg.Wait() then blocks until decremented down to zero
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure.
		// https://golang.org/doc/faq#closures_and_goroutines
		// the way in which official docs
		i := i

		// wrap worker in closure that tells WaitGroup that worker is done
		// this way the worker itself does not have to be aware of the concurrency
		// primitives involved in its execution
		// defer statement = defers the execution until surrounding function returns
		// i.e. when closure returns, wg.Done() is invoked
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	// block until WaitGroup counter goes back to 0; all workers are done
	wg.Wait()
}
