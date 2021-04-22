# EXERCISE

Create a new library
In it, create a function that returns Go version
Create a command and import your library
Call your function that returns Go version
Run your program
Note: The dot below runs all the files in the current directory.
go run . and so on.

## HINTS

Create your package function like this:

``` go
func Version() string {
    return runtime.Version()
}
```

## EXPECTED OUTPUT

It should print the current Go version on your system.

## SOLUTION

Create dir where this exercise will exist.

In same dir create library package `.go` file. Declare function/variables as
needed. Make sure to have capital letter for exported declarations.

Invoke `go mod init`, creates go.mod file for library package.

Create `cmd` dir for main program. Create main.go and add main function.

Import package using path. Invoke functions, use declarations, etc.

To run I didn't build, but could've, i.e. invoke `go run ./cmd/main.go`
