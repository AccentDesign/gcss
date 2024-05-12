package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Tables = []gcss.Style{
	{
		Selector: ".table",
		Props: gcss.Props{
			CaptionSide: props.CaptionSideBottom,
			FontSize:    fontSm,
			LineHeight:  leadingTight,
			Width:       props.UnitPercent(100),
		},
	},
	{
		Selector: ".table-caption",
		Props: gcss.Props{
			MarginTop: spacing4,
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
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
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
