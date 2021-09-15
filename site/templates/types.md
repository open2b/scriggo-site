{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: types{% end %}
{% Article %}

{% raw doc %}

# Types

Scriggo implements all [types of Go](https://golang.org/ref/spec#Types) but only the most common types
are documented here.

All values in Scriggo have a type and there are several basic types.

## bool

Boolean values are the `true` and `false` values. These are examples of boolean variable declarations:

```scriggo
{% var a bool %}     a is false
{% var b = true %}   b is true
{% c := false %}     c is false
```

## string

Strings are written with double quotes, as "hello", or with a grave accent, as \`hello\`. These are examples of string variable declarations:

```scriggo
{% var a string %}     a is an empty string
{% var b = "hello" %}  b is the string "hello"
{% c := "world" %}     c is the string "world"
```

To read the length in bytes of a string use the `len` function, instead to read the length in characters use the
`runeCount` function. The name of the `runeCount` function comes from "rune" which is the term used in Go to name the
characters. For a more accurate explanation you can read
[Strings, bytes, runes and characters in Go](https://go.dev/blog/strings).

## html

The _html_ type is a string type representing HTML code. Unlike the type string, _show_ does not escape a html value in 
an HTML context. For example:

```scriggo
<div>
    {{ "<b>this is bold</b>" }}           {# is escaped #}
    {{ html("<b>this is bold</b>") }}     {# is not escaped #}
</div>
```
<pre class="result">
&amp;lt;b&amp;gt;this is bold&amp;lt;/b&amp;gt;
&lt;b&gt;this is bold&lt;/b&gt;
</pre>

A value with type string can be converted to the html type only if it is an untyped constant, as `"<b>this is bold</b>"`. 
If it is a variable, it cannot be converted. On the other hand, a html type value can always be converted to the
string type. 

```scriggo
{%%
    var a html = "<b>this is bold</b>"  // "<b>this is bold</b>" can be converted to html
    var b string = a                    // a cannot be coverted to b
    var c html = b                      // ERROR: b is not an untyped constant
%%}
```

## markdown

The _markdown_ type is a string type representing Markdown code. Unlike the type string, _show_ does not escape a 
markdown value in a Markdown context. For example, supposing a Markdown context:

```scriggo 
{{ "# This is a title" }}
{{ markdown("# This Is A Title") }}
```
<pre class="result">
\# This is a title
# This is a title
</pre>

The markdown type, the html type and the _css_, _js_ and _json_ types are called format types. These format types
can be converted to the string type only if the value to convert is an untyped constant as seen in the example for the 
html type.

## int

Values with type `int` are integer numbers. These are examples of int variable declarations:

```scriggo
{% var a int %}  a is 0
{% var b = 5 %}  b is 5
{% c := 7 %}     c is 7
```

## slice

A slice is a numbered sequence of elements of the same type. For example `[]int{5, 2, 7, 9}` is a slice of `int`
values and `[]string{"hello", "ciao"}` is a slice of `string` values. A slice can be `nil` which is a value other than
an empty slice.

These are examples of slice variable declarations:

```scriggo
{% var a []int %}                a is a nil int slice
{% var b = []int{} %}            b is an empty int slice
{% var c = []int{3, 0, 7, 2} %}  c is an int slice with elements 3, 0, 7 and 2
{% d := []int{2, 9, 5} %}        d is an int slice with elements 2, 9 and 5
```

### Length

The `len` built-in function is used to read the length of a slice, ie the number of elements it contains.

```scriggo
{% var s = []int{"a", "b", "c"} %}
{{ len(s) }}
```
<pre class="result">3</pre>

### Accessing to an element of a slice

To access an element of a slice use square brackets `[` and `]` with its index starting from 0. For example:

```scriggo
{% var s = []string{"a", "b", "c"} %}
{% s[1] = "e" %}
{{ s[0] }} {{ s[1] }} {{ s[2] }}
```
<pre class="result">a e c</pre>

To access to the last element of a slice `s` you can write `s[len(s)-1]`.

Accessing a nil slice or accessing a index that does not exist is an error.

### Iterate over the elements of a slice

To iterate over the elements of an slice use the _for_ statement. For example:

```scriggo
{% var greetings = []string{"Hello", "Ciao", "你好"} %}

{% for greeting in greetings %}
  {{ greeting }}
{% end %}
```
<pre class="result">
  Hello
  Ciao
  你好
</pre>

### Slicing

Slicing is an operation that returns a portion of a slice. To do a slicing use square brackets `[` and `]` with the
initial and final indexes separated with a colon `:`. The final index is not included in the returned slice.

```scriggo
{% var s = []string{"a", "b", "c", "d", "e"} %}
{{ s[1:3] }}
```
<pre class="result">
b
c
</pre>

```scriggo
{% for x in s[2:5] %} {{ x }} {% end %}
```
<pre class="result">
 c  d  e 
</pre>

A slicing operation does not copy the elements, as a consequence the resulting slice refers the same elements of the
original slice.

```scriggo
{% var s = []int{0, 1, 2, 3, 4, 5} %}
{% var p = s[0:3] %}
{% s[0] = 5 %}
{{ p }}
```
<pre class="result">
5
1
2
</pre>

### Appends elements to a slice

The length of a slice cannot change, but as you can do a slicing to have a shorter slice, you can append to have a
longer slice. The append operation is done with the `append` built-in function:

```scriggo
{% var s = []int{0, 1} %}
{% s = append(s, 2, 3, 4, 5) %}
{{ s }}
```
<pre class="result">
0
1
2
3
4
5
</pre>

{% end raw doc %}
