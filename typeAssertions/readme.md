# Type assertions

For an expression x of interface type and a type T, the primary expression
`x.(T)` asserts that x is not nil and that the value stored in x is of type T.
The notation x.(T) is called a type assertion

If T is not an interface type, x.(T) asserts that the dynamic type of x is
identical to the type T. In this case, T must implement the (interface) type of
x; otherwise type assertion is invalid since it is not possible for x to store
a value of type T. If T is an interface type, x.(T) asserts that the dynamic
type of x implements the interface T.

If the assertion holds, the value of the expression is the value stored in x
and its type is T. If the assertion is false, a run-time panic occurs. In other
words, even though the dynamic type of x is known only at runtime, the type of
x.(T) is known to eb T in a correct program.

``` go
var x = interface{}=7   // x has dynamic type int and value 7
i := x.(int)            // i has type int and value 7

type I interface { m() }

func f(y I) {
  // so will panic occur in first example?
  s := y.(string)       // illegal: string does not implement I (missing method m)
  r := y.(io.Reader)    // r has type io.Reader and the dynamic type of y must
                        // implement both I and io.Reader
  // ...
}
```

``` go
// A type assertion used in an assignment or initialization of form:
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok interface{} = x.(T) // dynamic types of v and ok are T and  bool
// yields an additional untyped boolean value. The value of ok is true if
// assertion holds; otherwise false and value of v is zero type for T.
// no run-time panic occurs in this case.
```
