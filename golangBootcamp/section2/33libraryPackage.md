# Libraries

A Library package, i.e. package that does not belong to `package main` cannot be
executed.

Attempting to `go build somePackageName.go` doesn't do anything, as the package
can be imported directly into any of the `package main` files.

## Archive

To create an archive file, go to directory where library package is.
ex. `./demoLibrary`, invoke `go mod init` in said directory, then invoke
`go install` which adds an archive file to `pkg/github.com/...` directory

The purpose of an archive is a compiled file which reduces program compilation
time for imported packages, i.e. when `import "packageName"` is used, the
archive file is injected.

An archive file is a compressed file and it can be unpacked.

## Exporting

Generally, a package cannot access any declarations from another package, even
if the the other package is imported.

To export a package's declarations, their names must start with a capital
letter. There are no `public` or `private` keywords.

// all other paths were giving me import errors
// when importing a package it will automatically be searched under
// `$GOPATH/src` by the golang compiler
