package main

import "fmt"

/*
	Stringers
		
		One of the most ubiquitous interfaces is Stringer
		defined by the fmt package

			type Stringer interface {
				String() string
			}
		
		A Stringer is a type that can describe itself as a string.
		The fmt package (and many more)
		look for this interface to print values.

		https://golang.org/pkg/fmt/#Stringer
*/

type Person struct {
	Name string
	Age int
}

/* was not implemented in original code base */
type Stringer interface {
    String() string
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main(){
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a,z) // output: Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)

	/* was not implemented in original code base */
	testVar := a.String()
	fmt.Println(testVar) // outputK Arthur Dent (42 years)
}