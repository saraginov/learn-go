# Range over channels

We have seen how range provides iteration over basic data structures.
We can also use `range` syntax to iterate over values received from a channel.

It is possible to close a non-empty channel and still have the remaining values
be received.
