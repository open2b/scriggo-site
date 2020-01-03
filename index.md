# Scriggo

## What is Scriggo?

Scriggo is a fast embedded Go language interpreter.

## Supported syntaxes

Scriggo can support three different kinds of syntaxes.

The first (and the best known) is the **program syntax**, which is exactly the same syntax used by Go.

```go
package main

import "fmt"

func main() {
    fmt.Println("I'm Scriggo!")
}
```

a program can also be compiled and executed with the _package-less_ option,
which is useful in every context where it's important to avoid boilerplate:

```go
// This is a full and valid package-less Scriggo program!
fmt.Println("I'm Scriggo!")
```

the third syntax is the **template syntax**, that looks like that:

{% raw %}
```go
{% name := "Scriggo" %}
{% for i := 0; i < 20 ; i++ %}
    Welcome to the {{ name }} template!
{% end for %}
```
{% endraw %}

## How fast is Scriggo?

The execution of Scriggo ... TODO

See the [benchmarks]() for a comparison between Scriggo and other technologies.

# Documentation

Here you can find some useful resources that cover from the first steps with Scriggo as user to the most advanced details on the Scriggo internal implementation.

## Write Scriggo templates

For designers who want to work with the Scriggo template.

- [The Scriggo template](/doc/users/template.md) covers the usage of the Scriggo template starting from the simplest pages.

## Write Scriggo programs

For users who use Scriggo in an embedded context or through an interpreter.

- [The package-less option](/doc/users/package-less-option.html) covers the differences in the program syntax when the _package-less_option_ is on.

## Embed Scriggo in your application

If you want to develope an application that uses Scriggo, this is the right section for you.
It covers the Scriggo APIs and the `scriggo` command line tool.

- [Get started embedding Scriggo](/doc/users/get-started-embedding.html)
- [godoc.org/scriggo]() the Scriggo APIs documentation, hosted on godoc.org.
- [The scriggo command line tool]() introduces and describes the use of the command `scriggo`.

## Internals of Scriggo

For developers who want to develope Scriggo.
Covers the Scriggo internals, as well as the repository structure.

- [Scriggo internals](doc/developers/internals.md) makes an overview of the internals of Scriggo.
- [The Scriggo Virtual Machine](/doc/developers/vm.md)

