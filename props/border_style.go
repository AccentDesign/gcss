package props

type BorderStyle string

const (
	BorderStyleSolid  BorderStyle = "solid"
	BorderStyleDashed BorderStyle = "dashed"
	BorderStyleDotted BorderStyle = "dotted"
	BorderStyleDouble BorderStyle = "double"
	BorderStyleHidden BorderStyle = "hidden"
	BorderStyleNone   BorderStyle = "none"
)

func (b BorderStyle) String() string {
	return string(b)
}
