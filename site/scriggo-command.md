{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo command{% end %}
{% Article %}

# Scriggo command

Scriggo has a command line interface, the `scriggo` command, that allows to:

* [run a template file](#run-a-template-file)
* [serve templates with support for Markdown](#serve-a-template)
* [initialize an interpreter for Go programs](#initialize-an-interpreter)
* [generate the code for package importers](#generate-a-package-importer)

### Get the `scriggo` command

You can get the binary of the scriggo command from the [releases](https://github.com/open2b/scriggo/releases/) page.

Alternatively, you can install the command with the `go install` command as explained below.

Before installing the Scriggo command, <a href="https://golang.org/dl/">download and install Go</a>.

When Go is installed, open a terminal and execute:

```shell
$ go install github.com/open2b/scriggo/cmd/scriggo@latest
```

then test if `scriggo` can be executed:

```shell
$ scriggo version
scriggo version v0.53.2 (go1.17)
```

If the `scriggo` command is not found, you should add the directory where the command has been installed to your `PATH`.

`go` installs the `scriggo` command in the directory named by the `GOBIN` environment variable. If it is not set it
defaults to `$GOPATH/bin` or, if the `GOPATH` environment variable is not set, to `$HOME/go/bin`.


### Get help from command line

To get help from the command line run the following command:

```shell
$ scriggo help
```

## Run a template file

The Scriggo Run command runs a template file and its extended, imported and rendered files. All Scriggo builtins are
available in the template file.

The basic Run command takes this form:

```shell
$ scriggo run [-o output] file
```

For example:

```shell
scriggo run article.html
```

runs the file "article.html" as HTML and prints the result to the standard output. Extended, imported and rendered file
paths are relative to the directory of the executed file.

The `-o` flag writes the result to the named output file or directory, instead to the standard output.

Markdown is converted to HTML with the Goldmark parser with the options `html.WithUnsafe`, `parser.WithAutoHeadingID`
and `extension.GFM`.

### Complete syntax

The complete `scriggo run` command takes this form:

```shell
$ scriggo run [-o output] [-root dir] [-const name=value] [-format format] [-metrics] [-S n] file 
```

The `-o` flag writes the resulting code to the named output file or directory.

The `-root` flag sets the root directory to named directory instead of the file's directory.

The `-const` flag runs the template file with a global constant with the given name and value. name should be a Go
identifier and value should be a string literal, a number literal, true or false. There can be multiple `name=value`
pairs.

The `-format` flag forces render to use the named file format.

The `-metrics` flag prints metrics about execution time.

The `-S` flag prints the assembly code of the executed file and n determines the maximum length, in runes, of
disassembled `Text` instructions

    n > 0: at most n runes; leading and trailing white space are removed
    n == 0: no text
    n < 0: all text

### Examples

```shell
$ scriggo run index.html
```

```shell
$ scriggo run -const 'version=1.12 title="The ancient art of tea"' index.md
```

```shell
$ scriggo run -root . docs/article.html
```

```shell
$ scriggo run -format Markdown index
```

```shell
$ scriggo run -o ./public ./sources/index.html
```

## Serve a template

The Scriggo Serve command runs a web server and serves the template rooted at the current directory.
All Scriggo builtins are available in template files. It is useful to learn [Scriggo templates](/templates).

The basic Serve command takes this form:

```shell
$ scriggo serve
```

It renders HTML and Markdown files based on file extension.

For example, serving this request:

    http://localhost:8080/article

it renders the file "article.html" as HTML if exists, otherwise renders the file "article.md" as Markdown.

Serving a URL terminating with a slash:

    http://localhost:8080/blog/

it renders "blog/index.html" or "blog/index.md". 

Markdown is converted to HTML with the [Goldmark](https://github.com/yuin/goldmark) parser with the options
`html.WithUnsafe`, `parser.WithAutoHeadingID` and `extension.GFM`.

Templates are automatically rebuilt when a file changes.

### Complete syntax

The complete `scriggo serve` command takes this form:

```shell
$ scriggo serve [-S n] [--metrics] 
```

The `-S` flag prints the assembly code of the served file and n determines the maximum length, in runes, of
disassembled `Text` instructions

    n > 0: at most n runes; leading and trailing white space are removed
    n == 0: no text
    n < 0: all text

The `--metrics` flags prints metrics about execution time.

## Initialize an interpreter

The Scriggo Init command initializes an interpreter for Go programs. It creates the following files:

* go.mod
* go.sum
* main.go
* packages.go
* Scriggofile

The syntax is: 

```shell
$ scriggo init [dir]
```

where `dir` is the directory in which to create the files. If no argument is given, `scriggo init` uses the current directory.
If the directory already contains ".go" files or a "vendor" directory, the command fails. 

The command creates a Scriggofile with the instruction to create an importer for the Go standard library.

If the directory does not contain a `go.mod` file, the command creates it and as module path uses the directory name.

## Generate a package importer

The Scriggo Import command generate the code for a package importer. An importer is used by Scriggo to import a package
when an "import" declaration is executed.

The code for the importer is generated from the instructions in a Scriggofile. The Scriggofile should be in a Go module.

The basic Import command takes this form:

```shell
$ scriggo import [-o output]
```

For example:

```shell
$ scriggo import -o packages.go
```

generates the code for an importer, with instructions in a Scriggofile called "Scriggofile" in the current directory
and writes it into the file "packages.go".

### Complete syntax

The complete `scriggo import` command takes this form:

```shell
$ scriggo import [-f Scriggofile] [-v] [-x] [-o output] [module]
```

If an argument is given, it must be a local rooted path or must begin with a `.` or `..` element and it must be a module
root directory. Import looks for a Scriggofile named "Scriggofile" in that directory.

If no argument is given, the action applies to the current directory.

The `-f` flag forces import to read the given Scriggofile instead of the Scriggofile of the module.

The importer in the generated Go file have type `native.Importer` and it is assigned to a variable named "packages".
The variable can be used as an argument to the `Build` and `BuildTemplate` functions in the `scriggo` package.

To give a different name to the variable use the instruction `SET VARIABLE` in the Scriggofile:

    SET VARIABLE foo

The package name in the generated Go file is by default "main", to give a different name to the package use the
instruction `SET PACKAGE` in the Scriggofile:

    SET PACKAGE boo

The `-v` flag prints the imported packages as defined in the Scriggofile.

The `-x` flag prints the executed commands.

The `-o` flag writes the generated Go file to the named output file, instead to the standard output.

