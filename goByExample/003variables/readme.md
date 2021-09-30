# Variables

In Go, **variables** are explicitly declared and used by the compiler; e.g. to check type-correctness of function calls.

`var` declares 1 or more variables, i.e. multiple variables can be declared at once

Go will infer the type of initialized variables.
Variables declared without a corresponding initialization are **zero-valued**.
For example, the zero value for an `int` is `0`.

The `:=` syntax is shorthand for declaring and initializing a variable,
e.g. `f := "apple"` for `var f string = "apple"`

---

Poor explanation from <gobyexample.com>, reason being `=` is assignment operator
where as `:=` declaration and assignment, meaning we can't do the following

``` go
f:=1
f:=3
// but we can do
f:=1
f=3
```
