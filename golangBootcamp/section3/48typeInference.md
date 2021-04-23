# Initialization

``` go
package main
import "fmt"
func main() {
    var save bool = true
    fmt.Println(safe)
    
    var inferredBool = true
    fmt.Println(inferredBool)
}
```

Go deduces the type of the variable automatically if the variable is initialized
and declared at the same time.

## Removing var

``` go
func foo() {
    hello := "hello"
    // functionally the same as
    var hello2 = "hello2"
    // or 
    var hello3 string = "hello3"
}
```

Using the `varName := "value"` is called the short declaration statement.

Short declaration statements cannot be used without values.

## Can't short declare var in package/file level

short declaration cannot be used at packet/file level because every statement in
package/file level scope must start with a keyword!

``` go
package main

var var1 = true
var var2 bool = false
// short declaration below will cause error, 2 declarations above will not
var3 := "opa"

func main() {}
```

## Short declare multiple variables

Number of variables and number must match the number of values, nth variable
gets corresponding nth value.

Type inference works the same way for multiple variables as it does for a single
variable.

``` go
package main
func main() {
    safe, speed := true, 50
}
```

## Redeclaration

short declaration can be used with existing variables, this is known as redeclaration.

``` go
package main
func main() {
    var hello string
    hello, speed := "Hello World", 20
    
    // useful when using error
    file, err := os.Open("...") if err != nil
    errdata, err := ioutil.ReadAll(file) if err != nil {
        
    }
}
```

## When to use short declaration

If you do not know the initial value do not use short declaration.

Don't use short declaration for package scope.

``` go
// don'ts
score := 0 // Don't
// dos
var score int // 0 by default
```
