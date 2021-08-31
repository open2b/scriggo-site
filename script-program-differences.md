{% extends "/layouts/article.html" %}
{% macro Title %}Differences between a Scriggo script and a program{% end %}
{% Article %}

# Differences between a Scriggo script and a program

<ol>
  <li><a href="#a-simple-example">A simple example</a></li>
  <li><a href="#function-declarations" id="markdown-toc-function-declarations">Function declarations</a></li>
  <li><a href="#imports-in-scripts">Imports in scripts</a></li>
  <li><a href="#equivalent-program-for-a-script" id="markdown-toc-equivalent-program-for-a-script">Equivalent program for a script</a></li>
<  li><a href="#package-level-declarations-in-scripts">Package level declarations in scripts</a></li>
</ol>

One of the supported syntaxes of Scriggo is the **script** mode.

The script uses the same syntax of the programs (that is the same syntax of Go), with some exceptions:

1. The package name must not be included: the package is always `main`.
2. The function main must not be declared: every statement included in a script is implicitly declared in the main function (except for imports, see below).
3. First-level function declarations can be made using the package level function declaration syntax of Go, but the semantic is unchanged.


Let's see some examples.

## A simple example

```go
println("hello")
```

this script prints `hello` on the standard error using the `println` builtin.
Note that the `package` clause, just like the `func main() { .. }` declaration is absent. They are implicit.

## Function declarations

Now let's see a more complex example that includes a function declaration:

```go
func SayHello(lang string) {
    switch lang {
        case "italiano":
        println("Ciao!")
        case "english":
        println("Hello!")
        case "español":
        println("¡Hola!")
        default:
        println("Still studying " + lang + "...")
    }
}

for _, lang := range []string{"en", "it", "es", "jp"} {
    SayHello(lang)
}
```

the function declaration `func SayHello(..` is internally handled as `SayHello := func(..`.

As you may have guessed, a function that is still not declared cannot be called, so *the functions must be declared before they are used*.

```go

Answer() // <- ERROR!

func Answer() int {
    return 42
}

Answer() // 42
```

## Imports in scripts

A script can import other packages with the same syntax used by a program.

```go
import "fmt"
import . "math"

fmt.Println("Ohoho! " + MaxInt8)
```

## Equivalent program for a script

More in general, you can consider scripts as a subset of programs; so, a script can always be translated into the equivalent program:

```go
// [ SCRIPT IMPORTS ]

package main

func main() {

    // [ SCRIPT SOURCE CODE ]

}
```

## Package level declarations in scripts

Given that a script can only define the content of the body of the function `main`, you may be wondering: how can I declare and use package level declarations?
The answer is that package level declarations are handled by the host application that compiles and runs the script.

Among other things, the host is responsible of:

- choose what are the importable packages
- choose what are the package level declarations
- limit the use of some statement or the amount of available memory/time

To sum up, in scripts:

- it's not possible to declare package level declarations
- the available package level declarations and importable packages depends on who makes the script compile/run