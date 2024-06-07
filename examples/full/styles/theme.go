package styles

import (
	"bytes"
	"fmt"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"io"
	"slices"
)

type (
	Theme struct {
		MediaQuery        string
		Background        props.Color
		Foreground        props.Color
		Primary           props.Color
		PrimaryForeground props.Color
	}
)

var (
	lightTheme = &Theme{
		MediaQuery:        "@media (prefers-color-scheme: light)",
		Background:        variables.White,
		Foreground:        variables.Neutral900,
		Primary:           variables.Neutral900,
		PrimaryForeground: variables.White,
	}
	darkTheme = &Theme{
		MediaQuery:        "@media (prefers-color-scheme: dark)",
		Background:        variables.Neutral900,
		Foreground:        variables.Neutral100,
		Primary:           variables.White,
		PrimaryForeground: variables.Neutral900,
	}
)

// CSS Writes the CSS for the theme to the writer.
func (t *Theme) CSS(w io.Writer) error {
	var buf bytes.Buffer
	for _, style := range slices.Concat(
		t.Layout(),
		t.Buttons(),
	) {
		if err := style.CSS(&buf); err != nil {
			return err
		}
	}
	if buf.Len() > 0 {
		_, err := fmt.Fprintf(w, "%s{%s}", t.MediaQuery, buf.String())
		return err
	}
	return nil
}
