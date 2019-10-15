package main

import (
	"fmt"
)

func OurSqrt(x float64) float64{

	/*
		Similar to one point iteration OR successive substitution, 
		Open Methods Simple Fixed-Point Iteration x@i+1 = g(x@i) 
	*/ 

	/*
		x@i+1 = x@i - ( f(x@i) / f'(x@i) ) // Newton-Rhapson Applied Numerical Methods Steven C. Chapra
		f(x@i) = (x@i)^2 - c // where c is constant ( equivalent to x float64 = 2 in func main) 
		f'(x@i) = 2*(x@i) // power function

		thus
		x@i+1 = x@i - ((x@i)^2 - c / 2*(x@i))

		where x@0 != 0, because cannot divide by 0; 
		thus
		x@0 = 1

		in turn
		x@1 = x@0 - ((x@0)^2 - (x float64) / 2*(x@0))
		x@1 = 1 - ((1)^2 - 2) / (2 *1) = 1.5 

		x@2 = x@1 - ((x@1)^2 - (x float64) / 2*(x@1))
		x@2 = 1.5 - ((1.5)^2 - 2) / (2 * 1.5) = 1.41666666...

		consequently 
		break condition x@i+1 - x@i < 0.001; 
		in other words for x@i+1 - x@i > 0.001 {...keep looping}
	*/


	z := float64(1)
	zPrev := float64(0)
	
	for z - zPrev > 0.001 {
		zPrev = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z, zPrev)
	}
	
	return z
}

func main()  {
	fmt.Println(OurSqrt(2))
}