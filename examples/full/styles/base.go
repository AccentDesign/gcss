package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/variables"
)

// Base returns the base styles for the stylesheet.
func (ss *StyleSheet) Base() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				MinHeight: variables.FullScreenHeight,
			},
		},
	}
}

// Base returns the base styles for the theme.
func (t *Theme) Base() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				BackgroundColor: t.Background,
				Color:           t.Foreground,
			},
		},
	}
}
