# Channels

**Channels** are the pipes that connect concurrent goroutines. Values can be
sent into channels from one goroutine and receive those values into another
goroutine.

Create a new channel with `channelName := make(chan val-type)`. Channels are
typed by the values they convey.

Send a value into a channel using the channel `channelName <- data`

By default sends and receives block until both the sender and the receiver are
ready. This property allows us to wait at the end of our program for the "ping"
message without having to use any other synchronization
