# Differences between a Scriggo script and a program

One of the supported syntaxes of Scriggo is the **script** mode.

The script uses the same syntax of the programs (that is the same syntax of Go), with some exceptions:

1. The package name must not be included: the package is always `main`.
2. The function main must not be declared: every statement included in a script is implicitly declared in the main function (except for imports, see below).
3. First-level function declarations can be made using the package level function declaration syntax of Go, but the semantic is unchanged.


Let's see some examples. Note that all the following codes are valid scripts.

## A simple example

```go
println("hello")
```

this script prints `hello` on the standard error using the `println` builtin.
Note that the `package` clause, just like the `func main() { .. }` declaration is absent. They are implicit.

## Function declarations

Now let's see an example that includes function declarations:

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

SayHello("en")
SayHello("it")
```

the function declaration `func SayHello(..` is internally handled as `SayHello := func(..`.

As you may have guessed, a function that is still not declared cannot be called.

## Imports in scripts

A script can import other packages with the same syntax used by a program.

```go
import "fmt"

fmt.Println("Ohoho!")
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