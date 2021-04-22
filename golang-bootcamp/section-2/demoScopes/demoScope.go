package main

import "fmt"

const ok = true

func nope() {
	// main function does not have access to variables in nope function scope
	const notOk = false
	var bye = "Bye"

	fmt.Println(notOk, bye, '\n')
}

func main() {
	var hello = "Hello"

	fmt.Println(hello, ok, '\n')

	nope()
}
