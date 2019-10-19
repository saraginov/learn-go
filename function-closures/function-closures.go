package main

import "fmt"

/*
	Function closures

		Go functions may be closures.
		A closure is a function value that references variables from outside its body.
		The function may access and assign to the referenced variables;
		in this case the function is "bound" to the variables.

		For example, the adder function returns a closure
		Each closure is bound to its own sum variable

	Function closures are "similar" to shared state counters
*/

func adder() func(int) int{
	/*
		adder has no parameters,
		adder returns a function
			the function adder returns an int as an argument and returns an int
	*/

	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func main(){
	pos, neg := adder(), adder()

	for i := 0; i<10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}