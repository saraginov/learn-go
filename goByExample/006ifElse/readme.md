# If-Else

Branching with `if` and `else` in Go is straight-forward.

You can have an `if` statement without an `else`

A statement can precede conditionals; any variables declared in this statement are available in all branches.

``` go
for var x:=0, y:=3 {
  x = 3
  y = 10
  // i.e. variables declared in this statement are available in all branches
}
```

Parentheses `()` are not needed around conditions but code block must be enclosed by `{}`

There is no `ternary if` in Go, therefore you must use a full `if` statement even for basic conditions
