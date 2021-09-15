{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the using statement and itea {% end %}
{% Article %}

{% raw doc %}

# Using and itea

The _using_ statement is used to evaluate a block of content and then use it, via a predeclared identifier named _itea_, in a statement or declaration.

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

is evaluated and the result, represented by the identifier _itea_, is assigned to the _navigation_ variable.

The body of the using statement can contain any content, including other declarations and statements, and _itea_ can be
used like any other identifier. In the following example, the body also contains a show statement and _itea_ is passed
to a function:  

```scriggo
{% sendmail(from, to, wordwrap(itea)); using %}
Hello {{ name }}, thanks for your kind reply.
{% end %}
```

Given this macro declaration:  

```scriggo
{% macro Dialog(title string, content html) %}
<div class="dialog">
  <h1>{{ title }}</h1>
  <div class="content">{{ content }}</div>
</div>
{% end %}
```

you can call it like this:

```scriggo
{% show Dialog("You wrote", itea); using %}
  <p>You wrote:</p>
  <cite>{{ message }}</cite>
{% end %}
```

In the previous example itea must have type [html](types#html) in order to be passed to the `Dialog` macro. The type of
itea depends on the context of the using statement. In the example above, assuming the context is HTML, itea has the
type [html](types#html).

You can specify the type for itea among the format types _string_, _html_, _markdown_, _css_, _js_ and _json_.

Look at this example:

```scriggo
{% show itea; using markdown %}
# The ancient art of tea
The Ancient Art of Tea is a delightful glimpse into
the philosophy, history and culture of tea in China.
{% end %}
```
<pre class="result">
&lt;h1&gt;The ancient art of tea&lt;/h1&gt;
&lt;p&gt;The Ancient Art of Tea is a delightful glimpse into
the philosophy, history and culture of tea in China.&lt;/p&gt;
</pre>

The body is evaluated as Markdown because we have explicitly specified the type of itea by writing `using markdown`. If the context of the show statement is HTML, the show statement then converts the value of itea from markdown to html and shows it.

## Itea with a macro type

In some cases, you may not want to immediately evaluate the body of the _using_ statement.

Look at this example:

```scriggo
{% show Header(itea); using %}
  <title>My Site</title>
  {{ socials() }}
{% end %}
```

The body is evaluated and the result is passed to the `Header` macro. If `Header` doesn't use its argument, Scriggo still executes the body and calls the `socials` function.

To avoid this, you can change `Header` by passing them a macro instead of a string. `Header` will only call the macro if it wants to render it. Now you can write `using macro` instead of `using` only: 

```scriggo
{% show Header(itea); using macro %}
  <title>My Site</title>
  {{ socials() }}
{% end %}
```

In this case _itea_ represents a macro with body the body of the _using_ statement. The body is evaluated and the 
`socials` macro is called only when `Header` calls the macro.

### Parameters

The following `Users` macro shows the users that have been passed to them. To show each user, it calls the macro
`getUser`, also passed as an argument:

```scriggo
{% macro Users(users []User, getUser macro(user User) html) %}
<div class="users">
  <h2>Users</h2>
  {% for user in users %}
  <div class="user">{{ getUser(user) }}</div>
  {% end %}
</div>
{% end %}
```

You can put together the call to `Users` with the macro showing a user, as follows:

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
