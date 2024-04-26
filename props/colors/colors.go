package colors

import (
	"fmt"
	"image/color"
)

type Color struct {
	Keyword string
	RGBA    color.RGBA
}

func (c Color) String() string {
	if c.Keyword != "" {
		return c.Keyword
	}
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.RGBA.R, c.RGBA.G, c.RGBA.B, float32(c.RGBA.A)/255.0)
}

func (c Color) Alpha(a uint8) Color {
	c.RGBA.A = a
	return c
}

func RGBA(r, g, b, a uint8) Color {
	return Color{RGBA: color.RGBA{r, g, b, a}}
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
