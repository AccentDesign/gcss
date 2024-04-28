# gcss

CSS written in Go.

## Motivation

The `gcss` project is a unique approach to writing CSS using Go.
The motivation behind this project is not to replace CSS entirely,
but to provide a way to write CSS in Go.
This allows developers to leverage Go's features to generate CSS. For instance,
Go's `text/template` or `templ` package can be used to generate CSS based on some data.
This opens up new possibilities for dynamic and data-driven styling in web development.

## Next steps

The next steps for this project are to add more features to the CSS package.
This includes adding support for more CSS properties and maybe mixins.
What I don't want to do is to add support for all CSS functionality as some things are better in CSS, but I do want to be able to create 
a few UI components that are configurable using Go.

When I know what it all looks like I will write tests for it. But it's pretty basic stuff and I'm a pretty basic person so things should be ok.
## What I don't need

* I don't need UI libs that are written for the purpose of JS frameworks. 
* I don't need linters when I have Go's static typing.
* I don't need javascript to generate CSS.
* I don't need templates with 400 css classes in them.
* I don't need css with more variables in them than actual css properties.

## What I do need

* Go, HTMX and maybe a few manual css styles.

## Usage

just clone the repo for now and bugger about for yourself. If you have any good ideas I'm listening

## Docs

```bash
go install github.com/cosmtrek/air@latest
```
```bash
air -build.cmd 'go build -o ./tmp/main docs/docs/main.go'
```