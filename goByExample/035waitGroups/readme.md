# Wait Groups

WaitGroups can be used to wait for multiple goroutines to finish.

to declare a WaitGroup `var wg sync.WaitGroup`, where wg is variable

If a WaitGroup is explicitly passed into functions, it should be done by
pointer

todo [x] - what are concurrency primitives?
<https://medium.com/@anandpillai/synchronization-in-go-using-concurrency-primitives-a-case-study-535bb2a71c13>

## <https://golang.org/doc/faq#closures_and_goroutines>

Closures as goroutines

``` go
func main() {
  done := make(chan bool)
  values := []string{"a", "b", "c"}
  for _, v := range values {
    go func() {
      fmt.Println(v)
      done <- true
    }()
  }
  // wait for all goroutines to complete before exiting
  for _ = range values {
    <-done
  }
}
```

One might mistakenly expect to see a, b, c ass the output. What one will
probably see instead is c, c, c. This is because each iteration of the loop
uses the same instance of the variable v, so each closure shares that single
variable. When the closure runs, iit prints the value of v at the time
`fmt.Println()` is executed, but v may have been modified since the goroutine
was launched.

To help detect this and other problems before they happen, run `go vet`

``` go
// to bind the current value of v to each closure as it is launched, one must
// modify the inner loop to create a new variable each iteration. One way is to
// pass the variable as an arg to closure
for _, v := range values {
  go func(u string) {
    fmt.Println(u)
    done <- true
  }(v)
}
// or create new var
for _, v := range values {
  v := v // create a new v within local scope
  go func() {
    fmt.Println(v)
    done <- true
  }()
}
```
