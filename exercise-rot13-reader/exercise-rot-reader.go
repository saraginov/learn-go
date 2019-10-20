package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for i := 0; i < len(p); i++ {
		char := p[i]
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			p[i] += 13
		} else if (char > 'M' && char <= 'Z') || (char > 'm' && char <= 'z') {
			p[i] -= 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Printf("s: %v\n", s)
	r := rot13Reader{s} // io.Reader is &Reader
	fmt.Printf("r: %v\n", r)
	fmt.Printf("r.r: %v\n", r.r)
	written, err := io.Copy(os.Stdout, &r)
	fmt.Printf("\nwritten: %v err = %v \n", written, err)
}
