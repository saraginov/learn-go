package main

import (
	"fmt"
	"time"
)

// ticker should tick 3 times before we stop it
func main() {
	// use select builtin on channel to await values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// tickers can be stopped like timers. Once a ticker is stopped, it won't
	// receive any more values on its channel
	time.Sleep(1600 * time.Millisecond)
	// time.Sleep(2500 * time.Millisecond)
	ticker.Stop()
	done <- true
}
