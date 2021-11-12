# Errors

In Go it is idiomatic to communicate errors via an explicit, separate return
value. This contrasts with the exceptions used in other languages.
Makes it easy to see which function returns an error and to handle them using
the same constructs employed for other non-error tasks.

**By convention, errors are the last return value and have type error,**
**a built-in interface**

`errors.New("message")` constructs a basic error value with the given error
message

**A `nil` value in the error position indicates that there was no error!**

It is possible to use custom types as *errors* by implementing the `Error()`
method on them.

``` go
// custom error
type argError struct {
  arg  int
  prob string
}
// use custom type to explicitly represent an argument error
func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
```

## <https://go.dev/blog/error-handling-and-go>

Go uses **error** values to indicate an abnormal state
e.g. `os.Open()` returns a non-nill error when it fails to open a file.
func signature: `func Open(name string) (file *File, err error)`

``` go
// open file with os.Open(), call log if an error occurs
f, err := os.Open("filename.ext")
if err != nil {
  log.Fatal(err)
}
// do something with the open *File f
```

The **error** type is an interface. An error represents any value that can
describe itself as a string, like all built in types is predeclared in the
universe block.
The universe block encompasses all Go source text.
<https://golang.org/ref/spec#Blocks>

```go
type error interface {
  Error() string
}
```

most commonly used error implementation is the errors package's un-exported
**errorString** type.

``` go
// errorString is a trivial implementation of an error
type errorString struct {
  s string
}
func (e *errorString) Error() string {
  return e.s
}
// New returns an error that formats as given text
func New(text string) error {
  return &errorString{text}
}
// e.g. of errors.New()
func Sqrt(f float64) (float64, error) {
  if f < 0 {
    return 0, errors.New("math: square root of a negative number")
  }
  // implementation
}
// caller can access error string by calling error's Error() method or by
// printing it
f, err := Sqrt(-1)
if err != nil {
  // fmt package formats an error value by calling its Error() string method
  fmt.Println(err)
}
```

It is the error implementation's responsibility to summarize the context.
The error returned by `os.Open()` formats as
"open /etc/passwd: permission denied", not just "permission denied".
`fmt.Errorf()` is useful to summarize context. It formats a string according to
`Printf()`'s rules and returns it as an error created by `errors.New()`

``` go
func Sqrt(f float64) (float64, error) {
  if f < 0 {
    return 0, fmt.Errorf("math: square root of a negative number %g", f)
  }
  // implementation
}
```

In many cases `fmt.Errorf()` is good enough, but since *error* is an interface,
arbitrary data structures can be used as error values in order to allow callers
to inspect the details of the error.

``` go
// recovering invalid argument
type NegativeSqrtError float64
func (f NegativeSqrtError) Error() string {
  return fmt.Sprintf("math: square root of a negative number %g", float64(f))
}
// a sophisticated caller can then use a type assertion to check for a
// NegativeSqrtError and handle it, while callers that just pass the error to
// fmt.Println or log.Fatal will see no change in behaviour
func Sqrt(f float64) (float64, error) {
  if f < 0 {
    return 0, &NegativeSqrtError{f}
  }
  // implementation, and I don't know whether we would return 0 or maybe some
  // other number
}
```

the json package specifies a `SyntaxError` type which the json.Decode function
returns when it encounters a syntax error parsing a JSON blob

``` go
type SyntaxError struct {
  msg     string  // description of error
  Offset  int64   // error occurred after reading offset bytes
}
func (e *SyntaxError) Error() string { return e.msg }
// offset field isn't even shown in default error formatting
// but callers can use it to add file and line information to their error msgs
if err := dec.Decode(&val); err != nil {
  if serr, ok := err.(*json.SyntaxError); ok {
    line, col := findLine(f, serr.Offset)
    return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
  }
  return err
}
// simplified version of https://perkeep.org/
```

**error** interface requires only a `Error() string` method; specific
implementations might have additional methods
e.g. the [net](https://pkg.go.dev/net) package returns errors of type **error**,
some error implementations have additional methods defined by the **net.Error**
interface

``` go
type Error interface {
  error
  Timeout() bool // Is the error a timeout?
  Temporary() bool // Is error temp?
}
// client code can test for net.Error with type assertion and then distinguish
// transient network errors from permanent ones. e.g. web crawler might sleep
// and retry or give up
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
  time.Sleep(1e9)
}
if err != nil {
  log.Fatal(err)
}
```

### Simplifying repetitive error handling

In Go error handling is important. The language's design and conventions
encourage explicit error checking when error occurs(distinct from conventions in
other langs. of throwing exceptions and sometimes catching them).
This makes Go verbose, there are techniques to minimize repetitive error
handling

``` go
// consider an App Engine application with an HTTP handler that retrieves a
// records from db and formats it with template
// what is the App Engine? <https://cloud.google.com/appengine/docs/go/>
func init() {
  http.HandleFunc("/view", viewRecord)
}
// view record is the http handler
func viewRecord(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  // I don't know what appengine.NewContext does, and I have no intentions of
  // finding out at time of writing
  key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
  record = new(Record)
  if err := datastore.Get(c, key, record); err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if err := viewTemplate.Execute(w, record); err != nil {
    http.Error(w, err.Error(), 500)
  }
}
```
