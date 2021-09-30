# Arrays

In Go, an **array** is a numbered sequence of elements of a specific length.

`var a [5]int` creates an array that will hold exactly 5 int(s).
The type of elements and length are both part of the arrays type.
By default an array is zero-valued, which for int(s) means 0.

We can set a value at an index using the `array[index] = value` syntax,
and get a value with `array[index]`.

The built-in `len()` returns the length of an array.

Use `arrName := [n]int{1,3,4,5,6...n}` syntax to declare and initialize an array in one line.

Array types are 1 dimensional but multi-dimensional types can be composed to build
multidimensional data structures.

Slices are more common in Go.

I tried but I can't have multi-type data arrays such as `var twoDTwo [2]int[3]string`

## part 2

Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the
first array element (as would be the case in C). This means that when you assign
or pass around an array value you will make a copy of its contents. (To avoid the
copy you could pass a pointer to the array, but then that's a pointer to an array,
not an array). One way to think about arrays is as a sort of struct but with indexed
rather than named fields: a fixed size composite value.

An array literal can be specified as so: `b := [2]string{"Penn", "Teller"}`
Or you can have the compiler count the array elements for you: `b := [...]string{"Penn", "Teller"}`
in both cases the type of b is [2]string
