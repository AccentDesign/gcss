package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

var Buttons = []gcss.Style{
	{
		Selector: ".button",
		Props: gcss.Props{
			AlignItems:     props.AlignItemsCenter,
			BorderRadius:   variables.Size1H,
			Display:        props.DisplayInlineFlex,
			FontSize:       variables.Size4,
			FontWeight:     props.FontWeightMedium,
			Height:         variables.Size10,
			JustifyContent: props.JustifyContentCenter,
			LineHeight:     variables.Size5,
			PaddingTop:     variables.Size2,
			PaddingRight:   variables.Size4,
			PaddingBottom:  variables.Size2,
			PaddingLeft:    variables.Size4,
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
			Width:   variables.Size10,
		},
	},
	{
		Selector: ".button-icon:hover",
		Props: gcss.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
}
