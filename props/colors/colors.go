package colors

import (
	"fmt"
	"image/color"
)

type (
	Color struct {
		Keyword string
		RGBA    RGBA
	}
	RGBA color.RGBA
)

func (c Color) String() string {
	if c.Keyword != "" {
		return c.Keyword
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

func CurrentColor() Color {
	return Color{Keyword: "currentColor"}
}

func Inherit() Color {
	return Color{Keyword: "inherit"}
}

func Transparent() Color {
	return Color{Keyword: "transparent"}
}
