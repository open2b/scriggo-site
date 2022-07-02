{% extends "/layouts/doc.html" %}
{% macro Title string %}The Scriggo Template Specification{% end %}
{% Article %}

{% raw content %}

# The Scriggo Template Specification

<nav style="display: flex;">
    <dl>
        <dt><a href="#introduction">Introduction</a></dt>
        <dt><a href="#template-source-code-representation">Template source code representation</a></dt>
        <dt><a href="#lexical-elements">Lexical elements</a></dt>
        <dd><a href="#comments">Comments</a></dd>
        <dt><a href="#types">Types</a></dt>
        <dd><a href="#format-types">Format types</a></dd>
        <dd><a href="#macro-types">Macro types</a></dd>
        <dt><a href="#contexts">Contexts</a></dt>
        <dt><a href="#declarations">Declarations</a></dt>
        <dd><a href="#macro-declarations">Macro declarations</a></dd>
        <dd><a href="#predeclared-identifiers">Predeclared identifiers</a></dd>
        <dt><a href="#expressions">Expressions</a></dt>
        <dd><a href="#primary-expressions">Primary expressions</a></dd>
        <dd><a href="#default-expression">Default expression</a></dd>
        <dd><a href="#render-operator">Render operator</a></dd>
        <dd><a href="#contains-operators">Contains operators</a></dd>
        <dd><a href="#operators">Operators</a></dd>
        <dd><a href="#truthful-values">Truthful values</a></dd>
        <dd><a href="#extended-logical-operators">Extended logical operators</a></dd>
        <dd><a href="#conversions">Conversions</a></dd>
    </dl>
    <dl>
        <dt><a href="#statements">Statements</a></dt>
        <dd><a href="#go-statement">Go statement</a></dd>
        <dd><a href="#if-statement">If statement</a></dd>
        <dd><a href="#for-statement">For statement</a></dd>
        <dd><a href="#show-statement">Show statement</a></dd>
        <dd><a href="#using-statement">Using statement</a></dd>
        <dd><a href="#itea">Itea</a></dd>
        <dd><a href="#raw-statement">Raw statement</a></dd>
        <dt><a href="#template-files">Template files</a></dt>
        <dd><a href="#extends-declarations">Extends declarations</a></dd>
        <dd><a href="#import-declarations">Import declarations</a></dd>
        <dt><a href="#template-initialization-and-execution">Template initialization and execution</a></dt>
        <dt><a href="#package-unsafe">Package "unsafe"</a></dt>
    </dl>
</nav>

## Introduction

This is a reference manual for Scriggo templates.

The Scriggo template specification is based on the Go language specification, see <a href="https://go.dev/ref/spec">go.dev/ref/spec</a>. The Go specification applies, except as written in this specification.

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
### Keywords

The following keywords are added to the keywords of Go:

```
and          extends      not          render
contains     in           or           show
end          macro        raw          using
```

```
{{    }}    {%    %}
{%%   %%}   {#    #}
```

## Types

### Format types

A format type represents a sets of string values that can have a special treatment when used with `show` statement. The `string` type is a format type.

```
FormatType = "string" | "html" | "css" | "js" | "json" | "markdown" .
```

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

## Contexts

Templates are autoescaping

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

### Macro declarations

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

A distraction free declaration is a macro declaration where the body of the macro end at the end of the file.
A distraction free declaration does not have the `macro` keyword, does not have parameters and have an implicit
return value. The name of the macro in a distraction free declaration is an exported identifier.

```
DistractionFreeMacroDecl = "{%" DistractionFreeMacroName "%}" Content EOF .
DistractionFreeMacroName = exported_identifier .
```

A distraction free declaration can only be used in a file with an extends declaration.   

```
{ extends "layout.html" %}
{% Article %}

This is the article content,

{% a := 5 %} {# a is defined in the scope of the body of the Article macro #}

and more article content.
```

### Predeclared identifiers

