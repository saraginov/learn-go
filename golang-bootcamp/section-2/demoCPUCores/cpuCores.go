package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.NumCPU returns number of cores as int
	fmt.Println(runtime.NumCPU())
	// add more cpu and download ram xd
	fmt.Println(runtime.NumCPU() * 2)
}
