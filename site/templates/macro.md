{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the macro declaration{% end %}
{% Article %}

{% raw doc %}

# Macro

A macro is a piece of code that is given a name so that it can be called in another part of the template. A _macro_
declaration is used to declare a macro. For example:

```scriggo
{% macro title %}An awesome book{% end %}
```

declares a macro named _title_. A macro is a function and as such parenthesis are used to call it:

```scriggo
{{ title() }}
```
<pre class="result">An awesome book</pre>

In the body of the macro, you can access to, other than global variables, variables and macros declared in the same
file before the macro declaration. For example:

```scriggo
{% var product = "tablet" %}
{% macro title %}An awesome {{ product }}{% end %}
```
<pre class="result">An awesome tablet</pre>

A variable declared in a body of a macro is not visible outside the macro.

## Macro with parameters

A macro can have parameters as any other function. Parameters are declared in parentheses, separated by a comma and
the type must be indicated in addition to the name of the parameter.

The following declaration declares a macro, called "image", with three parameters:   

```scriggo
{% macro image(url string, width int, height int) %}
  <img src="{{ url }}" width="{{ width }}" height="{{ height }}">
{% end %}
```

A macro should be called with the same number of arguments, in the same order and with the same type of the parameters
in the declaration.

```scriggo
{{ image("picture.jpg", 400, 500) }}
```
<pre class="result">
&lt;img src="picture.jpg" width="400" height="500"&gt;
</pre>

A type can be omitted if the next parameter has the same type. The previous declaration can be rewritten as:

```scriggo
{% macro image(url string, width, height int) %}
  <img src="{{ url }}" width="{{ width }}" height="{{ height }}">
{% end %}
```

## Distraction free declaration

If a macro is declared in a file with an extends declaration, to be called in the extended file, and the macro has no
parameters, as in this example:

```scriggo
{% extends "layout.html" %}
{% macro Main %}

<ul>
  {% for product in products %}
  <li><a href="{{ product.URL }}">{{ product.Name }}</a></li>
  {% end for %}
</ul>

{% end macro %}
```

the declaration can be written in a special form called "distraction free". The following example is equivalent to the
previous one but uses a distraction free macro declaration:

```scriggo
{% extends "layout.html" %}
{% Main %}

<ul>
  {% for product in products %}
  <li><a href="{{ product.URL }}">{{ product.Name }}</a></li>
  {% end for %}
</ul>
```

In a distraction free declaration, the `macro` keyword and the terminating `{% end %}` or `{% end macro %}` are omitted
and the macro ends at the end of the file. It means that all the code after the declaration end up to the end of the
file is the body of the macro.

Note that in Scriggo the code `{% Ident %}`, where `Ident` is an exported identifier (the first letter is in uppercase),
is always a distraction free macro declaration.

## Call a macro declared in another file

To call a macro declared in another file, such as a file that declares macros useful in various parts of the template,
you have to import the file with an import declaration.

```scriggo
{% import "/imports/images.html" for Image %}
{{ Image("offer.png", 200, 200) }}
```
<pre class="result">&lt;img src="offer.png" width="200" height="200"&gt;</pre>

Macro declared in other files should have the first letter in uppercase in order for them to be imported.

{% end raw doc %}
