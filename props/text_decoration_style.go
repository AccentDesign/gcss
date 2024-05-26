package props

type TextDecorationStyle string

const (
	TextDecorationStyleSolid  TextDecorationStyle = "solid"
	TextDecorationStyleDouble TextDecorationStyle = "double"
	TextDecorationStyleDotted TextDecorationStyle = "dotted"
	TextDecorationStyleDashed TextDecorationStyle = "dashed"
	TextDecorationStyleWavy   TextDecorationStyle = "wavy"
)

func (t TextDecorationStyle) String() string {
	return string(t)
}
