package props

import "fmt"

type (
	Border struct {
		Width Unit
		Style BorderStyle
		Color Color
	}
	BorderStyle string
)

const (
	BorderStyleSolid  BorderStyle = "solid"
	BorderStyleDashed BorderStyle = "dashed"
	BorderStyleDotted BorderStyle = "dotted"
	BorderStyleDouble BorderStyle = "double"
	BorderStyleHidden BorderStyle = "hidden"
	BorderStyleNone   BorderStyle = "none"
)

func (b Border) String() string {
	return fmt.Sprintf("%s %s %s", b.Width.String(), b.Style.String(), b.Color.String())
}

func (b BorderStyle) String() string {
	return string(b)
}
