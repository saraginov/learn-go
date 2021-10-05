package main

import "fmt"

func intSeq() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var nextInt func() int = intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
