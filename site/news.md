{% extends "/layouts/page.html" %}
{% macro Title string %}Scriggo News{% end %}
{% Content %}

# News

## Released Scriggo v0.60.0
February 17, 2025

This release follows the release of Go 1.24 and introduces updates to ensure compatibility. The minimum supported Go version is now 1.23, while the Scriggo command is compiled with Go 1.24.

## Released Scriggo v0.59.1
January 27, 2025

This is a bug-fix release that addresses an [issue related to the call of method values](https://github.com/open2b/scriggo/issues/967).

## Released Scriggo v0.59.0
December 22, 2024

This release introduces LiveReload functionality to the [Scriggo Serve](/scriggo-command#serve-a-template) command, enabling automatic page reloads in the browser whenever the page’s template files are modified.

With LiveReload and the automatic template rebuild, changes are instantly reflected in the browser as soon as you save the files.

## Released Scriggo v0.58.1
December 16, 2024

This release focuses on bug fixes and updates dependencies. It addresses the following issues:

* **Resolved a race condition** occurring when a function or macro was called indirectly.
* **Fixed issues causing empty output** when invoking a macro, especially in cases of recursive calls.
* **Addressed template recompilation errors** in the Scriggo `serve` command when a render operator was indirectly used within a `show` statement.

## Released Scriggo v0.58.0
December 04, 2024

This release fixes the following issues:

* **Fixed a panic** in the Scriggo command when trying to import or extend a non-existent file.
* **Fixed rendering** issues when a non-Markdown file is rendered inside a Markdown context.

Additionally, files with the ".mdx" extension are now treated as Markdown files by default.

We’ve added the **MarkdownConverter** method to `native.Env`. A function or method, whether passed globally or imported into a template, can now use this method to retrieve the Markdown converter passed when building the template.

Finally, the `scriggo serve` command now enables support for footnotes in Markdown parsing.

### Changes

* [cmd/scriggo: fix panic when importing or extending a non-existent file](https://github.com/open2b/scriggo/commit/0fc1e83760a0110682347854a276d4bed984cc12)
* [cmd/scriggo: add extension.Footnote to goldmark](https://github.com/open2b/scriggo/commit/b8a4eedf047d51752f7c13d366378db104fe3b48) **(breaking change)**
* [internal/runtime: refactor macro calls and renderer implementation](https://github.com/open2b/scriggo/commit/3894f61234f4412bffc2085e9f5bd57d69181286)
* [internal/compiler: add .mdx as default Markdown file extension](https://github.com/open2b/scriggo/commit/0025e57955f322bb700fd85a1db71101937dc15c) **(breaking change)**
* [native,internal/runtime: introduce the Env.MarkdownConverter method](https://github.com/open2b/scriggo/commit/2c05377e66c0db077bc9ab9578e8fd750ed086e7) **(breaking change)**

**This version breaks the compatibility with previous versions.**

## Released Scriggo v0.57.1
October 22, 2024

This release only updates the GitHub Action files and the GoReleaser configuration for the release of Scriggo binaries.

## Released Scriggo v0.57.0
October 22, 2024

### Support for Go Compiler Versions 1.22 or Higher

This release adds support for Go compiler versions 1.22 or later.

Note that support for Go versions 1.21 or earlier is therefore removed, as functionality of Scriggo is no longer guaranteed for these compiler versions.

### Fixes

* [compiler/parser: fix panic on invalid default left expression](https://github.com/open2b/scriggo/commit/ce873dfde1e73cffd4dc7c5dce05b30229c46bd2)

### Dependency Updates

* [scriggo: update goldmark from v1.4.13 to v1.7.4](https://github.com/open2b/scriggo/commit/134381acd8708d4c6437c1cf33c0dc2ac1d2c347)
* [scriggo: update fsnotify from v1.5.4 to v1.7.0](https://github.com/open2b/scriggo/commit/caf8a0d4470d8e11347e13f60281861199c75692)

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