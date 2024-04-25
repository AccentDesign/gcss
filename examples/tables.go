package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/style"
)

var Tables = []style.Style{
	{
		Selector: ".table",
		Props: style.Props{
			CaptionSide: props.CaptionSideBottom,
			FontSize:    props.Unit{0.875, props.Rem},
			LineHeight:  props.Unit{1.25, props.Rem},
			Width:       props.Unit{100, props.Percent},
		},
	},
	{
		Selector: ".table-caption",
		Props: style.Props{
			Color:     gray600,
			MarginTop: props.Unit{1, props.Rem},
		},
	},
	{
		Selector: ".table-tr",
		Props: style.Props{
			BorderBottom: props.Border{
				Width: props.Unit{1, props.Px},
				Style: props.BorderStyleSolid,
				Color: gray200,
			},
		},
	},
	{
		Selector: ".table-tfoot-tr",
		Props: style.Props{
			BackgroundColor: gray100,
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
