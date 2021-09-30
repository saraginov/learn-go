# Slices

Slices are a key data type in Go, giving a more powerful interface to sequences
than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number
of elements).
To create an empty slice with non-zero length, use the builtin `make()`.
e.g. create a slice of strings of len 3 (initially zero-valued) `s := make([]string, 3)`

We can get and set values just like with arrays.

`len()` returns the length of the slice.

In addition to above basic operations, we can `append()` which returns a slice
containing one or more new values. Note that we need to accept a return value
from append as we get a new slice value.

Slices can also be `copy()`-ied. e.g. Create empty slice of same length as input:
`c := make([]string, len(s))`.

Slices support a **slice** operator with the syntax `slice[low:high]`.
e.g. `l := s[2:5]` get's a slice of elements s[2],s[3],s[4],
e.g. `l = s[:5]` slices up to, but excluding s[5]
e.g. `l = s[2:]` slices up from (and including) s[2]

Don't specify the length property of an array declaration do declare a slice in a single line
e.g. `t := []string{"g", "h", "i"}`

Slices can be composed into multi-dimensional data structures.
The length of the inner slices can vary, unlike with multi-dimensional arrays.

Note that while slices are different types than arrays, they are rendered similarly
by `fmt.Println()`.

## <https://go.dev/blog/slices-intro>

The type specification for a slice is `[]T`, where **T** is the type of the elements
of the slice. Unlike an array type, a slice type has no specified length

A slice literal is declared just like an array literal, except you leave out the
element count: `letters := []string{"a", "b", "c", "d"}`

A slice can be made with a built-in function called `make()`, which has the signature
`func make([]T, len, cap) []T`, where **T** stands for type of the slice to be
created.

The function `make()` takes a type, a length, and an optional capacity. When called,
`make()` allocates an array and returns a slice that refers to that array.

``` go
var s []byte
s = make([]byte], 5, 5)
// s == []byte{0, 0, 0, 0, 0}
```

When the capacity argument is omitted, it defaults to the specified length.
The most succinct version of the same code `s := make([]byte, 5)`

The length and capacity of a slice can be inspected using the build-in `len()` and
`cap()` functions.

The zero value of a slice is `nil`. The `len()` and `cap()` functions will both return 0 for a nil slice.

A slice can also be formed by "slicing" an existing slice or array.
Slicing is done by specifying a half-open range with two indices separated by a colon.
The start and end indices of a slice expression are optional;
they default to zero and the slice's length respectively
This is also the syntax to create a slice, given an array

``` go
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
c := b[1:4]
// c == []byte{'o', 'l', 'a'}
// c = b[:2] == []byte{'g', 'o'}
// c = b[2:] == []byte{'l', 'a', 'n', 'g'}
// c = b[:] == []byte{'g', 'o', 'l', 'a', 'n', 'g'}

x := [3]string{"hey", "hello", "hola"}
s := x[:] // slice referencing the storage of x
```

## Slice Internals

A slice is a descriptor of an array segment.
It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment)

The length is the number of elements referred to by the slice. The capacity is the number
of elements in the underlying array (beginning at the element referred to by the slice pointer).

Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
This makes slice operations as efficient as manipulating array indices.
Therefore, modifying the elements (not the slice itself) of a re-slice modifies the
elements of the original slice:

``` go
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:]
// e = len 2 cap 2 <- my interpretation
// e == []byte{'a','d'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```

A slice cannot be grown beyond its capacity.
Attempting to do so will cause a runtime panic, just as when indexing outside the bounds
of a slice or array.
Similarly slices cannot be re-sliced below zero to access earlier elements in the array.

``` go
s := make([]byte, 5)
s = s[2:4]
s = s[:cap(s)]
// error thrown by
// sTest = sTest[:4]
```

To increase the capacity of a slice one must create a new, larger slice and copy
the contents of the original slice into it.
This technique is how dynamic array implementations from other languages work
behind the scenes.

``` go
s := make([]byte, 5)
// double the capacity of s by making a new slice, t, copying the contents of
// s into t, and then assigning the slice value t to s
t := make([]byte, len(s), (cap(s)+1*2) // +1 in case cap(s) == 0
for i := range s {
  t[i] = s[i]
}
s = t
```

The looping piece of this common operation is made easier by the build-in copy
function.
As the name suggests, `copy()` copies data from a source slice to a destination slice.
It returns the number of elements copied.
`func copy(dst, src []T) int`
The `copy()` function supports copying between slices of different lengths (it will
copy ONLY up to the smaller number of elements). In addition, `copy()` can handle
source and destination slices that share the same underlying array, handling
overlapping slices correctly.

``` go
// simplify example above
t := make([]byte, len(s), (cap(s) + 1) * 2)
copy(t, s)
s = t
```

A common operation is to append data to to the end of a slice.
The `append()` function can grow the slice anr returns the updated slice value.

``` go
func AppendByte(slice []byte, data ...byte) []byte {
  m := len(slice) // 3
  n := m + len(data) // 6 
  if n > cap(slice) { // if necessary, reallocate
    // allocate double what's needed, for future growth
    newSlice := make([]byte, (n+1)*2)
    copy(newSlice, slice)
    slice = newSlice
  }
  slice = slice[0:n]
  copy(slice[m:n], data)
  return slice
}

test := []byte{'a', 'b', 'c'}
test = AppendByte(test, 'e', 'f', 'g')
// test == {'a', 'b', 'c', 'e', 'f', 'g'}
```

Depending on the characteristics of the program, it may be desirable to
allocate in smaller or larger chunks, or to put a ceiling on the size of a
reallocation.

Go provides a built-in `append()` function that's good for most purposes;
it has the signature `func append(s []T, x ...T) []T`
it appends the elements `x` to the end of the slice `s`, and grows the slice if
a greater capacity is needed.

``` go
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

Since the zero value of a slice (`nil`) acts like a zero-length slice, you
can declare a slice variable and then append() to it in a loop:

``` go
// Filter returns a new slice holding only the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
      if fn(v) {
        p = append(p,v)
      }
    }
    return p
```

## A possible gotcha

Re-slicing does not make a copy of the underlying array. The full array will be
kept in memory until it is no longer referenced.
Occasionally this can cause the program to hold all the data in memory when only a
small piece of it is needed.

``` go
var digitRegexp = regexp.MustCompile("[0-9]+")
func FindDigits(filename string) []byte {
  b, _ := ioutil.ReadFile(filename)
  return digitRegexp.Find(b)
}
```

The code behaves as advertised, but the returned `[]byte` points into an array
containing the entire file.
Since the slice references the original array, as long as the slice is kept around
the garbage collector can't release the array;
the few useful bytes of the file keep the entire contents in memory.

To fix the problem one can copy the interesting data to a new slice before
returning it

``` go
func CopyDigits (filename string) [] {
  b, _ := ioutil.ReadFile(filename)
  b = digitRegexp.find(b)
  c := make([]byte, len(b))
  copy(c,b)
  return c
}

// concise version would be
// var c []byte
// append(c, b)
```
