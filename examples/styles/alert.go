package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

var Alerts = []gcss.Style{
	{
		Selector: ".alert",
		Props: gcss.Props{
			BackgroundColor: props.ColorTransparent(),
			BorderStyle:     props.BorderStyleSolid,
			BorderWidth:     props.UnitPx(1),
			BorderRadius:    variables.Size1H,
			Padding:         variables.Size4,
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
