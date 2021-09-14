{% extends "/layouts/page.html" %} 
{% macro Title string %}Scriggo Frequently Asked Questions (FAQ){% end %}
{% Content %}

{% raw faq %}

# Frequently Asked Questions

<dl>
    <dt><a href="#origin-and-nature-of-the-project">Origin and nature of the project</a></dt>
    <dd><a href="#what-is-the-story-of-this-project">What is the story of this project?</a></dd>
    <dd><a href="#is-scriggo-open-source">Is Scriggo Open Source?</a></dd>
    <dd><a href="#what-is-the-relationship-between-scriggo-and-google">What is the relationship between Scriggo and Google?</a></dd>
    <dd><a href="#can-i-contribute-to-scriggo">Can I contribute to Scriggo?</a></dd>
    <dt><a href="#scriggo-and-go">Scriggo and Go</a></dt>
    <dd><a href="#can-you-extend-the-go-language-with-x-functionality">Can you extend the Go language with X functionality?</a></dd>
    <dd><a href="#can-i-use-the-shebang-in-go-files">Can I use the shebang in Go files?</a></dd>
    <dd><a href="#what-versions-of-go-does-scriggo-support">What versions of Go does Scriggo support?</a></dd>
    <dt><a href="#safety">Safety</a></dt>
    <dd><a href="#what-do-you-mean-that-scriggo-is-safe">What do you mean that Scriggo is safe?</a></dd>
    <dd><a href="#how-can-i-run-third-party-code-safely">How can I run third party code safely?</a></dd>
    <dt><a href="#versions">Versions</a></dt>
    <dd><a href="#how-does-versioning-work">How does versioning work?</a></dd>
    <dd><a href="#can-scriggo-be-used-in-production">Can Scriggo be used in production?</a></dd>
    <dt><a href="#performance">Performance</a></dt>
    <dd><a href="#why-is-scriggo-so-fast-and-allocates-less-than-other-interpreters-in-go">Why is Scriggo so fast and allocates less than other interpreters in Go?</a></dd>
    <dd><a href="#how-can-scriggo-perform-like-mainstream-languages">How can Scriggo perform like mainstream languages?</a></dd>
    <dd><a href="#there-is-still-room-for-optimization">There is still room for optimization?</a></dd>
    <dt><a href="#templates">Templates</a></dt>
    <dd><a href="#go-already-has-a-template-engine-why-another-one">Go already has a template engine, why another one?</a></dd>
    <dd><a href="#what-is-a-truthful-value-in-templates">What is a truthful value in templates?</a></dd>
    <dd><a href="#can-you-implement-x-functionality-in-templates-even-if-it-is-not-present-in-go">Can you implement X functionality in templates even if it is not present in Go?</a></dd>
    <dd><a href="#why-isnt-the--operator-supported-in-templates">Why isn't the ?: operator supported in templates?</a></dd>
    <dd><a href="#why-is-a-string-value-not-convertible-to-html-type-in-the-templates">Why is a string value not convertible to html type in the templates?</a></dd>
    <dd><a href="#what-is-the-predeclared-identifier-itea">What is the predeclared identifier itea?</a></dd>
    <dt><a href="#implementation">Implementation</a></dt>
    <dd><a href="#how-is-the-compiler-written">How is the compiler written?</a></dd>
    <dd><a href="#how-does-the-runtime-work">How does the runtime work?</a></dd>
    <dt><a href="#how-are-the-dependencies-in-the-gomod-file-used">How are the dependencies in the go.mod file used?</a></dt>
</dl>

## Origin and nature of the project

### What is the story of this project?

