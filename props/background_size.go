package props

import "strings"

type BackgroundSize string

func (b BackgroundSize) String() string {
	return string(b)
}

func BackgroundSizeWidth(unit Unit) BackgroundSize {
	return BackgroundSize(unit.String())
}

func BackgroundSizeDimension(width, height Unit) BackgroundSize {
	return BackgroundSize(width.String() + " " + height.String())
}

func BackgroundSizes(sizes ...BackgroundSize) BackgroundSize {
	var s []string
	for _, size := range sizes {
		s = append(s, string(size))
	}
	return BackgroundSize(strings.Join(s, ","))
}

func BackgroundSizeCover() BackgroundSize {
	return BackgroundSize("cover")
}

func BackgroundSizeContain() BackgroundSize {
	return BackgroundSize("contain")
}
