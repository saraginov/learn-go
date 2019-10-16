package main

import "fmt"

func main() {
	/*
		Arrays

			The type [n]T os am array of n values of type T
			The expression
				var a [10]int 
			declares a variable as as array of ten integers
			
			IMPORTANT!
			An array's length is part of its type, thus arrays cannot be resized.
			This seems limiting, but don't worry; 
			Go provides a convenient way of working with arrays 
	*/
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}