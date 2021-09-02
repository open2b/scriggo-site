{% extends "/layouts/article.html" %}
{% macro Title string %}Scriggofile{% end %}
{% Article %}

# Scriggofile

A Scriggofile is a text file with a specific format used by the `scriggo` command.
The `scriggo` command uses the instructions in a Scriggofile to create an
interpreter or to generate an importer to use in an existing application.

A Scriggofile defines which packages a program, script or template can import,
what exported declarations in a package can be imported and so on.  

## Format

The format of the Scriggofile is:

```
# A comment
INSTRUCTION arguments
```

A line starting with `#` is a comment, and the instructions are case
insensitive but for convention are written in uppercase (the syntax recalls
that used by Dockerfile). 

A Scriggofile must be encoded as UTF-8 and it should be named `Scriggofile`
or with the extension `.Scriggofile` as for `example.Scriggofile`.

## Intructions

The instructions are:

### IMPORT STANDARD LIBRARY 

Makes almost all the packages in the Go standard library importable in a program, script or template
executed by the application.

To view all packages imported execute from a console:

```
$ scriggo stdlib
```

### IMPORT &lt;package&gt;

Make the package with path `package` importable.
 
### IMPORT &lt;package&gt; INCLUDING &lt;A&gt; &lt;B&gt; &lt;C&gt;

As for `IMPORT <package>` but only the exported names `A`, `B` and `C` are imported.

### IMPORT &lt;package&gt; EXCLUDING &lt;A&gt; &lt;B&gt; &lt;C&gt;

As for `IMPORT <package>` but the exported names `A`, `B` and `C` are not imported.  

### IMPORT &lt;package&gt; AS &lt;as&gt;

As for `IMPORT <package>` but the path with which it can be imported is named `as`.
`INCLUDING` and `EXCLUDING` can be used as for the other forms of `IMPORT` at the end of
the instruction. Is not possible to use a path `as` that would conflict with a Go
standard library package path, even if this latter is not imported in the Scriggofile.
    
### IMPORT &lt;package&gt; AS main

Make the package with path `package` imported as the main package in a script or template.
It is the same as writing:

```
import . "<package>"
```
in a Go program.

`INCLUDING` and `EXCLUDING` can be used as for the other forms of `IMPORT` at the end
of the instruction.

### IMPORT &lt;package&gt; AS main NOT CAPITALIZED

As for `IMPORT <package> AS main` but the exported names in the package will be imported
not capitalized. For example a name `FooFoo` declared in the package will be imported in
the script or template as `fooFoo`.

### TARGET PROGRAMS SCRIPTS TEMPLATES

Indicates witch are the targets of the application. It will be able to execute only the
type of sources listed in the `TARGET` instruction. This instruction is only read by the
`build` and `install` commands.

### SET VARIABLE &lt;name&gt; 

Set the name of the variable to witch is assigned the value of type `scriggo.PackageLoader`
with the packages to import. By default the name is `packages`. This instruction is only
read by the `embed` command. 

### SET PACKAGE &lt;name&gt;

Set the name of the package of the generated Go source file. By default the name of the
package is `main`. This instruction is read only by the command `scriggo embed`.

### GOOS linux windows

Specifies the operating systems that will be supported by the built interpreter.
If the GOOS at the time the Scriggofile is parsed is not listed in the `GOOS` instruction,
the `init` and `embed` commands fail. If there is no `GOOS` instruction, all the
operating systems are supported.

To view possible GOOS values execute:
```
$ go tool dist list
```
