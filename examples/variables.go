package examples

import (
	"github.com/AccentDesign/gostyle/props/background"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/unit"
)

var (
	white                 = colors.RGBA(255, 255, 255, 255)
	backGround            = white
	primary               = colors.RGBA(17, 24, 39, 255)
	primaryForeground     = white
	secondary             = colors.RGBA(229, 231, 235, 255)
	secondaryForeground   = colors.RGBA(17, 24, 39, 255)
	destructive           = colors.RGBA(239, 68, 68, 255)
	destructiveForeground = white
	borderColor           = colors.RGBA(229, 231, 235, 255)
	muted                 = colors.RGBA(243, 244, 246, 255)
	mutedForeground       = colors.RGBA(75, 85, 99, 255)
	fontXs                = unit.Rem(0.75)
	fontSm                = unit.Rem(0.875)
	leadingNone           = unit.Rem(1)
	leadingTight          = unit.Rem(1.25)
	radius                = unit.Rem(0.375)
	radiusFull            = unit.Px(9999)
	spacing1              = unit.Rem(0.25)
	spacing2              = unit.Rem(0.5)
	spacing3              = unit.Rem(0.75)
	spacing4              = unit.Rem(1)
	spacing10             = unit.Rem(2.5)

	imageCaretUrl = background.ImageURL("data:image/svg+xml;charset=utf-8,%3Csvg aria-hidden='true' xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 10 6'%3E%3Cpath stroke='%236B7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='m1 1 4 4 4-4'/%3E%3C/svg%3E")
)
