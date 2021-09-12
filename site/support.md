{% extends "/layouts/page.html" %}
{% var Page = "support" %}
{% macro Title string %}Support{% end %}
{% macro Content html %}

<h1>Get Support</h1>

<div class="support">
    <div class="enterprise">

    <h2>Enterprise</h2>

    <p>Priority private support by core developers</p>

  </div>
  <div class="community">

    <h2>Community</h2>
    <p>Free public support by the community</p>
    <div class="channel github-discussions">
        <div><a href="https://github.com/open2b/scriggo/discussions">{{ render "partials/github-logo.html" }}</a></div>
        <div><a href="https://github.com/open2b/scriggo/discussions">GitHub Discussions</a></div>
    </div>

  </div>
</div>

{% end macro %}
