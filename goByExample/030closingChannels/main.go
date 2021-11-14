package main

import (
	"fmt"
)

// use jobs channel to comm work to be done from main() goroutine to a worker
// routine: when there are no more jobs for the worker, close the jobs channel
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// iife below is worker goroutine, it repeatedly receives from jobs with
	// j, more := <- jobs, in this special 2-value form of receive, the more value
	// will be false if jobs has een closed and all values in the channel have
	// already been received.
	go func() {
		for { // i.e. while or infinite loop without return in else statement
			fmt.Println("What happens inside the goroutine?")
			// why does for loop waits for j, more := <- jobs?
			// is the channel blocking?
			// "By default sends and receives block until both
			// the sender and the receiver are ready. This property allows us to wait
			// at the end of our program for the "ping" message without having to use
			// any other synchronization"
			j, more := <-jobs
			fmt.Println("j: ", j)
			fmt.Println("jobs: ", more)
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		// time.Sleep(3 * time.Second)
		// fmt.Println("what happens between a job being sent and here")
		// nothing makes no difference
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// I don't understand what the purpose of done is in this program...
	// I thought that the purpose is to force main() to wait for iife to finish
	// executing before program closes... but I don't think this is it,
	// with or without <-done, the program seems to behave the same... it wasn't
	// behaving like this for me yesterday but it is now so I was maybe doing
	// something wrong in those tests...
	// we await the worker using the synchronization approach we saw earlier
	<-done
	fmt.Println("does it change")
}
