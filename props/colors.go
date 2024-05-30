package props

import (
	"fmt"
	"image/color"
)

type Color struct {
	Keyword string
	Color   color.Color
}

// String returns a string representation of the color.
// If the color is a keyword, the keyword is returned.
// Otherwise, the color is returned as an rgba string.
func (c Color) String() string {
	if c.Keyword != "" {
		return c.Keyword
	}
	r, g, b, a := c.Color.RGBA()
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", r>>8, g>>8, b>>8, float32(a>>8)/255.0)
}

// Alpha returns a new color with the alpha channel set to the given value.
func (c Color) Alpha(a uint8) Color {
	r, g, b, _ := c.Color.RGBA()
	c.Color = color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: a}
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
	r1, g1, b1, a1 := c.Color.RGBA()
	r2, g2, b2, a2 := with.Color.RGBA()
	r := uint8(float64(r1>>8)*(1-percentage) + float64(r2>>8)*percentage)
	g := uint8(float64(g1>>8)*(1-percentage) + float64(g2>>8)*percentage)
	b := uint8(float64(b1>>8)*(1-percentage) + float64(b2>>8)*percentage)
	a := uint8(float64(a1>>8)*(1-percentage) + float64(a2>>8)*percentage)
	return Color{Color: color.RGBA{R: r, G: g, B: b, A: a}}
}

// ColorRGBA returns a new color with the given red, green, blue, and alpha values.
func ColorRGBA(r, g, b, a uint8) Color {
	return Color{Color: color.RGBA{R: r, G: g, B: b, A: a}}
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
