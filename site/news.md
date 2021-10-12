{% extends "/layouts/page.html" %}
{% macro Title string %}Scriggo News{% end %}
{% Content %}

# News

## Released Scriggo v0.53.1
October 12, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.53.1

### Optimizations

* cmd/scriggo: improve 'run' execution time using an I/O buffer
* compiler/emitter: reuse registers after emitting 'show' statement

### Fixes

* compiler/checker: fix conversion from untyped int constant to float64
* compiler/parser: fix parsing of literal numbers
* compiler/builder: fix 32-bit compilation of Scriggo


## Released Scriggo v0.53.0
October 04, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.53.0

**This version breaks the compatibility with previous versions.**

### Implement the 'scriggo run' command.

The new Run command executes a template file and its extended, imported and rendered files:

```shell
$ scriggo run index.html
```

All Scriggo builtins are available in template files. See the [scriggo command](https://scriggo.com/scriggo-command#run-a-template-file) for the complete syntax.

### Discard shebang line in templates (breaking change).

Discards a shebang line in a template file. In this way the template file can have a shebang as first line:

{% raw %}

```scriggo
#!/usr/bin/scriggo run
{% extends "layout.html" %}
{% Article %}
```

{% end raw %}

### Fix BuildTemplate to get the format from a FormatFS file system.

If a file system implements the FormatFS interface, Scriggo calls its Format method to get the file format.
This didn't work and this release fixes it.

### Other relevant commits

* all: improve documentation about contribution and testing
* compiler/checker: fix setValue in case of Scriggo defined complex type
* compiler/emitter: fix changeRegister in case of conversion


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