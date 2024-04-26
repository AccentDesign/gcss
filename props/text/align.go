package text

type Align string

const (
	AlignLeft    Align = "left"
	AlignRight   Align = "right"
	AlignCenter  Align = "center"
	AlignJustify Align = "justify"
	AlignStart   Align = "start"
	AlignEnd     Align = "end"
)

func (a Align) String() string {
	return string(a)
}
