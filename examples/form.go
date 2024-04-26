package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/style"
)

var Form = []style.Style{
	{
		Selector: ".input",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: background},
			Border: props.Border{
				Width: props.Unit{1, props.UnitPx},
				Style: props.BorderStyleSolid,
				Color: props.Color{RGBA: border},
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        props.Unit{40, props.UnitPx},
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.Width{Unit: props.Unit{100, props.UnitPercent}},
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: style.Props{
			BackgroundColor: props.Color{Keyword: props.ColorTransparent},
			BorderWidth:     props.Unit{0, props.UnitNone},
			FontSize:        fontSm,
			FontWeight:      props.FontWeight500,
		},
	},
	{
		Selector: ".input-label",
		Props: style.Props{
			FontSize:   fontSm,
			FontWeight: props.FontWeight500,
			LineHeight: leadingTight,
		},
	},
	{
		Selector: ".input-help",
		Props: style.Props{
			Color:    props.Color{RGBA: mutedForeground},
			FontSize: fontSm,
		},
	},
	{
		Selector: ".input-error",
		Props: style.Props{
			Color:    props.Color{RGBA: destructive},
			FontSize: fontSm,
		},
	},
	{
		Selector: ".select",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: background},
			Border: props.Border{
				Width: props.Unit{1, props.UnitPx},
				Style: props.BorderStyleSolid,
				Color: props.Color{RGBA: border},
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        props.Unit{40, props.UnitPx},
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.Width{Unit: props.Unit{100, props.UnitPercent}},
		},
	},
	{
		Selector: ".select:not([size])",
		Props: style.Props{
			Appearance:       props.AppearanceNone,
			PaddingRight:     spacing10,
			PrintColorAdjust: props.PrintColorAdjustExact,
			// background in `docs/css/unhandled.css`
		},
	},
}
