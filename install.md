{% extends "/layouts/article.html" %}
{% macro Title %}Install Scriggo{% end %}
{% Article %}

# Install `scriggo` command

Before installing the Scriggo command, <a href="https://golang.org/dl/">download and install Go</a>.

When Go is installed, open a terminal and execute:

```
$ go install github.com/open2b/scriggo/cmd/scriggo@latest
```

then test if `scriggo` can be executed:

```
$ scriggo version
scriggo version 0.48.0 (go1.17)
```

If the `scriggo` command is not found, you should add the directory where the command has been installed to your `PATH`.

`go` installs the `scriggo` command in the directory named by the `GOBIN` environment variable. If it is not set it
defaults to `$GOPATH/bin` or, if the `GOPATH` environment variable is not set, to `$HOME/go/bin`.
