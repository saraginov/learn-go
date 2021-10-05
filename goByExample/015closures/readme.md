# Closures

Go supports anonymous functions, which can form closures.
Anonymous functions are useful when we want to define a function inline
without having to name it.

``` go
func intSeq() func() int{
  i := 0
  return func() int {
    i++
    return i
  }
}
```

The function `intSeq` in the example above returns another function, which we
define anonymously in the body of `intSeq`.
The returned function **closes over** the variable `i` to form a closure.

``` go
nextInt := intSeq() 
fmt.Println(intSeq()) // 1
fmt.Println(intSeq()) // 2
fmt.Println(intSeq()) // 3
newInts := intSeq()
fmt.Println(intSeq()) // 1
```

Assigning the result of `intSeq()`(function) to `nextInt` captures its own `i`
value, which will be updated each time `nextInt()` is invoked.
