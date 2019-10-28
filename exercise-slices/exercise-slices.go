package main

import (
	"golang.org/x/tour/pic"
)

/*
	open import url (golang.org/x/tour/pic)
		|-> will notice pic.Show requires an input:
			f func(int, int) [][]uint8

	thus we can establish that we don't need to modify any of the starter code in func main(){...}
	and that we will pass func Pic(dx, dy int) [][]uint8 as the argument

	next, declare and initialize the data we need to return (returnSlice)
		there are 2 solutions from content covered so far
			i) 		nil slice
			ii) 	make(slice, len, cap)
		I chose make because "return a slice of length dy", implying length must == dy

	next, we need to generate the data for each slice in the return
		to do so iterate of dy OR len(returnSlice)
			for i := 0; i < dy; i++ {...}

	next, from the problem statement "each element (of returnSlice)... is a slice of dx 8-bit
	 unsigned integers"
		 we can only infer that each element of returnSlice is in turn a slice of
		 what we assume an unknown length and capacity
	thus
		we declare and initialize a nil slice to which we will append as many uint8
		values, as dx says there are
	declare and initialize :
		'var returnSliceElement []uint8'
		NOTE : each time we iterate over dy, we re-initialize returnSliceElement

	next, iterate over dx and append a uint8 value to returnSliceElement
		for j := 0; j < dx; j++ {
			returnSliceElement = append(returnSliceElement, uint8)
		}

	once we have iterated over the entirety of len(dx) and generated the relevant
	returnSliceElement, assign it to the return slice for the corresponding y
		returnSlice[i] = returnSliceElement

	next we need to determine the values for each uint8
		from the problem statement
			"The choice of image is up to you.
			Interesting functions include (x+y)/2, x*y, and x^y."
		there are 2 solutions we could've implemented
			i) generate two array(s)/slice(s), one of dy and one of dx length,
				where each element in the said "list" is a random int from the math package
				(as we are not constrained by the math expression(x+y)/2;
					thus x and y can be anything)
				then using the i and j indices retreive x and y from the data set
				|-> since we increment i and j by 1 each iteration, we can use the changing
					indices' value as x and y respectively
			ii) use dx & dy as x and y
				since dx and dy are static int(s) we end up with single color image
*/

// Pic is a function that generates a pattern using slices based on x to y expression
func Pic(dx, dy int) [][]uint8 {
	returnSlice := make([][]uint8, dy)

	for i := 0; i < dy; i++ {

		var returnSliceElement []uint8

		for j := 0; j < dx; j++ {
			//returnSliceElement = append(returnSliceElement, uint8((i+j)/2))
			//returnSliceElement = append(returnSliceElement, uint8(i*j))
			returnSliceElement = append(returnSliceElement, uint8(i^j))
		}

		/*sample output returnSliceElement := []uint{0,1,2,3,4,5,6,7,8,9....dx}*/
		returnSlice[i] = returnSliceElement
	}
	return returnSlice
}

func main() {
	pic.Show(Pic)
}
