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
	"github.com/AccentDesign/gostyle/style"
)

var Form = []style.Style{
	{
		Selector: ".input",
		Props: style.Props{
			BackgroundColor: backGround,
			Border: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: borderColor,
			},
			BorderRadius:  radius,
			Display:       display.Flex,
			FontSize:      fontSm,
			Height:        size10,
			LineHeight:    leadingTight,
			PaddingTop:    size2,
			PaddingRight:  size3,
			PaddingBottom: size2,
			PaddingLeft:   size3,
			Width:         unit.Percent(100),
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: style.Props{
			BackgroundColor: colors.Transparent(),
			BorderWidth:     unit.Raw(0),
			FontSize:        fontSm,
			FontWeight:      font.WeightMedium,
		},
	},
	{
		Selector: ".input-label",
		Props: style.Props{
			FontSize:   fontSm,
			FontWeight: font.WeightMedium,
			LineHeight: leadingTight,
		},
	},
	{
		Selector: ".input-help",
		Props: style.Props{
			Color:    mutedForeground,
			FontSize: fontSm,
		},
	},
	{
		Selector: ".input-error",
		Props: style.Props{
			Color:    destructive,
			FontSize: fontSm,
		},
	},
	{
		Selector: ".select",
		Props: style.Props{
			BackgroundColor: backGround,
			Border: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: borderColor,
			},
			BorderRadius:  radius,
			Display:       display.Flex,
			FontSize:      fontSm,
			Height:        size10,
			LineHeight:    leadingTight,
			PaddingTop:    size2,
			PaddingRight:  size3,
			PaddingBottom: size2,
			PaddingLeft:   size3,
			Width:         unit.Percent(100),
		},
	},
	{
		Selector: ".select:not([size])",
		Props: style.Props{
			Appearance:       appearance.None,
			PaddingRight:     size10,
			PrintColorAdjust: print.ColorAdjustExact,
			BackgroundImage:  background.ImageLayers(iconChevronDown),
			BackgroundPosition: background.PositionEdgeOffset(
				background.PositionEdgeItem{Edge: background.PositionEdgeRight, Unit: size3},
				background.PositionEdgeItem{Edge: background.PositionEdgeCenter},
			),
			BackgroundRepeat: background.RepeatNoRepeat,
			BackgroundSize:   background.SizeDimension(unit.Em(1), unit.Em(1)),
		},
	},
}
