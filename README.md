[![Test](https://github.com/AccentDesign/gcss/actions/workflows/go-test.yml/badge.svg)](https://github.com/AccentDesign/gcss/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/AccentDesign/gcss)](https://goreportcard.com/report/github.com/AccentDesign/gcss)
<a href="https://pkg.go.dev/github.com/AccentDesign/gcss"><img src="https://img.shields.io/badge/Documentation%20on-pkg.go.dev-blue.svg"/></a>

<img src="banner.jpg" alt="banner" style="width: 100%; height: auto;">

# gcss

CSS written in Pure Go.

No JS builders, no preprocessors, no linters, no frameworks, no classes, no variables, no overrides, no plugins, no dependencies, no javascript, no templates, no bs, no nothing.

Just Go.

## Motivation

This is really just a bit of fun and a way to write CSS in Go. I wanted to see if it was possible and it is with ease.
I wanted to find a way to easily control the CSS from the server side and not have to worry about pre-building the css to take variables and stuff.
I didnt want to use UI libraries that are written for JS frameworks and I didn't want to use preprocessors or linters that add more steps to the build process.

Could I just use CSS? Yes of course and I will, but I wanted to see if I could write CSS in Go as this is what is compiling the rest of the project.

## Gopher

No it looks nothing like the Go gopher, but it's a gopher and I like it. It's the best I could get from the LM without giving up, [ideogram.ai (1400097641)](https://ideogram.ai/g/E-5MQp7QTPO4uyF9PvERzw/3).
## Next steps

The next steps for this project are to add more features to the CSS package.
This includes adding support for more CSS properties when the need arises.
What I don't want to do is to add support for all CSS functionality as some things are better in CSS, but I do want to be able to create 
a few UI components that are configurable using Go.

## What I don't need

* I don't need UI libs that are written for the purpose of JS frameworks. 
* I don't need linters when I have Go's static typing.
* I don't need javascript to generate CSS.
* I don't need templates with 400 css classes in them.
* I don't need css with more variables in them than actual css properties.

## What I do need

* Go's static typing.

## Installation

```bash
go get github.com/AccentDesign/gcss
```

## Usage

### Basic usage

`gcss` defines a `Style` type that can be used to hold the properties for a specific css selector, eg:

```go
style := gcss.Style{
    Selector: "body",
    Props: gcss.Props{
        BackgroundColor: props.ColorRGBA(0, 0, 0, 128),
    },
}
```

The `CSS` function on the `Style` is used to write the style to a `io.Writer`:

```go
style.CSS(os.Stdout)
```

which gives you:

```css
body{background-color:rgba(0,0,0,0.50);}
```

That's all there is to it. But it's not very useful on it's own I hear you say.

### Multiple styles

Well you can then use that to define a `Styles` type that can be used to hold multiple `Style` types:

```go
type Styles []gcss.Style

func (s Styles) CSS(w io.Writer) error {
    // handle your errors
    for _, style := range s {
        style.CSS(w)
    }
    return nil
}

styles := Styles{
    {
        Selector: "body",
        Props: gcss.Props{
            BackgroundColor: props.ColorRGBA(0, 0, 0, 128),
        },
    },
    {
        Selector: "main",
        Props: gcss.Props{
            Padding: props.UnitRem(8.5),
        },
    },
}

styles.CSS(os.Stdout)
```

which gives you:

```css
/* formatted for visibility */
body{
    background-color:rgba(0,0,0,0.50);
}
main{
    padding:8.500rem;
}
```

### Need a bit more? what about a dark and light theme? keep the last example in mind and read on.

Define a `Theme` type that can be used to hold attributes for a specific theme, eg:

```go
type Theme struct {
    MediaQuery string
    Background props.Color
}

func (t *Theme) CSS(w io.Writer) error {
    // handle your errors
    fmt.Fprintf(w, "%s{", t.MediaQuery)
    for _, style := range t.Styles() {
        style.CSS(w)
    }
    fmt.Fprint(w, "}")
}

// Styles returns the styles for the theme.
// Can be any number of styles you want and any number of functions
// you just need them in the CSS function to loop over.
func (t *Theme) Styles() Styles {
    return Styles{
        {
            Selector: "body",
            Props: gcss.Props{
                BackgroundColor: t.Background,
            },
        },
    }
}
```

Then you can define a `Stylesheet` type that can be used to hold multiple `Theme` types:

```go
type Stylesheet struct {
    Dark  *Theme
    Light *Theme
}

func (s *Stylesheet) CSS(w io.Writer) error {
    // handle your errors
    s.Dark.CSS(w)
    s.Light.CSS(w)
    return nil
}
```

Finally, you can use the `Stylesheet` type to write the css to a `io.Writer`:

```go
styles := Stylesheet{
    Dark: &Theme{
        MediaQuery: "@media (prefers-color-scheme: dark)",
        Background: props.ColorRGBA(0, 0, 0, 255),
    },
    Light: &Theme{
        MediaQuery: "@media (prefers-color-scheme: light)",
        Background: props.ColorRGBA(255, 255, 255, 255),
    },
}

styles.CSS(os.Stdout)
```

gives you:

```css
/* formatted for visibility */
@media (prefers-color-scheme: dark) {
    body{
        background-color:rgba(0,0,0,1.00);
    }
}
@media (prefers-color-scheme: light) {
    body{
        background-color:rgba(255,255,255,1.00);
    }
}
```

Hopefully this will get you going. The rest is up to you.

* Maybe create a button function that takes a `props.Color` and returns a Style.
* Or add extra `Styles` to the `Stylesheet` to additionally include non themed styles.
* It's all about how you construct the `Stylesheet` and use the `gcss.Style` type.
* If I could have created a `Stylesheet` type that fits well any use case at all I would have, but there is a world of possibility, so I left it up to you.

## The benefits

* Total control of the CSS from the server side.
* CSS doesn't have mixins, but you can create a function that returns a `Style` type and reuse it.
* Keeps the css free of variables.
* Keeps html free of classes like `bg-gray-50 text-black dark:bg-slate-800 dark:text-white` and eliminates the need to remember to add the dark variant.
* I recently saw a button component on an html page 10 times with over 1800 characters in the class attribute of each. This is not maintainable nor debuggable.
* Keeps the css clean and easy to debug with no overrides like the above.

## Examples

For example usage see the [examples](./examples) directory that include:

* [CSS resets](./examples/css-resets) - A simple example collection of css resets.
* [Templ integration](./examples/integration-templ) - An example of how to load styles from gcss with the [templ](https://templ.guide) package.
* [Media queries](./examples/media-queries) - An example of how to use media queries.
* [Themed CSS using multiple HTTP handlers](./examples/themed-multiple-http-handlers) - An example of how to use multiple http handlers to serve different themes.
* [Themed CSS using a single HTTP handler](./examples/themed-single-http-handler) - An example of how to use a single http handler to serve different themes using media queries.
* [Write to a file](./examples/to-file) - An example of how to write to a file.
* [Write to an HTTP handler](./examples/to-http-handler) - An example of how to write to an http handler.
* [Write to stdout](./examples/to-stdout) - An example of how to write to stdout.

## Contributing

If you would like to contribute to this project, please open an issue or a pull request. We welcome all contributions and ideas.

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
<link rel="stylesheet" href="base.css">
<link rel="stylesheet" href="gcss-styles.css">
<link rel="stylesheet" href="utils.css">
```

Try to keep the specificity of the `gcss` styles to 1 by using single classes this will ensure any `tailwindcss` utilities
will be able to overwrite your styles where required.
