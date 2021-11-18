package main

import (
	"fmt"
	"time"
)

// workers will receive work on the jobs channel and send the corresponding
// results on results
func worker(id string, jobs <-chan int, results chan<- int) {
	// fmt.Println("worker", id, "started up")
	for j := range jobs {
		fmt.Println("worker", id, "starter job", j)
		// sleep a second per job to simulate an expensive task
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// in order to use pool of workers, send them work and collect their results
	// step 1: create jobs
	// runtime.GOMAXPROCS(1) => 1 or 4 or w.e doesn't change anything in results
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	codes := []string{"A", "B", "C", "D", "E", "F"}

	// step 2: start up workers, blocked until jobs are sent
	for w := 0; w < 3; w++ {
		go worker(codes[w], jobs, results)
	}

	// step 3: send jobs, followed by close channel to end indicate end of work
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect results of work, collecting all results ensures that worker
	// goroutines have finished. Alternative way to wait for multiple goroutines
	// is to use WaitGroup(s)
	for a := 1; a <= numJobs; a++ {
		<-results // ensure worker go
	}
}

// if worker is blocked until jobs are done, how is worker C starting job
// before worker A&B start up? or it happens so fast that they have to wait to
// write to stdout? or
// // worker C started up
// worker C starter job 1
// // worker A started up
// worker A starter job 2
// // worker B started up
// worker B starter job 3
// worker C finished job 1
// worker C starter job 4
// worker B finished job 3
// worker B starter job 5
// worker A finished job 2
// worker C finished job 4
// worker B finished job 5
