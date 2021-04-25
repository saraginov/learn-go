# Program input

## Getting command line args && Slices Introduction

`os package` allows you to access operating system functionalities.

Go can't access keyboard arguments by default.

Inside the os package, there is an exported variable `var Args []string` which
is a slice of string inputs.

Args is a slice of strings. A slice is a data structure that can store a series
of specified data types, in this case strings.

Each value in a slice is stored in an unnamed variable.

When `go run main.go hi yo` is invoked in cli, the strings 'hi' and 'yo' are
stored in a slice.

Using the `Args` variable we can access the cli arguments passed to the program
by using `Args[0]; Args[1]` and so on, i.e. the index.

The program stores the temporary path to the currently executing program at its
first index. At the second index it stores the first command line argument.
