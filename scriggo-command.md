{% extends "/layouts/article.html" %}
{% macro Title %}Scriggo command{% end %}
{% Article %}

# Scriggo command

Scriggo has a command line interface, the `scriggo` command, that allows to:

* run Go programs
* add and remove importable packages
* build scriggo commands with specific packages
* generate a package loader for an existing application

It is a single binary executable with no external dependencies. It includes the entire Go library in a executable that when compressed is only 5 megabyte.

If your use of the Scriggo command is not limited to run programs, before you start you must <a href="https://golang.org/dl/">download and install Go</a>.

See <a href="/doc/install">install Scriggo command</a> to install the `scriggo` command using precompiled binaries or compiling from source.

### Get help from command line

To get help from the command line run the following command:

```
$ scriggo help
```

## Run Go programs

A Scriggo command is a standalone interpreter and consequently, to run programs, does not require Go installed. You can give the
file of the program as the first argument to the `scriggo` command: 

```
$ scriggo file [arguments...]
```

where `file` is the Go program file to run and `arguments` are the arguments that are passed to the executed program as
`os.Arg[1]`, `os.Arg[2]` ecc... 

The exit status of the `scriggo` command is the exit status of the executed program.

If the name of the file could be interpreted as a `scriggo` subcommand, use `--` as in:

```
$ scriggo -- file [arguments...]
```

### Shebang

The file to execute can have a shebang as first line as in:

```go
#!/usr/bin/env scriggo

package main

func main() {
    println("Hello, World!")    
}
```

<blockquote>
  <p>Shebang is not allowed by the <a href="https://golang.org/ref/spec">Go specification</a>,
  but Scriggo runs Go program files starting with a shebang because is extremely
  useful when Go programs are not compiled to binary but interpreted as Scriggo does.</p>
</blockquote>  

To raise an error if the file has a shebang, use the <span style="white-space:nowrap">`-strict`</span> option:

```
$ scriggo -strict file [arguments...]
```

## Add and remove importable packages

An executed program can import only packages that have been explicitly compiled in a `scriggo` command. To list all the
packages compiled in a `scriggo` command run from command line:

```
$ scriggo list packages
```

It will prints to the standard output all the packages compiled in the executed `scriggo` command and that can be imported in a Go program.

To add or remove a package use the Rebuild command. The basic Rebuild command takes this form:

```
$ scriggo rebuild [add|remove] path[@version]
```

It creates the rebuilded Scriggo command in the current directory with the name `scriggo` or `scriggo.exe`, overwriting
an existing file with the same name.

### Add packages

To add packages use the `scriggo rebuild add` command:

```
$ scriggo rebuild add [-force] path[@version]
```

where `path` is the import path of the package and `version` is the version of the package's module. If the argument
`version` is missing, the version is the latest. You can give more packages as argument. For example to add the package
`foo` with version `v1.3` and the latest version of the package `bar` run from command line:
 
```
$ scriggo rebuild add foo@v1.3 bar
```

If a package already exists with a different version, it will be updated at the new version. If more packages exist
for the same module an error occurs. To force the update of all the packages of the same module to the given version use
the `-force` option, but the packages of the same module listed as argument must have the same version.

For packages of the Go standard library don't give the version because the version is implicitly the version of Go.
To add almost all the packages of the Go standard library use `stdlib` as package argument. 

### Remove packages

To remove packages use the `scriggo rebuild remove` command:

```
$ scriggo rebuild remove paths
```

where `paths` are the import paths of the packages to remove. For example to remove packages `foo` and `bar` run from command line:

```
$ scriggo rebuild remove foo bar
```

## Build Scriggo commands

The scriggo command uses the instructions in a [Scriggofile](scriggofile) to create the interpreter or generate the package loader.

The `scriggo build` command builds a standalone Go interpreter from a [Scriggofile](scriggofile) in a module.

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

## Generate a package loader

