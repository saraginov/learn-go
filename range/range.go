package main

import(
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	/*
		Range
			
			The range "form" of the for loop iterates over a slice or map.

			When ranging over a slice, two values are returned for each iteration
			The first is the index, 
			the second is a copy of the element at that index
	*/
	
	for i, v := range pow {
		 // i is index , v is value, ** is to the power of operator
		fmt.Printf("2**%d = %d\n", i, v)
	}

	/*
		You can skip the index or value by assigning to _.

			for i, _ := range someSplice {...}
			for _, value := range someOtherSplice {...}

		If only index is wanted, omit the second variable

			for i := range anotherSplice {...}
	*/
	powTwo := make([]int, 10) // == [0,1,2,3,4,5,6,7,8,9]
	for i := range powTwo {
		/*
			Overwrite pow[i] with 
			1 << uint(i)
				where uint is [0,1,2,...,9]
				where 1 in binary value is 000 000 000 000 000 000 000 000 000 000 01
			thus 
			1 << uint(0) == 000 000 000 000 000 000 000 000 000 000 01 << 0
				yields
					000 000 000 000 000 000 000 000 000 000 01 
						in decimal 
							1
			1 << uint(1) == 000 000 000 000 000 000 000 000 000 000 01 << 1
				yields
					000 000 000 000 000 000 000 000 000 000 10
						in decimal 
							2
			1 << uint(2) == 000 000 000 000 000 000 000 000 000 000 01 << 2
				yields
					000 000 000 000 000 000 000 000 000 001 00
						in decimal 
							4
			...
			1 << uint(9) == 000 000 000 000 000 000 000 010 000 000 00 << 9
				yields
					000 000 000 000 000 000 000 010 000 000 00
						in decimal 
							512
		*/
		uIntVal := uint(i)
		fmt.Printf("for i:%d uIntVal: %v", i, uIntVal)
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range powTwo {
		fmt.Printf("%d\n", value)
	}
}