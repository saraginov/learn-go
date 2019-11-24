package main

import (
	"fmt"
	"math"
)

func sqrt (x float64) string{
	/*
		Go's if statements are like its for loops;
		the expression need not be surrounded by paratheses (), but the {} are required
	*/
	if x < 0{
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64{
	/*
		If with a short statement

		Like for, the if statement can start with a short statement to execute before the condition.

		Variables declared by the statement are only in score until the end of the if

		Google  => (Try using v in the last return statement.) 
		You can't it's not in the scope... 
	*/
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func powTwo (x, n, lim float64) float64 {
	/*
		If and else

		Variables declared inside an if short statement are also available inside any of the else blocks
	*/
	if v := math.Pow(x, n); v < lim {
		return v		
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here as it is outside of if-else block
	return lim
}

func main () {
	fmt.Println(sqrt(2), sqrt(-4))

	/*
		Both calls to pow and powTwo return their results before the call to fmt.Println in main begins 
	*/
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20), // without comma, there is a syntax error, unexpected newline
	)

	fmt.Println(
		powTwo(3,2,10),
		powTwo(3,3,20), 
	)
}