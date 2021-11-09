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
    // we do not have to do something like Animal(Dog{}) in order to put a val
    // of type Dog into slice of Animal values because conversion is handled
    // automatically. Within the animals slice, each element is of Animal type,
    // but the different values have different underlying types
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
That means that a function that takes an interface{} value as a parameter, can
be supplied with any value(of any type),i.e.`func DoSomething(v interface{}) {}`
will accept any parameter whatsoever.
v is not of *any* type; v is of interface{} type; i.e. when we pass a variable
to the function the *Go Runtime* will perform a type conversion (if necessary),
and convert the value to an interface{} value.

Value gets converted because all values have exactly one type at runtime

An interface value is constructed of two words of data; one word is used to
point to a method table for the value's underlying type, and the other word
is used to point to the actual data being held by that value.
Note from self: I understand the second word to be a pointer to the data in
memory

## <https://research.swtch.com/interfaces>

Interfaces allow for duck typing, passing different type of variables like a
purely dynamic language like Python.

``` go
// define interface
type ReadCloser interface {
    Read(b []byte) (n int, err os.Error)
    Close()
}
// define function as taking ReadCloser, this function calls Read repeatedly to
// get all data that was requested and then calls Close
func ReadAndClose(r ReadCloser, buf []byte) (n int, err os.Error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    r.Close()
    return
}
// the code that calls ReadAndClose can pass a value of any type as long as it
// has Read and Close methods with the correct signatures. if wrong type value
// is passed, error thrown at compile time
```

Interfaces aren't restricted to static type checking. Checking whether an
interface has an additional method can be checked dynamically:

``` go
import "strconv"

type Stringer interface {
    String() string
}
// The value any has static type interface{}, meaning no guarantee of any
// methods at all; it could contain any type.
func ToString(any interface{}) string {
    // the "comma ok" assignment asks whether it is possible to convert `any` to
    // an interface value of type Stringer, which has the method String
    // if ok is set to True, the body of the statement calls the method to
    // obtain a string to return
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    // switch picks of a few basic types before giving up
    // stripped down version of what fmt package does
    // if above could be replaced by adding case Stringer: at the top of the
    // switch
    switch v := any.(type) {
        case int:
            return strconv.Itoa(v)
        case float:
            return strconv.Ftoa(v, 'g', -1)
    }
    return "???"
}
// a value of type Binary can be passed ToString, which will format it using the
// String method, even though the program never says that Binary intends to
// implement Stringer. There's no need: the runtime can see that Binary has a
// String method, so it implements Stringer, even if the author of Binary
// has never heard of Stringer
type Binary uint64
func (i Binary) String() string {
    return strconv.Uitob64(i.Get(),2)
}
func (i Binary) Get() uint64 {
    return uint64(i)
}
```

Implicit conversion are checked at compile time, explicit interface-to-interface
conversions can inquire about method sets at run time.

Interface values are represented as a two-word pair giving a pointer to
information about the type stored in the interface and a pointer to the
associated data
Assigning b to an interface value of type Stringer sets both words of the
interface value

## convert []T to []interface{}

```go
package main
import "fmt"
func PrintAll(vals [] interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}
func main() {
    names := []string{"stanley", "david", "oscar"}
    // PrintAll below yields error: cannot use names (type []string) as type
    // []interface {}
    PrintAll(names)
    // instead names need to be declared converted to []interfaces{}
    // this doesn’t come up very often, because []interface{} turns out to be
    // less useful than initially expected
    vals := make([]interfaces{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```

## Pointers and interfaces

Another subtlety of interfaces is that an interface definition does not
prescribe whether an implementor should implement the interface using a pointer
receiver or a value receiver

``` go
type Animal interface {
    Speak() string
}
type Dog struct {}
func (d Dog) Speak() string {
    return "Woof!"
}
type Cat struct {}
func (c *Cat) Speak() string {
    return "Meow!"
}
func main() {
    animals := []Animal{Dog{}, new(Cat)}
    // If you run the code below with animals := []Animal{Dog{}, Cat{}}
    // cannot use Cat literal (type Cat) as type Animal in array element... will
    // be thrown
    // it's not saying that the interface Animal demands that you define your
    // method as a pointer receiver
    // it's saying that we attempted to convert a Cat struct into an animal
    // interface, but only *Cat satisfies the interface
    // Can be fixed by passing a *Cat pointer to the Animal slice instead of
    // a Cat value, by using new(Cat) instead of Cat{}
    // new(Cat) is the same as &Cat{} i.e. address of Cat{}, I disagree with
    // Jordan, I prefer &Cat{}, but maybe this will change over time
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }

    animals = []Animal{&Dog{}, new(Cat)}
    // this also works, we don't need to change the type of receiver of Speak
    // method for Dog!
    // That is, a *Dog value can utilize the Speak method defined on Dog, but
    // a Cat value cannot access the Speak method defined on *Cat.
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

**A pointer type can access the methods of its associated value, but not vice**
**versa.**

Everything in Go is passed by value. Every time a func is invoked, data passed
into it is copied. In the case of a method with a value receiver, the value is
copied when calling the method.

`func (t T)MyMethod(s string) {...}` is a function of type `func(T, string)`;
method receivers are passed into the function by value just like any other
parameter
Not from self: I don't understand why the above is true, looks like an anon func
