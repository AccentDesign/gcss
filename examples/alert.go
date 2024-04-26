package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Alerts = []style.Style{
	{
		Selector: ".alert",
		Props: style.Props{
			BackgroundColor: props.Color{Keyword: props.ColorTransparent},
			BorderStyle:     props.BorderStyleSolid,
			BorderWidth:     unit.Px(1),
			BorderRadius:    radius,
			Padding:         spacing4,
		},
	},
	{
		Selector: ".alert-primary",
		Props: style.Props{
			BorderColor: props.Color{RGBA: primary},
			Color:       props.Color{RGBA: primary},
		},
	},
	{
		Selector: ".alert-destructive",
		Props: style.Props{
			BorderColor: props.Color{RGBA: destructive},
			Color:       props.Color{RGBA: destructive},
		},
	},
}
