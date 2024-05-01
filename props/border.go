package props

import (
	"fmt"
)

type (
	Border struct {
		Width Unit
		Style BorderStyle
		Color Color
	}
)

func (b Border) String() string {
	return fmt.Sprintf("%s %s %s", b.Width.String(), b.Style.String(), b.Color.String())
}
