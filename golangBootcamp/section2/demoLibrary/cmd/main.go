package main

// all other paths were giving me import errors
// when importing a package it will automatically be searched under
// `$GOPATH/src` by the golang compiler

// important: do not add '/' at the end of directory name, do not add printer.go
// or printer after the directory where our package is provided.

// I think that over time as golang has evolved, the conventions used by Inanc
// is no longer valid -> the only error free execution is from the one below
// further more I had to call go mod init a second time and I changed the dir
// names so that they do not have '-' in them.

// After further testing, it turns out that the '-' symbol in the dir names was
// irrelevant, the only problem was that `go mod init` needs to be invoked after
// every package update, or path update. MUST delete go.mod file in dir where
// library is for `go mod init` to work as intended

import (
	"fmt"

	printer "github.com/saraginov/learnGo/golangBootcamp/section2/demoLibrary"
)

// for some reason, not debugging now, the import bath above breaks

func main() {
	const shouldBe = printer.AMIDECLARED

	fmt.Println("const should be set to exported value true, actual val: ", shouldBe)

	printer.Hello()
}