The predeclared type identifiers _html_, _css_, _js_, _json_ and _markdown_, with the type _string_ are called format types:

```
Format types:
    string html css js json markdown
```

The format types are the only types that can be used as macro result type.

## Expressions

### Primary expressions

The _default expression_ is added to primary expressions.

```
PrimaryExpr =
	Operand |
	Conversion |
	MethodExpr |
	DefaultExpr |
	PrimaryExpr Selector |
	PrimaryExpr Index |
	PrimaryExpr Slice |
	PrimaryExpr TypeAssertion |
	PrimaryExpr Arguments .

DefaultExpr = ( identifier | identifier Arguments | Render ) "default" Expression .
```

### Default expression

The left expression `x` of the _default expression_

```
x default e
```

is _resolved_ if one of the following conditions applies:

* `x` is a predeclared identifier of a variable, constant or function
* `x` is a render expression and the file to render exists
* `x` is a call expression `M(a1, a2, … an)` with `M` an identifier representing a macro declared in the scope of a file that extends the file of the default expression

If the left expression `x` is resolved, `x default e` is replaced with `x` otherwise it is replaced with `e`.

If it is replaced based on its left expression, it must still be valid in its context if it were replaced based on its right expression.

If it is replaced based on its right expression and `x` is a function call, its arguments must be assignable to the empty interface type and if the last argument is followed by `...`, the type of the last argument must be assignable to a slice type `[]T`.


A default expression may only be used in the following cases:

1. Variable declaration as in `var a = x default e`.
2. Short variable declaration as in `a := x default e`.
3. Typed constant declaration as in `const c = x default e`.
4. Assignment as in `v = x default e`.
5. Show statement as  in `show x default e`.

In a declaration, the type of the declared variable or constant is the type it would have if the default expression were replaced with its right expression `e` only. In a constant declaration without a type and with `e` untyped, the declared constant is untyped with the same kind of `e`, and  also `x` must be untyped with the same kind of `e`.

In a declaration or assignment, if the right expression `x` of the default expression is a call or render expression, the type of the right expression `e` must be a format type.

For example:

``` 
var filters []Filter = filters default nil
```

```
{{ Head() default "" }}
```

```
{% show Head() default this using %} ... {% end %}
```

```
{{ render "specials.html" default "no specials" }}
```

### Render operator

For a string literal s, the expression

```
render s
```

renders the template file with path s, and it's evaluated to a string value with the result of the rendering. The type of the resulted value is the format type relative to the format of the rendered file.

```
{# s has type string if "codes.txt" has format Text #}
{% s := render "codes.txt" %}

{# the show expression has type html if "footer.html" has format HTML #}
{{ render "footer.html" }}
```

### Contains operators

Scriggo templates have <code>contains</code> and <code>not contains</code> operators. The following rules applies:

* For a value `v` of type `[]T` or `[n]T` and a value `x`, the expression `v contains x` yield an untyped boolean; `true` if an element of `v` is equals to `x`. `x` must be assignable to `T` and comparable to the values of `T`.

* For a map `v` with key values of type `T` and a value `x` of type `T`, the expression `v contains x` yield an untyped boolean; `true` if `v` has `x` as key.

* For string values `v` and `x`, the expression `v contains x` yield an untyped boolean; `true` if `x` is within `v`.

* For a string value `v` and a value `x` of type `rune`, the expression `v contains x` yield an untyped boolean; `true` if the Unicode code point `x` is within `v`.

* For a value `v` of type `[]byte` and a value `x` of type `rune`, the expression `v contains x` yield an untyped boolean; `true` if the Unicode code point `x` is within `v`.

If `!(v contains y)` is a valid expression evaluated to the value `b`, also `v not contains y` is a valid expression and it is evaluated to the value `!b`.

### Operators

The operators are extended with the operators `and`, `or`, `not`, `render`, `contains` and `not contains`.

