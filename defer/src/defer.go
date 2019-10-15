package main

import "fmt"

/*
	A defer statement defers the execution of a function until the surrounding function(s) return

	The deferred call's arguments are evaluated immediately, but the function call is not executed until surrounding function returns

	https://blog.golang.org/defer-panic-and-recover

	refer to func testingDefer() {...}
	
	all fmt.Println() get added in array A = []
	all defer fmt.Println() added in array B = [] IN REVERSE ORDER!

	every function "element" of array A gets called;
	once array A has been exhausted, and by exhausted I mean each function element has returned some value;
	every function "element" of array B gets called;

	Therefore the executional order of testingDefer is
		test 11111
		test 22222
		test 33333
		test 44444
		test 55555
		test 666666
		test 7777777
		world 33333
		world 222222
		world 11111

	TODO: 	Figure out if each "deferred" function waits for previous "deferred" function to execute before it gets called
			OR if it doesn't you can create race conditions
*/

func testingDefer(){
	fmt.Println("test 11111") // called 1st

	fmt.Println("test 22222") // 2nd

	defer fmt.Println("world 11111") //10th
	
	fmt.Println("test 33333") // 3rd 
	
	fmt.Println("test 44444") // 4th

	defer fmt.Println("world 222222") // 9th
	
	fmt.Println("test 55555") // 5th
	
	fmt.Println("test 666666") // 6th

	defer fmt.Println("world 33333") //8th
	
	fmt.Println("test 7777777") // 7th
}

func main() {
	testingDefer()

	/*
		Stacking defers

			Deferred function calls are pushed onto a stack,
			When a function returns, its deferred calls are executed in last-if-first-out order

			https://blog.golang.org/defer-panic-and-recover

	*/

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	/*
		Sample output of counting, loop, then done block above
			counting
			done
			9
			8
			7
			6
			5
			4
			3
			2
			1
			0
	*/
}