# Pointers

Go supports **pointers**, allowing values of variables and records to be passed
within the program by reference.

``` go
// zeroval has an int parameter, so arguments will be passed to it by value
// i.e. will get a copy of n distinct from the one in the calling function
func zeroval(n int) {
  n = 0
}
// zeroptr in contrast has an *int parameter, meaning that it takes an int
// pointer, i.e. the address of the of the value from the n variable in the
// calling function, also known as passing by reference
// the *n code in the function body **dereferences** the pointer from its memory
// address to the current value at that address
func zeroptr(n *int) {
  *n = 0
}
```

Assigning a value to a **dereferenced** pointer changes the value at the
referenced address.
The syntax `&varName` gives the memory address of `varName`, i.e. the pointer
to `varName`
