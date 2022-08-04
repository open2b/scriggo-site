{% extends "/layouts/page.html" %}
{% macro Title string %}Scriggo News{% end %}
{% Content %}

# News

## Released Scriggo v0.56.1
August 4, 2022

This is a bug-fix release.

#### Changes

* [builtin: fix invalid behavior of function 'HasSuffix'](https://github.com/open2b/scriggo/commit/7038b90584b0fbd5a026188d8920aa2040bd0bc4)

## Released Scriggo v0.56.0
August 4, 2022

### Support Go 1.19

This release officially supports [Go 1.19](https://go.dev/doc/go1.19).
### Changes

* [scriggo: require Go 1.18 as the minimum supported version](https://github.com/open2b/scriggo/commit/4cbb9b98f3e798dbfe872353e22a7341cd23f59e) (breaking change)

**This version breaks the compatibility with previous versions.**

## Released Scriggo v0.55.0
July 8, 2022

This release is focused on the addition of the [map selector expressions](/templates/specification#map-selector-expressions) in
templates, the addition of builtins for YAML as well as the removal of some deprecated features.

Also, `any` is now an alias for `interface{}`.

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.55.0

**This version breaks the compatibility with previous versions.**

### Changes

* [scriggo: use Go 1.18 to compile the scriggo command](https://github.com/open2b/scriggo/commit/18f1e95da9504884319c94c99996cf51971d6359)

* [scriggo: require Go 1.17 as the minimum supported version](https://github.com/open2b/scriggo/commit/a253ef56d113b352f4b7778854637f5ec2f16b55) **(breaking change)**

* [cmd/scriggo: add 'debug/buildinfo' and 'net/netip' packages to stdlib](https://github.com/open2b/scriggo/commit/220f028304dab8dfdd2965cc6bd78bab38fa49bf)

* [compiler: allow 'any' as 'interface{}' alias](https://github.com/open2b/scriggo/commit/30a9cf2d64c13cc1a50f21fe0b631dc5b99e16bc)

* [all: remove script support](https://github.com/open2b/scriggo/commit/f423f045394dc5364d9de0ec94413cc26af34a6c) **(breaking change)**

* [all: remove the dollar identifier](https://github.com/open2b/scriggo/commit/2ecde7b004f83694650dfa2a2638dea9ef6dee69) **(breaking change)**

* [compiler, runtime: extend the selector expression to maps in templates](https://github.com/open2b/scriggo/commit/de8250578ff06499b67eae0e672061407c2540f7)

* [builtin: add 'marshalYAML' and 'unmarshalYAML' builtins](https://github.com/open2b/scriggo/commit/46e3dfe8ae0b74bdfdc419bca7592a276182f344)

## Released Scriggo v0.54.0
June 28, 2022

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.54.0

**This version breaks the compatibility with previous versions.**

### Changes

* compiler/checker: don't consider 'make(...)' a constant function call **(breaking change)**

* runtime: implement range over channels

* all: add 'else' statement to the 'for in' and 'for range' statements **(breaking change)**

* scriggo: update dependencies

### Optimizations

* compiler/emitter, runtime: simplify 'for range' instruction with 'break'

### Fixes

* compiler/checker: fix terminating statements

* ast: add '(*Block).String' method

* ast/utils: don't panic '(*dumper).Visit' if a block has nil position

## Released Scriggo v0.53.5
January 07, 2022

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.53.5

### Fixes

* compiler/checker: fix const shift expression with untyped left operand

* compiler/emitter: fix panic involving operations on pointer values

* compiler/checker: return error in case of parenthesis around ident on :=

* compiler/checker: fix call to 'setValue'

* compiler/parser: fix error message on not existent path


## Released Scriggo v0.53.4
October 27, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.53.4

### Fixes

* compiler/emitter,runtime: fix addressability of values read by GetVar

  Invalid behavior occurred when assigning to slices, pointers, or maps at the package level.

* compiler/disassembler: fix panic disassembling 'MapIndex' operation

  Disassembling code with an index expression on a map resulted in a panic.


## Released Scriggo v0.53.3
October 21, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.53.3

### Fixes

* builtin: fix time tests

  Time tests depended on the local time zone and passed only in CET time zone.

* compiler/checker: fix panic on redeclared as imported package name

  If a native imported package name was redeclared, an internal error would occur instead of returning a 'redeclared as
  imported package name' error.


## Released Scriggo v0.53.2
October 20, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.53.2

### Optimizations

* compiler/emitter: optimize {% raw %}`{{ render .. }}`{% end %}

  This change optimizes the `show` statement with a `render` expression, generating fewer virtual machine instructions and fewer memory allocations.

### Fixes

* cmd/scriggo: fix package name in file generated by the Import command

  The scriggo command could generate an invalid Go file.

  Thanks to [@traetox](https://github.com/traetox) for the contribution.

* compiler,runtime: fix builtin `delete` when key value has Scriggo type

  Using the `delete` builtin with a type declared in the interpreted source code, a runtime panic occurred.


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