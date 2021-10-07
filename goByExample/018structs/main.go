package main

import "fmt"

type person struct {
	name string
	age  int
}

// It’s idiomatic to encapsulate new struct creation in constructor functions
func createPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// can safely return a pointer to local variable as a local variable will
	// survive the scope of the function
	// question: when is this data released from memory?
	return &p
}

func main() {
	fmt.Println(person{"Bob", 30})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Anne", age: 40})
	fmt.Println(createPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Access struct fields with a dot.

	sp := &s
	fmt.Println(sp.age)

	// You can also use dots with struct pointers
	// the pointers are automatically dereferenced.
	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s)

	var james person = person{"James", 30}
	fmt.Println(james)
}