Scriggo was initially (around 2018 when it still didn't have a name) the porting from Perl to Go of a template engine
that we, at Open2b, used for our ecommerce platform. While porting the platform to Go, we realized that the previous
template engine was too limited for the new platform. This is why we looked around for a valid alternative.

Maybe because Go is a young project compared to other languages or maybe because it is not a dynamic language and
therefore unwilling to be used as a scripting language, that we did not find an adequate solution to our needs.
Such as it could be, but just to name one, Jinja for Python.

We found that template engines available in Go either weren't exactly immediate to use for the audience we were
targeting, or they were too limited in the possibilities of integration with our platform, or they were too limited in
functionality. Many of the systems that we evaluated were not adequately supported, not to say almost left to
themselves.

For these reasons, we thought about the possibility of expanding the system in Go, which we had ported, in order to
add more advanced scripting functionalities. That was the moment when Scriggo began to take its current form.

Some choices, such as the use of the same code base for the template engine and for the Go interpreter and the creation
of a virtual machine, were two decisive choices, among many, to give Scriggo its clean design and its execution
efficiency.

Scriggo could have been publicly released in 2020, but a further year of use in production served us to make profound
improvements, mainly to the template engine, which have given it immediacy in use, making it a natural extension of
the Go language.

On September 2021, Scriggo was made public by Open2b as an open source project for the Go community and beyond,
with the hope that it will further contribute to its growth. Scriggo, which includes a template engine and a Go
embeddable interpreter, provides a stable base for creating applications, sites, and projects of various kinds.

### Is Scriggo Open Source?

Yes, it is distributed under a BSD license, the same as Go.

### What is the relationship between Scriggo and Google?

None, apart from the fact that the Go language, with which Scriggo is written and which it implements, is a Google
project.

### Can I contribute to Scriggo?

Of course, you are welcome.

## Scriggo and Go

### Can you extend the Go language with X functionality?

Scriggo, for programs, implements exclusively the Go language specification, so if it is not present in the Go
specification it cannot be implemented.

### Can I use the shebang in Go files?

No, because the shebang is not mentioned in the Go specification. If you need it, you could remove the shebang line from the
file when it is read by Scriggo.

### What versions of Go does Scriggo support?

Scriggo supports the latest two stable versions of Go. Therefore, it currently supports versions 1.16 and 1.17. When
version 1.18 is released, the next stable version of Scriggo will support versions 1.17 and 1.18.

## Safety

### What do you mean that Scriggo is safe?

The Scriggo source code uses only the standard Go library and does not use the "unsafe" and "syscall" packages.

Code executed by Scriggo, by default, cannot access external code such as packages, functions, variables, constants and
types (apart from the Go builtins) that have not been explicitly provided to it. Therefore, it is possible to restrict
access to the network, file system, environment variables, and the execution of other processes.

The `print` and `println` builtins, which print to the standard error by default, can be easily replaced. 

Executed code can be interrupted via a context. Therefore, it is not possible to write code that cannot be interrupted.

It is not possible instead to limit the memory that an executed code can allocate, there is no safe and efficient way
that can be implemented in Go. Moreover, the new types specified in the executed code allocate memory that is not
released. This is a known limitation of the Go reflect implementation. Specifying the same type does no result,
however, in a new memory allocation.

### How can I run third party code safely?

By running the process in a controlled environment such as a virtual machine.

## Versions

### How does versioning work?

Until version 1.0.0 is released, all versions will have the form "0._minor_._patch_". The minor will be increased if
the version introduces an incompatibility with the previous version. Patches are released for the latest minor only.
Once version 1.0.0 is released, patches for the current and previous minor will be released.

### Can Scriggo be used in production?

Scriggo v0.45.0 is currently used in production and the latest version will be soon. It is not yet in version 1.0 as
there are some part of the Go specification to implement, and we still want to be free to make small and limited
changes to the template specification and packages after receiving feedback from the community.

## Performance

### Why is Scriggo so fast and allocates less than other interpreters in Go?

Scriggo implements a virtual machine with registers. Go programs and templates are first compiled to bytecode and then executed by the Scriggo virtual machine.

This allows for excellent execution speed and a small memory allocation compared to other interpreters in Go.

### How can Scriggo perform like mainstream languages?

The Go static type checking made it possible to design efficient Scriggo virtual machine instructions even if they have
to call the Go reflect later.

### There is still room for optimization?

The currently generated bytecode is not yet as optimized as we would like. There is therefore still room for
optimization.

## Templates

### Go already has a template engine, why another one?

The [template engine](https://pkg.go.dev/html/template) that comes with the standard Go library was designed to be used
in simpler situations such as data formatting and is not, and never pretended to be, a template engine with an
integrated programming language. Moreover, it has a non-immediate understanding syntax for those coming from other 
templating languages like Jinja and Liquid.

### What is a truthful value in templates?

We wanted the template language to be easy to use even for non-programmers. A simple statement like the following:

``
{% if len(products) == 0 %} There are no products {% end %}
``

brings with it some concepts such as slices, the built-in `len` and the equality operator, which on the whole could 
discourage people approaching a template language without coding skills.

In Scriggo, therefore, the previous code can also be written in the following way:

```
{% if not products %} There are no products {% end %}
```

The condition of the _if_ statement can have any type and is _true_ if it is truthful. A _nil_ or empty slice is not
truthful and therefore in the context of the _if_ condition, `len(products) == 0` can be replaced with
`not products`.

`and`, `or` and `not` are three boolean operators, available in templates, which are based on the concept of
truthful values. For more details you can see the [documentation](/templates/operators#and-or-and-not) and the
[template specification](/templates/specification#truthful-values).

### Can you implement X functionality in templates even if it is not present in Go?
If the X functionality is templates specific and does not also concern the Go language, then it can be evaluated. If, on
the other hand, it could also concern the Go language, then it is less likely that it can be implemented as we prefer
to maintain as much consistency as possible with the Go language.

### Why isn't the ?: operator supported in templates?

The same reasons for implementing it in templates would also apply to the Go language, so it should first be added to
the Go specification.

In templates context, consider for example that the following code (not valid in Scriggo)

```
{{ products ? "there are " + sprintf("%d", len(products)) : "no products" }}
```

can be written in Scriggo as

```
{% if products %}there are {{ len(products) }}{% else %}no products{% end %}
```

### Why is a string value not convertible to html type in the templates?

Only untyped string constants are convertible to the _html_ type. The same is true for the _markdown_, _css_, _js_, and _json_ 
types. This helps to limit cross-site scripting (XSS) vulnerabilities that the template writer might unintentionally 
introduce.

### What is the predeclared identifier itea? 

The predeclared identifier _itea_ represents the body of a _using_ statement. It is the rendered body or a macro
with that body.

We were looking for a name that was as short as Go's _iota_ and was not in common use. Hence, "itea". If you are 
curious, itea is a genus of shrubs and small trees, commonly called sweetspires (Wikipedia). 

## Implementation

### How is the compiler written?

Scriggo's compiler is written in Go, entirely from scratch and therefore does not rely on any other previously written
compiler. The same compiler is able to compile both Go programs and Scriggo templates with the same code base. This
allows to make fixes, improvements and new features at the same time for the compilation of programs and templates.

The compiler consists of the parser, a completely handwritten and not self-generated recursive descent parser, the type
checker and the emitter. The final result of the compilation is a set of data structures, one for each Scriggo
function defined in the source code and one for each native function called by Scriggo functions. Templates have an
implicit _main_ function. Each Scriggo function then has its body compiled to bytecode.

### How does the runtime work?

The Scriggo runtime is a virtual machine with registers. Like the compiler, it is written from scratch. Many of the
virtual machine instructions rely on Go reflect functions to operate on type literal values, call native functions,
and implement the _select_ statement. This allows Scriggo to have excellent interoperability with Go's native values
and functions.

The virtual machine has 4 stacks of registers, one for integers, one for floating-points, one for strings, and one for
all other values. The stacks have a fixed initial size and can then grow when needed.

For more details on registers, and bytecode instructions you can see the [disassembler documentation](/disassembler).

### How are the dependencies in the go.mod file used?

The [github.com/open2b/scriggo](https://pkg.go.dev/github.com/open2b/scriggo) module has dependencies that are used
exclusively by the scriggo command. When using Scriggo embedded in your applications there are no dependencies apart
from the standard Go library.

{% end raw faq %}
