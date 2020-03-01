---
layout: article
---
# Get Started

In this get started you will learn how to run Go programs with the command line `scriggo` interpreter and will learn how
to start using Scriggo in applications and templates. 

It requires 10 minutes to be completed.

## Run a Go program

The `scriggo` command is a standalone interpreter that executes Go programs without requiring other software installed. 

Before you start working with Scriggo, <a href="/doc/install">install the Scriggo command</a> using precompiled binaries or compiling from source.

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

As an alternative you can add a shebang to the file:

```go
#!/usr/bin/env scriggo

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")    
}
``` 

and make the program's file executable:

```
$ chmod u+x hello.go
```  

Then execute the program:

```
$ ./hello.go
Hello, World!
```  

A program executed by the `scriggo` command can have access only to packages, variables, constants, functions and types
that has been explicitly compiled in the `scriggo` command.

See <a href="/doc/scriggo-command">scriggo command</a> for more details.

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

    // Load the program.
    program, err := scriggo.Load(src, nil, nil)
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

See <a href="/doc/package-importers">package importers</a> for more details.

## Use Scriggo in templates

The Scriggo template language supports inheritance, macros, includes, imports and autoescaping but most of all
it uses the Go language as the template scripting language. 

Scriggo templates can be written with plain text, HTML, CSS and JavaScript.

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
// Load and render a Scriggo template on the standard output.
package main

import (
    "os"
    "github.com/open2b/scriggo/template"
)

func main() {

    // src is the source code of the template file to render.
    src := []byte(`
    <!DOCTYPE html>
    <html>
    <head>Hello</head> 
    <body>
        {% who := "World" %}
        Hello, {{ who }}!
    </body>
    </html>
    `)

    // files is used to read the template files.
    files := template.MapReader{"index.html" : src}

    // Load the template.
    tmpl, err := template.Load("index.html", files, nil, template.LanguageHTML, nil)
    if err != nil {
        panic(err)
    }
 
    // Render the template on the standard output.
    err = tmpl.Render(os.Stdout, nil, nil)
    if err != nil {
        panic(err)
    }

}
```
{% endraw %}
{% raw %}

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

See the <a href="/doc/template">Scriggo template language</a> for more details.

<!--
The following is a more complex example of a Scriggo template:

{% raw %}
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
{% endraw %}

-->