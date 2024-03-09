# Framed TCP

A Go package for framing messages in a TCP stream.

## The interface
We expose the interface with `Receive()` and `Send()` for abstracting away sending and receiving framed messages from wrapped TCP connection.

## Implementations
Three concrete implementations are provided:
* Fixed length messages
* Messages with 4 byte big-endian integer prefix for length
* Messages delimeted by a byte sequence (like `\r\n`)
