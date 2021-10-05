package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fibOutOfMain(n int) int {
	if n < 2 {
		return n
	}
	return fibOutOfMain(n-1) + fibOutOfMain(n-2)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}
	// since fib was defined in main, Go knows which function to call with fib()
	fmt.Println(fib(7))

	var fibFromOutOfMain func(n int) int
	fibFromOutOfMain = fibOutOfMain
	fmt.Println(fibFromOutOfMain(7))
}
