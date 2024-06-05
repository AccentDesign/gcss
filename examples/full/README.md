# Full example with mutex

This example hase the following:

* A `StyleSheet` that generates global resets, base styles and includes themes.
* The `Stylesheet` has a `Mutex` that generates the CSS only once to avoid multiple builds.
* Both `body` and `buttons` have styles attached to the `StyleSheet` as well as the `Theme` to ensure the css is loaded in the most appropriate places.

## Example output

```css
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

body {
    min-height: 100vh;
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