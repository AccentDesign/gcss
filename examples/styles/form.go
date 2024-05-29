package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

var Form = []gcss.Style{
	{
		Selector: ".input",
		Props: gcss.Props{
			BackgroundColor: input,
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			BorderRadius:  variables.Size1H,
			Display:       props.DisplayFlex,
			FontSize:      variables.Size3H,
			Height:        variables.Size10,
			LineHeight:    variables.Size5,
			PaddingTop:    variables.Size2,
			PaddingRight:  variables.Size3,
			PaddingBottom: variables.Size2,
			PaddingLeft:   variables.Size3,
			Width:         props.UnitPercent(100),
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: gcss.Props{
			BackgroundColor: props.ColorTransparent(),
			BorderWidth:     props.UnitRaw(0),
			FontSize:        variables.Size3H,
			FontWeight:      props.FontWeightMedium,
		},
	},
	{
		Selector: ".input-label",
		Props: gcss.Props{
			FontSize:   variables.Size3H,
			FontWeight: props.FontWeightMedium,
			LineHeight: variables.Size5,
		},
	},
	{
		Selector: ".input-help",
		Props: gcss.Props{
			Color:    mutedForeground,
			FontSize: variables.Size3H,
		},
	},
	{
		Selector: ".input-error",
		Props: gcss.Props{
			Color:    destructive,
			FontSize: variables.Size3H,
		},
	},
	{
		Selector: ".select",
		Props: gcss.Props{
			BackgroundColor: input,
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			BorderRadius:  variables.Size1H,
			Display:       props.DisplayFlex,
			FontSize:      variables.Size3H,
			Height:        variables.Size10,
			LineHeight:    variables.Size5,
			PaddingTop:    variables.Size2,
			PaddingRight:  variables.Size3,
			PaddingBottom: variables.Size2,
			PaddingLeft:   variables.Size3,
			Width:         props.UnitPercent(100),
		},
	},
	{
		Selector: ".select:not([size])",
		Props: gcss.Props{
			Appearance:       props.AppearanceNone,
			PaddingRight:     variables.Size10,
			PrintColorAdjust: props.PrintColorAdjustExact,
			BackgroundImage:  iconChevronDown,
			BackgroundPosition: props.BackgroundPositionEdges(
				props.BackgroundPositionEdge{Position: props.BackgroundPositionRight, Unit: variables.Size3},
				props.BackgroundPositionEdge{Position: props.BackgroundPositionCenter},
			),
			BackgroundRepeat: props.BackgroundRepeatNoRepeat,
			BackgroundSize:   props.BackgroundSizeDimension(props.UnitEm(1), props.UnitEm(1)),
		},
	},
}
