{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the variables{% end %}
{% Article %}

{% raw doc %}

# Variables

Scriggo allows you to declare new variables in a template in addition to the global ones.

## Declaration

A variable to be used, must first be declared with the `var` keyword:

```scriggo
{% var welcome = "hello" %}
```

or with a short assignment:

```scriggo
{% welcome := "hello" %}
```

Variables declared in a template can be used in the same way as the global variables:

```scriggo
{% var amount = 50 %}

shipping is free for an amount of at least ${{ amount }} 
```
<pre class="result">shipping is free for an amount of at least $50</pre>

## Variable's type

Each variable has a type and its type is the type of the value assigned to it at declaration time. The type can also be
specified explicitly in the declaration. For example:

```scriggo
{% var a string %}    a has type string and value ""
{% var b int %}       b has type int and value 0
{% var c bool %}      c has type bool and value false
{% var e []string %}  e has type []string and value nil
```

The following declarations are like the previous ones but with an initial value:

```scriggo
{% var a = "hello" %}                    a has type string and value "hello"
{% var b = 5 %}                          b has type int and value 5
{% var c = true %}                       c has type bool and value true
{% var e = []string{"ciao", "hello"} %}  e has type []string and value []string{"hello", "ciao"}
```

## Assignment


Once a variable is declared, its value can be changed at any time with an assignment. The type, on the other hand,
cannot change, so only values of the same type can be assigned to the variable.

```scriggo
{% greatings = "hello" %}
```

There are several other assignment statements that combine the assignment with an operator:

<div class="assignments">

| Writing  |  is like writing  |
|----------|-------------------|
| a++      | a = a + 1         |
| a--      | a = a - 1         |
| a += b   | a = a + b         |
| a -= b   | a = a - b         |
| a *= b   | a = a * b         |
| a /= b   | a = a / b         |
| a %= b   | a = a % b         |

</div>

### Visibility

The global variables are visible in all template files and in any part of the code. Instead, the visibility of the
variables declared in the code depends on where they are declared:

* if a variable is declared in a file with an extends statement or in an imported file, and not in a block (like the
    body of a macro, function or statement), it is visible in all file.

* if a variable is declared in a block (like the body of a macro, function or statement), it is visible from its
    declaration up to the end of the block.

* in the other cases it is visible from its declaration to the end of the file. 

### Examples

In the following example, the variable `n` is visible both inside and outside the `if` statement. The variable `s` is
visible only in the body of the `if`.

```scriggo
{% var n = 3 %}
{% if n > 2 %}
  {% var s = "hello" %}
  {{ n }}{{ s }}
{% end if %}
{{ n }}
```

In this other example, the `n` variable is not visible in the "column.html" file.

```scriggo
{% n := 3 %}
{{ render "column.html" }}
```

A [macro](/templates/macro) can be used to make its value available in another file. In this other example, the value of
the variable `n` is passed as an argument to the `Column` macro which could be declared in another file:

```scriggo
{% n := 3 %}
{{ Column(n) }}
```

### Multiple declarations and assignments

Declarations and assignments with `=` allow you to declare and assign more variable at the same time.

This example:

```scriggo
{% var n, s = 5, "ciao" %}

{% n, s = 2, "hello" %}
```

is equivalent to this one:

```scriggo
{% var n = 5 %}
{% var s = "ciao" %}

{% n = 2 %}
{% s = "hello" %}
```

The short assignment is a bit more complex, therefore read the
[Go language documentation](https://golang.org/doc/effective_go#redeclaration) before using it with multiple variables.

Multiple assignment is useful for example for swap the values of two variables:

```scriggo
{% a, b = b, a %}
```

and multiple declaration and assignment are useful for assigning the return values of a function call that return two or
more values:

```scriggo
{% var d, err = parseDecimal("12.99") %}
```

{% end raw doc %}
