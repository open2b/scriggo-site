{% extends "/layouts/doc.html" %}
{% macro Title string %}Get Started{% end %}
{% Article %}

# Get Started With Templates

In the first step you have learned how to <a href="/get-started">embed Go programs</a>.
In this second step you will learn how to execute Scriggo templates in your applications.

It requires 10 minutes to be completed.

Before you start using Scriggo in a Go application you must <a href="https://golang.org/dl/">download and install Go</a>.

## Execute templates in your application

Scriggo, in templates, supports inheritance, macros, partials, imports and contextual autoescaping but most of all it uses the Go
language as the template scripting language. 

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
    "github.com/open2b/scriggo"
)

func main() {

    // Content of the template file to run.
    content := []byte(`
    <!DOCTYPE html>
    <html>
    <head>Hello</head> 
    <body>
        {% who := "World" %}
        Hello, {{ who }}!
    </body>
    </html>
    `)

    // Create a file system with the file of the template to run.
    fsys := scriggo.Files{"index.html": content}

    // Build the template.
    template, err := scriggo.BuildTemplate(fsys, "index.html", nil)
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

Execute `go mod tidy` and build the application directly from the `hello-template` directory.

```
$ go mod tidy
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

## Globals

A template executed by Scriggo, apart from the Go builtins, can only use explicitly passed globals with the `Globals`
option and can only import packages whose importer has been passed with the `Packages` option.

Scriggo provides some useful globals, in the <a href="https://pkg.go.dev/github.com/open2b/scriggo/builtin">github.com/open2b/scriggo/builtin</a>
package, that can be used in templates.

Replace the content of the `main.go` file with the following:

{% raw %}
```go
// Build and run a Scriggo template.
package main

import (
    "os"
    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/builtin"
    "github.com/open2b/scriggo/native"
)

func main() {

    // Content of the template file to run.
    content := []byte(`
    <!DOCTYPE html>
    <html>
    <head>Hello</head> 
    <body>
        {% who := "World" %}
        Hello, {{ replace(who, "World", "世界", 1) }}!
    </body>
    </html>
    `)

    // Create a file system with the file of the template to run.
    fsys := scriggo.Files{"index.html": content}

    // Allow to use the "replace" built-in in the template file.
    globals := native.Declarations{
        "replace": builtin.Replace,
    }
    opts := scriggo.Options{Globals: globals}

    // Build the template.
    template, err := scriggo.BuildTemplate(fsys, "index.html", opts)
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

Rebuild the application directly from the `hello-template` directory.

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
        Hello, 世界!
    </body>
    </html>
 
```

### Use your globals

The following code defines a global function called "printHello" and passes it to the `BuildTemplate` function.

```go
globals := native.Declarations{
    "printHello": func (who string) string { return fmt.Sprintf("Hello, %s!", who) },
}
opts := scriggo.Options{Globals: globals}
template, err := scriggo.BuildTemplate(fsys, "index.html", opts)
```

Globals can be functions, variables, constants, types and even packages.
