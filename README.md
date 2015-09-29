`$GOPATH` Janitor
==============

`gpj` is a simple, no-frills way to check your `$GOPATH` for unused libraries.

To install, simply `go get github.com/zmb3/gpj`.

### Example

In this example, we start with a fresh `$GOPATH` and `go get` a command with a
couple of dependencies.  We run `gpj` and don't get any output - this indicates
that we don't have any unused packages.  We finish by removing the source for
the command but leaving its dependencies behind.  `gpj` happily identifies the
unused libraries.

[![Example](https://asciinema.org/a/84hjfs1p4q4xxm7fbzxn2ot1v.png)](https://asciinema.org/a/84hjfs1p4q4xxm7fbzxn2ot1v)
