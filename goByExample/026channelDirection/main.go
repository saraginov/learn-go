package main

import "fmt"

// ping only accepts a channel for sending values. It would be a compile-time
// error to try to receive on this channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong accepts one channel for receives(pings) and a second for sends(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	// can I just do
	// pongs <- pings no I cannot because pings is an incoming channeland
	// pongs is a channel which takes a string, therefore pongs <- pings is
	// passing a channel to pongs instead of a string and it throws an error
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
