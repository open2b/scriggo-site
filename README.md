# Scriggo website

This repository is the complete scriggo.com website.

## Licensing

The code in this repository is licensed under a BSD license, the same of Scriggo. The content on the scriggo.com site,
except where otherwise noted, is licensed under the [Creative Commons Attribution 4.0 International license](https://creativecommons.org/licenses/by/4.0/).

The Scriggo logo is copyright © Open2b Software Snc.

Gopher logos are based on an original work by Renée French.

## Contributing

Contributions to the scriggo.com site are greatly appreciated, as new documentation, bug reports, any king of suggestion:

* if you find an issue on the site, [open an issue](https://github.com/open2b/scriggo/issues/new) on Scriggo.
* if you want to write new documentation, as "Switch from X to Scriggo", [discuss with us](https://github.com/open2b/scriggo/discussions) and after it is approved open a PR on this repository.
* if you use Scriggo at work, let us know.
### Running the website

The scriggo.com site is a Scriggo template. You can use the [scriggo command](https://scriggo.com/scriggo-command) to run it locally:

```
cd site
scriggo serve
```

### Build the website

To build the pages of the website, execute the following command in the root directory of the repository:

```
go run .
```

It will create a new directory named "public" with the compiled site pages. If the directory already exists, it will be deleted first.