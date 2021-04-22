package printer

import "fmt"

func printGreeting() {
	fmt.Println("un-exported")
}

func Greeting() {
	fmt.Println("WELCOME")
}
