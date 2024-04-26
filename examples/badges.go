package examples

import (
	"github.com/AccentDesign/gostyle/props/align"
	"github.com/AccentDesign/gostyle/props/border"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/display"
	"github.com/AccentDesign/gostyle/props/font"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Badges = []style.Style{
	{
		Selector: ".badge",
		Props: style.Props{
			AlignItems:    align.ItemsCenter,
			Display:       display.InlineFlex,
			BorderRadius:  radiusFull,
			FontSize:      fontXs,
			FontWeight:    font.Weight600,
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
			BackgroundColor: colors.Color{RGBA: primary},
			Color:           colors.Color{RGBA: primaryForeground},
		},
	},
	{
		Selector: ".badge-secondary",
		Props: style.Props{
			BackgroundColor: colors.Color{RGBA: secondary},
			Color:           colors.Color{RGBA: secondaryForeground},
		},
	},
	{
		Selector: ".badge-destructive",
		Props: style.Props{
			BackgroundColor: colors.Color{RGBA: destructive},
			Color:           colors.Color{RGBA: destructiveForeground},
		},
	},
	{
		Selector: ".badge-outline",
		Props: style.Props{
			BackgroundColor: colors.Transparent(),
			Border: border.Border{
				Color: colors.Color{RGBA: borderColor},
				Style: border.StyleSolid,
				Width: unit.Px(1),
			},
		},
	},
}
