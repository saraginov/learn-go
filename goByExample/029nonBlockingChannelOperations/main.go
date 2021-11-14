package main

import (
	"fmt"
)

func test(msgChannel chan<- string) {
	msgChannel <- "anotherTest"
}

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive. If a value is available on messages then select will
	// take the <-messages case with that value. If not it will immediately take
	// the default case.
	fmt.Println("no message received, should print immediately after this line")

	// does this iife send "test subject" to messages? -> no it did not change
	// anything...
	// go func(msg chan<- string) {
	// 	fmt.Println("hello")
	// 	msg <- "test subject"
	// }(messages)
	// trying to see if I call test synchronously what will happen
	// if I invoke test(messages) without go, I don't create a go routine and it
	// throws a panic when I do go run (all goroutines are asleep - deadlock)
	// when I add go keyword to create a goroutine, issue above is resolved, but
	// I still get no message received
	// go test(messages)
	// fmt.Println(<-messages)
	// this is weird... Println line above prints anotherTest, meaning messages
	// has a value so why does it not match select case below?
	// modifying go routine to match main.go in timeouts, it's not working either
	// go func() {
	// 	time.Sleep(2 * time.Second
	// 	messages <- "did you receive this?"
	// }()
	// fmt.Println(<-messages)
	// can't get it to work in select below, even if fmt.Println shows a string
	// which I can print, select always goes to default

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	fmt.Println("This I think will print after the select and before no msg sent")
	// non-blocking send. msg cannot be sent to the messages channel, because the
	// channel has no buffer and there is no receiver. Therefore the default case
	// is selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	fmt.Println("This should print after no message sent and before no activity")

	// we can use multiple cases above the default clause to implement a multi-way
	// non-blocking select. Here we attempt non-blocking receives on both messages
	// and signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signals", sig)
	default:
		fmt.Println("no activity")
	}
	fmt.Println("This should print after no activity?")
}
