package props

type TextAlign string

const (
	TextAlignLeft    TextAlign = "left"
	TextAlignRight   TextAlign = "right"
	TextAlignCenter  TextAlign = "center"
	TextAlignJustify TextAlign = "justify"
	TextAlignStart   TextAlign = "start"
	TextAlignEnd     TextAlign = "end"
)

func (t TextAlign) String() string {
	return string(t)
}
