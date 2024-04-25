package examples

import "github.com/AccentDesign/gostyle/props"

var (
	white   = props.RGBA{255, 255, 255, 255}
	gray100 = props.RGBA{245, 247, 250, 255}
	gray200 = props.RGBA{229, 231, 235, 255}
	gray600 = props.RGBA{75, 85, 99, 255}
	gray900 = props.RGBA{17, 24, 39, 255}
	red500  = props.RGBA{239, 68, 68, 255}
	red800  = props.RGBA{185, 28, 28, 255}

	radiusMd = props.Unit{0.375, props.UnitRem}

	spacing2 = props.Unit{0.5, props.UnitRem}
	spacing3 = props.Unit{0.75, props.UnitRem}
	spacing4 = props.Unit{1, props.UnitRem}
)
