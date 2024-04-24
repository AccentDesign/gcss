package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/style"
)

var Buttons = []style.Style{
	{
		Selector: ".button",
		Props: style.Props{
			AlignItems:     props.AlignItemsCenter,
			BorderRadius:   radiusMd,
			Display:        props.DisplayFlex,
			JustifyContent: props.JustifyContentCenter,
			PaddingTop:     spacing2,
			PaddingRight:   spacing4,
			PaddingBottom:  spacing2,
			PaddingLeft:    spacing4,
		},
	},
	{
		Selector: ".button-primary",
		Props: style.Props{
			BackgroundColor: gray900,
			Color:           white,
		},
	},
	{
		Selector: ".button-primary:hover",
		Props: style.Props{
			BackgroundColor: gray900.Alpha(230),
		},
	},
	{
		Selector: ".button-secondary",
		Props: style.Props{
			BackgroundColor: gray200,
			Color:           gray900,
		},
	},
	{
		Selector: ".button-secondary:hover",
		Props: style.Props{
			BackgroundColor: gray200.Alpha(204),
		},
	},
	{
		Selector: ".button-destructive",
		Props: style.Props{
			BackgroundColor: red500,
			Color:           white,
		},
	},
	{
		Selector: ".button-destructive:hover",
		Props: style.Props{
			BackgroundColor: red500.Alpha(230),
		},
	},
	{
		Selector: ".button-outline",
		Props: style.Props{
			BackgroundColor: props.ColorTransparent,
			BorderColor:     gray200,
			BorderStyle:     props.BorderStyleSolid,
			BorderWidth:     props.Unit{1, props.Px},
		},
	},
	{
		Selector: ".button-outline:hover",
		Props: style.Props{
			BackgroundColor: gray200.Alpha(230),
		},
	},
	{
		Selector: ".button-ghost",
		Props: style.Props{
			BackgroundColor: props.ColorTransparent,
		},
	},
	{
		Selector: ".button-ghost:hover",
		Props: style.Props{
			BackgroundColor: gray200.Alpha(230),
		},
	},
}
