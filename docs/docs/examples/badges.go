package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props/align"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/display"
	"github.com/AccentDesign/gcss/props/font"
	"github.com/AccentDesign/gcss/props/unit"
)

var Badges = []gcss.Style{
	{
		Selector: ".badge",
		Props: gcss.Props{
			AlignItems:    align.ItemsCenter,
			Display:       display.InlineFlex,
			BorderRadius:  radiusFull,
			FontSize:      fontXs,
			FontWeight:    font.WeightSemiBold,
			LineHeight:    leadingNone,
			PaddingTop:    size1,
			PaddingRight:  size3,
			PaddingBottom: size1,
			PaddingLeft:   size3,
		},
	},
	{
		Selector: ".badge-primary",
		Props: gcss.Props{
			BackgroundColor: primary,
			Color:           primaryForeground,
		},
	},
	{
		Selector: ".badge-secondary",
		Props: gcss.Props{
			BackgroundColor: secondary,
			Color:           secondaryForeground,
		},
	},
	{
		Selector: ".badge-destructive",
		Props: gcss.Props{
			BackgroundColor: destructive,
			Color:           destructiveForeground,
		},
	},
	{
		Selector: ".badge-outline",
		Props: gcss.Props{
			BackgroundColor: colors.Transparent(),
			Border: border.Border{
				Color: borderColor,
				Style: border.StyleSolid,
				Width: unit.Px(1),
			},
		},
	},
}
