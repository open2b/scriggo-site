{% extends "/layouts/page.html" %}
{% macro Title string %}Scriggo News{% end %}
{% Content %}

# News

## Released Scriggo v0.52.2
September 28, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.52.2

### Fix serialization of struct in JS and JSON contexts

Fix the serialization of struct field values in JS and JSON contexts.


## Released Scriggo v0.52.1
September 23, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.52.1

### Add unsafeconv as builtin in scriggo serve

Make the package [unsafeconv](/templates/builtins#unsafeconv) available as builtin package in templates rendered using `scriggo serve`.

### Fix scriggo version

Fix the version printed by `scriggo version` ([#882](https://github.com/open2b/scriggo/issues/882))


## Released Scriggo v0.52.0
September 22, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.52.0

**This version breaks the compatibility with previous versions.**

### Add scriggo command binaries.

Starting from this version, the binaries of the scriggo command are available on GitHub.

### Replace env.Exit with env.Stop (breaking change).

This change replaces the Exit method of the native.Env interface
with the Stop method. The Stop method has the same behaviour of
Exit but accepts an error instead of a status code.

The following code:

    env.Exit(0)
    env.Exit(2)

with this change, can be rewritten as

    env.Stop(nil)
    env.Stop(scriggo.NewExitError(2, nil))

### Add the unsafeconv package as built-in in templates.

### Other fixes and improvements.