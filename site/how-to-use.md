{% extends "/layouts/doc.html" %}
{% macro Title string %}How to use Scriggo{% end %}
{% Article %}

{% raw content %}

# How to use

This page explains how to use the Scriggo packages in Go programs to execute a template. If you want to learn the
template language, see the [templates](/templates) section instead.

### The first program

Let's start with this program:  

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

A template consists of one or more files that Scriggo reads from a [file system](https://pkg.go.dev/io/fs#FS). The
`scriggo.BuildTemplate` function gets a file system `fsys`, the path of the template file and builds a template
starting from this file. It returns a `*scriggo.Template` value that represents the compiled template.

To execute the template, you call the `Run` method on the compiled template. It writes the resulting code to the
`io.Writer` value passes as first argument. The `Run` method can be called several times even concurrently.

### Builtins

By default, a template code can only use the Go builtins, but you can create new builtins to use as globals in a
template.

The package [github.com/open2b/scriggo/builtin](https://pkg.go.dev/github.com/open2b/scriggo/builtin) has many useful
builtins that you can use in a template. So, let's start using them.

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

The `native.Declarations` type is used to represent variables, constants (typed and untyped), functions, types and
packages passed as globals when building a template, and can then be used in template code.

The following example exemplifies how to pass globals to a template. Note that the variable `v` is passed by reference.

```go
var v = "boo"
type Foo struct{ s string }

globals := native.Declarations{
    "v": &v,                                     // a variable named "v"
    "inc": func(i int) int { return i+1 },       // a function named "inc"
    "Foo": reflect.TypeOf(Foo{}),                // a type named "Foo"
    "limit" : 1024,                              // an int constant named "limit"
    "Pi": native.UntypedNumericConst("3.14159"), // an untyped numeric constant named "Pi"
    "colors": native.Package{                    // a package named "colors"
        Name: "colors",
        Declaration: native.Declarations{
            "Red": "#FF0000",
            "Yellow": "#FFFF00",
        },
    },
}
```

They can be used in the template:

```scriggo
{{ v }}   {{ inc(5) }}  {{ Foo{"foo"}.s }}  {{ limit }}  {{ Pi }}  {{ colors.Red }}
```

#### Pass a variable as global

If the template code assigns a new value to a global variable, the variable in the Go code will also have that value.

If you don't want this behavior, pass a typed nil pointer to the `BuildTemplate` function, then pass the initial value
of the variable to the `Run` method. In this way, each execution of the template will have a different variable:

```go
package main

import (
	"log"
	"os"

	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/native"
)

func main() {
	fsys := scriggo.Files{"index.txt": []byte("{{ v }}")}
	globals := native.Declarations{
		"v": (*string)(nil), // the initial value is passed to the Run method.
	}
	opts := &scriggo.BuildOptions{
		Globals: globals,
	}
	template, err := scriggo.BuildTemplate(fsys, "index.txt", opts)
	if err != nil {
		log.Fatal(err)
	}
	err = template.Run(os.Stdout, map[string]interface{}{"v": "foo"}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

{% end raw content %}