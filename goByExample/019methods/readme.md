# Methods

Go supports **methods** declared on **struct** types.

Methods for structs can be defined for either pointer or value receiver types.

Go automatically handles conversion between values and pointers for method
calls.
IMPORTANT: You may want to use a pointer receiver type to avoid copying on
method calls or to allow the method to mutate the receiving struct.
