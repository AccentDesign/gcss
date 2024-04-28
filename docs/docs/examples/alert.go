package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/unit"
)

var Alerts = []gcss.Style{
	{
		Selector: ".alert",
		Props: gcss.Props{
			BackgroundColor: colors.Transparent(),
			BorderStyle:     border.StyleSolid,
			BorderWidth:     unit.Px(1),
			BorderRadius:    radius,
			Padding:         size4,
		},
	},
	{
		Selector: ".alert-primary",
		Props: gcss.Props{
			BorderColor: primary,
			Color:       primary,
		},
	},
	{
		Selector: ".alert-destructive",
		Props: gcss.Props{
			BorderColor: destructive,
			Color:       destructive,
		},
	},
}
