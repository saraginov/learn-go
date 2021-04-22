# What is Scope?

Scope refers to accessibility, i.e. who can see what.

There are package, file, func and block scopes.

## Declarations

A unique name bound to a scope. Same name cannot be declared a second time
inside the same scope.

## Package Scope

``` go
package main
// import is file scoped, i.e. only visible in this file
import "fmt"

// const ok and func main(){} are package scoped, visible to all files which
// belong to the main package. Other packages cannot see the main package 
// declaration because it can't be imported, only executed after compiling
const ok = true
func main() {
    // anything enclosed in curly braces, {...} is considered to fall within
    // said block scope
    var hello = "Hello"
    fmt.Println(hello, ok)
}
```

## Variable visibility

Variables can only be accessed within their scope only after they have been
declared, i.e. no hoisting like in JS.
