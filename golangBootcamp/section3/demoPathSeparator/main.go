package main

import (
	"fmt"
	"path"
)

// path.Split() declaration
// func Split(path string) (dir, file string)
// means it takes path as input arg
// and returns 2 strings, dir and file

func main() {
	dir, file := path.Split("asset/agreement.pdf")
	fmt.Println(dir)
	fmt.Println(file)

	// if you don't care about dir for instance
	var newFile string
	_, newFile = path.Split("someRandomDir/style.css")
	fmt.Println(newFile)

	color := "green"
	// `"dark " + color` is an expression
	color = "dark " + color
	fmt.Println(color)

	speed := 100                        // int
	force := 2.5                        // float
	speed = int(float64(speed) * force) // potential loss of data, but just ex
	fmt.Println(speed)
}
