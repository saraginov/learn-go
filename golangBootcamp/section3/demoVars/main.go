package main

import "fmt"

func main() {
	var hello1 string = "hello1"

	var hello2 = "hello2"

	var hello3 string
	hello3 = "hello3"

	hello4 := "hello4"

	fmt.Println(hello1, hello2, hello3, hello4)

	var1, var2 := 1, "var2"

	fmt.Println(var1, var2)

	var1 = 30
	fmt.Println(var1)

	var1, var3 := 40, 50
	fmt.Println(var1, var3)

	// we must declare variable
	// var1 := 50
	fmt.Println(var1, var3)

	// sum should be 30.5
	sum := 27 + 3.5
	fmt.Println(sum)

	sup, sup2, sup3 := "sup", "", "sup3"
	fmt.Println(sup, sup2, sup3)
}
