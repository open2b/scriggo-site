{% extends "/layouts/article.html" %}
{% macro Title %}Internals{% end %}
{% Article %}

# Internals

<ol>
  <li>
    <a href="#an-overview-of-the-scriggo-implementation">An overview of the Scriggo implementation</a>
    <ol>
      <li>
        <a href="#the-compiler" id="markdown-toc-the-compiler">The compiler</a>
        <ol>
          <li><a href="#the-lexer" id="markdown-toc-the-lexer">The lexer</a></li>
          <li><a href="#the-parser" id="markdown-toc-the-parser">The parser</a></li>
          <li><a href="#the-type-checker">The type checker</a></li>
          <li><a href="#the-emitter">The emitter</a></li>
       </ol>
      </li>
      <li><a href="#the-runtime" id="markdown-toc-the-runtime">The runtime</a></li>
    </ol>
  </li>
</ol>

## An overview of the Scriggo implementation

Scriggo is composed by two main parts:

- the **compiler**
- the **runtime**

The **compiler**, in turn, is composed by:

- the **lexer**
- the **parser**
- the **type checker**
- the **emitter**

The **runtime** is implemented on the **Scriggo virtual machine**.

![internals_overview](/images/internals_overview.png)

### The compiler

The Scriggo compiler takes a source code and generates some bytecode ready to be executed by the virtual machine.

Let's take a look at the entire process.

#### The lexer

The lexer reads a source code and outputs a list of _tokens_.

![lexer](/images/lexer.png)


The lexer recognizes all the tokens used by Go plus the ones specific for the template, as {% raw %}`{%`, `%}` {% end %}, `macro` etc.. To get an overview of the Scriggo template syntax see [the Scriggo Template](/doc/template).

#### The parser

The **parser** takes the the list of tokens returned by the **lexer** and generates an AST (Abstract Syntax Tree).

![parser](/images/parser.png)

In the parser we can se the first main difference between the _programs_, the _scripts_ and the _templates_.
For example, the AST representing a program is composed by a tree, that contains a package, that contains a list of declarations.
On the other hand, a script or a template is a tree that holds a list of statements, and does not contain a package.

#### The type checker

The **type checker** performs several operations on the AST given by the **parser**:

1. reorders the declarations, if necessary
1. make a type check
1. transform the AST, preparing it for the emission

![typechecker](/images/typechecker.png)

The type checker is the first part of the compiler where the Scriggo code get in touch with the _host_; the declarations passed at compilation time are injected into the AST as _predefined values_.

Before starting with the type checking a _dependency analysis_ is performed on the AST. This analysis has two main purposes: the first is to catch any _initialization loop_, the second is to find an order for the declarations.

The source code of the type checker can be found in the `internal/compiler` package:

- [checker.go]() declares the type `typechecker` and some of its methods.
- [checker_assignment.go]() declares functions to perform the type checking of declarations and assignment.
- [checker_dependencies.go]() makes a dependency analysis on the given AST.
- [checker_package.go]() makes a type checking on a package.
- [checker_expressions.go]() performs the type checking of the expressions.
- [checker_statements.go]() declares functions to type check the statements; for example it exports the method `checkNodes`.

#### The emitter

The *emitter* takes the AST and _emits_ the code that will be intepreted by the virtual machine.

![emitter](/images/emitter.png)

Since the virtual machine knows nothing about programs, template and scripts, the emitter to make them converge into a list of instructions. For instance, the `show` statement in the templates (that renders a value) is emitted as a function call to a function that takes a writer and an expression as arguments.

The emitter relies on the **builder**, that knows the internal implementation of the virtual machine and exposes some utility function to the emitter hiding as much as possible the internals details.
Note that, without the use of the builder, the emission of instructions for the virtual machine would be very complex.

### The runtime

The runtime of Scriggo executes the byte code using its internal virtual machine.
For a detailed explanation of the virtual machine and to see the set of the bytecode instructions that it uses, see the section [Bytecode](bytecode).
