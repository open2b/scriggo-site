{% extends "/layouts/article.html" %}
{% macro Title %}Get Started{% end %}
{% Article %}

# Get Started

In this get started you will learn how to run Go programs with the command line `scriggo` interpreter and will learn how
to start using Scriggo in applications and templates. 

It requires 10 minutes to be completed.

## Run a Go program

The `scriggo` command is a standalone interpreter that executes Go programs without requiring other software installed. 

Before you start working with Scriggo, <a href="/install">install the Scriggo command</a> using precompiled binaries or compiling from source.

To run a Go program, as the following example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")    
}
```

you can give the file of the program as the first argument to the `scriggo` command:

```
$ scriggo hello.go
Hello, World!
```  

A program executed by the `scriggo` command can have access only to packages, variables, constants, functions and types
that has been explicitly compiled in the `scriggo` command.

See <a href="/scriggo-command">scriggo command</a> for more details.

## Use Scriggo in applications

Before you start using Scriggo in a Go application you must <a href="https://golang.org/dl/">download and install Go</a>.

Open a terminal and create a new directory for the application: 

```
$ mkdir hello
$ cd hello
```

Initialize a Go module in the previous created directory:

```
$ go mod init hello
```

Create a file `main.go` with the following source code: 

```go
package main

import (
    "github.com/open2b/scriggo"
)

func main() {

    // src is the source code of the program to run.
    src := strings.NewReader(`
        package main

        func main() {
            println("Hello, World!")
        }
    `)

    // Build the program.
    program, err := scriggo.Build(src, nil)
    if err != nil {
        panic(err)
    }
 
    // Run the program.
    _, err = program.Run(nil)
    if err != nil {
        panic(err)
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

> On Windows execute the application with the following command:
> ```
> $ hello.exe
> Hello, World!
> ``` 

### Import packages

The program executed by Scriggo in the previous example:

```go
package main

func main() {
    println("Hello, World!")
}
```

does not import packages and uses the builtin `println` to print the string on the standard error.

A program executed by Scriggo can have access only to packages, variables, constants, functions and types that are
explicitly provided through a package importer.

See <a href="/package-importers">package importers</a> for more details.

## Use Scriggo in templates

The Scriggo template language supports inheritance, macros, partials, imports and autoescaping but most of all
it uses the Go language as the template scripting language. 

Scriggo templates can be written with plain text, HTML, Markdown, CSS, JavaScript and JSON.

Open a terminal and create a new directory for the application: 

```
$ mkdir hello-template
$ cd hello-template
```

Initialize a Go module in the previous created directory:

```
$ go mod init hello-template
```

Create a file `main.go` with the following source code:

{% raw %}
```go
// Build and run a Scriggo template.
package main

import (
    "os"
    "github.com/open2b/scriggo/templates"
)

func main() {

    // Content of the template file to run.
    content := `
    <!DOCTYPE html>
    <html>
    <head>Hello</head> 
    <body>
        {% var who = "World" %}
        Hello, {{ who }}!
    </body>
    </html>
    `

    // File system used to read the template files. 
    fsys := templates.MapFS{"index.html" : content}

    // Build the template.
    template, err := templates.Build(fsys, "index.html", nil)
    if err != nil {
        panic(err)
    }
 
    // Run the template and print it to the standard output.
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        panic(err)
    }

}
```
{% end raw %}

Build the application directly from the `hello-template` directory.

```
$ go build
```

Execute the application:

```
$ ./hello-template

    <!DOCTYPE html>
    <html>
    <head>Hello</head> 
    <body>
        Hello, World!
    </body>
    </html>
 
```

See the <a href="/template">Scriggo template language</a> for more details.

{#
The following is a more complex example of a Scriggo template:

{% raw code %}
```html
{% extends "layout.html" %}
{% import "banners.html" %}
{% Head %}
    <title>Hello</title>
{% end %}
{% Body %}
    {% include "column.html" %} 
    <div>
      {% who := "World" %}
      Hello, {{ who }}!
    </div>
    {% show banners.Banner() %}
{% end %}
</html>
 ```
{% end raw code %}

#}