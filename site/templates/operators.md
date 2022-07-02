{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: operators{% end %}
{% Article %}

{% raw doc %}

# Operators

Scriggo offers several arithmetic, comparison and logic operators. Since Scriggo is strongly typed, each operator can
be used only with specific types, and for binary operators the two operands should have the same type.

As an exception, the three logical operators `and`, `or` and `not` accept any type, and the operands of the `contains`,
`not contains` and `shift` operators have specific types but can be different.

## Basic operators

These are the basic operators. They are the Go operators and therefore are also available in programs.

<div class="operators">

|     |                     |
|-----|---------------------|
| ==  | equal               |
| !=  | not equal           |
| <   | less                |
| <=  | less or equal       |
| >   | greater             |
| >=  | greater or equal    |
| \|\|| conditional OR      |
| &&  | conditional AND     |
| !   | NOT                 |
| +   | sum                 |
| -   | difference          |
| *   | product             |
| /   | quotient            |
| %   | remainder           |

</div>

For example:

```scriggo
{% if stock > 10 %}
  Wide availability
{% end %}
```

There are other operator, but they are not discussed in this documentation (see the
[Go specification](https://go.dev/ref/spec#Operators)). 

The following table shows which operators can be used with which type:

<div class="data-type-operators">

|  Type       |  Operators                       |
|-------------|----------------------------------|
| Numeric     | ==  !=  <  <=  >  >=  +  -  *  / |
| Integer     | %                                |
| String      | ==  !=  <  <=  >  >=  +          |
| Boolean     | ==  !=  !  && \|\|             |

</div>

## And, Or and Not

Scriggo allows to combine any expression with the `and`, `or` and `not` operators and their evaluation returns true or
false based on the truth value of their operands.

```scriggo
{% if promotion and stock %} promotion with immediate availability {% end %}

{% if not price %} login to view price {% end %}
```

The Go language does not have `and`, `or` and `not`, but they are available in Scriggo templates because they make
simple `if` statement conditions easier for non-programmers to read and write.

As an alternative you can use the Go logical operators `&&`, `||` and `!` that want boolean values as operands. For
example, the above examples can be written as:

```scriggo
{% if promotion != nil && stock > 0 %} promotion with immediate availability {% end %}

{% if price == 0  %} login to view price {% end %}
```

### "true" or "false"

For the `and`, `or` and `not` operators all values are true, except for the following values which are false:

<div class="data-type-false">

|  Type       |  "false" value  |
|-------------|-----------------|
| bool        | false           |
| int         | 0               |
| float64     | 0.0             |
| string      | ""              |
| []string    | nil or empty    |
| Product     | Product{}       |

</div>

See the [Scriggo template specification](/templates/specification#truthful-values) for more details.

## Contains

The _contains_ and _not contains_ boolean operators indicates if a slice contains or not contains an element, if a map
contains or not contains a key, if a string contains or not contains a substring or a rune (character).

These operators have the following form:

```scriggo
a contains b
a not contains b
```

The _contains_ operators are specific to Scriggo templates and are not present in the Go language.

### Examples

Checks if a slice of strings contains a specific string:

```scriggo
{% var colors = []string{"red", "blue", "yellow", "green"} %}
{% if colors contains "yellow" %}
  colors contains "yellow"
{% end %}
```
<pre class="result">  colors contains "yellow"</pre>

Checks if a string contains a substring:

```scriggo
{% if product.Name contains "bundle" %}
  the product's name contains the word "bundle"
{% end %}
```
<pre class="result">  the product's name contains the word "bundle"</pre>

Checks if a map contains a key:

```scriggo
{% var nameOf = map[int]string{1: "one", 4: "four", 7: "seven"} %}
{% if nameOf contains 7 %}
  nameOf contains the key 7
{% end %}
```
<pre class="result">  nameOf contains the key 7</pre>

{% end raw doc %}
