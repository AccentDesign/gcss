package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

// Resets returns the resets styles for the stylesheet.
func (ss *StyleSheet) Resets() Styles {
	return Styles{
		{
			Selector: "*,::after,::before,::backdrop,::file-selector-button",
			Props: gcss.Props{
				Border: props.Border{
					Width: variables.Size0,
					Style: props.BorderStyleSolid,
				},
				BoxSizing: props.BoxSizingBorderBox,
				Margin:    variables.Size0,
				Padding:   variables.Size0,
			},
		},
		{
			Selector: "html,:host",
			Props: gcss.Props{
				FontFamily: props.FontFamilySans,
				LineHeight: props.UnitRaw(1.5),
			},
			CustomProps: []gcss.CustomProp{
				{Attr: "-webkit-text-size-adjust", Value: "100%"},
				{Attr: "tab-size", Value: "4"},
			},
		},
		{
			Selector: "body",
			Props: gcss.Props{
				LineHeight: props.UnitInherit(),
			},
		},
		// more resets
	}
}
