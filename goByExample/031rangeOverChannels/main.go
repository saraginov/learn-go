package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// queue <- "three" // throws a fatal error: all goroutines are asleep -
	// deadlock! HOWEVE if I clear the queue as shown below, then I can pass
	// third string to the channel
	fmt.Println("clearing queue", <-queue)
	queue <- "three"
	close(queue)

	// iterate over 2 values in queue channel
	// Q: how is memory impacted by this? it seems that memory(ram) is limited
	// and so far all knowledgebases for channels and routines pretend as if it
	// memory is infinite and there is very little talk of what happens in as size
	// of channel is increased, also what happens if
	// range below iterates over each element as it's received from queue.
	// because channel queue is closed, iteration terminates after receiving the
	// 2 elements
	// It is possible to close a non-empty channel and still have the remaining
	// values be received
	for elem := range queue {
		fmt.Println(elem)
	}
}
