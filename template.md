{% extends "/layouts/article.html" %}
{% macro Title %}Scripts{% end %}
{% Article %}

{% raw code %}

# Scriggo templates

Scriggo is a modern and powerful template engine for Go, supporting inheritance,
macros, partials, imports and autoescaping but most of all it uses the Go language as the
template scripting language. 

```
{% extends "layout.html" %}
{% import "banners.html" %}
{% macro Body %}
    <ul>
      {% for _, product := range products %}
      <li><a href="{{ product.URL }}">{{ product.Name }}</a></li>
      {% end %}
    </ul>
    {{ render "pagination.html" }}
    {{ Banner() }}
{% end %}
```

Scriggo template files can be written in plain text, HTML, Markdown, CSS, JavaScript and JSON.

## Expressions

Scriggo supports all Go expressions. To evaluate an expression and render its result, put the expression between `{{` and `}}`:

```
{{ expression }}
```
 
Examples:

```
Page has title "{{ capitalize(page.Title) }}" and shows {{ len(products) + 1 }} products
```

`{{ }}` escapes the printed value based on its context.

In addiction, Scriggo templates support the `render` operator to render a template file, the `contains` operator and the `default` expression statement.

## Statements

Scriggo supports in templates almost all Go statements. A statement is delimited between `{%` and `%}`:

```
{% statement %}
```

Examples:

```
{% var a = 5 %}

{% a = 7 %}

{% a, ok = b.(int) %}

{% defer func() { } %}

{% go func() { } %}  
```

`if`, `switch`, `select` and `for` statements require a closing `end`.

Example:

```
{% if n > 5 %} n is greater than 5 {% end %}  
```

`end` can be followed by the statement's name that it closes.

Example:

```
{% for i, product := range products %}
  {{ i }}. {{ product.Name }}
{% end for %}
```

In addiction, Scriggo supports in templates six new statements: `extends`, `import`,
`macro`, `show`, `raw` and `using`.

### Functions

Function declarations are by design not supported by Scriggo in templates, use instead what is more appropriate between a macro:

```
{% macro image(url string) %}<img src="{{ url }}">{% end %}
``` 

and a function literal:

```
{% inc := func(i int) int { return i + 1 } %}
```

## Simpler syntax

Scriggo also support a simpler and relaxed syntax for a gentle introduction to templates.

The `for` `range` statement:

```
{% for _, product := range products %}
```

can be written simply as:

```
{% for product in products %}
```

Furthermore the condition of the `if` statement can have any type. If it valuates to the zero value of its type, it is false, otherwise it is true. 

So the following statements:

```
{% if price > 0 %} ... {% end %}
{% if producer != nil %} ... {% end %}
```

can be written simply as:

```
{% if price %} ... {% end %}
{% if producer %} ... {% end %}
```

As a special case, if the condition has an interface type, the condition is false if the interface is `nil` or its dynamic
value is the zero value for its type.

Finally there are tree new operators: `and`, `or` and `not`. They are relaxed forms of the `&&`, `||` and `!` operators
where the operands, as the `if` condition, can have any type.

So the statement:

```
{% if discount > 0 && banners != nil %} ... {% end %}
```

can be written simply as:

```
{% if discount and banners %} ... {% end %}
```

### extends

The statement `extends` extends a template file from another one, the parent, implementing template inheritance.
A parent file contains calls to macros declared in the extended files. 

Example:

```
{# parent.html #}
{{ Banner() }}
```

```
{% extends "parent.html" %}
{% macro Banner %} ... {% end %}
```

This is a more complete example of a parent file:

```
<!DOCTYPE html>
<html>
<head>
  <title>{{ Title() }}</title>
</head>
<body>
  {{ Header() }}
  <div>
    {{ Column() }}
    {{ Body() }}
  </div>
  {{ Footer() }}
</body>
</head>
```

and an extended file:

