package main

import (
	"fmt"
	"io"
	"strings"
)

/*
	Readers

		The io package specifies the io.Reader interface,
		which represents the read end of a stream of data.

		The GO standard library contains many implementations of these interfaces,
		including files, network connections, compressors, ciphers and others.

		The io.Reader interface has a Read method:

			func (T) Read(b []byte) (n int, err error)

		Read populates the given byte slice with data
		and returns the number of bytes and an error value

		It returns an io.EOF error when the stream ends.

		The example below creates a strings.Reader and consumes its output 8bytes at a time


		IMPORTANT:-TODO:

			I do not yet understand how and where the pointer to b is implemented
			and even in https://golang.org/src/io/io.go
			how the original Read(p []byte) (n int, err error) {} methods which never
			assign a value to, and are not passed a pointer to p can change the value of p
			outside the scope of the nested function

			ie.
				r := strings.NewReader("Hello Reader!")
				b := make([]byte, 8) // => [0 0 0 0 0 0 0 0 0 0] and type == []uint8
				for {
					n, err := r.Read(b)
						// each time we call Read b gets overwritten, but there is no pointer to b
						// and b is not assigned a return either...

						VERY_IMPORTANT:
						how does 'Read reads up to len(p) bytes into p'? *******!!!!

						// the action is equivalent to declaring and initializing
						// var scalar int in func main, passing scalar to v.Scale(scalar)
						// and overriding its value inside of the method declaration
						// in the example below
						// https://tour.golang.org/methods/8

						// func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
						// so it makes sense for r to be updated,
						// and I can't find where b would be assigned as a property of r
						// in the golang source code
					if err == io.EOF {
						break
					}
				}


*/

func main() {
	// https://golang.org/src/strings/reader.go L154
	r := strings.NewReader("Hello Reader!")

	b := make([]byte, 8)

	c := &b

	// fmt.Printf("r = %v\n", r)
	fmt.Printf("b = %v , type = %T\n ", b, b)
	fmt.Printf("c = %v , type = %T\n ", c, c)

	for {
		fmt.Println("running for loop")
		// https://golang.org/src/strings/reader.go L38
		n, err := r.Read(b)

		/*
			to be able to read ASCII Code byte values for b
			https://web.stanford.edu/class/cs101/bits-bytes.html
		*/
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)

		fmt.Printf("b[:n] = %q\n", b[:n])

		fmt.Printf("c = %v , type = %T\n ", c, c)

		if err == io.EOF {
			break
		}
	}

}
