package main

import "fmt"

func main() {

	/*
		Go has pointers, 
		A pointer holds the memory address of a value

		The type *T is a pointer to a T value. Its zero value is nil

		ex.
			var p *int

		The & operator generates a pointer to its operand
		
		ex.
			i := 42
			p = &i

		The * operator denotes the pointer's underlying value

		fmt.Println(*p) // read i thorugh the pointer p
		*p = 21 // set i through pointer p

		This is known as "dereferencing" and "indirecting"

		Unlike C, Go has no pointer arithmetic

	*/
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p) // read i through the poiinter
	*p = 21 // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j
}