package props

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

// Alpha returns a new color with the alpha channel set to the given value.
func (c Color) Alpha(a uint8) Color {
	c.RGBA.A = a
	return c
}

// Mix returns a new color that is a mixture of the two colors.
// The percentage parameter is a float between 0 and 1
// and determines how much of the second color to mix with the first color.
func (c Color) Mix(with Color, percentage float64) Color {
	if percentage < 0 {
		percentage = 0
	}
	if percentage > 1 {
		percentage = 1
	}
	r := uint8(float64(c.RGBA.R)*(1-percentage) + float64(with.RGBA.R)*percentage)
	g := uint8(float64(c.RGBA.G)*(1-percentage) + float64(with.RGBA.G)*percentage)
	b := uint8(float64(c.RGBA.B)*(1-percentage) + float64(with.RGBA.B)*percentage)
	return Color{RGBA: color.RGBA{r, g, b, 255}}
}

// ColorRGBA returns a new color with the given red, green, blue, and alpha values.
func ColorRGBA(r, g, b, a uint8) Color {
	return Color{RGBA: color.RGBA{r, g, b, a}}
}

func ColorCurrentColor() Color {
	return Color{Keyword: "currentColor"}
}

func ColorInherit() Color {
	return Color{Keyword: "inherit"}
}

func ColorTransparent() Color {
	return Color{Keyword: "transparent"}
}
