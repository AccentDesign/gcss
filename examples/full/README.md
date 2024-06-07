# Full example with mutex

This example hase the following:

* A `StyleSheet` that has has global resets, base styles, media queries for devices, themes and a `Mutex` that generates the CSS only once to avoid multiple builds.
* The `main` element has base styles and media queries for different screen sizes.
* Both `body` and `buttons` have base styles as well as themed for the likes of background and foreground variations.

## Usage

For an example of adding form styles.

1. Create a new `form.go` file. and add the following code (or add what you need):

```go
// Form returns the styles for buttons for the base stylesheet.
func (ss *StyleSheet) Form() Styles {
	return Styles{}
}

// Form returns the styles for the layout for the media.
func (m *Media) Form() Styles {
	return Styles{}
}

// Form returns the styles for the layout for the theme.
func (t *Theme) Form() Styles {
	return Styles{}
}
```

2. Then add the functions to the `CSS` function in `StyleSheet`, `Media` and `Theme`.


## HTML

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1" name="viewport" />
    <meta content="Example starting point written in gcss." name="description" />
    <title>Theme</title>
    <link rel="stylesheet" href="/stylesheet.css">
</head>
<body>
<main>
    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
    <div>
        <button class="button button-primary">Click me</button>
    </div>
</main>
</body>
</html>
```

## CSS

```css
/* formatted for visual clarity */

/* resets */
*, ::after, ::before, ::backdrop, ::file-selector-button {
    border: 0 solid;
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

html, :host {
    font-family: ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
    line-height: 1.5;
    -webkit-text-size-adjust: 100%;
    tab-size: 4;
}

body {
    line-height: inherit;
}

/* base */
body {
    min-height: 100vh;
}

main {
    display: grid;
}

.button {
    align-items: center;
    border-radius: 0.375rem;
    display: inline-flex;
    font-size: 0.875rem;
    font-weight: 500;
    height: 2.500rem;
    justify-content: center;
    line-height: 1.250rem;
    padding-bottom: 0.500rem;
    padding-left: 1.000rem;
    padding-right: 1.000rem;
    padding-top: 0.500rem;
}

/* media */
@media (max-width: 768px) {
    main {
        padding: 2.000rem;
        row-gap: 1.500rem;
    }
}

@media (min-width: 769px) {
    main {
        padding: 4.000rem;
        row-gap: 2.000rem;
    }
}

/* themes */
@media (prefers-color-scheme: light) {
    body {
        background-color: rgba(255, 255, 255, 1.00);
        color: rgba(23, 23, 23, 1.00);
    }

    .button-primary {
        background-color: rgba(23, 23, 23, 1.00);
        color: rgba(255, 255, 255, 1.00);
    }
}

@media (prefers-color-scheme: dark) {
    body {
        background-color: rgba(23, 23, 23, 1.00);
        color: rgba(245, 245, 245, 1.00);
    }

    .button-primary {
        background-color: rgba(255, 255, 255, 1.00);
        color: rgba(23, 23, 23, 1.00);
    }
}

```