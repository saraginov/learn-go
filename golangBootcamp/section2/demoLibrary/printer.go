package printer

import "fmt"

func hello() {
	fmt.Println("Un-exported Hello")
}

func Hello() {
	fmt.Println("Exported Hello")
}

const AMIDECLARED = true
