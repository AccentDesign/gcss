package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/style"
)

var Form = []style.Style{
	{
		Selector: ".input",
		Props: style.Props{
			BackgroundColor: white,
			Border: props.Border{
				Width: props.Unit{1, props.UnitPx},
				Style: props.BorderStyleSolid,
				Color: gray200,
			},
			BorderRadius:  radiusMd,
			Display:       props.DisplayFlex,
			FontSize:      props.Unit{0.875, props.UnitRem},
			Height:        props.Unit{40, props.UnitPx},
			LineHeight:    props.Unit{1.25, props.UnitRem},
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.Unit{100, props.UnitPercent},
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: style.Props{
			BackgroundColor: props.ColorTransparent,
			BorderWidth:     props.Unit{0, props.UnitNone},
			FontSize:        props.Unit{0.875, props.UnitRem},
			FontWeight:      props.FontWeight500,
		},
	},
	{
		Selector: ".input-label",
		Props: style.Props{
			FontSize:   props.Unit{0.875, props.UnitRem},
			FontWeight: props.FontWeight500,
			LineHeight: props.Unit{1.25, props.UnitRem},
		},
	},
	{
		Selector: ".input-help",
		Props: style.Props{
			Color:    gray600,
			FontSize: props.Unit{0.875, props.UnitRem},
		},
	},
	{
		Selector: ".input-error",
		Props: style.Props{
			Color:    red800,
			FontSize: props.Unit{0.875, props.UnitRem},
		},
	},
}
