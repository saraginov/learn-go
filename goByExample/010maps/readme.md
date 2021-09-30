# Maps

**Maps** are Go's built-in associative data type(sometimes called hashes or dicts
in other languages)

To create an empty map, use the built-in `make()`, `make(map[key-type]val-type)`

Set key/value pairs using typical `name[key] = val` syntax

Printing a map with `fmt.Println()` will show all of its key value pairs

Get a value for a key with `name[key]`.

The built-in `len()` returns the number of key/value pairs when called on a map
The built-in `delete()` removes key value pairs from a map e.g. `delete(mapName, "keyValue")`
