package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/style"
)

var Tables = []style.Style{
	{
		Selector: ".table",
		Props: style.Props{
			CaptionSide: props.CaptionSideBottom,
			FontSize:    fontSm,
			LineHeight:  leadingTight,
			Width:       props.Width{Unit: unit.Percent(100)},
		},
	},
	{
		Selector: ".table-caption",
		Props: style.Props{
			MarginTop: props.Margin{Unit: spacing4},
		},
	},
	{
		Selector: ".table-tr",
		Props: style.Props{
			BorderBottom: props.Border{
				Width: unit.Px(1),
				Style: props.BorderStyleSolid,
				Color: props.Color{RGBA: border},
			},
		},
	},
	{
		Selector: ".table-tfoot-tr",
		Props: style.Props{
			Color:           props.Color{RGBA: mutedForeground},
			BackgroundColor: props.Color{RGBA: muted},
		},
	},
	{
		Selector: ".table-td, .table-th",
		Props: style.Props{
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			TextAlign:     props.TextAlignLeft,
		},
	},
}
