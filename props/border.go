package props

import (
	"strings"
)

type (
	Border struct {
		Width Unit
		Style BorderStyle
		Color Color
	}
)

func (b Border) String() string {
	var parts []string

	if b.Width != (Unit{}) {
		parts = append(parts, b.Width.String())
	}

	if b.Style != "" {
		parts = append(parts, b.Style.String())
	}

	if b.Color != (Color{}) {
		parts = append(parts, b.Color.String())
	}

	return strings.TrimSpace(strings.Join(parts, " "))
}
