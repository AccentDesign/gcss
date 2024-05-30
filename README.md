[![Test](https://github.com/AccentDesign/gcss/actions/workflows/go-test.yml/badge.svg)](https://github.com/AccentDesign/gcss/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/AccentDesign/gcss)](https://goreportcard.com/report/github.com/AccentDesign/gcss)
<a href="https://github.com/AccentDesign/gcss-theme"><img src="https://img.shields.io/badge/Example%20on-Github-blue.svg"/></a>

# gcss

CSS written in Pure Go.

No JS builders, no preprocessors, no linters, no frameworks, no classes, no variables, no overrides, no plugins, no dependencies, no javascript, no templates, no bs, no nothing.

Just Go.

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

## Installation

```bash
go get github.com/AccentDesign/gcss
```

## Usage

There are multiple ways you can use `gcss` in your project. For examples of property values, see the `style_test.go` file.

### File

write to a file:

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

### Http handler

write using an http handler:

```go
var styles = []gcss.Style{
    {
        Selector: ".button",
        Props: gcss.Props{
            BackgroundColor: props.ColorRGBA(17, 24, 39, 255),
        },
    },
}

http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/css")
    for _, style := range styles {
        if err := style.CSS(w); err != nil {
            panic(err)
        }
    }
})
```

### Theming

Example of theming using `gcss`, the below could write to 2 separate css files or route handlers for dark and light themes.

```go
package theme

import (
    "io"
    "github.com/AccentDesign/gcss"
    "github.com/AccentDesign/gcss/props"
    "github.com/AccentDesign/gcss/variables"
)

// Theme is a struct that holds the theme properties.
type Theme struct {
    Background    props.Color
    Color         props.Color
}

// WriteCSS writes the css for the theme to the writer.
func (t *Theme) WriteCSS(w io.Writer) {
    for _, e := range t.styles() {
        if err := e.CSS(w); err != nil {
            panic(err)
        }
    }
}

// styles returns the styles for the theme.
func (t *Theme) styles() []gcss.Style {
    return []gcss.Style{
        {
            Selector: "body",
            Props: gcss.Props{
                BackgroundColor: t.Background,
                Color:           t.Color,
            },
        },
    }
}

// NewDarkTheme creates a new dark theme.
func NewDarkTheme() *Theme {
    return &Theme{
        Background: variables.Zinc900,
        Color:      variables.Zinc50,
    }
}

// NewLightTheme creates a new light theme.
func NewLightTheme() *Theme {
    return &Theme{
        Background: variables.Zinc50,
        Color:      variables.Zinc900,
    }
}
```

Then switch between them using media queries in your HTML.

```html
<!-- Light theme -->
<link rel="stylesheet" href="light.css" media="(prefers-color-scheme: light)" />

<!-- Dark theme -->
<link rel="stylesheet" href="dark.css" media="(prefers-color-scheme: dark)" />
```

The benefit of this:

* Keeps the css free of variables
* Keeps html free of classes like `bg-gray-50 text-black dark:bg-slate-800 dark:text-white` and eliminates the need to remember to add the dark variant
* Keeps the css clean and easy to debug with no overrides like the above
* Allows for easy theming based on server side logic

## Mix it up with other CSS frameworks

You can mix `gcss` with other CSS frameworks like `tailwindcss` for example:

separate the css files into base and utils:

```css
/* base.css */
@tailwind base;
```

```css
/* utils.css */
@tailwind utilities;
```

Then add the `gcss` styles in between in your html:

```html
<link rel="stylesheet" href="base.css" />
<link rel="stylesheet" href="gcss-styles.css" />
<link rel="stylesheet" href="utils.css" />
```

Try to keep the specificity of the `gcss` styles to 1 by using single classes this will ensure any `tailwindcss` utilities
will be able to overwrite your styles where required.
