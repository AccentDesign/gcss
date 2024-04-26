package border

import (
	"fmt"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/unit"
)

type (
	Border struct {
		Width unit.Unit
		Style Style
		Color colors.Color
	}
)

func (b Border) String() string {
	return fmt.Sprintf("%s %s %s", b.Width.String(), b.Style.String(), b.Color.String())
}
