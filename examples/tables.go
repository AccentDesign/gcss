package examples

import (
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/table"
	"github.com/AccentDesign/gcss/props/text"
	"github.com/AccentDesign/gcss/props/unit"
	"github.com/AccentDesign/gcss/style"
)

var Tables = []style.Style{
	{
		Selector: ".table",
		Props: style.Props{
			CaptionSide: table.CaptionSideBottom,
			FontSize:    fontSm,
			LineHeight:  leadingTight,
			Width:       unit.Percent(100),
		},
	},
	{
		Selector: ".table-caption",
		Props: style.Props{
			MarginTop: size4,
		},
	},
	{
		Selector: ".table-tr",
		Props: style.Props{
			BorderBottom: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: borderColor,
			},
		},
	},
	{
		Selector: ".table-tfoot-tr",
		Props: style.Props{
			Color:           mutedForeground,
			BackgroundColor: muted,
		},
	},
	{
		Selector: ".table-td,.table-th",
		Props: style.Props{
			PaddingTop:    size2,
			PaddingRight:  size3,
			PaddingBottom: size2,
			PaddingLeft:   size3,
			TextAlign:     text.AlignLeft,
		},
	},
}
