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

## Printf

``` go
func main() {
    var brand = "Google"
    fmt.Printf("%q\n", brand)
    total, success, fail := 2350, 543, 433
    fmt.Printf("total: %d success: %d / %d\n", ops, lk, fail)
}
```

### Escape Sequence

``` go
// \n \" \\ 
// \\n prints \n
fmt.Printf("hi\nhi") // 5 char len, '\n' escape sequence counts as 1 char
fmt.Printf("hi\\n\"hi\"") // hin"hi"
```

### Printing type

``` go
// can print type of any value, whether variable or not
var speed int
fmt.Printf("%T\n", speed)
fmt.Printf("%T\n", 26.3)
```

### Printing values && Argument Index

``` go
var (
    planet = "venus"
    distance = 261
    orbital = 224.701
    hasLife = false
)
// %v can print any value but not type safe
// printing values
fmt.Printf("Planet: %v\nDistance: %v millions kms\n Orbital Period: %v days\n",
    planet, distance, orbital);
fmt.Printf("Does %v have life? %v\n", planet, hasLife)
fmt.Printf("Planet: %s\nDistance: %d millions kms\d Orbital Period: %f days\n",
    planet, distance, orbital);
fmt.Printf("Does %s have life? %t\n", planet, hasLife)
// Argument Index
fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)
fmt.Printf("%s is %d away. Think! %[2]s kms! %[1]d OMG!\n", planet, distance)
// Precision
fmt.Printf("Orbital Period: %.0f days\n", orbital) // prints 225
fmt.Printf("Orbital Period: %.1f days\n", orbital) // prints 225.7
fmt.Printf("Orbital Period: %.2f days\n", orbital) // prints 225.70
fmt.Printf("Orbital Period: %.4f days\n", orbital) // prints 225.7010
```

## Pass name as cli args

``` go
package main
import (
    "fmt"
    "os"
)
func main() {
    if len(os.Args > 1) {
        fmt.Printf("Your name is ")
        for i:=0; i<len(os.Args); i++ {
            fmt.Printf("%v", os.Args[i] )
        }
        fmt.Printf("\n")
    } else {
        fmt.Println("No name provided\n")
    }
}
```
