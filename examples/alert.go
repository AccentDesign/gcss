package examples

import (
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/unit"
	"github.com/AccentDesign/gcss/style"
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
