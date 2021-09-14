{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the using statement and itea {% end %}
{% Article %}

{% raw doc %}

# Using and itea

The using statement can be used in combination with other statements and declarations to render a piece of code and
use the rendered code via the predeclared identifier _itea_.

In this example:

```scriggo
{% navigation := itea; using %}
  <li><a href="/">Home</a>
  <li><a href="/about-us">About Us</a>
{% end using %}
```

the body of the _using_ statement:

```scriggo
  <li><a href="/">Home</a>
  <li><a href="/about-us">About Us</a>
```

is rendered and the result, represented by the predeclared identifier _itea_, is assigned to the _navigation_ variable.

Therefore, the body of the _using_ statement is rendered first, and after that the statement or declaration that
precedes the _using_ statement is evaluated with the _itea_ identifier representing the rendered value.

In this other example, _itea_ is used as argument for a function call:  

```scriggo
{% sendmail(from, name, wordwrap(itea)); using %}
Hello {{ name }}, thanks for your kind reply.
{% end %}
```

### One more example

Given this macro declaration:  

```scriggo
{% macro Dialog(title string, content html) %}
<div class="dialog">
  <h1>{{ title }}</h1>
  <div class="content">{{ content }}</div>
</div>
{% end %}
```

You can show the dialog as follows:

```scriggo
{% show Dialog("You wrote", itea); using %}
  You wrote:
  {{ message }}
{% end %}
```

Without the _using_ statement, you would have had to write instead:

```scriggo
{% macro content %}
  You wrote:
  {{ message }}
{% end %}

{{ Dialog(title, content()) }}
```

## Using macro

In some cases one would not want to immediately evaluate the body of the _using_ statement.

Look at this example:

```scriggo
{% show Header(itea); using %}
  <title>My Site</title>
  {{ socials() }}
{% end %}
```

The body is rendered and its value is passed to the `Header` macro. If in certain circumstances you implement the
`Header` macro so that it doesn't use its argument, the body was rendered anyway and the `socials` macro was called
anyway.

In this case you could change `Header` to take a macro instead of a string value, so `Header` calls the macro only if it
wants to render it. Now you can write `using macro` instead of `using` only: 

```scriggo
{% show Header(itea); using macro %}
  <title>My Site</title>
  {{ socials() }}
{% end %}
```

In this case _itea_ represents a macro with body the body of the _using_ statement. Only if and when `Header` calls
the macro, the body is rendered and consequently the `socials` macro id called.

Below we will look at a more complex use case.

### Using macro with parameters

The following macro `Users` renders a slice of `User` using a macro passed as argument to render each single user:   

```scriggo
{% macro Users(users []User, one macro(user User) html) %}
<div class="users">
  <h2>Users</h2>
  {% for user in users %}
  <div class="user">{{ one(user) }}</div>
  {% end %}
</div>
{% end %}
```

You can call this macro to show all the users in this way:

```scriggo
{% show Users(users, itea); using macro(user User) %}
<dl>
  <dt>First name</dt>
  <dd>{{ user.FirstName }}</dd>
  <dt>Lat name</dt>
  <dd>{{ user.LastName }}</dd>
</dl>
{% end %}
```

## Learn more about _using_

Now that you have learned, with examples of increasing complexity, how _using_ works, you can find out more by reading
the template specification of the [using statement](/templates/specification#using-statement).

{% end raw doc %}
