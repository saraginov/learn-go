package main

import(
	"fmt"
	"math"
)

/*
	Function Values

		Functions are values too.
		They can be passed around just like other values/variables

		Function values may be used as function arguments and return values.

*/
func compute(fn func(float64, float64) float64)	float64{
	/*
		compute takes in a function as argument
			the parameter function has an alias 'fn'
			the func 'fn' has 2 parameters of its own, which are 2 float64(s) 
			and returns a float64
		compute returns a float64 value
	*/

	return fn(3,4)
}

func main(){

	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y) // c^2 = x^2 + y^2 == Pythagorean theorem
	}
	fmt.Println(hypot(5,12)) // expected output: 13

	fmt.Println(compute(hypot))	// expected output: 5
	fmt.Println(compute(math.Pow)) //expected output: 81 Math.pow == 3*3*3*3
}