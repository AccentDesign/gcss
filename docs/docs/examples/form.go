package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props/appearance"
	"github.com/AccentDesign/gcss/props/background"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/display"
	"github.com/AccentDesign/gcss/props/font"
	"github.com/AccentDesign/gcss/props/print"
	"github.com/AccentDesign/gcss/props/unit"
)

var Form = []gcss.Style{
	{
		Selector: ".input",
		Props: gcss.Props{
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
		Props: gcss.Props{
			BackgroundColor: colors.Transparent(),
			BorderWidth:     unit.Raw(0),
			FontSize:        fontSm,
			FontWeight:      font.WeightMedium,
		},
	},
	{
		Selector: ".input-label",
		Props: gcss.Props{
			FontSize:   fontSm,
			FontWeight: font.WeightMedium,
			LineHeight: leadingTight,
		},
	},
	{
		Selector: ".input-help",
		Props: gcss.Props{
			Color:    mutedForeground,
			FontSize: fontSm,
		},
	},
	{
		Selector: ".input-error",
		Props: gcss.Props{
			Color:    destructive,
			FontSize: fontSm,
		},
	},
	{
		Selector: ".select",
		Props: gcss.Props{
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
		Props: gcss.Props{
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
