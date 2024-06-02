package main

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"io"
	"os"
)

type (
	Styles []gcss.Style
	Media  struct {
		Query  string
		Styles Styles
	}
	Stylesheet struct {
		Styles Styles
		Medias []Media
	}
)

// WriteCSS writes the CSS for the media query to the writer
func (m Media) WriteCSS(w io.Writer) error {
	if _, err := io.WriteString(w, m.Query); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "{"); err != nil {
		return err
	}
	for _, style := range m.Styles {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "}"); err != nil {
		return err
	}
	return nil
}

// WriteCSS writes the CSS for the stylesheet to the writer
func (ss Stylesheet) WriteCSS(w io.Writer) error {
	// Write the base styles first
	for _, style := range ss.Styles {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	// Write the media queries next
	for _, media := range ss.Medias {
		if err := media.WriteCSS(w); err != nil {
			return err
		}
	}
	return nil
}

var (
	base = Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				Margin: props.UnitRaw(0),
			},
		},
	}
	screen736 = Media{
		Query: "@media only screen and (max-device-width: 736px)",
		Styles: Styles{
			{
				Selector: "main",
				Props: gcss.Props{
					Padding: props.UnitRaw(0),
				},
			},
		},
	}
	stylesheet = Stylesheet{
		Styles: base,
		Medias: []Media{screen736},
	}
)

// This is just a basic idea of how you could structure your CSS
// the goal hear is just to wrap the css how you wish with what ever you wish
// construct your stylesheet to suit your needs.
// The end goal is just to call CSS on each style with the object to write to.
func main() {
	file, err := os.Create("media.css")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := stylesheet.WriteCSS(file); err != nil {
		panic(err)
	}
}
