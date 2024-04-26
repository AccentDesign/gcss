package border

type (
	Style string
)

const (
	StyleSolid  Style = "solid"
	StyleDashed Style = "dashed"
	StyleDotted Style = "dotted"
	StyleDouble Style = "double"
	StyleHidden Style = "hidden"
	StyleNone   Style = "none"
)

func (s Style) String() string {
	return string(s)
}
