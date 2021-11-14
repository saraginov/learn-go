package main

import (
	"fmt"
	"time"
)

func main() {
	// timer will wait 2 seconds
	fmt.Println("about to declare and init timer1")
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("after timer1 and before channel opeartion")

	// we wait 2 seconds for <-timer1.C before continuing, does C refer to chan?
	// <-timer1.C blocks on timer's channel C until it sends a value indicating
	// that timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// to just wait, time.Sleep can be used, however timer is useful because it
	// can be cancelled before it fires
	timer2 := time.NewTimer(time.Second)
	fmt.Println("about to run go func()")
	go func() {
		fmt.Println("first line inside go func()")
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	fmt.Println("after go func() and before stop2")
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("after sleep")
	// without time.Sleep, goroutine never executes, i.e. program closed before
	// routine can do its thing
}
