package main

import (
	"fmt"
)

/* ---------------------------------------------------------------------------------------------
	Type assertions

		A type assertion provides access to an interface value's underlying concrete value

			t := i.(T)

		This statement asserts that the interface value i 
		holds the concrete type T 
		and assigns the underlying T value to the variable T
		
		If i does not hold a T, the statement will trigger a panic
		
		IMPORTANT:

			To test whether an interface holds a specific value,
			a type assertion can return two values:
				i) 	the underlying value
				ii) a boolean value that reports whether the assertion succeeded
		
				t, ok := i.(T)

			If i holds a T, then t will be the underlying value 
			and ok will be true.

			If not, ok will be false 
			and t will be the zero value of type T
			and NO PANIC occurs

			NOTE:
				Similarity between this syntax and that of reading from a map
				ie
				var someMapVar = make(map[string]int)
				elem, ok := someMapVar[someKey]
					if someMapVar has the key someKey, ok == true else false 
*/

func typeAssertions(){
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // output: hello

	s,ok := i.(string)
	fmt.Println(s, ok) // output: hello true

	f,ok := i.(float64)
	fmt.Println(f, ok) // output: 0 false
	fmt.Println(f) // output: 0

	// f = i.(float64) // panic
	/*
		panic: interface conversion: interface {} is string, not float64
	*/
}

/* ---------------------------------------------------------------------------------------------
	Type switches

		A type switch is a construct that permits several type assertions in a series.

		A type switch is like a regular switch statement,
		but the cases in a type switch specify types (not values),
		and those values are compared against the type of the value held by the given interface value 

			switch v := i.(type) {
			case T:
				// here v has type T
			case S: 
				// here v has type S
			default:
				// no match; here v has the same type as i	
			}
		
		The declaration in a type switch has the same syntax as a type assetion i.(T),
		but the specific type T is replaced with the keyword type.

		The switch statement above,
		tests whether the interface value i holds a value of type T or S.
		In each of the T and S cases, the variable v will be of type T or S respectively
		and hold the value held by i.
		In the default case (no match), the variable v is of the same interface type and value as i

*/
func do(i interface{}){
	/* 
		Cannot have switch valueOfI, ok := i.(type) {...}
		because switch must have a single value condition to evaluate 
		and it throws syntax error: valueOfI, ok := i.(type) used as value
			... yes I tested, I was curious to see what error would be thrown
	*/
	switch v := i.(type){
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// %q a double-quoted string safely escaped with Go syntax
		fmt.Printf("%q is %v bytes long\n", v, len(v)) 
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeSwitches(){
	do(21)		// output:
	do("hello")	// output: hello is 5 bytes long
	do(true)	// output: I don't know about type bool!
}

//---------------------------------------------------------------------------------------------
func main(){

	typeAssertions()
	typeSwitches()

}