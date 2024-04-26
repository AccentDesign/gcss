package examples

import (
	"github.com/AccentDesign/gostyle/props/border"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Alerts = []style.Style{
	{
		Selector: ".alert",
		Props: style.Props{
			BackgroundColor: colors.Transparent(),
			BorderStyle:     border.StyleSolid,
			BorderWidth:     unit.Px(1),
			BorderRadius:    radius,
			Padding:         size4,
		},
	},
	{
		Selector: ".alert-primary",
		Props: style.Props{
			BorderColor: primary,
			Color:       primary,
		},
	},
	{
		Selector: ".alert-destructive",
		Props: style.Props{
			BorderColor: destructive,
			Color:       destructive,
		},
	},
}
