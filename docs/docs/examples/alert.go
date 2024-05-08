package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Alerts = []gcss.Style{
	{
		Selector: ".alert",
		Props: gcss.Props{
			BackgroundColor: props.ColorTransparent(),
			BorderStyle:     props.BorderStyleSolid,
			BorderWidth:     props.UnitPx(1),
			BorderRadius:    radius,
			Padding:         spacing4,
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
