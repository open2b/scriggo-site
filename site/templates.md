{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates{% end %}
{% Article %}

{% raw code %}

# Scriggo templates

Scriggo templates supports inheritance, macros, partials, imports and contextual autoescaping but most of all it uses
the Go language as the template scripting language.

## Syntax

The following is a minimal example of a template file written with Scriggo:

```
{% extends "layout.html" %}
{% import "banners.html" %}
{% macro Body %}
    <ul>
      {% for product in products %}
      <li><a href="{{ product.URL }}">{{ product.Name }}</a></li>
      {% end %}
    </ul>
    {{ render "pagination.html" }}
    {{ Banner() }}
{% end %}
```

The double curly braces:

```
{{ product.Name }}
```

shows the content of a variable and in general the result of the evaluation of an expression.

Single curly braces with a percent sign:

```
{% if product.Stock > 10 %} good availability {% end %}
```

are statements or declarations that allow you to define the layout, execute repeated or condition-based code, and
declare and assign variables (and also define constants and types).

### Multiline statements

Single curly braces with a double percent sign:

```
{%%
    for product in products {
        if product.Stock == 0 {
            continue
        }
        if product.Stock > 10 {
            show "good availability: "
        }
        show product.Name, "\n"
    }
%%}
```

is called a multiline statement and can contain multiple statements and declarations. The syntax is the same as the Go
language, apart from some specific Scriggo constructs.

### Comments

Single curly braces with a single hash sign:

```scriggo
{# review the following code #}
```

are comments. Comments are discarded when the template is rendered and consequently are not present in the resulted
code. Comments in Scriggo can be nested:

```scriggo
{#
    {# show the promotion #}
    Promotion: {{ promotion }}
#}
``` 

## Statements and declarations

There are several statements and declarations that allow you to structure your template code in different parts that
can be reused, as [show](/templates/show-render#show), [macro](/templates/macro),
[extends](/templates/extends-and-import#extends) and [import](/templates/extends-and-import#import).

The [if](/templates/if-for-switch#if) statement allows you to execute a block of code based on a condition and the
[for](/templates/if-for-switch#for) statement allows you to repeatedly execute a block of code, such as a list of 
articles. There is also the [switch](/templates/if-for-switch#switch) statement.

The var keyword allows you to declare new [variables](/templates/variables) and the
[assignment](/templates/variables#assignment) operators, as =, allow you to change their value.

## How learn more about templates

To learn more about templates continue to read the documentation and refer to the [specification](/templates/specification) for more details.  

{% end raw code %}
