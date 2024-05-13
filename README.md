[![Test](https://github.com/AccentDesign/gcss/actions/workflows/go-test.yml/badge.svg)](https://github.com/AccentDesign/gcss/actions/workflows/go-test.yml)

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

## What I don't need

* I don't need UI libs that are written for the purpose of JS frameworks. 
* I don't need linters when I have Go's static typing.
* I don't need javascript to generate CSS.
* I don't need templates with 400 css classes in them.
* I don't need css with more variables in them than actual css properties.

## What I do need

* Go, HTMX and maybe a few manual css styles.

## Usage

There are multiple ways you can use `gcss` in your project. For examples of property values, see the `style_test.go` file.

Write to a css file:

```go
var styles = []gcss.Style{
    {
        Selector: ".button",
        Props: gcss.Props{
            BackgroundColor: props.ColorRGBA(17, 24, 39, 255),
        },
    },
}

f, err := os.Create("styles.css")
if err != nil {
    panic(err)
}

for _, style := range styles {
    if err := style.CSS(f); err != nil {
        panic(err)
    }
}
```

Or as an example using a gin route:

```go
var styles = []gcss.Style{
    {
        Selector: ".button",
        Props: gcss.Props{
            BackgroundColor: props.ColorRGBA(17, 24, 39, 255),
        },
    },
}

router.Handle("GET", "/styles.css", func(c *gin.Context) {
    c.Writer.Header().Set("Content-Type", "text/css")
    for _, style := range styles {
        if err := style.CSS(c.Writer); err != nil {
            panic(err)
        }
    }
})


```

## Examples

CSS:
```bash
go run examples/styles/cli/main.go export -directory "examples/static/css"
```

Hugo:
```bash
cd examples
hugo server -D
```