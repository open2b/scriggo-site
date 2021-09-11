{% extends "/layouts/doc.html" %}
{% macro Title string %}Switch from Jinja to Scriggo templates{% end %}
{% Article %}
{% raw document %}

# Switch from Jinja to Scriggo

This page is targeted to users who are already familiar with designing templates in
Python/Jinja but have a limited/no experience with Go/Scriggo.

Here is provided a overview of the similarities and the differences between the two
templating languages.

> Some examples on this page are taken from [the Jinja
> documentation](https://jinja.palletsprojects.com/en/3.0.x/templates/) (copyright
> 2007 Pallets).


- [From Jinja to Scriggo](#from-jinja-to-scriggo)
  - [Syntax](#syntax)
  - [Type system](#type-system)
  - [Escaping](#escaping)
  - [Tests](#tests)
  - [Comments](#comments)
  - [Expressions](#expressions)
    - [Literals](#literals)
      - [Strings](#strings)
      - [Lists](#lists)
      - [Dictionaries](#dictionaries)
      - [Tuples](#tuples)
    - [If Expression](#if-expression)
    - [Math and comparisons](#math-and-comparisons)
    - [Logic](#logic)
    - [Operator `in`](#operator-in)
    - [Operator `+` (plus)](#operator--plus)
    - [Operator `~` (tilde)](#operator--tilde)
    - [Methods](#methods)
  - [List of control structures](#list-of-control-structures)
    - [For](#for)
    - [If](#if)
    - [Block Assignments](#block-assignments)
    - [Calls](#calls)
    - [Filters](#filters)
    - [Assignments](#assignments)
  - [Template inheritance](#template-inheritance)
  - [Blocks and macros](#blocks-and-macros)
  - [Including other template files](#including-other-template-files)
  - [Importing other template files](#importing-other-template-files)
  - [List of Builtin filters](#list-of-builtin-filters)
    - [`abs`](#abs)
    - [`attr`](#attr)
    - [`length`](#length)
  - [HTML Escaping](#html-escaping)
    - [Automatic contextual escaping](#automatic-contextual-escaping)
  - [Continue to learn](#continue-to-learn)

## Syntax

Jinja and Scriggo share the same syntax for enclosing statements, rendering
expressions and commenting.

- `{% .. %}` for statements
- `{{ .. }}` for expressions
- `{# .. #}` for comments

Some statements need an _end_ tag to be closed.

For example:

```scriggo
{% if condition %}
    ...
{% end %}
```

Note that in Scriggo, unlike Jinja, the statement name is optional and, if provided,
it must be separated by the `end` keyword:

```scriggo
{% if condition %}
    ...
{% end if %}
```

## Type system

One of the biggest differences between Jinja and Scriggo is the type system.

Jinja is based on Python, thus it offers a dynamically-typed template system, while
Scriggo (which is based on Go) is statically typed.

This means that, when writing templates in Scriggo, most of the errors will be catch
at compile time.

For example, writing this code in Scriggo:

```scriggo
{% var name string = "hello" %}
{% name = 23 %}                     {# ← WRONG! #}
```

results in a _compilation error_ like this:

```
cannot use 23 (type untyped int) as type string in assignment
```

## Escaping

Like in Jinja, the easiest way to output a string that is part of the Scriggo syntax
is putting it inside a string literal and showing it:

```
{{ "{{" }}
```

For bigger sections, Scriggo makes available the _raw_ block that acts the same as
the _raw_ block in Jinja:

```scriggo
{% raw %}
    <ul>
    {% for item in seq %}
        <li>{{ item }}</li>
    {% end for %}
    </ul>
{% end raw %}
```

> Note that the _raw_ block is closed by `{% end raw %}`, not `{% endraw %}`.

## Tests

Since Scriggo implements a full programming language, tests can be performed using
any valid expression.

For example this test in Jinja:

```jinja
{% if loop.index is divisibleby 3 %}
```

can be rewritten in Scriggo simply using arithmetic and comparison operators:

```scriggo
{% if loop.index % 3 == 0 %}
```

## Comments

In Scriggo, just like Jinja, you can write commented code enclosing it between `{#`
and `#}`. Commented portions of the template won't be rendered by Scriggo.

Note that Scriggo, unlike Jinja, supports nested comments:

```scriggo
{#  {# This is a nested comment! #}  #}
```

## Expressions

### Literals

#### Strings

Strings, in Scriggo, can be enclosed between double quotation marks or backticks.

For example, these are valid string literals in Scriggo:

```scriggo
{{ "hello" }}
{{ `hi!` }}
```

The single quotation mark is used to denote a _rune_, thus it can only enclose one
character:

```scriggo
{% var mychar = 's' %}
{{ mychar }}
```

#### Lists

Scriggo has a typed list type called _slice_, where the type must be specified at
compilation-time. Also, slice elements must be enclosed between `{` and `}` instead
of `[` and `]`.

For example this code in Jinja:

```jinja
{% set l = [1, 2, 3] %}
```

can be written in Scriggo as:

```scriggo
{% var l = []int{1, 2, 3} %}
```

#### Dictionaries

Scriggo has a typed dictionary type called _map_, where both the key and the value
types must be specified at compile-time.

So, this dictionary:

```jinja
{% set data = {'dict': 'of', 'key': 'and', 'value': 'pairs'} %}
```

can be written as:

```scriggo
{% var data = map[string]string{"dict": "of", "key": "and", "value": "pairs"} %}
```

Note that maps in Scriggo, just like in Go, do not keep any information about keys
ordering.

#### Tuples

Scriggo has no equivalent for tuples. In general, you may want to use slices for
sorted values, and maps or structs for values with an associated key.

For example, consider this tuple in Jinja, representing a date:

```jinja
{% set birthday = (1, 5, 2020) %}
```

This can be rewritten in Scriggo using a slice:

```scriggo
{% var birthday = []int{1, 5, 2020} %}
```

or, if you want named values, you can use a map:

```scriggo
{% var birthday = map[string]int{"day": 1, "month": 5, "year": 2020} %}
```

or a struct:

```scriggo
{% type date struct { day, month, year int } %}
{% var birthday = date{day: 1, month: 5, year: 2020} %}
```

### If Expression

While Jinja does have _if expressions_, Scriggo doesn't.

In Scriggo you can use an _if statement_ to assign to a variable conditionally:

```scriggo
{% var title = "bar" %}
{% if 2 + 2 == 4 %}
    {% title = "foo" %}
{% end %}
<h1>{{ title }}</h1>
```

### Math and comparisons

Scriggo and Jinja shares a big part of their operators.

For example this template code is valid both in Jinja and Scriggo:

```
{{ 3 + 10 }} {{ 32 - 2 }} {{ 5 / 2.5 }} {{ 11 % 2 }}
{{ 24 == (2 + 5) }} {{ 1 < 4 }}
```

### Logic

Scriggo, just like Jinja, supports the `and`, `or` and `not` logic operators, but
 with a difference: while in Jinja (like in Python) the `and` and `or` expressions
 are evaluated as one of the two operands, in Scriggo they are always evaluated to a
 boolean value.

Consider this:

```
{{ 0 or "hello" }}
```

Jinja renders:

```
hello
```

while Scriggo renders:

```
true
```

The behavior of the `not` operator is the same in Jinja and Scriggo, thus it always
evaluates to a boolean.

### Operator `in`

While Jinja has the _in_ operator:

```jinja
{{ 1 in [1, 2, 3] }}
{{ 3 not in [1, 2, 3] }}
```

Scriggo has the _contains_ and _not contains_ operators (note that the operands are
reverted respect to the _in_ operator):

```scriggo
{{ []int{1, 2, 3} contains 1 }}
{{ []int{1, 2, 3} not contains 3 }}
```

### Operator `+` (plus)

The operator `+` in Scriggo, just like in Jinja, can also be used to concatenate
strings.

```scriggo
{% var name = "hello" %}
{{ name + ", " + "world! "}}
```

### Operator `~` (tilde)

Jinja has the _tilde operator_ that converts the operands in strings and concatenates them:

```jinja
{% set name = "FooBar" %}
{{ "Hello " ~ name ~ "!" }}
```

In Scriggo this can be achieved using a builtin function like `sprintf`:

```scriggo
{% var name = "FooBar" %}
{{ sprintf("Hello %s!", name) }}
```

If the operands are strings, you can use the [operator `+`](#operator--plus).

### Methods

In Scriggo, just like in Jinja, you can call methods on values using the syntax:

```
{{ value.Method(a1, a2) }}
```

## List of control structures

### For

Scriggo supports the _for loop_ syntax used in Jinja, plus the ones used in Go.

For example this code is valid both in Scriggo and Jinja:

```
<h1>Members</h1>
<ul>
{% for user in users %}
  <li>{{ user.username }}</li>
{% endfor %}
</ul>
```

### If

Scriggo has the _if_ statement with optional _else_.
The condition can be any expression of any type.

Here are some examples:

```scriggo
{% if 2 + 2 == 4 %}
    Well!
{% else %}
    Bad...
{% end if %}
```

```scriggo
{% var primes = []int{2, 3, 5, 7, 11, 13} %}
{% var n = 5 %}
{% if primes contains n %}
    {{ n }} is prime.
{% end if %}
```

### Block Assignments

Scriggo provides a mechanism similar to the block assignments in Jinja, through the
`using` statement and the `itea` builtin identifier.

So, for example, this code in Jinja:

```jinja
{% set navigation %}
    <li><a href="/">Index</a>
    <li><a href="/downloads">Downloads</a>
{% endset %}
```

can be reproduced in Scriggo as:

```scriggo
{% var navigation = itea; using %}
    <li><a href="/">Index</a>
    <li><a href="/downloads">Downloads</a>
{% end using %}
```

Since `itea` is an identifier, it can be used in any expression, so it can also be
passed as argument to builtin functions as you use filters in the block assignments
in Jinja:

```jinja
{% set reply|upper %}
    You wrote:
    {{ message }}
{% endset %}
```

becomes in Scriggo:

```scriggo
{% var reply = toUpper(itea); using string %}
    You wrote:
    {{ message }}
{% end using %}
```

### Calls

Call blocks can be reproduced in Scriggo using the `using` and `show` statements and
the `itea` builtin identifier.

For example this code in Jinja:

```jinja
{% macro render_dialog(title) %}
    <div>
        <h2>{{ title }}</h2>
        <div class="contents">
            {{ caller() }}
        </div>
    </div>
{% endmacro %}

{% call render_dialog('Hello World') %}
    This is a simple dialog rendered by using a macro and
    a call block.
{% endcall %}
```

can be reproduce in Scriggo as:

```scriggo
{% macro render_dialog(title string, contents macro() html) %}
    <div>
        <h2>{{ title }}</h2>
        <div class="contents">
            {{ contents() }}
        </div>
    </div>
{% end macro %}

{% show render_dialog("Hello World", itea); using macro() html %}
    This is a simple dialog rendered by using a macro and
    a call block.
{% end using %}
```

> The `{% show .. %}` statement is equivalent to `{{ .. }}`, but the former is
> mandatory when used in combination with the `using` statement.

### Filters

Filters can be rewritten in Scriggo using builtin functions and syntactic constructs.

See [builtin filters](#list-of-builtin-filters) for some examples.

### Assignments

In Jinja, you can assign values to variables using the `set` statement:

```jinja
{% set navigation = "x.html" %}
{% set navigation = "y.html" %}
{% set key, value = call_something() %}
```

In Scriggo (just like in Go) there are two main differences:

- variables must be declared before being assigned
- when declaring a variable, the `var` keyword is used instead of `set`

So the example above becomes:

```scriggo
{% var navigation = "x.html" %}
{% navigation = "y.html" %}
{% var key, value = call_something() %}
```

> Scriggo (like Go) also supports the short variable declaration operator `:=`, which
> is not covered in this document.

## Template inheritance

Just like Jinja, Scriggo provides a mechanism to implement the template inheritance.

The _extends_ declaration can be used in Scriggo to extend another file, just like in Jijnja:

```
{% extends "layout/default.html" %}
```

## Blocks and macros

Both the _blocks_ and the _macros_ of Jinja can be implemented in Scriggo using the flexible system of _macros_.

For example, consider this template file in Jinja that extends another file:

```jinja
{% extends "base.html" %}
{% block title %}Index{% endblock %}
{% block head %}
    {{ super() }}
    <style type="text/css">
        .important { color: #336699; }
    </style>
{% endblock %}
{% block content %}
    <h1>Index</h1>
    <p class="important">
      Welcome to my awesome homepage.
    </p>
{% endblock %}
```

In Scriggo it becomes:

```scriggo
{% extends "base.html" %}
{% macro Title %}Index{% end macro %}
{% macro Head(super macro()) %}
    {{ super() }}
    <style type="text/css">
        .important { color: #336699; }
    </style>
{% end macro %}
{% macro Content() %}
    <h1>Index</h1>
    <p class="important">
      Welcome to my awesome homepage.
    </p>
{% end macro %}
```

Macros in Scriggo, just like macros in Jinja, can accept parameters.

So, this template in Jinja:

```jinja
{% macro input(name, value='', type='text', size=20) %}
    <input type="{{ type }}" name="{{ name }}" value="{{
        value|e }}" size="{{ size }}">
{% endmacro %}

{{ input("Car", "Red & Blue", size=30) }}
```

can be written in Scriggo as:

```scriggo
{% macro input(name, value, typ string, size int) %}
    {% if typ == "" %}{% typ = "text" %}{% end if %}
    {% if size == 0 %}{% size = 20 %}{% end if %}
    <input type="{{ typ }}" name="{{ name }}" value="{{
        value }}" size="{{ size }}">
{% end macro %}

{{ input("Car", "Red & Blue", "", 30) }}
```

As you may have noticed, there are some main differences:

- Scriggo macros declarations, just like function declarations in Go, does not
  provide a default mechanism for parameters, so their value must be explicitly set
  in the body of the macro.
- The types of the parameters, in Scriggo, must be explicitly declared. The compiler
  will check if the types of the argument passed to the macro call match the types of
  the parameters.
- Scriggo, just like Go, does not have named parameters.

## Including other template files

Jinja uses the _include_ statement:

```jinja
{% include 'header.html' %}
```

Scriggo uses a more flexible system making available the _render_ operator that reads
a template file and returns its rendered content.

So, the example above may be written as:

```scriggo
{{ render "header.html" }}
```

The value returned by _render_ can be used in any expression:

```scriggo
{% var content = render "myfile.html" %}
File is {{ len(content) }} bytes long.
```

## Importing other template files

Scriggo has the _import_ declaration, just like Jinja, but with some differences.

There are several form of _import_ available.

To import a template file specifying the declarations to import:

```scriggo
{% import "path/to/file.html" for Title, Content %}
{{ Title() }}
{{ Content() }}
```

To import every (exported) declaration from a template file:


```scriggo
{% import "path/to/file.html" %}
{{ Title() }}
{{ Content() }}
```

To import a template file, prepending a prefix to the imported declarations:

```scriggo
{% import myfile "path/to/file.html" %}
{{ myfile.Title() }}
{{ myfile.Content() }}
```

## List of Builtin filters

To get an overview about filters, see the [corresponding section](#filters).

Here are just some examples:

### `abs`

Jinja:

```jinja
{{ -3|abs }}
```

Scriggo has the builtin function `abs`:

```scriggo
{{ abs(-3) }}
```

### `attr`

Jijna

```jinja
{{  foo|attr("bar") }}
```

In Scriggo you can access to attributes using a selector, just like in Jinja:

```
{{  foo.bar }}
```

Note that if `foo` has no attribute `bar`, it results in a compilation error.

> In general, attributes names in Scriggo begins with an uppercase letter: that means
> that such field is exported. In Scriggo, just like in Go, only exported fields can
> be accessed by other packages from the one that declared such type.

### `length`

Scriggo offer the builtin function `runeCount` that returns the number of characters
of a string.

```scriggo
{{ runeCount("hellò") }}
```

renders to

```
5
```

Scriggo also offers the builtin function `len` that returns the number of elements in
slices and maps and the number of bytes in strings.

```scriggo
{{ len([]int{10, 20, 30}) }}
{{ len("hellò") }}
```

renders to:

```
3
6
```

## HTML Escaping

In Jinja the escaping can be achieved both manually (using the
[e](https://jinja.palletsprojects.com/en/3.0.x/templates/#jinja-filters.escape)
filter) or automatically.

In Scriggo the code is always escaped before being rendered in a given context.

For example:

```scriggo
{{ "<b>Not real bold...</b>" }}
```

renders as:

```
&lt;b&gt;Not real bold...&lt;/b&gt;
```

This happens because the rendered value has type [string](/templates/types#strings), so it must be escaped before being
rendered in a _HTML_ context.

For this reason, if you write something like:

```scriggo
{{ html("<b>Real bold!</b>") }}
```

this is rendered as:

```
<b>Real bold!</b>
```

because the value type is explicitly set to the `html` type, so Scriggo won't escape it.

### Automatic contextual escaping

Scriggo also offers an automatic contextual escaping.
For example consider this:

```scriggo
{% var value = "< is more" %}
<p>{{ value }}</p>
<a href="{{ value }}"></a>
```

It is rendered as:

```
<p>&lt; is more</p>
<a href="%3c is more"></a>
```

As you may notice, the string `< is more` is automatically escaped using two
different algorithms:

* in the context of HTML it becomes: `&lt; is more`.
* in the context of an attribute value, that is automatically detected by the Scriggo
  parser, it is escaped as `%3c is more`.

## Continue to learn

Now you can checkout this resources:

- [Templates](/templates), which provides a more depth overview of the Scriggo
  template.
- [Get started with templates](/get-started-2#execute-templates-in-your-application),
  which provides a tutorial on embedding Scriggo templates in your Go application.
- [The Scriggo Template Specification](/templates/specification), which expands the [Go
  language specification](https://golang.org/ref/spec) covering every aspect of the
  Scriggo template in detail.

{% end raw document %}