{% extends "/layouts/article.html" %}
{% macro Title %}Get Started{% end %}
{% Article %}

Get Started
# How to use the `scriggo` command to create an interpreter

In this second part of get started we will learn how create an interpreter with the `scriggo` command. An interpreter is a standalone program that can execute Go programs, scripts and templates without requiring Go installed.

It required 10 minutes to be completed.

## Install the `scriggo` command

Open a terminal and execute:

```
$ go install github.com/open2b/scriggo/cmd/scriggo
```

Go will install the `scriggo` command in the `bin` subdirectory of the Go path.

To know what is the Go path, execute the command:

```
$ go env GOPATH
```

For convenience, add the Go path's `bin` subdirectory to your `PATH`:

```
$ export PATH=$PATH:$(go env GOPATH)/bin
```

Test if the `scriggo` command can be executed:

```
$ scriggo version
```

## Scriggofile
