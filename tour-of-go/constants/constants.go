package main

import(
	"fmt"
)

/*
	Constants:

		constants are declared like variables, but with the const keyword.

		Constans can be character, string, boolean or numeric values,

		Constants CANNOT be declared using the short declaration syntax :=
*/

const Pi = 3.14


/*
	Numeric Constants:

		Numeric constants are high-precision values.

		An untyped constant takes the type needed by its context.

		!!! An int can store at maximum a 64-bit integer and sometimes less
*/

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// ie, binary number that is 1.....000 
	Big = 1 << 100 // equivalent to 10,000,000,000 ^ 10,
	// Shift is right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	
	/*
		from intro, try printing needInt(Big), 
		
		I assume because the int can hold 64bits at most, it won't be able to accurately cast the int as a float in binary
		Whether it will crash or Go will throw an inacurate cast value I do not know, yet...

		VSC Go extension says
		constant 1267650600228229401496703205376 overflows int, which verifies my hypothesis, and I think it will run, 
			just in production it wouldn't be something we should rely on
		
		never mind, there is compile error
		./compile20.go:22:21: constant 1267650600228229401496703205376 overflows int
	*/
	// fmt.Println(needInt(Big))
	

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}