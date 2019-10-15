package main

import (
	"fmt"
	"runtime"
	"time"
)

func main () {
	fmt.Print("Go runs on ")

	/*
		A switch statement is a shorter way to write a sequence of if - else statements.

		It runs the first case whose value is equal to the condition expression.

		Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs
		the selected case not all the cases that follow.

		In effect, the break statement that is needed at the end of each case in those languages
		is provided automatically in Go.

		Another important difference is that Go's switch cases need not be constants, 
		and the values involved not be integers 
	*/

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		//plan9, windows....
		fmt.Printf("%s.\n", os)
	}

	/*
		.Now returns current local time ; .Weekday returns the day of the week specified by t 	
		Calculations on time are autoformatted, 
	*/
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	
	fmt.Printf("%s \n" ,today) // prints today's day, ie Monday, Tuesday, Wednesday, etc. 
	fmt.Print(today+3) // today + 0  is today's index, if today == Monday, today + 2 == Wednesday
	fmt.Print("\n")

	/*
		Switch evaluation order
		
			Switch cases evaluate cases from top-to-bottom, stopping when a case succeeds
			ex.
			switch i {
			case 0:
			case f():
			}
			does not call f if i==0
	*/
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2 :
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away.")
	}

	/*
		Switch with no condition
		
			Switch without a condition is the same as switch true {...}

			This construct can be a clean way to write long if-then-else chains
	*/

	t := time.Now()
	switch {
	case t.Hour() <14:
		fmt.Println("Good Morning!")
	case t.Hour() <17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}