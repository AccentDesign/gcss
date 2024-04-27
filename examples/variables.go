package examples

import (
	"github.com/AccentDesign/gcss/props/background"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/unit"
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
	size1                 = unit.Rem(0.25)
	size2                 = unit.Rem(0.5)
	size3                 = unit.Rem(0.75)
	size4                 = unit.Rem(1)
	size10                = unit.Rem(2.5)

	iconChevronDown = background.ImageURL("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGZpbGw9Im5vbmUiIHN0cm9rZT0iY3VycmVudENvbG9yIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiIHN0cm9rZS13aWR0aD0iMiIgZD0ibTE5IDlsLTcgN2wtNy03Ii8+PC9zdmc+")
)
