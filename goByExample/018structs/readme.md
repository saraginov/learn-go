# Structs

Go's **structs** are typed collections of fields.
They are useful for grouping data together to form records.

``` go
// declare struct
type person struct {
  name string
  age int
}

// create new struct
newPerson := person{"James", 30}
otherNewPerson := person{name: "James", age: 30}
noAge := person{name: "James"} // age will default to int 0
```

It’s idiomatic to encapsulate new struct creation in constructor functions.
