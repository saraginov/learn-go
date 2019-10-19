package main

import (
	"fmt"
	"math"
)

/* ---------------------------------------------------------------------------------------------
Interfaces
	An interface type is defined as a set of method signatures
	A value of interface type can hold any value that implements those methods
*/

// Abser is
type Abser interface{ Abs() float64 }

// MyFloat is
type MyFloat float64

// Vertex is
type Vertex struct{ X, Y float64 }

// Abs is
func (f MyFloat) Abs() float64 {
	/*
		this method signature Abs,
		is a value receiver,
		it does not have paramters,
		and it returns a float64
	*/
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Abs is
func (v *Vertex) Abs() float64 {
	/*
		this method signature Abs,
		is a pointer receiver,
		it does not have paramters,
		and it returns an float64
	*/
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interfaces() {
	var a Abser                             // Abs() float64 type == interface
	fmt.Printf("a: %v, a type: %T\n", a, a) // output: a: <nil>, a type: <nil>
	f := MyFloat(-math.Sqrt2)               // -1.414213562, type == MyFloat
	v := Vertex{3, 4}                       // {3, 4} type == Vertex

	a = f                                   // a MyFloat implements Abser; a == -1.414213562, type == MyFloat, MyFloat type == float64
	fmt.Printf("a: %v, a type: %T\n", a, a) //output: a: -1.4142135623730951, a type: main.MyFloat
	fmt.Println(a.Abs())                    // output: 1.4142135623730951

	a = &v                                  // a *Vertex implements Abser
	fmt.Printf("a: %v, a type: %T\n", a, a) // output: a: &{3 4}, a type: *main.Vertex
	fmt.Println(a.Abs())                    // output: 5

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.

	/*
		error: cannot use v(type Vertex) as type Abser in assignment:
			Vertex does not implement Abser (Abs method has pointer receiver)

		SUMMARY:
			identical error to:

				var varTest int
				varTest = "hello world"
	*/

	// a = v
	// fmt.Println(a.Abs())

	/*
		HOWEVER,

		bottom line works because
		remember
		methods with value receivers, take either a value or a pointer as the receiver
	*/
	fmt.Println(v.Abs())
}

/*---------------------------------------------------------------------------------------------
Interfaces are implemented implicitly

	A type implements an interface by implementing its methods.
	There is no explicit declaration of intent, no "implements" keyword

	Implicit interfacese decouple the definition of an interface from its implementation,
	which could then appear in any package without prearrangement
*/

// I is
type I interface{ M() }

// T is
type T struct{ S string }

/*
	This method 'M()' means:
	type T implements the interface I,
	but we don't need to explicitly declare that it does so.
*/

// M is
func (t T) M() {
	fmt.Println(t.S)
}

func interfacesAreImplicit() {
	/*
		when we initialize M(),
		we declare that M() receives a value (value receiver) t of type T;
		and because when type I is initialized and declared,
		we assign M() as its value;
		it follows that when we initialize an instance of type I
		we can pass to it whatever we can pass to M()
		thus
		we can pass it a value of type T to any instance of I
	*/
	var i I = T{"hello"}
	i.M()
}

/* ---------------------------------------------------------------------------------------------
Interface values

	Under the hood, interface values can be thought of as a tuple value and a concrete type:

		(value, type)

	An interface value holds a value of a specific underlying concrete type.

	Calling a method on an interface value
	executes the method of the same name on its underlying type
		ex. in interfacesAreImplicit:

			var i I = T{"hello"}
			i.M()

			// T{"Hello"} is the interface's (i I) value
			// i.M() executes M() for value receiver (t T), ie its underlying type
*/

// IValues is
type IValues interface{ MValues() }

// TValues is
type TValues struct{ S string }

// MValues is
func (t *TValues) MValues() {
	fmt.Println(t.S)
}

// FValues is
type FValues float64

// MValues is
func (f FValues) MValues() {
	fmt.Println(f)
}

func describe(i IValues) {
	fmt.Printf("(%v %T)\n", i, i)
}

func interfaceValues() {
	/*
	 declare an instance of type IValues of name i,
	 remember default init value of an interface is nil
	*/
	var i IValues
	describe(i) // output: (<nil>, <nil>)
	/*
		i.MValues() results in runtime error!

		panic: runtime error: invalid memory address or nil pointer dereference
			[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48ea0c]
	*/
	// i.MValues()

	i = &TValues{"Some String Value"}
	describe(i) // output: ( &{Some String Value} *interfaceValues.TValues )
	i.MValues() // output: Some String Value

	i = FValues(math.Pi)
	describe(i) //output: ( 3.141592653589793... interfaceValues.FValues)
	i.MValues() // output: 3.141592653589793...
}

/* ---------------------------------------------------------------------------------------------
Interface values with nil underlying values

	If the concrete value inside the interface itself is nil,
	the method will be called with a nil receiver.

		NOTE:
			In interfaceValues() above
			before var i IValues is assigned a value
			I added describe(i)
			and i.MValues()
			when if interface method i.MValues() was called
			with a nil value as pointer/value receiver
			a runtime error returned
			thus it imples nil values must be explicitly handled

	In some languages this would trigger a null pointer exception,
	but in Go it is common to write methods that
	handle being called with a nil receiver
	as shown below in interfaceNil()
	when the method MNil() is called

	NOTE:
		!!! An interface value that holds a nil concrete value is itself non-nil
*/

// INil is
type INil interface{ MNil() }

// TNil is
type TNil struct{ S string }

// MNil is
func (t *TNil) MNil() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describeNil(i INil) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfacesNilUnderlyingValue() {
	var i INil

	var t *TNil // t is a pointer to a TNil value, its zero value is nil
	i = t
	describeNil(i) // output: (<nil>, *interfacesNil.TNil)
	i.MNil()       // output:<nil>

	i = &TNil{"hello from i type &tNill"}
	describeNil(i) // output:(&{hello from i type &tNill}, *interfacesNil.TNil)
	i.MNil()       // output: hello from i type &tNill
}

/* ---------------------------------------------------------------------------------------------
Nil interface values

	A nil interface value holds neither value nor concrete type

	Calling a method on a nil interface is
	a run-time error
	because
	there is no type
	inside the interface tuple
	to indicate which
	concrete method to call
*/

// NillInterface is
type NillInterface interface{ NillM() }

func describeNillInterface(i NillInterface) {
	fmt.Printf("(%v %T)\n", i, i)
}

func nilInterfaceValue() {
	var i NillInterface
	describeNillInterface(i) // output: (<nil>, <nil>)

	/*
		panic: runtime error: invalid memory address or nil pointer dereference
			[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48e5af]
	*/
	// i.NillM()
}

/* ---------------------------------------------------------------------------------------------
The empty interface

	The interface type that specifies zero methods
	is know as the
	empty interface

		type SomeInterface interface {}

	An empty interface may hold values of any type
	Every type implements at least zero methods

	Empty interfaces are used by code that handles values of unknown type
	For example
		fmt.Print
		takes any number of arguments of type interface{}
*/

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v %T)", i, i)
}

func emptyInterfaces() {
	var i interface{}
	describeEmptyInterface(i) // output: (<nil>, <nil>)

	i = 42
	describeEmptyInterface(i) // output: (42, int)

	i = "empty interface string"
	describeEmptyInterface(i) // output: (hello, string)
}

//---------------------------------------------------------------------------------------------
func main() {

	interfaces()
	interfacesAreImplicit()
	interfaceValues()
	interfacesNilUnderlyingValue()
	emptyInterfaces()
}
