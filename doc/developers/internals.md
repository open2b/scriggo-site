{% raw %}
# Internals
{: .no_toc}

1. TOC
{:toc}

## An overview of the Scriggo implementation

Scriggo is composed by two main parts:

- the **compiler**
- the **virtual machine**

The **compiler**, in turn, is composed by:

- the **lexer**
- the **parser**
- the **type checker**
- the **emitter**

![internals_overview](/images/internals_overview.png)

### The compiler

The Scriggo compiles takes a source code and generates some byte code ready to be executed by the Scriggo virtual machine.

Let's take a look at the entire process.

#### The lexer

The lexer reads a source code and outputs a list of _tokens_.


For example, given the following source code:

```
a := 45 + f()
```

the lexer returns the tokens: `a`, `:=`, `45`, `+`, `f`, `(`, `)`.

#### The parser

The **parser** takes the the list of tokens returned by the **lexer** and generates an AST (Abstract Syntax Tree).

#### The type checker

The **type checker** performs several operations on the AST given by the **parser**:

1. reorders the declarations, if necessary
1. make a type check
1. transform the AST, preparing it for the emission

#### The emitter

The *emitter* takes the AST and _emits_ the code that will be intepreted by the virtual machine.

### The virtual machine

The Scriggo virtual machine takes the output of the Scriggo compiler and executes it. For a detailed explanation of the virtual machine and to see the set of instructions that it uses, see the section [Virtual Machine](/doc/developers/vm.md).

{% endraw %}