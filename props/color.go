package props

import (
	"fmt"
	"image/color"
)

type (
	Color interface {
		String() string
	}
	ColorKeyword string
	ColorRGBA    color.RGBA
)

const (
	ColorCurrentColor ColorKeyword = "currentColor"
	ColorInherit      ColorKeyword = "inherit"
	ColorTransparent  ColorKeyword = "transparent"
)

func (c ColorKeyword) String() string {
	return string(c)
}

func (c ColorRGBA) String() string {
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.R, c.G, c.B, float32(c.A)/255.0)
}

func (c ColorRGBA) Alpha(a uint8) Color {
	c.A = a
	return c
}
