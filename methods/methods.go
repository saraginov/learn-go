package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X,Y float64
}

/*
	Methods

		Go does not have classes. 
		However, you can define methods on type(s)

		A method is a function with a special receiver argument

		The receiver appears in its own argument list between 
		the func keyword and the method name

		In this example, the Abs() method has a receiver of type
		Vertex named v
*/
func (v Vertex) Abs() float64 { // method
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	Methods are Functions

		a method is just a function with a receiver argument.

		AbsTwo is written as a regular function with no change in functionality 
*/
func AbsTwo(v Vertex) float64 { // method written as function with identical functionality
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	We can declare a method on non-struct types, too

	In AbsThree we use a numeric type float64

	IMPORTANT :
		You can only declare a method with a receiver whose type is defined in the same
		package as the method.
		You cannot delcare a method with a receiver whose type is defined in another
		package (including built-in types such as int)
*/
type MyFloat float64

func (f MyFloat) AbsThree() float64 {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}


func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())

	fmt.Println(AbsTwo(v))

	f := MyFloat(-math.Sqrt2) // 2 does not need to be wrapped in () 
	fmt.Println(f.AbsThree()) // output: 1.414213.... 
	fmt.Println(f) // output: -1.414213.... 
	fmt.Printf("f: %v", f) // output: f: -1.414213.... 
}