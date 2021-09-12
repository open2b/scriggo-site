{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo templates: comments{% end %}
{% Article %}

{% raw doc %}

# Comments

If you want to add a comment or to comment a piece of code in a file, enclose the text between `{#` and `#}`:

```scriggo
{# this is a comment #}

{# TODO: review the following code.

    {% for product in products %}
        {{ product.Name }}
    {% end %}
#}
```

Unlike HTML comments, as `<!-- comment -->`, Scriggo comments, as `{# comment #}`, are discarded when the template is
rendered and consequently are not present in the rendered code.

Scriggo comments can be nested, for example:

```scriggo
{#
    {# show the promotion #}
    Promotion: {{ promotion }}
#}
```

{% end raw doc %}