```
{% extends "layout.html" %}
{% macro Title %}Awesome article{% end %}
{% macro Header %}
  ...
{% end %}
{% macro Body %}
  ...
{% end %}
{% macro Footer %}
  ...
{% end %}
```

### macro

The statement `macro` defines a macro, which is a way to repeat frequently-used template code. A macro can have
input parameters, like a function, but have a single unnamed return parameter with type `string`, `html`, `markdown`, `css`, `js`, or `json`. In a macro declaration the return parameter type can be omitted, in this case it is inferred by the context. A macro without input parameters can be declared and called without the parentheses. 

Examples:

```
{% macro title %}A wonderful book{% end %}

{% macro image(url string, width, height int) html %}
  <img src="{{ url }}" width="{{ width }}" height="{{ height }}">
{% end %}
```

A macro is a function and as such it is called with the parentheses:

```
{{ name() }}
 
{% show image("picture.jpg", 400, 500) %}
```

### Partials

Partials are template files whose contents can be rendered in other files with the `render` operator:

```
{{ render "partials/footer.html" }}
```

The name of the partial file can be a relative or an absolute path. The partial file does not inherit the variables
declared in the file in which it is rendered.

### import

The statement `import` imports the exported declarations in a template file in the current file scope. The exported declarations
are the macro, variable, constant and type declarations with the first letter of the declaration's name in uppercase.

If a template file with the given path does not exist, `import` tries to import the declarations from a package with the given path.

Scriggo supports in templates all the supported Go `import` forms plus a new "import-for" form that allows to indicate the names to import. For example:

```
{% import "elements.html" for Banners, Menus %}
{{ Banners() }}
{{ Menus() }}
```

The Go `import` form without a package name is only used to import packages and not template files:

```
{% import "fmt" %}
```

Import statements can be only at the beginning of a file and can be preceded only by an extends statement.   

Examples:

```
{% import "catalog.html" %}
<html>
<head>
  <title>Catalog</title>
</head>
<body>
  {{ Products() }}
</body>
</html>
```

```
{% extends "layout.html" %}
{% import elems "elements.html" %}

{% macro Header %}
  <div>{{ elems.Banners }}</div>
{% end %}

{% macro Body %}
  ...
{% end %}
```

### raw

The raw statement is used to render a content without executing Scriggo code:

```
{% raw %} {{ price }} {% end %}
```

Terminate the `raw` statement with `{% end %}`, `{% end raw %}` or `{% end raw example %}` where `example` can be any identifier.

```
{% raw example %}
Use {% raw %} ... {% end raw %} to render a content as is.
{% end raw example %}
```

It is rendered as `Use {% raw %} ... {% end raw %} to render a content as is`.

### using

The `using` statement can be used between `{%` and `%}` in conjunction with show statements, expression statements, send statements, assignments and variable declarations to capture the rendering of a content into a value.

The value is assigned to the predeclared identifier `itea` and has the type that follow `using`. The type can be a format type (`string`, `html`, `markdown`, `css`, `js` or `json`) or a macro type. If it is omitted, it is inferred by the context.

```
{% var link = itea; using %}
  <a href="{{ url }}">{{ name }}</a>
{% end %}

{% reply(itea); using string %}
  Dear {{ name }},
  thanks for your interest.
{% end using %}
```

When a using statement has a macro type, the content if rendered every time the macro is called:

```
{% show Head(itea); using macro %}
    <title>{{ Title() default "" }} - Awesome site</title>
    <link href="style.css">
{% end using %}
```
```
{% extends "layout.html" %}
{% macro Title %}Contact us{% end %}
{% macro Head(content macro() html) %}
    {{ content() }}
    <script src="contact-form.js"></script>
{% end macro %}
```

### Comments

Text between `{#` and `#}` is a comment and it is discarded during the rendering.

```
{# comment #}
```

Example:

```
{# show the user's name #}
{{ user.FirstName }} {{ user.LastName }}

{# this code requires a review: 
  {% for _, user := range users %}
     ...
  {% end %}
#}
```

{% end raw code %}
