package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

// Layout returns the styles for the layout for the base stylesheet.
func (ss *StyleSheet) Layout() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				MinHeight: variables.FullScreenHeight,
			},
		},
		{
			Selector: "main",
			Props: gcss.Props{
				Display: props.DisplayGrid,
			},
		},
	}
}

// Layout returns the styles for the layout for the media.
func (m *Media) Layout() Styles {
	switch m.MediaType {
	case Mobile:
		return Styles{
			{
				Selector: "main",
				Props: gcss.Props{
					Gap:     variables.Size6,
					Padding: variables.Size8,
				},
			},
		}
	case Desktop:
		return Styles{
			{
				Selector: "main",
				Props: gcss.Props{
					Gap:     variables.Size8,
					Padding: variables.Size16,
				},
			},
		}
	default:
		return Styles{}
	}
}

// Layout returns the styles for the layout for the theme.
func (t *Theme) Layout() Styles {
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
