# Importing packages

Given some file `example.go`, that has multiple imports:

``` go
package main

// adds print.go, format.go
// fmt belongs to file scope, which is why it can be accessed in the main func
import "fmt"
// declares errors.go in scope of main package
import "errors"
// declares time.go in scope of main package
import "time"

func main() {
    
}
```

## File Scope

Names are visible throughout the file in which they are declared.

Each file needs to independently import the packages refers because importing
packages takes place at the file scope, meaning if 2 files belong to the same
package, then importing a library in 1 file does not make the library's contents
available in the other file.

## Renaming imported packages

To rename package imports, add a name between the `import` keyword and the name
of the package path.

``` go
package main
// importing package fmt as f
import "fmt"
import f "fmt"
// useful when we need to import multiple packages with the same name
func main() {
    f.Print('Hello')
}
```
