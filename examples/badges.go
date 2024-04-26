package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Badges = []style.Style{
	{
		Selector: ".badge",
		Props: style.Props{
			AlignItems:    props.AlignItemsCenter,
			Display:       props.DisplayInlineFlex,
			BorderRadius:  radiusFull,
			FontSize:      fontXs,
			FontWeight:    props.FontWeight600,
			LineHeight:    leadingNone,
			PaddingTop:    spacing1,
			PaddingRight:  spacing3,
			PaddingBottom: spacing1,
			PaddingLeft:   spacing3,
		},
	},
	{
		Selector: ".badge-primary",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: primary},
			Color:           props.Color{RGBA: primaryForeground},
		},
	},
	{
		Selector: ".badge-secondary",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: secondary},
			Color:           props.Color{RGBA: secondaryForeground},
		},
	},
	{
		Selector: ".badge-destructive",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: destructive},
			Color:           props.Color{RGBA: destructiveForeground},
		},
	},
	{
		Selector: ".badge-outline",
		Props: style.Props{
			BackgroundColor: props.Color{Keyword: props.ColorTransparent},
			Border: props.Border{
				Color: props.Color{RGBA: border},
				Style: props.BorderStyleSolid,
				Width: unit.Px(1),
			},
		},
	},
}