```
binary_op  = "||" | "&&" | rel_op | add_op | mul_op | "or" | "and" | "contains" | "not contains" .
unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" | "not" | "render" .
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

### Truthful values

A value `x` of type `T` is _truthful_ unless one of the following conditions applies:

* `x` is the zero value of `T`,
* `x` is an empty slice or map,
* `T` is an interface and the dynamic value of `x` is not truthful,
* `T` is a struct or pointer to struct that implements `interface{ IsTrue() bool }` and `x.IsTrue()` is false.

Implementation restriction: A compiler can not assume that `x.IsTrue()` evaluates to the same value for the same value
of `x`.

### Extended logical operators

Extended logical operators applies to any value and yield an untyped boolean value or an untyped boolean constant if both the operands are constants. The right operand is evaluated conditionally.

```
and    conditional AND    p and q  is  "if p is truthful then q is truthful else false"
or     conditional OR     p or q   is  "if p is truthful then true else q is truthful"
not    NOT                not p    is  "if p is not truthful"
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
If the expression evaluate to truthful value, the "if" branch is executed, otherwise, if present,
the "else" branch is executed.

### For statement

The "for" statement is extended with a fourth form and an "else" block.

```
ForStmt  = "for" ( [ Condition | ForClause ] Block | [ RangeClause | InClause ] Block [ "else" Block ] ) .
InClause = Identifier "in" Expression .
```

A "for" statement with an "in" clause is the same of a "for" statement with a "range" clause with the first identifier only and the short assignment.

```
for i in s { }

// is the same of

for i := range s { } 
```

A "for" statement with an "else" block, the "else" block is executed if the "for" block is not executed.

### Show statement

The _show_ statement formats the expressions, according to the expression type and the context. If _show_ is used
in a macro, the formatted values are appended to the macro response value, otherwise the formatted values are printed
to the template output.

In content, the show statement can be also written using the short form.

Implementation restriction: a Scriggo templates implementation may not parse the short form. If not parsed, {{
and }} have no special meaning.

```
ShowStmt      = "show" Expression { "," Expression } .
ShortShowStmt = "{{" Expression "}}" .
```

The _show_ statement cannot be used in a function literal.

Examples:

```
{% show "the document has ", count, " characters" %}

{%%
    if price > 0 {
        show "price: ", price
    }
%%} 

{% macro Product(product Product) %}
    Name:  {{ product.Name }}
    Price: {{ product.Price }}
{% end %}

{% func(s string) {
    show s // invalid, cannot be used in a function literal.
}() %}
```

### Using statement

A using statement renders its content and then executes its affected statement. Within the affected statement the
predeclared identifier `itea` represents the rendered content.

```
UsingStmt    = "{%" AffectedStmt ";" "using" [ IteaType ] "%}" Content "{%" "end" [ "using" ] "%}" .
AffectedStmt =  VarDecl | ExpressionStmt | SendStmt | Assignment | ShowStmt .
IteaType     = ( FormatType | "macro" [  Parameters ] [ FormatType ] ) .
```

If a type is present, `itea` has that type. Otherwise, `itea` has the format type corresponding to the context of the
using statement. If the type is a macro type, the result type can be omitted. If omitted, the result type is the format
type corresponding to the context of the using statement.

```
{# there is a type, itea has type markdown #}
{% show itea; using markdown %}## {{ title }}{% end %}

{# if the context is HTML, itea has type html #}
{% var a = itea; using %}<a href="{{ src }}">{{ description }}</a>{% end %}

{# if the context is Text, itea has type macro(string) string #}  
{% names(itea); using macro(name string) %}Name: {{ name }}{% end %}
```

If the type of `itea` is a format type, `itea` is a string value containing the rendered content.

If the type of `itea` is a macro type, `itea` is a macro with content the content of the using statement. In the content
are defined the parameter of the macro if any.


### Itea

