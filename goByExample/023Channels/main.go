package main

import "fmt"

func main() {
	messages := make(chan string)
	// below we will receive "ping" to the message channel declared above, from
	// a new goroutine
	go func() {
		messages <- "ping"
	}()

	// the dataToStore := <-channelName receives the value from the channel
	msg := <-messages
	fmt.Println(msg)
}
