package main

import (
	"fmt"
	"os"
)

func main() {
	var l = len(os.Args) - 1

	fmt.Println("There are", l, "people!")
	for i := 0; i < l; i++ {
		n := os.Args[i+1]
		fmt.Println("Hello great", n, "!")
	}

	if l > 0 {
		fmt.Println("Nice to meet you all!")
	}
}
