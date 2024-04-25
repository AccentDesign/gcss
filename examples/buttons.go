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
			BackgroundColor: props.Color{RGBA: gray900},
			Color:           props.Color{RGBA: white},
		},
	},
	{
		Selector: ".button-primary:hover",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: gray900.Alpha(230)},
		},
	},
	{
		Selector: ".button-secondary",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: gray200},
			Color:           props.Color{RGBA: gray900},
		},
	},
	{
		Selector: ".button-secondary:hover",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: gray200.Alpha(204)},
		},
	},
	{
		Selector: ".button-destructive",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: red500},
			Color:           props.Color{RGBA: white},
		},
	},
	{
		Selector: ".button-destructive:hover",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: red500.Alpha(230)},
		},
	},
	{
		Selector: ".button-outline",
		Props: style.Props{
			BackgroundColor: props.Color{Keyword: props.ColorTransparent},
			Border: props.Border{
				Width: props.Unit{1, props.UnitPx},
				Style: props.BorderStyleSolid,
				Color: props.Color{RGBA: gray200},
			},
		},
	},
	{
		Selector: ".button-outline:hover",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: gray200.Alpha(230)},
		},
	},
	{
		Selector: ".button-ghost",
		Props: style.Props{
			BackgroundColor: props.Color{Keyword: props.ColorTransparent},
		},
	},
	{
		Selector: ".button-ghost:hover",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: gray200.Alpha(230)},
		},
	},
}
