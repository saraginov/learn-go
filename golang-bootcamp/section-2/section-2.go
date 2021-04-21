package main

/*
	package main is special, in that
	it allows Go to create an executable package
*/

/*
	import keyword makes another package(library) available
	after import {"fmt"}, package functions become available

	go doc -src fmt Println
	console prints comment description and function itself
*/
import "fmt"

/*
	func main () {} is special, in that
	it is the entry point to an application
	thus it is the first function the Go-lang runtime will execute
		\-> golang cannot call Println() by itself

	functions starting with capital letters,
	ex. func Hello (){} are exported
*/
func main() {
	/*
		golang supports Unicode characters in string literals
		and also in source-code :
	*/
	fmt.Printf("Hello Gopher!")
}

/*
	What is $GOPATH?

		$GOPATH is an environment variable which points to a directory,
		where the downloaded and your own Golang files are stored

		$GOPATH has 3 dirs
			- src: short for source, contains packages written or downloaded by us
			- pkg: short for package, contains binary compiled files
				\-> Golang uses compilation & linking of packages to make builds faster
			- bin: short for binary, contains compiled executable Go programs
				\-> calling go install sourceFile.go, executables are stored here
*/
