package main

import "fmt"

/*
	A struct is a collection of fields (I assume short for structure)
*/
type Vertex struct {
	X int
	Y int
	// can be written as X, Y int
}

/*
	Struct Literals

	A struct literal denotes a newly allocated struct value by listing the values of its fields
	You can list just a subset of fields by using the Name: suntax 
		IMPORTANT: Order of named fields is irrelevant

	The special prefix & returns a pointer to the struct value
*/
var (
	v1 = Vertex{1,2} 	// has Type Vertex
	v2 = Vertex{X:3} 	// Y:0 is implicit
	v3 = Vertex{}		// X:0 and Y:0 
	p = &Vertex{11,12}	// has type *Vertex
)

func main() {
	/*
		Expected return {1 2} {3 0} {0 0} &{11 12}
	*/
	fmt.Println(v1, v2, v3, p)


	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	/* Struct fields are accessed using a dot */
	v.X = 4

	fmt.Println(v.X)

	/*
		Pointers to struct(s)

		Struct fields can be accessed through a struct pointer.

		To access the field X of a struct when we have the struct pointer p we could write (*p).X

		IMPORTANT!
		However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference!

	*/
	vTwo := Vertex{1, 2}
	p := &vTwo
	p.X = 1e9 // Go permits us to access pointer without explicit dereference; (*p).X = 1e9
	fmt.Println(vTwo)


}