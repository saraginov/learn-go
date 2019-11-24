package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
	Here we see the Abs and Scale methods rewritten as functions
*/
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64){
	/*
		func ScaleFunc (v *Vertex, f float64)
		functions with a pointer argument must take a pointer
	*/
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleTwo(v Vertex, f float64){
	/*
		Removing the * from Vertex in the scale method as shown in ScaleTwo
			changes the behavior:
				cannot use &v argument type *Vertex for parameter of type Vertex 
			to compile must change arg passed to Scale to from (&v,f) to (v,f) 
				only performs data manipulation on local copy of var v in ScaleTwo
	*/
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) ScalePointerIndirection(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func ValArgFunc(v Vertex) float64 {
	/*
		Functions that take a value argument must take a value of that specific type
		in this case type == Vertex
	*/
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) ValReceiverFunc() float64{
	/*
		methods with value receivers take either a value or a pointer 
		as the receiver when they are called: 
	*/
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v)) // output: 50

	/*
		ScaleTwo(&v, 1/10) // cannot use &v argument type *Vertex for parameter of type Vertex  
	*/

	ScaleTwo(v, 10) // only performs data manipulation on local copy of var v in ScaleTwo
	fmt.Println(Abs(v)) // output: 50 

	/*
		Methods and pointer indirection

			Comparing ../pointer-receivers/src/method-pointers.go to identically named methods above
			we can conclude:
				
				!!!  functions with pointer argument must take a pointer!
				
				ex.
					var v Vertex, f int
					func (v *Vertex) ScalePointer(f float){...}
					ScalePointer(v, f) 	// Compile error!
					ScalePointer(&v, f) // OK

				IMPORTANT :

					!!! 
						while methods with pointer receivers take either value or a pointer
						as a receiver when they are called:

					ex.
						var v Vertex
						v.ScalePointer(5) 	// OK
						p := &v
						p.ScalePointer(10) 	// OK	
						
					For the statement v.ScalePointer(5),
					even though v is a value and not a pointer
	*************	the method with the pointer receiver is called automattically!!!	*************
					That is, as convenience,
					Go interprets the statement v.ScalePointer(5) as (&v)ScalePointer(5)
					since the ScalePointer method has a pointer receiver
	*/
	vIndirection := Vertex{3,4}
	vIndirection.ScalePointerIndirection(2)
	Scale(&v, 10) // functions with a pointer argument must take a pointer
	// ScaleFunc(v, 10) 	// Compile error!

	p := &Vertex{4,3}
	p.ScalePointerIndirection(3)
	Scale(p, 8) // since p declared as pointer when initialized, no error!

	fmt.Println(v, p) // output: {60 80} &{96 72}

	/*
		Methods and pointer indirection P2

			Equivalent thing happens in the reverse direction.

			Functions that take a value argument, must take a value of that specific type:

			ex.
				func AbsFunc(v Vertex) float64 {...}
				var v Vertex
				fmt.Println(AbsFunc(v)) 	// OK
				fmt.Println(AbsFunc(&v)) 	// Compile error!

			while methods with value receivers, take either a value or a pointer as the receiver
			when they are called:

			ex. 
				func (v Vertex) Abs() float64 {...}
				var v Vertex
				fmt.Println(v.Abs()) // OK
				p := &v
				fmt.Println(p.Abs()) // OK

			In the case of p.Abs()
			the call is interpreted as (*p).Abs()

	*/

	vValReceiver := Vertex{3,4}
	fmt.Println(vValReceiver.ValReceiverFunc())
	/*
		functions that take a value argument must take a value of that specific type

		fmt.Println(ValArgFunc(*vValReceiver)) // invalid indirect of vValReceiver (type Vertex)
		fmt.Println(ValArgFunc(&vValReceiver)) // cannot use &vValReceiver (type *Vertex) as
		 type Vertex in argument to ValArgFunc
	*/
	fmt.Println(ValArgFunc(vValReceiver))

	pValReceiver := &vValReceiver
	fmt.Println(pValReceiver.ValReceiverFunc())
	/*
		fmt.Println(ValArgFunc(pValReceiver)) // cannot use pValReceiver (type *Vertex) 
		as type Vertex in argument to ValArgFunc
	*/

	fmt.Println(ValArgFunc(*pValReceiver)) 
	/*
		*pValReceiver denotes the pointer's underlying value,
		therefore type of *pValReceiver is Vector, 
		where as type of pValReceiver is *Vector
	*/
}