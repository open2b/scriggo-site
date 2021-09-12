{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: types{% end %}
{% Article %}

{% raw doc %}

# Types

Scriggo implements all [types of Go](https://golang.org/ref/spec#Types) but only the most common types
are documented here.

All values in Scriggo have a type and there are several basic types.

## Booleans

Boolean values are the `true` and `false` values. A Boolean variable can be declared in one of the following ways:

```scriggo
{% var a bool %}     a is false
{% var b = false %}  b is false
{% var c = true %}   c is true
{% d := true %}      d is true
{% e := false %}     e is false
```

## Strings

Strings are written with double quotes, as "hello", or with a grave accent, as `hello`. A string variable can be
declared in one of the following ways:

```scriggo
{% var a string %}     a is an empty string
{% var b = "" %}       b is an empty string
{% var c = "hello" %}  c is the string "hello"
{% d := "hello" %}     d is the string "hello"
```

To read the length in bytes of a string use the `len` function, instead to read the length in characters use the
`runeCount` function. The name of the `runeCount` function comes from "rune" which is the term used in Go to name the
characters. For a more accurate explanation you can read
[Strings, bytes, runes and characters in Go](https://go.dev/blog/strings).

## Integers

Values with type `int` are integer numbers (in a range from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
on 64-bit architectures). An `int` variable can be declared in one of the following ways:

```scriggo
{% var a int %}  a is 0
{% var b = 0 %}  b is 0
{% var c = 5 %}  c is 5
{% d := 7 %}     d is 7
```

## Slices

A slice is a numbered sequence of elements of the same type. For example `[]int{5, 2, 7, 9}` is a slice of `int`
values and `[]string{"hello", "ciao"}` is a slice of `string` values. A slice can be `nil` which is a value other than
an empty slice.

A `slice` variable can be declared in one of the following ways:

```scriggo
{% var a []int %}                a is a nil int slice
{% var b []int = nil %}          b is a nil int slice
{% var c = []int{} %}            c is an empty int slice
{% var d = []int{3, 0, 7, 2} %}  d is an int slice with elements 3, 0, 7 and 2
{% e := []int{} %}               e is an empty int slice
{% f := []int{2, 9, 5} %}        f is an int slice with elements 2, 9 and 5
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

Accessing a nil slice or index that does not exist is an error.

### Iterate over the elements of a slice

To iterate over the elements of a slice use the for statement. For example:

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
{% for e in s[2:5] %}
{{ e }}
{% end %}
```
<pre class="result">
c
d
e
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
