package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	// because n has no type specified during compilation, n gets converted to
	// float 64 in the line below
	fmt.Println(math.Sin(n))

	// can I use it as an int thereafter? or will I get an error
	fmt.Println(math.Pow10(n))
	// it can be used, and is converted to appropriate type in each invocation
}
