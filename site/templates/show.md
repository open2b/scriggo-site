{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the show statement{% end %}
{% Article %}

{% raw doc %}

# Show

The _show_ statement followed by one or more expressions, evaluates them and displays the results:

```scriggo
{% show 5 + 2, " = ", 7 %}
{% show price %}
```
<pre class="result">
7 = 7
$29,99
</pre>

A show statement with a single expression can also be written as follows:

```
{{ 5 + 2 }}
{{ price }}
```
<pre class="result">
7
$29,99
</pre>

This second form is the most used because it is easier to write and read, you can still use one or the other indifferently.

Between `{%%` and `%%}` the show statement can be used as follows:

```
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

The show statement shows the expressions based on the context in which it is used:

```
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

{% end raw doc %}
