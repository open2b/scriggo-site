{% extends "/layouts/page.html" %}
{% macro Title string %}Scriggo News{% end %}
{% Content %}

# News

## Released Scriggo v0.52.0
September 22, 2021

For more details and binary releases see https://github.com/open2b/scriggo/releases/tag/v0.52.0

**This version breaks the compatibility with previous versions.**

#### Add scriggo command binaries.

Starting from this version, the binaries of the scriggo command are available on GitHub.

#### Replace env.Exit with env.Stop (breaking change).

This change replaces the Exit method of the native.Env interface
with the Stop method. The Stop method has the same behaviour of
Exit but accepts an error instead of a status code.

The following code:

    env.Exit(0)
    env.Exit(2)

with this change, can be rewritten as

    env.Stop(nil)
    env.Stop(scriggo.NewExitError(2, nil))

#### Add the unsafeconv package as built-in in templates.

#### Other fixes and improvements.