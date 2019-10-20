package main

import (
	"fmt"
	"strings" //Package strings implements simple functions to manipulate UTF-8 encoded strings.
)

// var outSideOfFunc []int = []int{0,1,2,3,4,5}
var outSideOfFunc = []int{0, 1, 2, 3, 4, 5}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	/*
		Slices

			An array has a fixed size.
			A slice on the other hand, is dynamically-sized,
			flexible view into the elements of an array.

			In practice, slices are much more common than arrays.

			The type []T is a slice with elements of type T

			A slice is formed by specifying two indices, a low and a high bound,
			separated by a colon (:)

			a[low : high]

			This selects a half-open range which includes the first element,
			but excludes the last one

			ex. create slice (called a) with elements 1 through 3
				a [1:4]
	*/

	var s []int = primes[1:4]
	fmt.Println(s)

	/*
		Slices are like references to arrays

			A slice does not store any data, it just describes a section of an underlying array

			Changing the elements of a slice modifies the corresponding elements of its underlying array

			Other slices that share the same underlying array will see those changes.
	*/

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2] // pointer by reference a = ["John", "Paul"]
	b := names[1:3] // b = ["Paul", "George"]

	fmt.Println(a, b)

	/*
		makes b = ["XXX", "George"], inturn a = ["John", "XXX"],
		names = ["John", "XXX","George", "Ringo",]
	*/
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/*
		Slice literals

			A slice literal is like an array literal without the length

			This is an array literal var someArray[3]bool = [3]bool{true, false, true}
			OR someArray[3]bool{true, false, true}

			And this creates the same array as above, then builds a slice that references it:
			var someSlice[]bool = []bool{true, false, true}
			OR someSlice := []bool{true, false, true}

			NOTE :
				Slices and Arrays are different
				var someSlice[]bool = [5]int{1,2,3,4,5,} IS INVALID
	*/

	var someSlice []bool = []bool{true, false, true}
	fmt.Println(someSlice)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sTwo := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sTwo)

	/*
		Slice Defaults

			When slicing, you may omit the high or low bounds to use their defaults instead

			The default us zero for the low bound and the length of the slice for the high bound.
				I think typo in docs and length of slice should say length initializing array
				If the difference between an array and a slice is length how can a slice have a default length?

			For the array

				var a [10]int

			the following slice expressions are equivalent
				a[0:10]
				a[:10]
				a[0:]
				a[:]
	*/

	sThree := []int{2, 3, 5, 7, 11, 13}
	sThree = sThree[1:4]
	fmt.Println(sThree) // output: [3,5,7]
	sThree = sThree[:2]
	fmt.Println(sThree) // output: [3,5,]
	sThree = sThree[1:]
	fmt.Println(sThree) // output: [5,]

	sFour := []int{12, 13, 15, 17, 111, 113}
	sFour = sFour[:]
	fmt.Println(sFour) // output: [12,13,15,17,111,113]

	/*
		Slice length and capacity

			A slice has both a length and capacity.

			The length of a slice is the number of elements it contains

			The capacity of a slice is the number of elements in the underlying array,
			counting from the first element in the slice.

			The length and capacity of a slice can be obtained using the expressions len(s) and cap(s)

			You can extend a slice's length by re-slicing it,
			provided it has sufficient capacity.

		CAPACITY :
			given we have some slice var sliceX []int = []int{1,2,3,4,5,6}
				where len=6, cap=6

			when we re-slice sliceX, where sliceX = sliceX[low:high]
			capacity is determined by subtracting 'low' from the original sliceX capacity (cap=6)

			for example
				sliceX = sliceX[0:high], capacity will remain equal to 6
				'high' only impacts the length of the slice
				sliceX = slice[1: high], capacity will now be 6-1 = 5
				sliceX = slice[3: high], capacity will now be 5-3 = 2

			Extend Slice beyond capacity:
				high MUST be less than or equal to capacity,
				else
					panic: runtime error: slice bounds out of range [low int: high int] with capacity int
	*/

	sFive := []int{2, 3, 5, 7, 11, 13}
	printSlice(sFive)
	// output: len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	sFive = sFive[:0]
	printSlice(sFive)
	// output: len=0 cap=6 []

	// Extend its length.
	sFive = sFive[:4]
	printSlice(sFive)
	// output: len=4 cap=6 [2 3 5 7]

	// Drop its first two values
	sFive = sFive[2:]
	printSlice(sFive)
	// output: len=2 cap=4 [5 7]

	// Recreates slice
	sFive = sFive[:4]
	printSlice(sFive)
	// output: len=4 cap=4 [5 7 11 13]

	// Slice 7,11
	sFive = sFive[1:3]
	printSlice(sFive)
	// output: len=2 cap=3 [7 11]

	/*
		Nil Slices

			The zero value of a slice is nil

			A nil slice has a length and capacity of 0 and has no underlying array
	*/

	var sSix []int
	fmt.Println(sSix, len(sSix), cap(sSix))
	if sSix == nil {
		fmt.Println("nil!")
	}

	/*
		Creating a slice with make

			Slices can be created with the built-in make function;
			this is how you create dynamically-sized arrays

			The make function allocates a zeroed array and returns a slice that refers to that array

				a := make([]int, 5) // len(a)=5

			To specify a capacity, pass a third argument to make:

				b := make([]int, 0, 5 ) // len(b)=0 cap(b)=5

				b = b[:cap(b)] // len(b)=5 cap(b)=5
				b = b[1:] // len(b)=4 cap(b)=4
	*/

	aTwo := make([]int, 5)
	printSliceWithString("aTwo", aTwo)
	// expected output: aTwo len=5 cap=5 [0,0,0,0,0]

	bTwo := make([]int, 0, 5)
	printSliceWithString("bTwo", bTwo)
	// expected output: bTwo len=0 cap=5 []

	cTwo := bTwo[:2]
	printSliceWithString("cTwo", cTwo)
	// expected output: cTwo len=2 cap=5 [0,0]
	dTwo := cTwo[2:5]
	printSliceWithString("dTwo", dTwo)
	// expected output: dTwo len=3 cap=3 [0,0,0]

	/*
		Slices of Slices

			Slices can contain any type, including other slices
	*/

	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players "take" turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	/*
		expected output:
			X _ X
			O _ X
			_ _ O
	*/

	/*
		Appending to a slice

			It is common to append new elements to a slice,
			Go provides a built-in function called "append"

				https://golang.org/pkg/builtin/#append

			func append(s []T, vs ...T)[]T

				The first parameter s of append is a slice of type T,
				the rest are T values to append to the slice

				The resulting value of append is a slice containing
				all the elements of the original slice
				plus the provided values

			If the backing array of s is too small to fit all the given values
			a bigger array will be allocated
			The returned slice will point to the newly allocated array.

			https://blog.golang.org/go-slices-usage-and-internals
	*/

	var sSeven []int
	printSlice(sSeven) // expected output: len=0 cap=0 [] ; [] == nil yields true

	// append works on nil slices
	sSeven = append(sSeven, 0)
	printSlice(sSeven) // expected output: len=1 cap=1 [0]

	// slice can grow as needed
	sSeven = append(sSeven, 1)
	printSlice(sSeven) // expected output: len=2 cap=2 [0, 1]

	// we can add more than one element at a time
	sSeven = append(sSeven, 2, 3, 4)
	printSlice(sSeven) // expected output: len=5 cap=5 [0, 1, 2, 3, 4]

	/*
		IMPORTANT :
			given sSeven == len=2 cap=2 [0, 1]
			sSeven = append(sSeven, 2, 3, 4, 5, 6, 7, 8)
			VS.
			sSeven = append(sSeven, 2, 3, 4, 5, 6, 7, 8, 9,) // with or without comma following last element

			both will yield a capacity of 10, however len will be 9 and 10 respectively

			similarly, with identical starting condition
			sSeven = append(sSeven, 2, 3, 4, 5, 6,)
			VS.
			sSeven = append(sSeven, 2, 3, 4, 5, 6, 7,)
			both will yield a capacity of 8, however len will be 7 and 8 respectively
	*/

	j := s[:10]
	printSlice(j)
	// expected output: len=10 cap=10 [0 1 2 3 4 5 6 7 9 0]

	k := s[:]
	printSlice(k)
	// expected output: len=9 cap=10 [0 1 2 3 4 5 6 7 9]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceWithString(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
