package main

import "fmt"

func main() {
	sum := 0
	/*
		Go has only one looping construct, the FOR loop

		The basic for loop has three components separated by semicolons:

			i) the init statement: executed before the first iteration

			ii) the condition expression: evaluated before every iteration

			ii) the post statement: executed at the end of every iteration

		The init statement will often be a short variable declaration, and the variables declared there
		are visible only in the scope of the for statement.
		
		The loop will stop iterating once the boolean condition evaluates to false

		Note: Unlike other languages like C, Java, or Javascript there are no parantheses surrounding
		the three components of the for statement and the braces { } are always required 
	*/
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)


	/*
		The init and post statements are optional
	*/
	sumTwo := 1
	for ; sumTwo < 1000; {
		sumTwo += sumTwo
	}
	fmt.Println(sumTwo)

	/*
		For is Go's "while" loop
	*/
	sumThree := 1
	for sumThree < 1000 {
		sumThree += sumThree
	}
	fmt.Println(sumThree)

	/*
		Forever = INFINITE Loop

		If loop conditions are imitted, the loop, loops forever
		an infinite loop is compactly expressed 
		
		for{
		
		}

		sum := 1
		for {
			sum += sum
			fmt.Println(sum)
		}

		After I ran the above loop, after max int size was reached, 
			last three prints were:
				2305843009213693952
				4611686018427387904
				-9223372036854775808
		subsequent prints of sum were 0, which to me implies that once sum could no longer be evaluated
		it was re-initialized to Go numeric default value
	*/

}