package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/props/background"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Form = []style.Style{
	{
		Selector: ".input",
		Props: style.Props{
			BackgroundColor: props.Color{RGBA: backGround},
			Border: props.Border{
				Width: unit.Px(1),
				Style: props.BorderStyleSolid,
				Color: props.Color{RGBA: border},
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        unit.Px(40),
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.Width{Unit: unit.Percent(100)},
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: style.Props{
			BackgroundColor: props.Color{Keyword: props.ColorTransparent},
			BorderWidth:     unit.None(0),
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
			BackgroundColor: props.Color{RGBA: backGround},
			Border: props.Border{
				Width: unit.Px(1),
				Style: props.BorderStyleSolid,
				Color: props.Color{RGBA: border},
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        unit.Px(40),
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.Width{Unit: unit.Percent(100)},
		},
	},
	{
		Selector: ".select:not([size])",
		Props: style.Props{
			Appearance:       props.AppearanceNone,
			PaddingRight:     spacing10,
			PrintColorAdjust: props.PrintColorAdjustExact,
			BackgroundImage:  background.ImageLayers(imageCaretUrl),
			BackgroundPosition: background.PositionEdgeOffset(
				background.PositionEdgeItem{Edge: background.PositionEdgeRight, Unit: spacing3},
				background.PositionEdgeItem{Edge: background.PositionEdgeCenter},
			),
			BackgroundRepeat: background.RepeatNoRepeat,
			BackgroundSize:   background.SizeDimension(unit.Em(0.75), unit.Em(0.75)),
		},
	},
}
