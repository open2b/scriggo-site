{% extends "/layouts/article.html" %}
{% macro Title string %}Get Started{% end %}
{% Article %}

# Get started

This documentation will help you get started with Scriggo.

<ol id="markdown-toc">
  <li><a href="#download-and-install">Download and install</a></li>
  <li><a href="#your-first-scriggo-application">Your first Scriggo application</a></li>
</ol>

## Download and install

Open a terminal and run:

```
go get github.com/open2b/scriggo/cmd/scriggo
```

## Your first Scriggo application

This section will help you create your first application that embeds Scriggo.

First of all, create a new Go module by creating an empty directory and adding a `go.mod` file on it.
The `go.mod` file will need to include the Scriggo module as a dependency:

```
module scriggo-get-started

go 1.13

require scriggo
```

Now add a source file to your module named `main.go`, that contains:

```go
package main

import "scriggo"

func main() {

    src := `
        package main

        func main() {
            println("Hello, I'm an embedded Scriggo!")
        }
    `

    loader := scriggo.MapStringLoader{"main": src}

    program, err := scriggo.Load(loader, nil)
    if err != nil {
        panic(err)
    }

    err = program.Run(nil)
    if err != nil {
        panic(err)
    }
}
```

now run the source code with:

```
go run main.go
```

if you see the output:

```
Hello, I'm an embedded Scriggo!
```

congratulations, you have just wrote your first application that uses Scriggo!

Now let's see the source code above in detail.

The line

```go
import "scriggo"
```

imports the package `scriggo`, that is the main package of the module **scriggo**(TODO). There are other packages that you can import from such module, for example:

- `scriggo/template` gives you access to the Scriggo template
- `scriggo/runtime` let you import declarations that can be used to work with the runtime. (TODO)

In this section we will focus on the package `scriggo`, that is the one we have imported in our example.

TODO