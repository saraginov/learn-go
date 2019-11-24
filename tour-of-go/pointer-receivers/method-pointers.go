package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	Pointer Receivers

		Can declare methods with pointer receivers.

		This means the receiver type has a literal syntax *T for some type T.
		IMPORTANT:
			T cannot itself be a pointer such as *int

		Methods with pointer receivers can modify the value to which the receiver points
			Scale modifies Vertex v values
		Since methods often need to modify their receiver, 
		pointer receivers are more common than value receivers.

		With a value receiver,
		a method operates on a copy of the original value
		Identical behavior as for any other function argument

*/
func (v *Vertex) Scale(f float64){
	/*
		Scale method is defined on *Vertex,
		the receiver is v and the receiver pointer is *Vertex
	*/

	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleTwo(f float64){
	/*
		value receiver
		removing the * from the declaration removes access to memory address
		v is now a local scope variable that only ScaleTwo has access to
	*/

	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v) // output: {3,4}
	fmt.Println(v.Abs()) // output: 5

	v.Scale(10) // v is the receiver
	fmt.Println(v) // output: {30,40}
	fmt.Println(v.Abs()) // output: 50

	// value receiver does not set the value of v in memory address
	v.ScaleTwo(10)
	fmt.Println(v) // output: {30,40}
}

/*
	Regular Function vs. Pointer Receiver vs. Value Receiver
		|-> how to modify original memory address value using 
			different types of functions
		|-> 
			i)		differences between approaches
					a. value receiver:
						operates on copy of received value
					b. regular function:
						if we passed the original variable value, like value receiver
						regular function would also operate on a copy
					c. pointer receiver:
						operates on memory address value
			ii) 	computational cost/efficiency
					a. value receiver && pointer receiver:
						if objective of function is to modify original value,
						these operations are MORE costly as we copy the variable value twice
						once when we call method/function
						and second time when we assign new/modified value to original variable
					b. pointer receiver:
						performs data manipulation directly in memory address
			iii)	optimization:
						if data set for Type which method is defined is very large,
						copying it would be far too expensive than using a pointer receiver.
			iv)		concurency :
						Value receivers are concurrency safe, 
						while pointer receivers are not concurrency safe.

	ex.
		// define new Type T
		type T struct {val1, val2 int}
		
		// regular function
		func FuncName (arg1, arg2 int) T {
			var newSomeVar T
			newSomeVar.val1 = arg1
			newSomeVar.val2 = arg2
			return newSomeVar
		}			

		// value receiver:
		func (v T) MethodNameVal(arg1 int) T {
			var newSomeVar T
			newSomeVar.val1 = v.val1 * arg1
			newSomeVar.val2 = v.val2 * arg1
			return newSomeVar
		}

		// pointer receiver:
		func (v *T) MethodNamePoint(arg1, arg2 int) {
			v.val1 = arg1
			v.val2 = arg2
		}

		var outsideVar = T {123, 321}

		func main() {
			//---------------- someVar ----------------
			var someVar T
			someVar.val1 = 11
			someVar.val2 = 22
			fmt.Println(someVar) // output: {11 22}
			
			someVar = FuncName(33, 44)
			fmt.Println(someVar) // output: {33 44}		
			
			someVar = someVar.MethodNameVal(10)
			fmt.Println(someVar) // output: {330 440}
			
			someVar.MethodNamePoint(77,88)
			fmt.Println(someVar) // output: {77 88}

			//---------------- outsideVar ----------------
			fmt.Println(outsideVar) // output: {123 321}
			
			outsideVar = FuncName(333, 444)
			fmt.Println(outsideVar) // output: {333 444}		
			
			outsideVar = outsideVar.MethodNameVal(100)
			fmt.Println(outsideVar) // output: {33300 44400}
			
			outsideVar.MethodNamePoint(777,888)
			fmt.Println(outsideVar) // output: {777 888}
		}
*/