package props

import (
	"fmt"
	"image/color"
)

type (
	Color struct {
		Keyword ColorKeyword
		RGBA    RGBA
	}
	ColorKeyword string
	RGBA         color.RGBA
)

const (
	ColorCurrentColor ColorKeyword = "currentColor"
	ColorInherit      ColorKeyword = "inherit"
	ColorTransparent  ColorKeyword = "transparent"
)

func (c Color) String() string {
	if c.Keyword != "" {
		return string(c.Keyword)
	}
	return c.RGBA.String()
}

func (c RGBA) String() string {
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.R, c.G, c.B, float32(c.A)/255.0)
}

func (c RGBA) Alpha(a uint8) RGBA {
	c.A = a
	return c
}
