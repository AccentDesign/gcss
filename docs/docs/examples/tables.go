package examples

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/overflow"
	"github.com/AccentDesign/gcss/props/table"
	"github.com/AccentDesign/gcss/props/text"
	"github.com/AccentDesign/gcss/props/unit"
)

var Tables = []gcss.Style{
	{
		Selector: ".table",
		Props: gcss.Props{
			CaptionSide: table.CaptionSideBottom,
			FontSize:    fontSm,
			LineHeight:  leadingTight,
			Width:       unit.Percent(100),
		},
	},
	{
		Selector: ".table-caption",
		Props: gcss.Props{
			MarginTop: size4,
		},
	},
	{
		Selector: ".table-tr",
		Props: gcss.Props{
			BorderBottom: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
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
			PaddingTop:    size2,
			PaddingRight:  size3,
			PaddingBottom: size2,
			PaddingLeft:   size3,
			TextAlign:     text.AlignLeft,
		},
	},
	{
		Selector: ".table-scroller",
		Props: gcss.Props{
			OverflowX: overflow.Auto,
		},
	},
}
