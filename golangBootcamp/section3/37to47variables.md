# Data Types And Variables

Every variable in Go has a type.

## Basic/Primitive Data Types

### numeric

``` go
int  // integer literals, e.g. -1 0 27
float64 // float literals, e.g. -.5 0. 1.
bool // predeclared constants, e.g. true false
string // string literal  e.g. "hi there" each char is 1-4 bytes
// it shows limits of understanding, from Inanc... all characters are 4 bytes
// or w.e smallest ram division on OS is, if binary less than 4 digits, preface
// considered padding... just like unions and structs
```

Integer literals, are denoted by data type `int`. Literal means value itself.

## Declaring Variables

Variables MUST be declared before they can be used.

``` go
// general syntax
var nameOfVar identifier
var someNum int
```

Variable names can start with a letter (upper/lower case, exported vs local), an
underscore, `_`, e.g. `_myVar`, or a unicode character.

A name cannot start with a number or punctuation. Cannot contain punctuation
whatsoever.

Obviously can;t use reserved keywords as variable names.

Once a variable's type has been declared, it cannot be changed!

## zero values

boolean: false
numeric: 0
string: ""
pointer: nil

## Unused variable error

All variables declared inside block scopes must be used,
however variables declared at the file scope do not have to be,
because it is not clear to the compiler at what point in the program's life
the programmer will use file scoped variable (i.e. another package may
refer to it now or in the future).

``` go
package main
import "fmt"

// will not cause error
var hello int

func main() {
    // secondHello will cause error
    var secondHello string        
}
```

use the `_` character for a blank-identifier, normally used for skipping return
values.

## Multiple declarations

``` go
package main
import "fmt"
func main() {
    var (
        speed int
        heat  float64
        off   bool
        brand string
    )
    var firstName, lastName string = "", ""
    fmt.Printf("%q %q\n", firstName, lastName)
    
    var otherSpeed, velocity int
    fmt.Println(speed, velocity)
    
    // ...
}
```

``` go
    var i int
    var i8 int8
    var i16 int16
    var i32 int32
    var i64 int64

    // float types
    var f32 float32
    var f64 float64

    // complex types
    var c64 complex64
    var c128 complex128

    // bool type
    var b bool

    // string types
    var s string
    var r rune  // also a numeric type
    var by byte // also a numeric type
```
