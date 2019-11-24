package main

import (
	"fmt"
	"math"
)

/*
	In Go, a name is exported if it begins with a capital letter.
	For example, Pizza is an exported name, as is Pi, which is exported from the math package/

	on the other hand pizza and pi do not start with a capital letter, hence they are not exported.


	When importing a package, you can refer only to its exported names.
	Any "un-exported" names are not accessible from outside the package.
*/
func main() {
	/*
		math.pi throws an error when compiing
		# command-line-arguments
		./compile3.go:9:14: cannot refer to un-exported name math.pi
		./compile3.go:9:14: undefined: math.pi
	*/
	// fmt.Println(math.pi)

	fmt.Println(math.Pi)
}
