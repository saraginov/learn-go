# What is GO PATH

``` go
// to list GO PATH
go env GOPATH
// OR
go env
// to list all GO environment variables
```

usage: `go env [-json] [-u] [-w] [var ...]`.
Env prints Go environment information.

By default env prints information as a shell script
(on Windows, a batch file).
The -json flag prints the environment in JSON format
instead of as a shell script.

The -u flag requires one or more arguments and unsets
the default setting for the named environment variables,
if one has been set with 'go env -w'.

The -w flag requires one or more arguments of the
form NAME=VALUE and changes the default settings
of the named environment variables to the given values.
