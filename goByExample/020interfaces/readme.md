# Interfaces

**Interface**(s) are named collections of method signatures.

``` go
type geometry interface{
  area() float64
  perim() float64
}
```

To implement an interface in Go, we just need to implement all of its methods.

If a variable has an interface type, then we can call methods that are
in the named interface.

## <https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go>

An interface is a set of methods but also a type.

A core concept in Go’s type system;
instead of designing our abstractions in terms of what kind of data our types
can hold,
we design our abstractions in terms of what actions our types can execute.

Any type which implements all methods for a specific interface is said to
**satisfy** the interface.
There is no implements keyword in Go;
whether or not a type satisfies an interface is determined automatically.

We can create a slice for the interface type, and all struct types which
satisfy the interface can be added/included in said slice.

``` go
type Animal interface {
    Speak() string
}
type Dog struct {}
func (d Dog) Speak() string {
    return "Woof!"
}
type Cat struct {}
func (c Cat) Speak() string {
    return "Meow!"
}
type Llama struct {}
func (l Llama) Speak() string {
    return "?????"
}
type JavaProgrammer struct {}
func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}
func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

The interface{} type, the empty interface, is the interface that has no methods.
Since there is no implements keyword, all types implement at least zero methods,
and satisfying an interface is done automatically, all types satisfy the empty
interface.
That means that a function that takes an interface{} value as a parameter,
that function can be supplied that function with any value. So, this function:
`func DoSomething(v interface{}) {}` will accept any parameter whatsoever.
v is not of any type; it is of interface{} type.
When passing a value into the DoSomething function, the Go runtime will perform
a type conversion (if necessary), and convert the value to an interface{} value.
All values have exactly one type at runtime

## <https://research.swtch.com/interfaces>
