package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	/*
		fibonacci series
			initial condition:
			f@0 = 0 ; f@1 = 1
			f@n = f@n-1 + f@n-2, where n>1
	*/
	currentFibNum, prevFibNum, minusTwoFibNum, i := 0, 0, 0, 0

	return func() int{
		if i == 1 {
			prevFibNum = currentFibNum
			currentFibNum = 1
		} else {
			minusTwoFibNum = prevFibNum
			prevFibNum = currentFibNum
			currentFibNum = prevFibNum + minusTwoFibNum
		}
		i++
		return currentFibNum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
	"Shorthand" solution:

		func fibonacci() func() int {
			nextFibNum, plusTwoFibNum := 0, 1
			return func() int{
				currentFibNum := nextFibNum
				nextFibNum, plusTwoFibNum = plusTwoFibNum, nextFibNum + plusTwoFibNum
				return currentFibNum
			}
		}

		Albeit computationally slightly more efficient, I prefer my original solution,
		because it solves for f@n:
			f@n = f@n-1 + f@n-2, for n > 1
			where 
				f@0 = 0 
				f@1 = 1
				f@2 = f@1 + f@0 = 1 + 0 = 1
				f@3 = f@2 + f@1 = 1 + 1 = 2
				f@4 = f@3 + f@2 = 1 + 2 = 3
				f@5 = f@4 + f@3 = 3 + 2 = 5
				f@6 = f@5 + f@4 = 5 + 3 = 8
		on the other hand the "Shorthand" solution solves for f@n+2 and satisfies a re-arranged
		recurrence relation of the Fibonacci series:
			f@n+2 = f@n+1 + f@n, for n+2 > 1 => n > -1
			where
				f@0 = 0 
				f@1 = 1
				f@2 = f@1 + f@0 = 1 + 0 = 1
				f@3 = f@2 + f@1 = 1 + 1 = 2
				f@4 = f@3 + f@2 = 1 + 2 = 3
				f@5 = f@4 + f@3 = 3 + 2 = 5
				f@6 = f@5 + f@4 = 5 + 3 = 8
		On iteration i=9 shorthand solution solves for f@11, where as original for f@9
*/