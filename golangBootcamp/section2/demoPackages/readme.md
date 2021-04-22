# Readme

The hi.go and bye.go files contain function declarations and definitions.

Because the package clause is main for all 3, when we compile the directory all
3 are compiled.

``` go
// using wildcard
go build ./ *.go // remove space between / and *
// or
go build main.go hi.go bye.go
// problem with first approach is that go names the executable bye instead of
// main; typing out all the file names seems to work correctly
```
