package main

import (
	"fmt"
	"time"
)

/*
	Channels are a typed conduit through which you can send and receive values with the channel operator, <-

	ch <- v 		// Send v to channel ch
	v := <- ch		// Receive from ch, and assign value to v

	The data flows in the direction of the arrow

	Like maps and slices, channels must be created before use:

	ch := make(chan int)

	By default, sends and receives block until the other side is ready.
	This allows goroutines to synchronize without explicit locks or condition variables.

	The 'sum' example below sums the numbers in a slice, distributing the work between two goroutines.
	Once both goroutines have completed their computation, it calculates the final result.

	IMPORTANT:
		Could not replicate behaviour in browser when running /tour exec.

		However when running locally, no guarantee that goroutines run and relay int data to channel
		in exact same order every time !!! (channels function)

		Especially evident when adding the third go routine for z

		I added separate channels for each goroutine and the results were always consistent/ identical

		NOTE:
			Using independent channels, it seems goroutines are blocked from changing execution order!
*/

func sum(s []int, c chan int) {
	sum := 0

	fmt.Println(s)
	// fmt.Println(c)

	for _, v := range s {
		sum += v
		fmt.Println(v)
		fmt.Println(sum)
	}
	// fmt.Printf("%v %#v	%T	%d", c, c, c, c)

	c <- sum // send sum to 'c'
}

func channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	// fmt.Println(s[len(s)/2:])

	// c is typed conduit which can send receive values with <- operator
	// c := make(chan int)
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go sum(s[:len(s)/2], c1) // s[:3] = [7,2,8], c is chan int
	go sum(s[len(s)/2:], c2) // s[3:] = [-9, 4, 0], c is chan int
	// go sum(s[:len(s)], c)
	go sum(s[:len(s)], c3)

	// fmt.Printf("%v %#v	%T	%d", c, c, c, c)
	// x, y := <-c, <-c //receive from 'c',
	x, y, z := <-c1, <-c2, <-c3 //receive from 'c',
	// x, y, z := <-c, <-c, <-c //receive from 'c',
	// x is return from first sum goroutine
	// y is return from second sum goroutine
	// goroutines appear to be executed bottom->top order

	// How do x and y know which value in the c channel is the one intended for them?
	// Is there an automatic buffer when not set?

	// fmt.Println(x, y, x+y, z)
	fmt.Println(x, y, x+y, z)
}

/*
	Buffered Channels

		Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel

		someChannel := make(chan int, 1000) // => where 1000 is the buffer size in bytes?

		Sends the buffered channel block only when the buffer is full.

			Receives block when the buffer is empty.
*/

func bufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	/*
		Overfilling the buffer causes the goroutine to stop

		Running it locally, ran out of memory froze pc as an infinite loop would,
		 in browser when running /tour it executes first for loop
		 at some point while second loop executes it freezes there as well
	*/

	// overfilling buffer using a for loop

	// for i := 0; i < 1000; i++ {
	// 	ch <- i
	// 	fmt.Println(<-ch)
	// }

	// for j := 0; j < 1000; j++ {
	// 	ch <- j
	// 	fmt.Println(<-ch)
	// }
}

/*
	Range and Close

		A sender can close a channel to indicate that no more values will be sent.

		Receivers can test whether a channel has been closed by assigning a second
		parameter to the receive expression: after

		v, ok := <-ch

			ok is false if there are no more values to receive and the channel is closed.

		The loop for i := range c receives values from the channel repeatedly until
		it is closed.

		NOTE:
			Only the sender should close a channel, never the receiver. Sending on a closed
			channel will cause a panic.

			Channels are not like files! You don't usually need to close them.
			Closing channels is only necessary when the receiver must be told there are no
			more values coming, such as to terminate a range loop
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func rangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) // cap for channel c is 10 bytes, thus n in func fibonacci is 10

	for i := range c {
		fmt.Println(i)
	}

}

/*
	Select

	The select statement lets a goroutine wait on multiple communication operations.

	A select blocks until one of its cases can run, then executes that case.

	It chooses at random if multiple are ready!
*/

func secondFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("running an iteration of for loop in Fibonacci....")
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func runningSelect() {
	/*
		The order of operations in this example is really peculiar,
		in that both the IIFE below and secondFibonacci run once the first time
		and then the time it takes to execute an iteration of the for loop
		in each function yields an A, B, B, A, A, B, B, A, A, ...
		(identical behaviour in browser)
	*/
	c := make(chan int)
	quit := make(chan int)

	go func() { // IIFE
		for i := 0; i < 10; i++ {
			fmt.Println("running iteration of for loop in go func IIFE")
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	secondFibonacci(c, quit)
}

/*
	The Select default

	The default case in a select is run if no other case is ready.

	Use a default case to try a send or receive without blocking:

		select {
		case i := <-c:
			// use i
		default:
			// receiving from c would block
		}
*/

func defaultSelect() {
	/*
		tick every second, every half a second print ...
		BOOM! after 5 seconds
	*/
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	channels()
	fmt.Println("CHANNELS --------------------------------------- END")

	bufferedChannels()
	fmt.Println("BUFFERED CHANNELS --------------------------------------- END")

	rangeAndClose()
	fmt.Println("RANGE AND CLOSE  --------------------------------------- END")

	runningSelect()
	fmt.Println("SELECT  --------------------------------------- END")

	defaultSelect()
	fmt.Println(" DEFAULT SELECT  --------------------------------------- END")
}
