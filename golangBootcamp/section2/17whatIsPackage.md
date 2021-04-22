# What is a package?

A group of go source files that belong to the same directory is considered a package.

Only source files belonging to the main package can be executed.

Use other package names to create libraries.

``` go
package main
// the package main is called a package clause
// each file may only have 1 package clause

```

Library packages can be imported but not executed, where as an executable package can be executed but not imported.

Library Packages can have any name.
