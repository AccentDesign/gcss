package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props/align"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/display"
	"github.com/AccentDesign/gcss/props/font"
	"github.com/AccentDesign/gcss/props/justify"
	"github.com/AccentDesign/gcss/props/unit"
)

var Buttons = []gcss.Style{
	{
		Selector: ".button",
		Props: gcss.Props{
			AlignItems:     align.ItemsCenter,
			BorderRadius:   radius,
			Display:        display.InlineFlex,
			FontSize:       fontSm,
			FontWeight:     font.WeightMedium,
			Height:         size10,
			JustifyContent: justify.ContentCenter,
			LineHeight:     leadingTight,
			PaddingTop:     size2,
			PaddingRight:   size4,
			PaddingBottom:  size2,
			PaddingLeft:    size4,
		},
	},
	{
		Selector: ".button-primary",
		Props: gcss.Props{
			BackgroundColor: primary,
			Color:           primaryForeground,
		},
	},
	{
		Selector: ".button-primary:hover",
		Props: gcss.Props{
			BackgroundColor: primary.Alpha(230),
		},
	},
	{
		Selector: ".button-secondary",
		Props: gcss.Props{
			BackgroundColor: secondary,
			Color:           secondaryForeground,
		},
	},
	{
		Selector: ".button-secondary:hover",
		Props: gcss.Props{
			BackgroundColor: secondary.Alpha(204),
		},
	},
	{
		Selector: ".button-destructive",
		Props: gcss.Props{
			BackgroundColor: destructive,
			Color:           destructiveForeground,
		},
	},
	{
		Selector: ".button-destructive:hover",
		Props: gcss.Props{
			BackgroundColor: destructive.Alpha(230),
		},
	},
	{
		Selector: ".button-outline",
		Props: gcss.Props{
			BackgroundColor: colors.Transparent(),
			Border: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: borderColor,
			},
		},
	},
	{
		Selector: ".button-outline:hover",
		Props: gcss.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
	{
		Selector: ".button-ghost",
		Props: gcss.Props{
			BackgroundColor: colors.Transparent(),
		},
	},
	{
		Selector: ".button-ghost:hover",
		Props: gcss.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
	{
		Selector: ".button-icon",
		Props: gcss.Props{
			BackgroundColor: colors.Transparent(),
			Border: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: borderColor,
			},
			Padding: unit.Initial(), // if sizes are added this can be removed as the base padding will be set by the sizes
			Width:   size10,
		},
	},
	{
		Selector: ".button-icon:hover",
		Props: gcss.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
}
