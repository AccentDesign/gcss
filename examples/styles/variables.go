package styles

import (
	"github.com/AccentDesign/gcss/props"
)

var (
	primary               = props.ColorRGBA(17, 24, 39, 255)
	primaryForeground     = props.ColorRGBA(255, 255, 255, 255)
	secondary             = props.ColorRGBA(229, 231, 235, 255)
	secondaryForeground   = props.ColorRGBA(17, 24, 39, 255)
	destructive           = props.ColorRGBA(239, 68, 68, 255)
	destructiveForeground = props.ColorRGBA(255, 255, 255, 255)
	muted                 = props.ColorRGBA(243, 244, 246, 255)
	mutedForeground       = props.ColorRGBA(75, 85, 99, 255)
	input                 = props.ColorRGBA(253, 253, 254, 255)
	borderColor           = props.ColorRGBA(203, 213, 225, 255)
	fontXs                = props.UnitRem(0.75)
	fontSm                = props.UnitRem(0.875)
	fontMd                = props.UnitRem(1)
	leadingNone           = props.UnitRem(1)
	leadingTight          = props.UnitRem(1.25)
	radius                = props.UnitRem(0.375)
	radiusFull            = props.UnitPx(9999)
	spacing1              = props.UnitRem(0.25)
	spacing2              = props.UnitRem(0.5)
	spacing3              = props.UnitRem(0.75)
	spacing4              = props.UnitRem(1)
	spacing10             = props.UnitRem(2.5)

	iconChevronDown = props.BackgroundImageURL("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGZpbGw9Im5vbmUiIHN0cm9rZT0iY3VycmVudENvbG9yIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiIHN0cm9rZS13aWR0aD0iMiIgZD0ibTE5IDlsLTcgN2wtNy03Ii8+PC9zdmc+")
)
