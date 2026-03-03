{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo command{% end %}
{% Article %}

# Scriggo command

Scriggo has a command line interface, the `scriggo` command, that allows to:

- [Run a template file](#run-a-template-file)
  - [Complete syntax](#complete-syntax)
  - [Examples](#examples)
- [Serve a template](#serve-a-template)
  - [Complete syntax](#complete-syntax-1)
- [Build a template](#build-a-template)
  - [Complete syntax](#complete-syntax-2)
  - [Examples](#examples-1)
- [Initialize an interpreter](#initialize-an-interpreter)
- [Generate a package importer](#generate-a-package-importer)
  - [Complete syntax](#complete-syntax-3)

### Get the `scriggo` command

You can get the binary of the scriggo command from the [releases](https://github.com/open2b/scriggo/releases/) page.

Alternatively, you can install the command with the `go install` command as explained below.

Before installing the Scriggo command, <a href="https://go.dev/dl/">download and install Go</a>.

When Go is installed, open a terminal and execute:

```shell
$ go install github.com/open2b/scriggo/cmd/scriggo@latest
```

then test if `scriggo` can be executed:

```shell
$ scriggo version
scriggo version v0.61.0 (go1.24)
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

The `-o` flag writes the result to the named output file or directory, instead of the standard output.

Markdown is converted to HTML with the Goldmark parser with the options `html.WithUnsafe`, `parser.WithAutoHeadingID`,
`extension.GFM`, and `extension.Footnote`.

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
`html.WithUnsafe`, `parser.WithAutoHeadingID`, `extension.GFM`, and `extension.Footnote`.

When a template file changes, the templates are automatically rebuilt, and with LiveReload, the page
in the browser is automatically reloaded.

To disable LiveReload, use the `--disable-livereload` flag:

```shell
$ scriggo serve --disable-livereload
```

### Complete syntax

The complete `scriggo serve` command takes this form:

```shell
$ scriggo serve [-S n] [--metrics] [--disable-livereload]
```

The `-S` flag prints the assembly code of the served file and n determines the maximum length, in runes, of
disassembled `Text` instructions

    n > 0: at most n runes; leading and trailing white space are removed
    n == 0: no text
    n < 0: all text

The `--metrics` flags prints metrics about execution time.

The `--disable-livereload` flag disables LiveReload, preventing automatic page
reloads in the browser.

## Build a template

The Scriggo Build command processes the template rooted at the current directory and writes the generated files to the
`public` directory by default. If the `public` directory already exists, it does nothing and returns an error.

Only files with extension `.html` and `.md` are processed as templates; other files, such as CSS and JavaScript files,
are copied as-is. Directories whose names start with an underscore (`_`), and files or directories whose names start
with a dot (`.`), are skipped but can still be referenced in template files.

The basic Build command takes this form:

```shell
$ scriggo build [dir]
```

If a directory `dir` is specified, the template rooted at that directory is built instead of the current directory.

For example:

```shell
$ scriggo build src
```

processes the template rooted at the `src` directory and writes the generated files to the `public` directory. HTML
and Markdown files are processed as templates; all other files are copied unchanged, resulting in a complete static
site ready for deployment.

The `-o` flag specifies an alternative output directory instead of the default `public`.

The `-const` flag builds the template with a global constant with the given name and value. `name` should be a Go
identifier and `value` should be a string literal, a number literal, `true` or `false`. There can be multiple
`name=value` pairs.

The `-llms` flag generates two outputs for each Markdown template file that extends an HTML layout: an HTML version
and a Markdown version. Both share the same path but use different file extensions. The Markdown output is intended
for consumption by LLMs.

For example, if the file `page.md` extends `layout.html`, the build process generates `page.html` by extending
`layout.html`, and also `page.md` by extending `layout.md`. The file `layout.md` must exist in the same directory
as `layout.html`.

The `-llms` flag requires a base URL as its argument, used to rewrite relative link destinations in the generated
Markdown files by prefixing them with the provided base URL. Link destinations that are absolute, or consist only
of a query string or a fragment, are left unchanged.

### Complete syntax

```shell
$ scriggo build [-llms url] [-const name=value] [-o output] [dir]
```

### Examples

```shell
$ scriggo build src
```

```shell
$ scriggo build -o dist src
```

```shell
$ scriggo build -llms https://docs.example.com/ src
```

```shell
$ scriggo build -const 'version=1.12 title="The ancient art of tea"'
```

```shell
$ scriggo build -o ../public
```

```shell
$ scriggo build -o /var/www site
```

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

The Scriggo Import command generates the code for a package importer. An importer is used by Scriggo to import a package
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

The importer in the generated Go file has type `native.Importer` and it is assigned to a variable named "packages".
The variable can be used as an argument to the `Build` and `BuildTemplate` functions in the `scriggo` package.

To give a different name to the variable use the instruction `SET VARIABLE` in the Scriggofile:

    SET VARIABLE foo

The package name in the generated Go file is by default "main", to give a different name to the package use the
instruction `SET PACKAGE` in the Scriggofile:

    SET PACKAGE boo

The `-v` flag prints the imported packages as defined in the Scriggofile.

The `-x` flag prints the executed commands.

The `-o` flag writes the generated Go file to the named output file, instead of the standard output.

