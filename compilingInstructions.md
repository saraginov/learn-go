# Compiling Instructions

``` go
go build [-o output] [-i] [build flags] [packages]
// ex.
go build -o ../bin/execFileNameWeWant
```

Source: <https://golang.org/cmd/go/>

Build compiles the packages named by the import paths, along with their
dependencies, but it does not install the results.

If the arguments to build are a list of .go files from a single directory,
build treats them as a list of source files specifying a single package.

When compiling packages, build ignores files that end in '_test.go'.

When compiling a single main package, build writes the resulting executable to
an output file named after the first source file ('go build ed.go rx.go' writes
'ed' or 'ed.exe') or the source code directory ('go build unix/sam' writes 'sam'
or 'sam.exe'). The '.exe' suffix is added when writing a Windows executable.

When compiling multiple packages or a single non-main package, build compiles
the packages but discards the resulting object, serving only as a check that the
packages can be built.

The -o flag forces build to write the resulting executable or object to the
named output file or directory, instead of the default behavior described in the
last two paragraphs. If the named output is a directory that exists, then any
resulting executables will be written to that directory.

The -i flag installs the packages that are dependencies of the target.

The build flags are shared by the build, clean, get, install, list, run, and
test commands:
