package examples

import (
	"github.com/AccentDesign/gostyle/props/border"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/margin"
	"github.com/AccentDesign/gostyle/props/table"
	"github.com/AccentDesign/gostyle/props/text"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/props/width"
	"github.com/AccentDesign/gostyle/style"
)

var Tables = []style.Style{
	{
		Selector: ".table",
		Props: style.Props{
			CaptionSide: table.CaptionSideBottom,
			FontSize:    fontSm,
			LineHeight:  leadingTight,
			Width:       width.Width{Unit: unit.Percent(100)},
		},
	},
	{
		Selector: ".table-caption",
		Props: style.Props{
			MarginTop: margin.Margin{Unit: spacing4},
		},
	},
	{
		Selector: ".table-tr",
		Props: style.Props{
			BorderBottom: border.Border{
				Width: unit.Px(1),
				Style: border.StyleSolid,
				Color: colors.Color{RGBA: borderColor},
			},
		},
	},
	{
		Selector: ".table-tfoot-tr",
		Props: style.Props{
			Color:           colors.Color{RGBA: mutedForeground},
			BackgroundColor: colors.Color{RGBA: muted},
		},
	},
	{
		Selector: ".table-td,.table-th",
		Props: style.Props{
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			TextAlign:     text.AlignLeft,
		},
	},
}
