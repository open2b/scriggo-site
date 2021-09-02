{% extends "/layouts/article.html" %}
{% macro Title %}Scriggo command{% end %}
{% Article %}

# Scriggo command

Scriggo has a command line interface, the `scriggo` command, that allows to:

* serve Scriggo templates with support for Markdown
* initialize an interpreter for Go programs
* generate a package importer for an existing application

Se how to <a href="/install">install the scriggo command</a>.

### Get help from command line

To get help from the command line run the following command:

```
$ scriggo help
```

## Initialize an interpreter

The Scriggo command allows to initialize an interpreter for Go programs. The syntax is: 

```
$ scriggo init [dir]
```

where `dir` is a (generally empty) directory. If no argument is given, `scriggo init` uses the current directory.

After 

If the name of the file could be interpreted as a `scriggo` subcommand, use `--` as in:

```
$ scriggo -- file [arguments...]
```

Currently, the `scriggo` command can run only a single package file.

## Build Scriggo commands

The scriggo command includes the Go standard library. If you want to include other packages, you can build a custom scriggo command. The scriggo command uses the instructions in a [Scriggofile](scriggofile) to create a custom scriggo command or generate a package importer.

The `scriggo build` command builds a custom scriggo command from a [Scriggofile](scriggofile) in a Go module.

The basic build command takes this form:

```
$ scriggo build [module]
```

Executables are created in the current directory. To build and install the executables in
the `bin` subdirectory of the Go path, use instead the install command:
 
```
$ scriggo install [module]
```

If the argument `module` is given, it can be a module path or a directory path:

* If the argument is a module path, the module is downloaded from its repository
and the build command looks for a [Scriggofile](scriggofile) named `Scriggofile` in its root.
A module argument can have a version as in `foo.boo@v2.1.0`. If no version is
given the latest version of the module is used.

* If the argument is a directory path, it must be rooted or must begin with
a `.` or `..` element and the directory must be the root of a module. The build
command looks for a [Scriggofile](scriggofile) named `Scriggofile` in that directory.

* If no argument is given, the action applies to the current directory.

The name of the executable is the last element of the module's path or
directory path. For example if the module's path is `boo/foo` the name of the
executable will be `foo` or `foo.exe`.

The interpreter is build from the instructions in the [Scriggofile](scriggofile). For example:

```
$ scriggo build github.com/example/foo
```

will build an interpreter named `foo` or `foo.exe` from the instructions in
the file at `github.com/example/foo/Scriggofile`.

In this other example:

```
$ scriggo build ./boo
```

the command will build an interpreter named `boo` or `boo.exe` from the
instructions in the [Scriggofile](scriggofile) `./boo/Scriggofile`.

### Complete syntax

The complete `scriggo build` command takes this form:

```
$ scriggo build [-f Scriggofile] [-w] [-v] [-x] [-work] [-o output] [module] 
```

With the exception of the flag `-o`, install has the same parameters as build:

```
$ scriggo install [-f Scriggofile] [-w] [-v] [-x] [-work] [module] 
```

The `-f` flag forces the command to read the given [Scriggofile](scriggofile) instead of the
Scriggofile of the module. For example:

```
$ scriggo build -f boo.Scriggofile github.com/example/foo
```

will build an interpreter named `foo` or `foo.exe` from the instructions in the file at `boo.Scriggofile`.

The `-w` flag omits the DWARF symbol table.

The `-v` flag prints the imported packages as defined in the [Scriggofile](scriggofile).

The `-x` flag prints the executed commands.

The `-work` flag prints the name of a temporary directory containing a work
module used to build the interpreter. The directory will not be deleted
after the build.

The `-o` flag forces build to write the resulting executable to the named output
file, instead in the current directory.

## Generate a package importer

