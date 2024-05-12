package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Buttons = []gcss.Style{
	{
		Selector: ".button",
		Props: gcss.Props{
			AlignItems:     props.AlignItemsCenter,
			BorderRadius:   radius,
			Display:        props.DisplayInlineFlex,
			FontSize:       fontMd,
			FontWeight:     props.FontWeightMedium,
			Height:         spacing10,
			JustifyContent: props.JustifyContentCenter,
			LineHeight:     leadingTight,
			PaddingTop:     spacing2,
			PaddingRight:   spacing4,
			PaddingBottom:  spacing2,
			PaddingLeft:    spacing4,
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
			BackgroundColor: props.ColorTransparent(),
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
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
			BackgroundColor: props.ColorTransparent(),
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
			BackgroundColor: props.ColorTransparent(),
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			Padding: props.UnitInitial(), // if sizes are added this can be removed as the base padding will be set by the sizes
			Width:   spacing10,
		},
	},
	{
		Selector: ".button-icon:hover",
		Props: gcss.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
}
