# Constants

Go supports **constants** of character, string, boolean, and numeric values.

`const` declares a constant value/variable.
A `const` statement can appear anywhere a `var` can.

Constant expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until it's given one, such as by an explicit conversion.

A number can be given a type by using it in a context that requires one, such as
a variable assignment or function call.
e.g. `math.Sin()` expects a `float64`
