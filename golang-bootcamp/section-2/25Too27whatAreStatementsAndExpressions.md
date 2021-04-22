# Statements and Expressions

## Statements

Instructions that tell Go to execute something.

Unless a semi-colon(;) is used only 1 statement should be made per line.

``` go
package main
import "fmt"
func main() {
    // all function invocations are statements, string literals are expressions
    fmt.Print("Hello "); fmt.Print(" World\n");
    // or
    fmt.Print("Bye ")
    fmt.Print("World\n")
    
    if 5 > 1 {
        fmt.Println("5 is bigger than 1\n")
    }
}
```

Statements can change the execution flow of code. Go implicitly adds semicolons
during compile time.

## Expressions

Expression is a code block which computes one or more values.

An expression can return more than 1 value.

Expressions should be used with or within a statement, expressions cannot be
used as standalone code.

Some statements like func calls can also act like expressions.

``` go
package main
import "fmt"
func main() {
    fmt.Print("hello" + "!")
}
```

## Importing multiple packages as an expression

``` go
package main
import (
    "fmt"
    "runtime"
)
func main() {
    // any code of block that returns a value is considered an expression
    // in the code below fmt.Println(...) is the statement and runtime.NumCPU()
    // is an expression.
    fmt.Println(runtime.NumCPU())
}
```

## Operator

An operator is used to combine expressions.
