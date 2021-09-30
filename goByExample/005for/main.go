package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("i = " + strconv.Itoa(i))
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("j = " + strconv.Itoa(j))
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			// breaks on first loop
			break
		}
		fmt.Println("n= " + strconv.Itoa(n))
	}
}
