package main

import (
	"fmt"
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
	fmt.Println(s[len(s)/2:])

	c := make(chan int)
	go sum(s[:len(s)/2], c) // s[:3] = [7,2,8], c is chan int
	go sum(s[len(s)/2:], c) // s[3:] = [-9, 4, 0], c is chan int

	fmt.Printf("%v %#v	%T	%d", c, c, c, c)
	x, y := <-c, <-c //receive from 'c',
	// x is return from first sum goroutine
	// y is return from second sum goroutine
	// goroutines appear to be executed bottom->top order

	fmt.Println(x, y, x+y)
}

/*
	Buffered Channels

*/
func main() {
	channels()

}
