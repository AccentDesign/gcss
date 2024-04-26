package examples

import (
	"github.com/AccentDesign/gostyle/props/appearance"
	"github.com/AccentDesign/gostyle/props/background"
	"github.com/AccentDesign/gostyle/props/border"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/display"
	"github.com/AccentDesign/gostyle/props/font"
	"github.com/AccentDesign/gostyle/props/print"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/props/width"
	"github.com/AccentDesign/gostyle/style"
)

var Form = []style.Style{
	{
		Selector: ".input",
		Props: style.Props{
			BackgroundColor: colors.Color{RGBA: backGround},
			Border: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: colors.Color{RGBA: borderColor},
			},
			BorderRadius:  radius,
			Display:       display.Flex,
			FontSize:      fontSm,
			Height:        unit.Px(40),
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         width.Width{Unit: unit.Percent(100)},
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: style.Props{
			BackgroundColor: colors.Transparent(),
			BorderWidth:     unit.Raw(0),
			FontSize:        fontSm,
			FontWeight:      font.Weight500,
		},
	},
	{
		Selector: ".input-label",
		Props: style.Props{
			FontSize:   fontSm,
			FontWeight: font.Weight500,
			LineHeight: leadingTight,
		},
	},
	{
		Selector: ".input-help",
		Props: style.Props{
			Color:    colors.Color{RGBA: mutedForeground},
			FontSize: fontSm,
		},
	},
	{
		Selector: ".input-error",
		Props: style.Props{
			Color:    colors.Color{RGBA: destructive},
			FontSize: fontSm,
		},
	},
	{
		Selector: ".select",
		Props: style.Props{
			BackgroundColor: colors.Color{RGBA: backGround},
			Border: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: colors.Color{RGBA: borderColor},
			},
			BorderRadius:  radius,
			Display:       display.Flex,
			FontSize:      fontSm,
			Height:        unit.Px(40),
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         width.Width{Unit: unit.Percent(100)},
		},
	},
	{
		Selector: ".select:not([size])",
		Props: style.Props{
			Appearance:       appearance.None,
			PaddingRight:     spacing10,
			PrintColorAdjust: print.ColorAdjustExact,
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
