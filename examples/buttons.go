package examples

import (
	"github.com/AccentDesign/gostyle/props/align"
	"github.com/AccentDesign/gostyle/props/border"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/display"
	"github.com/AccentDesign/gostyle/props/justify"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Buttons = []style.Style{
	{
		Selector: ".button",
		Props: style.Props{
			AlignItems:     align.ItemsCenter,
			BorderRadius:   radius,
			Display:        display.Flex,
			JustifyContent: justify.ContentCenter,
			PaddingTop:     spacing2,
			PaddingRight:   spacing4,
			PaddingBottom:  spacing2,
			PaddingLeft:    spacing4,
		},
	},
	{
		Selector: ".button-primary",
		Props: style.Props{
			BackgroundColor: primary,
			Color:           primaryForeground,
		},
	},
	{
		Selector: ".button-primary:hover",
		Props: style.Props{
			BackgroundColor: primary.Alpha(230),
		},
	},
	{
		Selector: ".button-secondary",
		Props: style.Props{
			BackgroundColor: secondary,
			Color:           secondaryForeground,
		},
	},
	{
		Selector: ".button-secondary:hover",
		Props: style.Props{
			BackgroundColor: secondary.Alpha(204),
		},
	},
	{
		Selector: ".button-destructive",
		Props: style.Props{
			BackgroundColor: destructive,
			Color:           destructiveForeground,
		},
	},
	{
		Selector: ".button-destructive:hover",
		Props: style.Props{
			BackgroundColor: destructive.Alpha(230),
		},
	},
	{
		Selector: ".button-outline",
		Props: style.Props{
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
		Props: style.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
	{
		Selector: ".button-ghost",
		Props: style.Props{
			BackgroundColor: colors.Transparent(),
		},
	},
	{
		Selector: ".button-ghost:hover",
		Props: style.Props{
			BackgroundColor: secondary.Alpha(230),
		},
	},
}
