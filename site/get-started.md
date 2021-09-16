{% extends "/layouts/doc.html" %}
{% macro Title string %}Get Started{% end %}
{% Article %}

# Get Started

In this get started you will learn how to embed Go programs in your applications and, in the next step, how to
<a href="/get-started-2"> execute Scriggo templates</a>.

1. [Embed Go programs in your application](#embed-go-programs-in-your-application)
2. [Execute templates in your application](/get-started-2)

It requires 10 minutes to be completed.

Before you start using Scriggo in a Go application you must <a href="https://golang.org/dl/">download and install Go</a>.

## Embed Go programs in your application

Open a terminal and create a new directory for the application:

```
$ mkdir hello
$ cd hello
```

Initialize a Go module in the previous created directory:

```
$ go mod init hello
```

Get the scriggo package:

```
$ go get github.com/open2b/scriggo
```

Create a file `main.go` with the following source code: 

```go
package main

import (
    "github.com/open2b/scriggo"
    "log"
)

func main() {

    // src is the source code of the program to run.
    src := []byte(`
        package main

        func main() {
            println("Hello, World!")
        }
    `)

    // Create a file system with the file of the program to run.
    fsys := scriggo.Files{"main.go": src}

    // Build the program.
    program, err := scriggo.Build(fsys, nil)
    if err != nil {
        log.Fatal(err)
    }
 
    // Run the program.
    err = program.Run(nil)
    if err != nil {
        log.Fatal(err)
    }

}
```

Build the application directly from the `hello` directory.

```
$ go build
```

Execute the application:

```
$ ./hello
Hello, World!
```

### Import packages

The program executed by Scriggo in the previous example:

```go
package main

func main() {
    println("Hello, World!")
}
```

does not import packages and uses the built-in `println` to print the string on the standard error.

A program executed by Scriggo can have access only to packages, variables, constants, functions and types that are
explicitly provided through an importer.

To import packages and the relative exported names, you have to pass an importer to the `Build` function. The code
of the importer can be coded manually or can be generated from a Scriggofile with the Import command.

#### The Import command

The `scriggo import` command allows to easily create an importer that imports the Go standard packages and other packages
according to the instructions in a <a href="scriggofile">Scriggofile</a>.

Create a file called "Scriggofile" with the following contents and put it in the module directory:

```
IMPORT STANDARD LIBRARY
IMPORT github.com/fatih/color
```

The Scriggo Import command reads the instructions in the Scriggofile and generates the source code of an importer that
imports the packages of the Go standard library and the `github.com/fatih/color` package.

Get the `github.com/fatih/color` module:

```
$ go get github.com/fatih/color
```

Execute `scriggo import` command in the module directory:

```
$ scriggo import -o packages.go
```

then replace the content of the `main.go` file with the following:

```go
package main

import (
    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/native"
    "log"
)

var packages native.Packages

func main() {

    // src is the source code of the program to run.
    src := []byte(`
        package main

        import (
            "fmt"
            "github.com/fatih/color"
        )
 
        func main() {
            fmt.Println("Here you are go")
            color.Red("Roses are red")
            color.Blue("Violets are blue")
        }
    `)

    // Create a file system with the file of the program to run.
    fsys := scriggo.Files{"main.go": src}

    // Use the importer in the packages variable.
    opts := &scriggo.BuildOptions{Packages: packages}

    // Build the program.
    program, err := scriggo.Build(fsys, opts)
    if err != nil {
        log.Fatal(err)
    }
 
    // Run the program.
    err = program.Run(nil)
    if err != nil {
        log.Fatal(err)
    }

}
```

Note the following lines in the `main.go` file:

```go
opts := &scriggo.BuildOptions{Packages: packages}
program, err := scriggo.Build(fsys, opts)
```

The importer, in the `packages` variable, is passed to the `Build` function. The `packages` variable is initialized in
the generated `packages.go` file.

#### Manually create an importer

The following code defines a package with path "acme.inc/colors" and name "colors" and passes it to the `Build` function.
The colors package exports the "Red" constant and the "Print" function.  

```go
packages := native.Packages{
    "acme.inc/colors" : native.Package{
        Name: "colors",
        Declarations: native.Declarations{
    	    "Red"   : "#C0392B",
    	    "Print" : func(color string) { fmt.Printf("The color is %s", color) },
        },
    },
}
opts := &scriggo.BuildOptions{Packages: packages}
program, err := scriggo.Build(fsys, opts)
```

The colors package can then be imported into the embedded program:

```go
package main

import "acme.inc/colors"

func main() {
	colors.Print(colors.Red)
}
```

{#
An importer must implement the `native.Importer` interface. In this example we use the `native.Packages` type that
implements that interface.

A package to be imported must implement the `native.Package` interface. In this example we use the `native.DeclarationsPackage` type
that implements that interface.

See <a href="/package-importers">package importers</a> for more details.

#}

## Continue with Scriggo templates

Continue the get started guide with the <a href="/get-started-2">Scriggo templates</a>.