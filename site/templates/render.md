{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: the render operator{% end %}
{% Article %}

{% raw doc %}

# Render

The _render_ operator, when used between `{{` and `}}`:

```scriggo
{{ render path }}
```

renders the template file with the named path and shows its content. A file rendered with render is called _partial_
because contains only a part of the content to render:

```scriggo
{{ render "/partials/promotion.html" }}
```

The path of the partial file can also be relative to the file that contains the render operator:

```scriggo
{{ render "../header.html" }}
```

When a partial file is rendered, the code of the file does not see the variables declared in the file that contain the
render operator, and it does not see the other declarations in this file. For this purpose, macros with parameters are
used, such as:

```scriggo
{{ Image("picture.jpg", 400, 500) }}
```

See [macro](macro) for details.

## More general use

Render is a Scriggo operator and as such it can be used in expressions, not only between {{ and }}, like any other
operator. Its evaluation returns a string. For example, it is possible to assign its evaluation for later use:

```scriggo
{% var email = render "email.html" %}
```

or use it as part of another expression:

```scriggo
{%%
    if discount {
        show "Offer: " + render "offer.html"
    }
%%}
```

## Non-existent files

If the file to render does not exist, a compilation error occurs. On the other hand, if you want to handle the case in
which the file does not exist, you can use a _default_ expression:

```scriggo
{% promo := render "extra.html" default "not found" %}
```

If a template file with path "extra.html" exists, `promo` will contain the result of its rendering, otherwise it will be
`"not found"`.

{% end raw doc %}
