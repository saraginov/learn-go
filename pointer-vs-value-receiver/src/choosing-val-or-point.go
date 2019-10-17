package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
	Choosing a value or pointer receiver

		There are two reasons to use pointer receiver:

			reason 1: 	So that the method can modify the value that its receiver points to.

			reason 2: 	To avoid copying the value on each method call.
						This can be more efficient if the receiver is a large struct
			
			Both Scale() and Abs() methods have receiver type *Vertex, 
			even though Abs() doesn't modify the receiver
			
			In general, all methods on a given type should have value or pointer receivers,
			NOT a mixture of both
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
} 

func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := &Vertex{3,4}
	fmt.Printf("Before scaling: %v, Abs: %v\n", v, v.Abs()) // output: Before scaling: &{X:3 Y:4}, Abs: 5

	v.Scale(6)
	fmt.Printf("After scaling: %v, Abs: %v\n", v, v.Abs()) // output: After scaling: &{X:18 Y:24}, Abs: 30
}