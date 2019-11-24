package main

import (
	"fmt"
	"math"
)

func main(){

	/*
		Type Interface:

			When declaring a variable without specifying an explicit type (either by using the := or var = expression syntax),
			the variable's type is inferred from the value on the right hand side

			When the right hand side of the declaration is typed, the new variable is of that same type.

			var i int
			j := i // j is an int

			But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex218
			depending on the precision of the constant
			i := 42 // int
			f := 3.142 // float64
			g := 0.867 + 0.5i // complex128
	*/

	a := 42
	b := 42.12345
	c := 42.123 + 69i

	fmt.Printf("a has a value of %v and is of type %T", a, a)
	fmt.Printf("b has a value of %v and is of type %T", b, b)
	fmt.Printf("c has a value of %v and is of type %T", c, c)

	/*
		The expression T(v) converts the value v to type T.

		Some numeric conversions:

		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)

		or more simply:
		i := 42
		f := float64(i)
		u := uint(f)

		Unlike in C, in Go assignment between items of different type requires an explicit conversion.
		
		Try removing the float64 or uint conversions in the example and see what happens

		./compile15.go:11:6: cannot use f (type float64) as type uint in assignment
	*/

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	
	fmt.Println(x,y,z)
}