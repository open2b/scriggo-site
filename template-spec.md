{% extends "/layouts/article.html" %}
{% macro Title string %}The Scriggo Template Specification{% end %}
{% Article %}

{% raw content %}

# The Scriggo Templates Specification

## Introduction

This is a reference manual for Scriggo templates.

The Scriggo templates specification is based on the Go language specification, see <a href="https://golang.org/ref/spec">golang.org/spec</a>. The Go specification applies, apart from what written in this specification.

## Template source code representation

Template source code between _{%_ and _%}_, _{%%_ and _%%}_ and _{{_ and _}}_ is Unicode text encoded in <a href="https://en.wikipedia.org/wiki/UTF-8">UTF-8</a>.

Implementation restriction: A template tool, as an editor or interpreter, may require all the template file content to be encoded in UTF-8.

## Lexical elements

### Comments

Line comments and general comments can only appear between _{%_ and _%}_ and between _{%%_ and _%%}_. In addition there is a third form:

3. Template comments start with the character sequence _{#_ and stop with the first subsequent not paired character sequence _#}_. Template comments cannot appear between _{%_ and _%}_, _{%%_ and _%%}_ and between _{{_ and _}}_.

```
{# this is a template comment #}

{# this {# is also a #} comment #}
```

## Types

### Format types

A format type represents a sets of string values that can have a special treatment when used with `show` statement. The `string` type is a format type.

```
string    The string type. The strings in this set are not escaped in a Text context.
html      The set of all strings that are not escaped in a HTML context.
css       The set of all strings that are not escaped in a CSS context.
js        The set of all strings that are not escaped in a JS context.
json      The set of all strings that are not escaped in a JSON context.
markdown  The set of all strings that are not escaped in a Markdown context.
```

### Macro types

There is the new _macro type_.

```
TypeLit = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
          SliceType | MapType | ChannelType | MacroType .
```

_Macro types_ are alias for function types with a single return parameter having a predefined format type (_string_, _html_, _css_, _js_, _json_, _markdown_).

Aliases of format types cannot be used as the return type in a type macro literal.

```
macro() markdown           // alias for func() markdown
macro(int, bool) string    // alias for func(int, bool) string

type foo = string

macro() foo              // not valid, foo is an alias of the string type  
macro([]int) int         // not valid, int is not a format type
macro([]int)             // not valid, has no return parameters
macro() (string, css)    // not valid, has too many return parameters
```

## Declarations

There are no function and method declarations in Scriggo templates.

```
Declaration = ConstDecl | TypeDecl | VarDecl .
```

A constant, type and variable declaration can also be written between _{%_ and _%}_.

```
{% var n = 5 %}

{%
    const (
        n = 5
        m = 7
    )
%}

{% type T struct{ s string } %}
```

In addition there are macro declarations. A macro declaration binds an identifier to a macro.

At part of the different declaration syntax, macros are functions. A macro declaration can only be written using _{%_ and _%}_.

```
MacroDecl   =
    "{%" "macro" MacroName [ Parameters ] [ MacroResult ] "%}"
    Content
    "{%" "end" [ "macro" ] "%}" .
MacroName   = identifier .
MacroResult = "string" | "html" | "css" | "js" | "json" | "markdown" .
```

If a macro declaration does not have a result type, it is inferred from the context of the declaration.

```
Template format   Macro result type

Text              string
HTML              html
CSS               css
JS                js
JSON              json
Markdown          markdown
```

Macro declarations are only allowed in the previous template contexts. 

```
{% macro Title %}A beautiful title{% end %}

{% macro Image(src string, with, height int) html %}
  <a href="{{ src }}" with="{{ width }}" height="{{ height }}">
{% end macro %}
```

Endless macros are macros whose content extends to the end of the file. An endless macro declaration does not have the `macro` keyword,
does not have parameters and have an implicit return value. The name of an endless macro is an exported identifier.

```
EndlessMacroDecl = "{%" EndlessMacroName "%}" Content EOF .
EndlessMacroName = exported_identifier .
```

```
{ extends "layout.html" %}
{% Article %}

This is the article content,

{% a := 5 %} {# a is defined in the scope of the body of the Article macro #}

and more article content.
```

### Predeclared identifiers

The predeclared type identifiers _html_, _css_, _js_, _json_ and _markdown_ that with the type _string_ are called format types:

```
Format types:
    string html css js json markdown
```

The format types are the only types that can be used as macro result type. 

## Expressions

### Contains operators

Scriggo templates have <code>contains</code> and <code>not contains</code> operators. The following rules applies:

* For a value `v` of type `[]T` or `[n]T` and a value `x`, the expression `v contains x` yield an untyped boolean; `true` if an element of `v` is equals to `x`. `x` must be assignable to `T` and comparable to the values of `T`.

* For a map `v` with key values of type `T` and a value `x` of type `T`, the expression `v contains x` yield an untyped boolean; `true` if `v` has `x` as key.

* For string values `v` and `x`, the expression `v contains x` yield an untyped boolean; `true` if `x` is within `v`.

* For a string value `v` and a value `x` of type `rune`, the expression `v contains x` yield an untyped boolean; `true` if the Unicode code point `x` is within `v`.

* For a value `v` of type `[]byte` and a value `x` of type `rune`, the expression `v contains x` yield an untyped boolean; `true` if the Unicode code point `x` is within `v`.

If `!(v contains y)` is a valid expression evaluated to the value `b`, also `v not contains y` is a valid expression and it is evaluated to the value `!b`. 

### Operators

The operators are extended with the logical operators `and`, `or` and `not`.

```
binary_op  = "||" | "&&" | rel_op | add_op | mul_op | "and" | "or" .
unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" | "not" .
```

#### Operator precedence

```
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=  contains  not contains
    2             &&  and
    1             ||  or
```

### The logical zero value

A value `x` of type `T` is _logically zero_ if `x` is the zero value of `T`. The following exceptions applies:
* if `x` is an empty slice or an empty channel, `x` is logically zero,
* if `x` is a non nil interface and `z` is its dynamic value, `x` is logically zero if `z` is logically zero.

### Extended logical operators

Extended logical operators applies to any value and yield an untyped boolean value or an untyped boolean constant if both the operands are constants. The right operand is evaluated conditionally.

```
and   conditional AND   p and q  is  "if p is logically zero then false else q is logically zero" 
or    conditional OR    p or q   is  "if p is logically zero then q is not logically zero else false"
not   NOT               not p    is  "if p is logically zero"
```

### Conversions

A value `x` with type `html`, `css`, `js`, `json` or `markdown` cannot be converted to any other type. The following exception apply:

* if `x` has type `markdown`, `x` can be converted to the `html` type. The conversion transforms the string according to Markdown but the transformation is implementation-dependent. If `x` is constant, the transformation is done at compile-time.  

```
const c1 css = "red"
var c2 js = "var a = 1;"
var c3 markdown = "# title"
const c4 = `{"a":1}`

css(c1)     // a css value can be converted to css
html(c2)    // invalid; a js value cannot be converted html
html(c3)    // a markdown value can be converted to html
json(c4)    // a untyped constant value can be converted to json
```

## Statements

### Go statement

Implementation restriction: an implementation of Scriggo templates can disallow the "go" statement. When disallowed
a compile-time error occurs if used.

### If statement

The expression of the "if" statement is not limited to boolean types but can have any type.
If the expression does not evaluate to a logically zero value, the "if" branch is executed, otherwise, if present,
the "else" branch is executed.  

### For statement

The "for" statement is extended with a fourth form

```
ForStmt  = "for" [ Condition | ForClause | RangeClause | InClause ] Block .
InClause = Identifier "in" Expression .
```

A "for" statement with a "in" clause is the same of a "for" statement with a "range" clause with the first identifier only and the short assignment.

```
for i in s { }

// is the same of

for i := range s { } 
```

## Template files

A template file is a name-less package. Template files in the same directory are different name-less packages.

A template file imported with an "import" statement is a name-less package and can only contain "import" statements and declarations.
A template file with an "extends" statement is a name-less package and can only contain, after the "extends" statements, "import" statements and
declarations.

## Packages

Imported template files and template files with the "extends" statement are packages without a "package" clause and without name.

They contain only "import" statements and variable, constant, type and macro declarations between _{%_ and _%}_ or _{%%_ and _%%}_. Outside these markers, is allowed only bytes [TODO].

## Import-for declarations

The "import" statement has a new "import-for" form.

```
ImportDecl = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
ImportSpec = [ [ "." | PackageName ] ImportPath | ImportPath "for" { Identifier "," } ] .
```

In the "import-for" form, each identifier is an exported identifier of the package within the importing source file.

```
// Import the Banners and Menus names.
import "imp/components.html" for Banners, Menus

// Import the Index and HasPrefix names.
import "strings" for Index, HasPrefix
 ```
The <code>ImportPath</code> is first interpreted as a template file path. If a template file with this path does not
exist, it is interpreted as a package path.

### Package `"unsafe"`

There is not the "unsafe" package of the Go language in Scriggo templates.

{% end raw content %}