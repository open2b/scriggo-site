{% extends "/layouts/doc.html" %}
{% macro Title string %}What is Scriggo{% end %} 
{% Article %}

# What is Scriggo?

Scriggo is a fast, modern and secure Go language template engine and embeddable interpreter.

* Fast, a very fast embeddable pure Go language runtime.
* Modern and powerful Template Engine with Go as scripting language.
* Native support for Markdown in templates.
* Secure by default. No access to packages unless explicitly enabled.
* Easy to embed and to interop with any Go application.

To start using Scriggo just get it:

<div class="get-scriggo">go get <span class="scriggo-path">github.com/open2b/scriggo</span></div>

Scriggo imported into a Go application has no dependencies other than the Go standard library.

## Template Engine

Scriggo is the most powerful template engine for Go:

* Support inheritance, macros, partials, imports and contextual autoescaping.
* Native support for Markdown, inside or without HTML.
* Use the Go as advanced template scripting language.
* Static type checking to catch errors as early as possible, and not at runtime. 
* Allow importing Go packages in templates.
* Templates compiled to bytecode and executed on the Scriggo virtual machine.

Scriggo has a simple and familiar syntax that is quick to learn whether you know Go or not yet. Documentation does not
assume you already know Go, but by reading the Scriggo documentation you can also learn Go.

## Embeddable Go Interpreter

Scriggo is also a fast embeddable interpreter for Go programs. It lets you build and run Go programs directly in your
application, nothing else. It does not require any Go installation to build and run Go programs.

* One of the fastest Go embeddable interpreters.
* Programs compiled to bytecode and executed on the Scriggo virtual machine.
* Easy integration with your Go application.
* Pure Go and no dependencies beyond the Go standard library.

## Where run Scriggo applications 

Scriggo can run on servers, desktops, mobile devices and also on serverless platforms
like <a href="https://aws.amazon.com/lambda/">Amazon Lambda</a>,
<a href="https://cloud.google.com/appengine/">Google App Engine</a> and
<a href="https://azure.microsoft.com/services/functions/">Azure Functions</a>. It can also be compiled to Web Assembly
to run Go programs and templates in browsers or in a sandboxed environment in your application. 
