package props

import (
	"image/color"
	"testing"
)

func TestColor_String(t *testing.T) {
	testCases := map[Color]string{
		Color{Keyword: "currentColor"}:     "currentColor",
		Color{Keyword: "inherit"}:          "inherit",
		Color{Color: color.Gray{0}}:        "rgba(0,0,0,1.00)",
		Color{Color: color.Gray{255}}:      "rgba(255,255,255,1.00)",
		Color{Color: color.Gray16{0}}:      "rgba(0,0,0,1.00)",
		Color{Color: color.Gray16{0xffff}}: "rgba(255,255,255,1.00)",
		ColorRGBA(62, 131, 248, 255):       "rgba(62,131,248,1.00)",
		ColorRGBA(0, 0, 0, 0):              "rgba(0,0,0,0.00)",
		ColorRGBA(255, 255, 255, 255):      "rgba(255,255,255,1.00)",
	}
	for c, expected := range testCases {
		if c.String() != expected {
			t.Errorf("expected %v, got %v", expected, c.String())
		}
	}
}

func TestColor_Alpha(t *testing.T) {
	c := ColorRGBA(62, 131, 248, 255)
	expected := ColorRGBA(62, 131, 248, 128)
	if c.Alpha(128) != expected {
		t.Errorf("expected %v, got %v", expected, c.Alpha(128))
	}
}

func TestColor_MixToWhite(t *testing.T) {
	base := ColorRGBA(62, 131, 248, 255)
	with := ColorRGBA(255, 255, 255, 255)
	testCases := map[float64]Color{
		-1:  ColorRGBA(62, 131, 248, 255),
		0:   ColorRGBA(62, 131, 248, 255),
		0.1: ColorRGBA(81, 143, 248, 255),
		0.2: ColorRGBA(100, 155, 249, 255),
		0.3: ColorRGBA(119, 168, 250, 255),
		0.4: ColorRGBA(139, 180, 250, 255),
		0.5: ColorRGBA(158, 193, 251, 255),
		0.6: ColorRGBA(177, 205, 252, 255),
		0.7: ColorRGBA(197, 217, 252, 255),
		0.8: ColorRGBA(216, 230, 253, 255),
		0.9: ColorRGBA(235, 242, 254, 255),
		1:   ColorRGBA(255, 255, 255, 255),
		1.1: ColorRGBA(255, 255, 255, 255),
	}
	for percentage, expected := range testCases {
		mixed := base.Mix(with, percentage)
		if mixed != expected {
			t.Errorf("%v expected %v, got %v", percentage, expected, mixed)
		}
	}
}

func TestColor_MixToBlack(t *testing.T) {
	base := ColorRGBA(62, 131, 248, 255)
	with := ColorRGBA(0, 0, 0, 255)
	testCases := map[float64]Color{
		-1:  ColorRGBA(62, 131, 248, 255),
		0:   ColorRGBA(62, 131, 248, 255),
		0.1: ColorRGBA(55, 117, 223, 255),
		0.2: ColorRGBA(49, 104, 198, 255),
		0.3: ColorRGBA(43, 91, 173, 255),
		0.4: ColorRGBA(37, 78, 148, 255),
		0.5: ColorRGBA(31, 65, 124, 255),
		0.6: ColorRGBA(24, 52, 99, 255),
		0.7: ColorRGBA(18, 39, 74, 255),
		0.8: ColorRGBA(12, 26, 49, 255),
		0.9: ColorRGBA(6, 13, 24, 255),
		1:   ColorRGBA(0, 0, 0, 255),
		1.1: ColorRGBA(0, 0, 0, 255),
	}
	for percentage, expected := range testCases {
		mixed := base.Mix(with, percentage)
		if mixed != expected {
			t.Errorf("%v expected %v, got %v", percentage, expected, mixed)
		}
	}
}
