package main

import(
	"fmt"
)

/*
	A function (func) can take in zero or more arguments

	in func add below, there are 2 arguments both of which are of type int

	NOTICE: the type comes after the variable name.
*/
func add(x int, y int) int {
	return x + y
}

/*
	When two or more consecutive named function parameters share a type, type can be omitted from all but the last.
*/
func addSixInts(a, b, c, d, e, f int) int {
	return a + b + c + d + e + f
}

/*
	A function can return any number of results

	The swap function returns two strings
*/
func swap(x, y string) (string, string){
	return y, x;
}

/*
	Go's return values may be named. If so, they are treated as variables defined at the top of the function

	These names should be used to document the meaning of return values

	A return statement without arguments returns the named return values. This is known as a "naked" return.

	Naked return statements should be used only in short functions, otherwise they can harm readability in longer functions.
*/
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x

	return
}

func main (){
	fmt.Println(add(42,13))
	fmt.Println(addSixInts(1,2,3,4,5,6,))

	a, b := swap("this was passed in first", "this was passed in second")
	fmt.Println(a,b)

	fmt.Println(split(17))
}