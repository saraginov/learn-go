# Multiple Return Functions

Go has built-in support for **multiple return values**.
This feature is often used in idiomatic Go, for example to return both the result
and error values from a function.

``` go
func vals() (int, int){
  return 1, 2
}
```

The `(int, int)` in the function signature above specifies that the function
returns 2 ints.

## Multiple assignments

``` go
// using the vals() function from above
a, b := vals()
```
