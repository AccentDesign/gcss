package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

var Tables = []gcss.Style{
	{
		Selector: ".table",
		Props: gcss.Props{
			CaptionSide: props.CaptionSideBottom,
			FontSize:    variables.Size3H,
			LineHeight:  variables.Size5,
			Width:       props.UnitPercent(100),
		},
	},
	{
		Selector: ".table-caption",
		Props: gcss.Props{
			MarginTop: variables.Size4,
		},
	},
	{
		Selector: ".table-tr",
		Props: gcss.Props{
			BorderBottom: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
		},
	},
	{
		Selector: ".table-tfoot-tr",
		Props: gcss.Props{
			Color:           mutedForeground,
			BackgroundColor: muted,
		},
	},
	{
		Selector: ".table-td,.table-th",
		Props: gcss.Props{
			PaddingTop:    variables.Size2,
			PaddingRight:  variables.Size3,
			PaddingBottom: variables.Size2,
			PaddingLeft:   variables.Size3,
			TextAlign:     props.TextAlignLeft,
		},
	},
	{
		Selector: ".table-scroller",
		Props: gcss.Props{
			OverflowX: props.OverflowAuto,
		},
	},
}
