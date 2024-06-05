package styles

import (
	"fmt"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"io"
)

type Media struct {
	Query  string
	Styles Styles
}

// CSS writes the media query to the writer.
func (m *Media) CSS(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s{", m.Query); err != nil {
		return err
	}
	for _, style := range m.Styles {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}
	return nil
}

// Media returns the media queries for the stylesheet.
func (ss *StyleSheet) Media() []Media {
	return []Media{
		{
			Query: "@media screen and (max-width: 768px)",
			Styles: Styles{
				{
					Selector: "main",
					Props: gcss.Props{
						Display: props.DisplayGrid,
						Gap:     variables.Size6,
						Padding: variables.Size8,
					},
				},
			},
		},
		{
			Query: "@media screen and (min-width: 769px)",
			Styles: Styles{
				{
					Selector: "main",
					Props: gcss.Props{
						Display: props.DisplayGrid,
						Gap:     variables.Size8,
						Padding: variables.Size16,
					},
				},
			},
		},
	}
}
