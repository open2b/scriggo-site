{% extends "/layouts/doc.html" %}
{% macro Title string %}What is Scriggo{% end %} 
{% Article %}

# What is Scriggo

Scriggo is a fast, modern and secure Go language template engine and embeddable runtime.

As template engine, it allows to use the Go language as a scripting language inside templates.

Scriggo can be easily included in any application and does not require Go installed on the client or server because
Scriggo implements the Go programming language with its own compiler and virtual machine written from scratch.

Scriggo is written in pure Go, so it can run on servers, desktops, mobile devises and also on serverless platforms
like <a href="https://aws.amazon.com/lambda/">Amazon Lambda</a>,
<a href="https://cloud.google.com/appengine/">Google App Engine</a> and
<a href="https://azure.microsoft.com/services/functions/">Azure Functions</a>.

It's fast. Executes code as fast as mainstream languages like Ruby, Python and Lua. 

## Why use Scriggo