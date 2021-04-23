# Assignment

Changing variable values after declaration.

``` go
package main
import "fmt"
func main() {
    var speed int
    fmt.Println(speed) // prints 0
    speed = 100
    fmt.Println(speed) // prints 100
    speed = speed - 25
    fmt.Println(speed) // prints 75
}
```

Can't assign variables of different types, even numeric types cannot be mixed
and matched either.

All literals are constants and do not have types which is why

``` go
var force float64
force = 33
```

works and does not throw an error

## Multiple Assignments

Multiple assignment works just like multiple short declaration.

``` go
package main
import ("fmt"; "time")

func main() {
    var (
        speed int
        now time.Time
    )
    speed, now = 100, time.Now()
    
    var (
        currentSpeed = 100
        prevSpeed    = 50
    )
    currentSpeed, prevSpeed = prevSpeed, currentSpeed
}
```

Readability can suffer greatly when assigning values to more than 3 variables in
a single line.

## Function return assignment to multiple variables

``` go
package main
import "fmt"

func main() {
    _, b := multi()
    fmt.Println(b)
}

func multi() (int, int) {
    return 5, 4
}
```
