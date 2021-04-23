# Type Casting

Similar to macro conversion in C. Use `type(value)` where type is the name of
the type.

``` go
func main() {
    speed := 100 // int
    force := 2.5 // float
    // like in C, float * int returns a float
    // speed = speed * force
    // will yield an error, and we cannot change speed type, so
    // speed = int(speed * force) // potential loss of data, but just ex
    // so much for being a "smart" language
    speed = int(float64(speed) * force)
    // need to cast speed to float in order to be able to multiply it by force
    // then cast it back to int to be able to store it in speed
}
```

Must cast all types all types, event int32 to int.

Type casting does not alter original variable type.

Converting an int to a string, converts the num to ascii character code
translation.
