{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the show statement and the render operator{% end %}
{% Article %}

{% raw doc %}

# Show and Render

The [show](#show) statement renders values and shows them, while the [render](#render) operator renders a template file.
The render operator is usually used in a show statement to show the rendered file.

## Show

The _show_ statement, followed by one or more expressions, renders the evaluation of expressions and shows the results.

```scriggo
{% show 5 + 2, " = ", 7 %}
{% show price %}
```
<pre class="result">
7 = 7
$29.99
</pre>

A show statement with a single expression can also be written as follows:

```scriggo
{{ 5 + 2 }}
{{ price }}
```
<pre class="result">
7
$29.99
</pre>

This short form, for a single expression, is the most used because it is easier to write and read, but you can use
one or the other indifferently.

Between `{%%` and `%%}` the show statement can be used as follows:

```scriggo
{%%
  var value = 55
  if value < 100 {
      show "value is ", value
  } else {
      show "too large"
  }
%%}
```
<pre class="result">value is 55</pre>

The show statement shows a value based on its type and the current context:

```scriggo
{% var greeting = "hello" %}

<div>{{ greeting }}</div>

<script>
  var a = {{ greeting }};
  var b = '{{ greeting }} world';
</script>
```
<pre class="result">
&lt;div&gt;hello&lt;/div&gt;

&lt;script&gt;
  var a = "hello";
  var b = 'hello world';
&lt;/script&gt;
</pre>

## Render

The _render_ operator, when used with a _show_ statement, for example between `{{` and `}}`:

```scriggo
{{ render path }}
```

renders the template file with the named path and shows its content. A file rendered with render is called _partial_
because contains only a part of the content that will be rendered:

```scriggo
{{ render "/promotion.html" }}
```

The path of the partial file can also be relative to the file that contains the render operator:

```scriggo
{{ render "../header.html" }}
{{ render "socials.html" }}
```

When a partial file is rendered, the code of the file does not see the variables declared in the file that contain the
render operator, and it does not see the other declarations in this file. For this purpose, macros with parameters are
used, such as:

```scriggo
{{ Image("picture.jpg", 400, 500) }}
```

See the [macro statement](macro) for details.

### More general use

Render is a Scriggo operator and as such it can be used in any expression like any other operator, not only with a
_show_ statement. For example, it is possible to assign its evaluation for later use:

```scriggo
{% var email = render "email.html" %}
```

or use it as part of another expression:

```scriggo
{%%
    if discount {
        show "Offer: " + render "offer.html"
    }
%%}
```

Its evaluation results in a string value. The specific type of the string depends on the format of the rendered file.
For example, if the rendered file is in HTML format, the resulting string will have the type _html_.

### Non-existent files

If the file to render does not exist, a compilation error occurs. On the other hand, if you want to handle the case in
which the file does not exist, you can use a _default expression_:

```scriggo
{% promo := render "extra.html" default "oops!" %}
```

If a template file with path "extra.html" exists, `promo` will contain the result of its rendering, otherwise it will be
`"oops!"`.

To show a rendered template file but do nothing if it does not exist, you can write:

```scriggo
{{ render "specials.html" default "" }}
```

You can use any expression as right operand of a _default expression_, for example you can call a function or macro:

```scriggo
{{ render "specials.html" default notify("specials.html does not exist") }}
```

or you can render another template file:

```scriggo
{{ render "specials.html" default render "no-specials.html" }}
```

{% end raw doc %}
