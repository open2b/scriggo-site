{% extends "/layouts/doc.html" %}
{% macro Title string %}How to use Scriggo{% end %}
{% Article %}

{% raw content %}

# How to use

This page explains how to use the Scriggo packages to execute templates. If you want to learn the
template language, see the [templates](/templates) section instead.

* [Template files](#template-files) 
* [Build and run templates](#build-and-run-templates)
* [Use builtins](#use-builtins)
* [Pass variables to templates](#pass-variables-to-templates)
* [Use other types of globals](#use-other-types-of-globals)
* [Use Markdown](#use-markdown)
* [Import Go packages](#import-go-packages)
* [Do not parse {{ ... }}](#do-not-parse---)
* [Execution environment (env)](#execution-environment-env)
* [Implement IsTrue](#implement-istrue)
* [Allow "go" statement](#allow-go-statement)
* [Handle build and run errors](#handle-build-and-run-errors)

<div style="margin-top: 2rem;"></div>

### Template files

A template consists of one or more files that Scriggo reads from a [file system](https://pkg.go.dev/io/fs#FS). You can 
use any file system as the files embedded in the executable with the [//go:embed](https://pkg.go.dev/embed) directive or
read from a directory as the file system returned by the [os.DirFS](https://pkg.go.dev/os#DirFS) function.

For the examples of this documentation we use a simple file system,
[scriggo.Files](https://pkg.go.dev/github.com/open2b/scriggo#Files), whose files are read from a map.

By default, Scriggo gets the format of a file from the filename extension:

| Format     | Extension                        |
|------------|----------------------------------|
| HTML       | .html                            |
| Markdown   | .md .mkd .mkdn .mdown .markdown  |
| CSS        | .css                             |
| JavaScript | .js                              |
| JSON       | .json                            |
| Text       | all other extensions             |

If the file system implements the `scriggo.FormatFS` interface:

```go
type FormatFS interface {
    fs.FS
    Format(name string) (Format, error)
}
```

it calls the `Format` method to get the file format.

### Build and run templates

A template is first compiled and then executed, even concurrently by multiple goroutines. The compilation parses, type 
checks and emits the template bytecode. Execution is fast because it executes the bytecode on a virtual machine.

To compile a template, you pass the file system and the filename to the
[scriggo.BuildTemplate](https://pkg.go.dev/github.com/open2b/scriggo#BuildTemplate) function, then call the 
[Run](https://pkg.go.dev/github.com/open2b/scriggo#Template.Run) method on the returned compiled template.

The following program compiles and runs `{{ "hello" }}`:

```go
package main

import (
    "log"
     "os"

    "github.com/open2b/scriggo"
)

func main() {
    fsys := scriggo.Files{"index.txt": []byte(`{{ "hello" }}`)}
    template, err := scriggo.BuildTemplate(fsys, "index.txt", nil)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Use builtins

By default, a template code can only use the Go builtins. To allow the template code to use other builtins, you can pass
them as globals to the BuildTemplate function.

Scriggo, with the package [github.com/open2b/scriggo/builtin](https://pkg.go.dev/github.com/open2b/scriggo/builtin),
provides useful builtins ready-to-use. You can use all of them or just some. For example, the following program uses the 
min and max builtins:

```go
package main

import (
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/builtin"
    "github.com/open2b/scriggo/native"
)

func main() {
    fsys := scriggo.Files{"index.txt": []byte(`{{ min(3, 4) }} {{ max(8, 5) }}`)}
    globals := native.Declarations{
        "min": builtin.Min,
        "max": builtin.Max,
    }
    opts := &scriggo.BuildOptions{
        Globals: globals,
    }
    template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Pass variables to templates

You can pass a variable as a global to the template, as you previously passed the Scriggo builtins, but variables are 
passed by reference. Scriggo supports several use cases for passing a variable to a template, but let's start with 
simpler case. Look at this example:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/native"
)

func main() {
    var who = "World"
    fsys := scriggo.Files{"index.txt": []byte(`Hello {{ who }} {% who = "Scriggo" %}`)}
    opts := &scriggo.BuildOptions{
        Globals: native.Declarations{"who": &who},
    }
    template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\nwho is %q", who)
}
```

Note that we passed the address of the _who_ variable. If you execute this example, it prints:

```go
Hello World
who is "Scriggo"
```

How you can see, the value assigned to the _who_ variable, in the Go code, has been modified by the template code. If
you run the template concurrently, the goroutines that execute the template could access to the same variable at the
same time, so it may be necessary to use a synchronization mechanism.

If you want each execution to access its own variable, pass the variable pointer to the Run method instead of
BuildTemplate, with a different variable for each execution.

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/native"
)

func main() {
    fsys := scriggo.Files{"index.txt": []byte(`Hello {{ who }} {% who = "Scriggo" %}`)}
    opts := &scriggo.BuildOptions{
        Globals: native.Declarations{"who": (*string)(nil)},
    }
    template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
    if err != nil {
        log.Fatal(err)
    }
    // Use a different variable for each call to Run, for example using a closure.
    var who = "World"
    err = template.Run(os.Stdout, map[string]interface{}{"who": &who}, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\nwho is %q", who)
}
```

Note that we passed `(*string)(nil)` to BuildTemplate as global. This way the compiler know the type of the variable. 
Then we passed the variable pointer to the Run method. If you execute the program, it prints again:

```go
Hello World
who is "Scriggo"
```

The printed value is not changed because the variable is still "shared" between the Go code and template code. But no
between different executions, if a different variable was passed for each of them.

If you don't need to "share" the variable at all, pass the variable's value to the Run method instead of the pointer:

```go
err = template.Run(os.Stdout, map[string]interface{}{"who": "World"}, nil)
```

In this case if the template code assigns a new value to the variable, no variables are changed in the Go code. But be 
aware that if the variable is for example a slice, a map or a pointer value, the template code can still change the
slice and map elements and the pointed value.

#### How to pass a pointer type

Passing a pointer type is the same as passing any other type, there is no special cases. To pass a variable `v` with
type `*T`, write:

```go
opts := &scriggo.BuildOptions{
    Globals: native.Declarations{"v": (**T)(nil)},
}
```

To "share" the variable with the Go code, pass a pointer to the variable: 

```go
v := &T{}
err = template.Run(os.Stdout, map[string]interface{}{"v": &v}, nil)
```

Otherwise, pass the value of the variable:

```go
v := &T{}
err = template.Run(os.Stdout, map[string]interface{}{"v": v}, nil)
```

### Use other types of globals

You can also pass functions, constants (typed and untyped), types and even packages to a template as globals. For
functions, we have already seen the example of Scriggo's min and max builtins.

The following code exemplifies all the types of globals that you can pass to a template:

```go
native.Declarations{
    "articles": &articles,                       // a variable named "articles"
    "inc": func(i int) int { return i+1 },       // a function named "inc"
    "Foo": reflect.TypeOf(Foo{}),                // a type named "Foo"
    "limit" : 1024,                              // an int constant named "limit"
    "Pi": native.UntypedNumericConst("3.14159"), // an untyped numeric constant named "Pi"
    "separator": native.UntypedStringConst("/"), // an untyped string constant named "separator"
    "True": native.UntypedBooleanConst(true),    // an untyped boolean constant named "True"
    "colors": native.Package{                    // a package named "colors"
        Name: "colors",
        Declarations: native.Declarations{
            "Red": "#FF0000",
            "Yellow": "#FFFF00",
        },
    },
}
```

All of them can then be used in the template, for example:

```scriggo
 {{ len(articles) }}   {{ inc(5) }}      {{ Foo{"foo"}.Boo }}   {{ limit }}
 {{ Pi }}              {{ separator }}   {{ True }}             {{ colors.Red }}
```

Except variables, as seen above, the template code cannot modify a global passed with a native.Declarations value.

### Use Markdown

Scriggo can run template files containing only Markdown and return the rendered Markdown code.

With Scriggo you can also, for example, run a Markdown file that extends an HTML file, run an HTML file that
render a Markdown file and convert a value with _markdown_ type to the _html_ type. To do this you need to pass a
[converter](https://pkg.go.dev/github.com/open2b/scriggo#Converter) to the BuildTemplate function to convert Markdown
to HTML.

For example, you can use <a href="https://github.com/yuin/goldmark" target="_blank">Goldmark</a> as a converter.

```go
package main

import (
    "io"
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/yuin/goldmark"
)

func main() {
    fsys := scriggo.Files{
        "index.html": []byte(`{{ markdown("# The Ancient Art Of Tea") }}`),
    }
    md := goldmark.New()
    opts := &scriggo.BuildOptions{
        MarkdownConverter: func(src []byte, out io.Writer) error {
            return md.Convert(src, out)
        },
    }
    template, err := scriggo.BuildTemplate(fsys, "index.html", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

If you run this program, it prints:

```go
<h1>The Ancient Art Of Tea</h1>
```

There any many options you can use with Goldmark, for example the _scriggo serve_ command uses these options:

```go
goldmark.New(
    goldmark.WithRendererOptions(html.WithUnsafe()),        // do not remove any HTML code
    goldmark.WithParserOptions(parser.WithAutoHeadingID()), // add ids to the headings
    goldmark.WithExtensions(extension.GFM))                 // allow GitHub Flavored Markdown
```

### Import Go packages

You have already seen how to [pass a package as global](#use-other-types-of-globals) to a template. In this case the
template code not have to import the package to use it.

You can also allow the import of packages. To import packages and their exported names, you need to pass an importer to 
the BuildTemplate function. The importer code can be also be
[generated with the scriggo import command](#import-packages-with-the-import-command). 

The following program defines a package with path "acme.inc/colors" and name "colors" and passes it to the 
`BuildTemplate` function. The colors package exports the "Red" constant and the "Name" function.

```go
package main

import (
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/native"
)

func main() {
    fsys := scriggo.Files{
        "index.txt": []byte(`
            {% import "acme.inc/colors" %}
            {{ colors.Red }} is called {{ colors.Name(colors.Red) }}
        `),
    }
    packages := native.Packages{
        "acme.inc/colors" : native.Package{
            Name: "colors",
            Declarations: native.Declarations{
                "Red": "#C0392B",
                "Name": func(color string) string {
                    if color == "#C0392B" {
                        return "Red"
                    }
                    return ""
                },
            },
        },
    }
    opts := &scriggo.BuildOptions{Packages: packages}
    template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

By importing packages, the types [CombinedPackage](https://pkg.go.dev/github.com/open2b/scriggo/native#CombinedPackage) 
and [CombinedImporter](https://pkg.go.dev/github.com/open2b/scriggo/native#CombinedImporter) may be useful.

#### Import packages with the import command 

The `scriggo import` command allows to easily create an importer that imports the Go standard packages and other
packages according to the instructions in a [Scriggofile](/Scriggofile).

For example, create a file called "Scriggofile" with the following content and put it the root directory of a Go
module:

```scriggofile
IMPORT STANDARD LIBRARY
```

The Scriggo Import command reads the instructions in the Scriggofile and generates the source code of an importer that
imports the packages of the Go standard library.

Execute `scriggo import` command in the module directory:

```shell
$ scriggo import -o packages.go
```

Be patient and wait for its conclusion, it takes several seconds.

The code in the `packages.go` file assigns a
[native.Packages](https://pkg.go.dev/github.com/open2b/scriggo/native#Packages) value to the _packages_ variable with the
packages indicated in the Scriggofile.

Then, you can pass the _packages_ variable to the BuildTemplate function:

```go
package main

import (
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/native"
)

var packages native.Packages

func main() {
    fsys := scriggo.Files{"index.txt": []byte(`
    {%%
        import "net/http"

        r, err := http.Get("https://example.com/")
        if err != nil {
            show "error: ", err
        } else {
            show "the server has responded with status ", r.Status
        }
    %%}
    `)}
    opts := &scriggo.BuildOptions{Packages: packages}
    template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Do not parse {{ ... }}

If you use a client-side template that already uses the `{{ ... }}` syntax, you can disable this syntax in Scriggo
using the _NoParseShortShowStmt_ option with the BuildTemplate function:

```go
opts := &scriggo.BuildOptions{
    NoParseShortShowStmt: true,
}
```

Using this option, Scriggo does not parse the short show statements, and then you can use this syntax client-side. In Scriggo, you can continue to use the `{% show ... %}` syntax.

### Execution environment (env)

The execution environment, or simply env, is a value created at each template execution whose type implements the 
[native.Env](https://pkg.go.dev/github.com/open2b/scriggo/native#Env) interface. A function or method, passed as global 
or imported into a template, receives env as first argument if its first parameter has type
[native.Env](https://pkg.go.dev/github.com/open2b/scriggo/native#Env).

You can use env to get the [context](https://pkg.go.dev/context#Context) passed to the Run method, exit the execution,
raise a fatal error, print with the _print_ and _println_ builtins used in the template, and get the caller's path 
relative to the root of the template. 

For example, the following program passes a builtin named _exit_ to the template. When called, it ends the execution of
the template.

```go
package main

import (
    "log"
    "os"

    "github.com/open2b/scriggo"
    "github.com/open2b/scriggo/native"
)

// Env terminates the template execution with status code.
func Exit(env native.Env, code int) {
    env.Exit(code)
}

func main() {
    fsys := scriggo.Files{"index.txt": []byte(`{% exit(0) %}{{ "not displayed" }}`)}
    opts := &scriggo.BuildOptions{
        Globals: native.Declarations{"exit": Exit},
    }
    template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
    if err != nil {
        log.Fatal(err)
    }
    err = template.Run(os.Stdout, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

Note that template code does not see the env parameter of the Exit function and calls the exit builtin as `exit(0)`.

### Implement IsTrue

The IsTrue method can be implemented by struct and pointer to struct types that can be truthful or not truthful.
See the [template specification](/templates/specification#truthful-values) for the details on truthful values.

Given this type:

```go
type Slide struct {
    Images []Image
    Timing int
}
```

a template writer would expect a Slide value to be truthful if it contains images and not truthful if it does not. So
she can write:

```scriggo
{% if slide %}
<div class="slide">
  {% for image in slide.Images %}{{ image }}{% end %}
</div>
{% end if %}
```

To do this, add the IsTrue method to the Slide (or *Slice) type:  

```go
func (s Slide) IsTrue() bool {
    return len(s.Images) > 0
}
```

### Allow "go" statement

By default, the _go_ statement is not allowed, the compilation of a template with a _go_ statement fails. You can allow
the _go_ statement using the AllowGoStmt option with the BuildTemplate function:

```go
opts := &scriggo.BuildOptions{
    AllowGoStmt: true,
}
```

Note that currently the goroutines started by the template code do not terminate when the execution of the template
ends. This behavior may change in a future version of Scriggo.

### Handle build and run errors

How handle errors depends on the context in which templates are compiled and executed. Therefore, the way in which
errors are handled may vary from application to application.

#### Build errors

The [BuildTemplate](https://pkg.go.dev/github.com/open2b/scriggo#BuildTemplate) function:

```go
template, err := scriggo.BuildTemplate(fsys, "index.md", opts)
if err != nil {
    if errors.Is(err, fs.ErrNotExist) {
        // handle do not exist error.
    }
    if err, ok := err.(*scriggo.BuildError); ok {
        // handle compilation error.
    }
    // handle other errors returned by the file system methods,
    // by the convert function and other internal errors.
}
```

If the file to be compiled does not exist, it returns an error satisfying `errors.Is(err, fs.ErrNotExist)`, while if a
file to extend, import or render does not exist, it returns a
[*scriggo.BuildError](https://pkg.go.dev/github.com/open2b/scriggo#BuildError) value.

A [*scriggo.BuildError](https://pkg.go.dev/github.com/open2b/scriggo#BuildError) value is an error in the template code
such as a syntax error, a type checking error, a cycle error and a limit exceeded error. You may show this error to
the template author, so that she can fix it.

Other errors are unexpected errors returned by a file system method, the convert function and other internal errors.

#### Run errors

The [Run](https://pkg.go.dev/github.com/open2b/scriggo#Template.Run) method:

```go
err = template.Run(os.Stdout, nil, nil)
if err != nil {
    if err, ok := err.(*scriggo.PanicError); ok {
        // handle a panic in the template code.
    }
    if err, ok := err.(*scriggo.ExitError); ok {
        // handle a call to env.Exit with a not zero code.
    }
    // handle other errors returned by the out.Write method,
    // by the convert function, and the Err method of a context
	// when it is cancelled.
}
```

If the template code calls the panic builtin or a panic is raised from a global function or method, and this panic is
not recovered, the Run method returns a [*scriggo.PanicError](https://pkg.go.dev/github.com/open2b/scriggo#PanicError)
value. You may panic with the value `err.Error()`, log the error, show the error on the console or return it to a
browser.

If the [Env.Exit](https://pkg.go.dev/github.com/open2b/scriggo/native#Env.Exit) method is called with a not zero code,
Run returns a [*scriggo.ExitError](https://pkg.go.dev/github.com/open2b/scriggo#ExitError) value. You may exit with
this code, log the exit error, show the exit error on the console or return it to a browser.

Other errors are unexpected errors returned by the output writer, the convert function, the error of a cancelled context
and other internal errors.

The Run method panics if the [Env.Fatal](https://pkg.go.dev/github.com/open2b/scriggo/native#Env.Fatal) method is
called.

{% end raw content %}