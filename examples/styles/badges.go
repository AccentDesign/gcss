package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

var Badges = []gcss.Style{
	{
		Selector: ".badge",
		Props: gcss.Props{
			AlignItems:    props.AlignItemsCenter,
			Display:       props.DisplayInlineFlex,
			BorderRadius:  variables.SizeFull,
			FontSize:      variables.Size3H,
			FontWeight:    props.FontWeightMedium,
			LineHeight:    variables.Size4,
			PaddingTop:    variables.Size1,
			PaddingRight:  variables.Size3,
			PaddingBottom: variables.Size1,
			PaddingLeft:   variables.Size3,
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
