package main

import(
	"fmt"
)

/*
	The var statement declares a list of variables; as in function argument lists, the type is last.
	
	A var statement can be at package or function level.
*/

var c, python, java bool

/*
	A var declaration can include initializers, one per variable.

	If an initializer is present, the type can be omitted; the variable will take the type of the initializer
*/
var i, j int = 1, 2 

/*
	Short Variable Declarations:

		Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

		Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available!
*/

func main() {
	var k int

	var cTwo, pythonTwo, javaTwo = true, false, "NO!"

	l := 3 // short variable declarations

	/*
		default init value for int is 0
		default init value for string is ""
		default init value for bool is false
	*/
	fmt.Println(k, c, python, java)

	fmt.Println(i, j, cTwo, pythonTwo, javaTwo)

	fmt.Println(l)
}