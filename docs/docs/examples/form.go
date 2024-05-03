package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Form = []gcss.Style{
	{
		Selector: ".input",
		Props: gcss.Props{
			BackgroundColor: backGround,
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        size10,
			LineHeight:    leadingTight,
			PaddingTop:    size2,
			PaddingRight:  size3,
			PaddingBottom: size2,
			PaddingLeft:   size3,
			Width:         props.UnitPercent(100),
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: gcss.Props{
			BackgroundColor: props.ColorTransparent(),
			BorderWidth:     props.UnitRaw(0),
			FontSize:        fontSm,
			FontWeight:      props.FontWeightMedium,
		},
	},
	{
		Selector: ".input-label",
		Props: gcss.Props{
			FontSize:   fontSm,
			FontWeight: props.FontWeightMedium,
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
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        size10,
			LineHeight:    leadingTight,
			PaddingTop:    size2,
			PaddingRight:  size3,
			PaddingBottom: size2,
			PaddingLeft:   size3,
			Width:         props.UnitPercent(100),
		},
	},
	{
		Selector: ".select:not([size])",
		Props: gcss.Props{
			Appearance:       props.AppearanceNone,
			PaddingRight:     size10,
			PrintColorAdjust: props.ColorAdjustExact,
			BackgroundImage:  iconChevronDown,
			BackgroundPosition: props.BackgroundPositionEdgeOffset(
				props.BackgroundPositionEdgeItem{Edge: props.BackgroundPositionEdgeRight, Unit: size3},
				props.BackgroundPositionEdgeItem{Edge: props.BackgroundPositionEdgeCenter},
			),
			BackgroundRepeat: props.BackgroundRepeatNoRepeat,
			BackgroundSize:   props.BackgroundSizeDimension(props.UnitEm(1), props.UnitEm(1)),
		},
	},
}
