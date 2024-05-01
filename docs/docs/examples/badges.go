package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Badges = []gcss.Style{
	{
		Selector: ".badge",
		Props: gcss.Props{
			AlignItems:    props.AlignItemsCenter,
			Display:       props.DisplayInlineFlex,
			BorderRadius:  radiusFull,
			FontSize:      fontXs,
			FontWeight:    props.FontWeightSemiBold,
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
			BackgroundColor: props.ColorTransparent(),
			Border: props.Border{
				Color: borderColor,
				Style: props.BorderStyleSolid,
				Width: props.UnitPx(1),
			},
		},
	},
}
