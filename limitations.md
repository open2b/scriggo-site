{% extends "/layouts/article.html" %}
{% macro Title string %}Limitations{% end %}
{% Article %}

{% raw %}

# Limitations

## Missing features

These limitations are features that Scriggo currently lacks but that are
under development. To check the state of a limitation please refer to the
GitHub issues linked in the list below.

* methods declarations ([#458](https://github.com/open2b/scriggo/issues/458))
* interface types definition ([#218](https://github.com/open2b/scriggo/issues/218))
* assigning to non-variables in 'for range' statements ([#182](https://github.com/open2b/scriggo/issues/182))
* importing the "unsafe" package from Scriggo ([#288](https://github.com/open2b/scriggo/issues/288))
* importing the "runtime" package from Scriggo ([#524](https://github.com/open2b/scriggo/issues/524))
* labeled continue and break statements ([#83](https://github.com/open2b/scriggo/issues/83))
* compilation of non-main packages without importing them ([#521](https://github.com/open2b/scriggo/issues/521))

For a comprehensive list of not-yet-implemented features
see the list of [missing features on Github](https://github.com/open2b/scriggo/labels/missing-feature).

## Limitations due to maintain the interoperability with Go official compiler `gc`

* types defined in Scriggo are not correctly seen by the `reflect` package.
    This manifests itself, for example, when calling the function
    `fmt.Printf("%T", v)` where `v` is a value with a Scriggo defined type.
    The user expects to see the name of the type but `fmt` (which internally
    relies on the package `reflect`) prints the name of the type that wrapped
    the value in `v` before passing it to gc.

* unexported fields of struct types defined in Scriggo are still accessible
    from native packages with the reflect methods. This is caused by the
    reflect methods that does not allow, by design, to change the value of an
    unexported field, so they are created with an empty package path. By the
    way, such fields (that should be unexported) can not be changed without
    the reflect and have a particular prefix to avoid accidental accessing.

* in a structure, types can be embedded but, apart from interfaces, if they
    have methods then they must be the first field of the struct. This is a
    limitation of the StructOf function of reflect.
    See [golang/go#15924](https://github.com/golang/go/issues/15924).

* a select supports a maximum of 65536 cases.

* Native packages can be imported only if they have been precompiled into the
    Scriggo interpreter/execution environment.
    Also see the commands `scriggo import` and `scriggo init`.

* types are not garbage collected. See issue [golang/go#28783](https://github.com/golang/go/issues/28783).

## Arbitrary limitations

These limitations have been arbitrarily added to Scriggo to enhance
performances:

* 127 registers of a given type (integer, floating point, string or general) per function
* 256 function literal declarations plus unique functions calls per function
* 256 types available per function
* 256 unique predefined functions per function
* 16384 integer values per function
* 256 string values per function
* 16384 floating-point values per function
* 256 general values per function

{% end raw %}