Within the affected statement of a [using statement](#using-statement), the predeclared identifier `itea` represents the
value resulted from the rendering of the macro's content. It can have a format type or a macro type.

### Raw statement

In a _raw_ statement the content is rendered as is, the tokens _{%_, _%}_, _{%%_, _%%}_, _{{_, _}}_, _{#_ and _#}_ have no special meaning.

```
RawStmt         = MarkedRawStmt | UnmarkedRawStmt .
MarkedRawStmt   = "{%" "raw" Marker [ Tag ] "%}" Content "{%" "end" "raw" Marker "%}" .
UnmarkedRawStmt = "{%" "raw" [ Tag ] "%}" Content "{%" "end" [ "raw" ] "%}" .
Marker          = letter { letter | unicode_digit } .
Tag             = string_lit .
```

The following code

```
{% raw %}
   {% if discount %}
     promotion
   {% end if %}
{% end %}
```

is rendered as

```
   {% if discount %}
     promotion
   {% end if %}
```

Marked raw statements ends at the first occurrence of `{% end raw marker %}` where `marker` is its marker. For example,
the following code

```
{% raw doc %}
   Use a raw statement to render unparsed text as
   {% raw %} # title {% end %}
   {% raw cycle %}
     {% for item in items %} {{ item }} {% end %}
   {% end raw cycle %}
{% end raw doc %}
```

is rendered as

```
   Use a raw statement to render unparsed text as
   {% raw %} # title {% end %}
   {% raw cycle %}
     {% for item in items %} {{ item }} {% end %}
   {% end raw cycle %}
```

A _raw_ statement can have a string literal _tag_. As for struct tags, an empty tag string is equivalent to an absent tag.

```
{% raw `lang:"javascript"` %} var a = 5; console.log(a); {% end %}

{# an empty tag is equilavent to an absent tag #}
{% raw "" %} <b! data {% end %}
```

## Template files

Templates are constructed linking together template files. A template file has a path, rooted at the template root, and a format: _Text_, _HTML_, _Markdown_, _CSS_, _JS_ or _JSON_. The format of a file is implementation dependant.

## Extends declarations

An extends declaration states that the _extended_ file depends on functionality of the file containing the declaration
and enables access to exported identifiers of this file. The exports names the path of the extended file.

```
ExtendsDecl = "extends" ExtendsPath
ExtendsPath = string_lit .
```

The `ExtendsPath` can be relative to the directory of the file containing the declaration or can be rooted at the root of the template. An extends declaration must be the first declaration in the file.

For a file containing an extends declaration, the following rules applies:

* it can only contain declarations,
* it cannot contain more that one extends declaration,
* outside {% and %}, {%% and %%} and {# and #}, it can only contain white space: spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A).
* it cannot be imported with the import declaration or rendered with the render operator.

## Import declarations

The `ImportPath` of an import declaration is first interpreted as the path of a template file. If a template file with
this path does not exist, it is interpreted as a package path.

If `ImportPath` is interpreted as a template file, the Go form without a package name is the same of the form with the
explicit period.

```
// if "header.html" resolves to a template file,
// the following import declarations are equivament.

import "header.html"
import . "header.html"
```

A new form of import declaration is added with an explicit list of exported identifiers. Only these identifiers are
imported within the importing source file.

```
// Import the Banners and Menus names.
import "imp/components.html" for Banners, Menus

// Import the Index and HasPrefix names.
import "strings" for Index, HasPrefix
```

```
ImportDecl    = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
ImportSpec    = [ [ "." | PackageName ] ImportPath | ImportForSpec ] .
ImportForSpec = ImportPath "for" Identifier { "," Identifier } .
```

## Template initialization and execution

The template execution starts from a template file. If the file contains an extends declaration, the execution starts
from the extended file. If the extended file extends another file, the execution starts from that other file, and so on
as long as an extended file does not extend any other files.

### Package `"unsafe"`

There is not the "unsafe" package of the Go language in Scriggo templates.

{% end raw content %}
