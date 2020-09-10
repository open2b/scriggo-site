---
layout: article
---

{% raw %}

# Scriggo Template

Scriggo Template is a modern and powerful template engine for Go, supporting inheritance,
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
    {% show "pagination.html" %}
    {% show Banner() %}
{% end %}
```

Scriggo Template's files can be written in plain text, HTML, CSS, JavaScript and JSON.

## Expressions

Scriggo Template supports all Go language expressions. To evaluate an expression and render its result, put the expression between `{{` and `}}`:

```
{{ expression }}
```
 
Examples:

```
{{ page.Title }}

{{ len(products) }}

{{ capitalize(name) }}

{{ n + 2 }}

{{ <-ch }}

{{ func(i int) { return i + i }(7) }}  

{{ map[string]int{ "a" : 2, "b" : 7 } }} 
```

`{{ }}` escapes the printed value based on its context.

## Statements

Scriggo Template supports almost all Go language statements. A statement is delimited between `{%` and `%}`:

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

In addiction, Scriggo Template supports four new statements that are used for template layouts: `extends`, `import`,
`macro` and `show`.

### Functions

Function declarations are by design not supported by Scriggo Template, use instead what is more appropriate between a macro:

```
{% macro image(url string) %}<img src="{{ url }}">{% end %}
``` 

and a function literal:

```
{% inc := func(i int) int { return i + 1 } %}
```

## Simpler syntax

Scriggo Template also supports a simpler and relaxed syntax for a gentle introduction to templates.

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
{% show Banners %}
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
  <title>{% show Title %}</title>
</head>
<body>
  {% show Header %}
  <div>
    {% show Column %}
    {% show Body %}
  </div>
  {% show Footer %}
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
parameters, like a function, but no return values. A macro without parameters can be declared and called without the parentheses. 

Examples:

```
{% macro title %}A wonderful book{% end %}

{% macro image(url string, width, height int) %}
  <img src="{{ url }}" width="{{ width }}" height="{{ height }}">
{% end %}
```

The statement `show` is used to call a macro:

```
{% show name %}
 
{% show image("picture.jpg", 400, 500) %}
```

### Partials

Partials are template files whose contents can be shown in other files with a `show` statement:

```
{% show "partials/footer.html" %}
```

The name of the partial file can be a relative or an absolute path. The partial file does not inherit the variables
declared in the file in which it is shown.

### import

The statement `import` imports the exported declarations in a file in the current file scope. The exported declarations
are the macro, variable, constant and type declarations with the first letter of the declaration's name in uppercase.

An import name the path, relative or absolute, of the imported file and can name an identifier to be used as prefix for access the declarations: 

```
{% import "path" %}
{% import prefix "path" %}
```

Examples: 

```
{% import "elements.html" %}
{% show Banners %}
```

```
{% import elems "elements.html" %}
{% show elems.Banners %}
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
  {% show Products %}
</body>
</html>
```

```
{% extends "layout.html" %}
{% import elems "elements.html" %}

{% macro Header %}
  <div>{% show elems.Banners %}</div>
{% end %}

{% macro Body %}
  ...
{% end %}
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

{% endraw %}
