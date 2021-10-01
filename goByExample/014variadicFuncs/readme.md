# Variadic Functions

Variadic functions can be called with any number of trailing arguments.
For example, `fmt.Println()` is a variadic function.

An example of a function signature that will take an arbitrary number of `ints`
as arguments `func sum(nums ...int){}`; also this function doesn't return
anything.

Variadic function can be called the usual way with individual arguments.
If there are multiple args in a slice, apply them to a variadic function
using `func(slice...)`
