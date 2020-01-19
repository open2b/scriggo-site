---
layout: article
---
# Install Scriggo

Before installing the Scriggo command, <a href="https://golang.org/dl/">download and install Go</a>.

When Go is installed, open a terminal and execute:

```
$ go install github.com/open2b/scriggo/cmd/scriggo
```

`go` will install the `scriggo` command in the `bin` subdirectory of the Go path.

To know what is the Go path, execute the command:

```
$ go env GOPATH
```

For convenience, add the Go path's `bin` subdirectory to your `PATH`:

```
$ export PATH=$PATH:$(go env GOPATH)/bin
```

Then test if `scriggo` can be executed:

```
$ scriggo version
```
