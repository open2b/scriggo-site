{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the if, for and switch statements{% end %}
{% Article %}

{% raw doc %}

# If, For and Switch

The [if statement](#if) is used to execute code based on a condition, the [for statement](#for) is used to iterate over
a [slice](/templates/types#slice) (but also maps, strings and channels) and the [switch statement](#switch) allows a
variable to be tested for equality against a list of values.

## If

The `if` statement executes its body if the condition, that can have any type, is valuated to true. For example the
following statement:

```scriggo
{% if stock > 10 %}
  Wide availability
{% end %}
```

shows "Wide availability" if _stock_ is greater than 10.

> Note that in programs, unlike in templates, the _if_ condition can only be a boolean value.

A condition is false for the values `false`, `0`, `0.0`, `""` and `nil`, for nil or empty slices and maps. Apart from [some
special cases](/template-spec#truthful-values), it is true for all other values.

For example:

```scriggo
""               // is false, it is an empty string
"0"              // is true
true             // is true
[]string{}       // is false, it is an empty slice
map[int]string{} // is false, it is an empty map
0.1              // is true 
[]int{1,2}       // is true
``` 

If you want to execute one code or another based on a condition, you can write:

```scriggo
{% if stock > 10 %}
  Wide availability
{% else %}
  Limited availability
{% end %}
```
<pre class="result">  Wide availability</pre>

Scriggo also supports the _else if_ form:

```scriggo
{% if stock > 10 %}
  Wide availability
{% else if stock > 0 %}
  Limited availability
{% else %}
  Not available
{% end %}
```
<pre class="result">  Wide availability</pre>

## For

The `for` statement iterates over a [slice](/templates/types#slice) (but also a maps, a strings and a channel ).
For example:

```scriggo
{% for article in articles %}
  <div>{{ article.Title }}</div>
{% end %}
```

It assigns an element to a variable in each iteration. In the previous example, the `article` variable is implicitly
declared and has the same type of the elements of `articles`. It is only visible in the body of the `for` statement.

This _for_ statement form can be used only in templates, but you can also use all the other forms of the _for_
statement that can be used in programs.

### for range

The _for range_ form is similar to the _for in_ form, but it also allows you to assign the element index, starting
from zero:

```scriggo
{% for i, article := range articles %}
  <div>{{ i+1 }}. {{ article.Title }}</div>
{% end %}
```

### for condition

The _for_ statement can be used with a condition. In this case the body is executed as long as the condition is true.

```scriggo
{% for i < len(articles) %}
  <div>{{ i+1 }}. {{ articles[i].Title }}</div>
  {% i++ %}
{% end %}
```

As with the _if_ statement, the condition of the _for_ statement can have any type. 

### for loop

The _for_ statement alone, without a condition, executes its body until it is interrupted with a _break_ statement.

```scriggo
{% for %}
  {% var v = value() %}
  {% if not v %}{% break %}{% end %}
{% end %}
```

### for init; condition; post

The _for init; condition; post_ form has three parts separated by a semicolon. The first part declares or assigns to a
variable before the first iteration. The second part is the condition and the last part usually increases a variable 
before the next iteration.

The iteration continues as long as the condition is true.

```scriggo
{% for i := 0; i < len(articles); i++ %}
  <div>{{ i+1 }}. {{ articles[i].Name }}</div>
{% end %}
```

### continue

The _continue_ statement can be used in the body of a _for_ statement to terminate the current iteration and
continue with the next one.

For example, the following code shows only products with a price:

```scriggo
{% for product in products %}
  {% if not product.Price %}
    {% continue %}
  {% end if %}
  <div>{{ product.Name }}</div>
{% end %}
```

### break

The _break_ statement can be used to stop the execution of a _for_ statement.

For example, the following code shows only `1 2 3`.

```scriggo
{% for n in []int{1, 2, 3, 4, 5} %}
  {% if n > 3 %}
    {% break %}
  {% end if %}
  {{ n }}
{% end %}
```

## Switch

The _switch_ statement can be used as an alternative to the _if_ statement when the condition is not simply true or
false. This statement executes a _case_ based on the value of its condition.

```scriggo
{% switch department.Name %}
{% case "Tappeti" %}
  Rugs and rugs on offer for yuor home
{% case "Cuscini", "Federe" %}
  Offers for the bedroom
{% default %}
  Specials
{% end %}
```
<pre class="result">Offers for the bedroom</pre>

Case values must have the same type of the condition. The cases are evaluated in order and only the first case that has
the same value as the condition is executed. If there are no cases with that value and there is a _default_ case, the
default case is executed.

A switch may or may not have a condition. If it doesn't, the case expressions must be boolean.

```scriggo
{% switch %}
{% case stock > 10 %}
  Available
{% case name == "Cuscini" %}
  Offers for the bedroom
{% default %}
  Specials
{% end %}
```

{% end raw doc %}
