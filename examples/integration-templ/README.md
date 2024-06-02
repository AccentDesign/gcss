# Templ Integration

This is an example of how to integrate the awesome [Templ](https://templ.guide) with `gcss`.
It's on the basic side, but it should give you a good idea of how to get started.

I will include a better example in the future, but for now, this should be enough to get you started.

## Notes

The Stylesheet type is there till we decide if it will be in the package itself or not.
Though this is unlikely to change till there is a clear need for it.

When you render the template, you will see that it renders two buttons that share the same style. 
This is done through the usage of the [Once](https://templ.guide/syntax-and-usage/render-once) handler provided by templ
and allows you to selectively render styles when they are required.
