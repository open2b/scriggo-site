{% extends "/layouts/doc.html" %}
{% macro Title string %}What is Scriggo{% end %} 
{% Article %}

# What is Scriggo?

Scriggo is the world’s powerful template engine and Go embeddable interpreter.

* Fast, a very fast embeddable pure Go language interpreter.
* Modern template engine with Go as scripting language.
* Native support for Markdown in templates.
* Secure by default. No access to packages unless explicitly enabled.
* Easy to embed and to interop with any Go application.

To start using Scriggo just get it:

<div class="get-scriggo">go get <span class="scriggo-path">github.com/open2b/scriggo</span></div>

Scriggo imported into a Go application has no dependencies other than the Go standard library.

## Template Engine

Scriggo is the world’s powerful template engine:

* Support inheritance, macros, partials, imports and contextual autoescaping.
* Native support for Markdown, inside or without HTML.
* Use the Go language as advanced template scripting language.
* Static type checking to catch errors as early as possible, and not at runtime. 
* Allow importing Go packages in templates.
* Templates are compiled to bytecode and executed on a fast virtual machine.

Scriggo templates have a simple and familiar syntax that is quick to learn, whether you know Go or not yet.
Documentation does not assume you already know Go, but by reading the Scriggo documentation you can also learn the Go 
language.

## Embeddable Go Interpreter

Scriggo is also a fast embeddable interpreter for Go programs. It lets you build and run Go programs directly in your
application, nothing else. It does not require any Go installation to build and run Go programs.

* One of the fastest Go embeddable interpreters.
* Programs compiled to bytecode and executed on a fast virtual machine.
* Easy integration with your Go application.
* Written in pure Go and no dependencies beyond the Go standard library.

## Where run Scriggo applications

Scriggo can run on servers, desktops, mobile devices and on serverless platforms like
<a href="https://aws.amazon.com/lambda/">Amazon Lambda</a>,
<a href="https://cloud.google.com/appengine/">Google App Engine</a> and
<a href="https://azure.microsoft.com/services/functions/">Azure Functions</a>. It can also be compiled to Web Assembly
to run Go programs and templates in browsers or in a sandboxed environment embedded in your application.
