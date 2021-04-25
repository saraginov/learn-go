package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v \n", os.Args)
	fmt.Printf("%#v \n", os.Args)
	fmt.Println(os.Args)
	var someNum int
	someNum = len(os.Args)
	fmt.Printf("%d \n", someNum)

	for i := 0; i < len(os.Args); i++ {
		var value string
		value = os.Args[i]
		fmt.Printf("%v \n", value)
	}
}
