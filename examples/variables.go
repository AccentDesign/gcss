package examples

import (
	"github.com/AccentDesign/gostyle/props"
	"github.com/AccentDesign/gostyle/props/background"
)

var (
	white                 = props.RGBA{255, 255, 255, 255}
	backGround            = white
	primary               = props.RGBA{17, 24, 39, 255}
	primaryForeground     = white
	secondary             = props.RGBA{229, 231, 235, 255}
	secondaryForeground   = props.RGBA{17, 24, 39, 255}
	destructive           = props.RGBA{239, 68, 68, 255}
	destructiveForeground = white
	border                = props.RGBA{229, 231, 235, 255}
	muted                 = props.RGBA{243, 244, 246, 255}
	mutedForeground       = props.RGBA{75, 85, 99, 255}
	fontXs                = props.Unit{0.75, props.UnitRem}
	fontSm                = props.Unit{0.875, props.UnitRem}
	leadingNone           = props.Unit{1, props.UnitRem}
	leadingTight          = props.Unit{1.25, props.UnitRem}
	radius                = props.Unit{0.375, props.UnitRem}
	radiusFull            = props.Unit{9999, props.UnitPx}
	spacing1              = props.Unit{0.25, props.UnitRem}
	spacing2              = props.Unit{0.5, props.UnitRem}
	spacing3              = props.Unit{0.75, props.UnitRem}
	spacing4              = props.Unit{1, props.UnitRem}
	spacing10             = props.Unit{2.5, props.UnitRem}

	imageCaretUrl = background.ImageURL("data:image/svg+xml;charset=utf-8,%3Csvg aria-hidden='true' xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 10 6'%3E%3Cpath stroke='%236B7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='m1 1 4 4 4-4'/%3E%3C/svg%3E")
)
