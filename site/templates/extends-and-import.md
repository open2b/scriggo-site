{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the extends and import declarations{% end %}
{% Article %}

{% raw doc %}

# Extends and Import

The [extends](#extends) and [import](#import) declarations allow you to structure the template in order to write some
code once and then call it in different parts of the template.

## Extends

The _extends_ declaration allows a file, that contains some content, to extend another file that defines its layout.
Such file is called extended file or layout file.

The following is an example of layout file:

```scriggo
<!DOCTYPE html>
<html>
<head>
  <title>{{ Title() }}</title>
</head>
<body>
  {{ Header() }}
  {{ Body() }}
</body>
</head>
```

It contains show statements that call the macros `Title`, `Header` and `Body`. Note that a call to a macro looks
like a function call, and this is because macros are functions. Declaration syntax is different.

These three macros are not declared in the file itself, but they will be declared in files that extend the layout file:

```scriggo
{% extends "/layouts/standard.html" %}
{% macro Title %}Awesome product{% end %}
{% macro Header %}
 <header>An awesome product</header>
{% end %}
{% macro Body %}
  This awesome product is...
{% end %}
```
<pre class="result">
&lt;!DOCTYPE html&gt;
&lt;html&gt;
&lt;head&gt;
  &lt;title&gt;Awesome product&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
  &lt;header>An awesome product&lt;/header&gt;
  This awesome product is...
&lt;/body&gt;
&lt;/head&gt;
</pre>

The extends declaration should be at the beginning of the file, before others declarations, and the file declares macros
called but not declared in the layout file. These macros should have the first letter in uppercase to be visible and
callable in the extended file.

The page can have other declarations, such as variables and macros, used in the same file, but it cannot have any content
other than the one contained within the macros.

The previous example used only macro with no parameters, but there is nothing preventing the use of macro with
parameters.

## Import

The _import_ declaration imports declarations from another file into the current file, so you can access its variables
or call its macros. Only the exported declarations (declarations with the first letter in upper case) are imported.

Imported files can only contain declarations and cannot have content, other than that inside the macros. They define
macros and variables that can be referred in other parts of the template by importing their file.

For example, if the file "images.html" defines the following macro:

```scriggo
{% macro Image(url string, width, height int) %}
  <img src="{{ url }}" width="{{ width }}" height="{{ height }}">
{% end %}
```

after importing the file, it is possible to call the macro:

```scriggo
{% import "images.html" %}
{{ Image("offer.png", 200, 200) }}
```
<pre class="result">&lt;img src="offer.png" width="200" height="200"&gt;</pre>

Import declarations should be at the beginning of the file, before anything else, but after an extends declaration if
it is present.

### Import for

As an alternative, you can also indicate the names you want to import into the current file from the other file. In
this case only these names will be imported.

```scriggo
{% import "images.html" for Image, Banner %}
{{ Image("offer.png", 200, 200) }}
```

### Import prefix

As an alternative, the import declaration allows you to indicate an identifier to be used as a prefix to access the
declarations in the imported file.

```scriggo
{% import images "images.html" %}
{{ images.Image("offer.png", 200, 200) }}
```

{% end raw doc %}
