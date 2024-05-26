package props

type TextDecorationLine string

const (
	TextDecorationLineNone        TextDecorationLine = "none"
	TextDecorationLineUnderline   TextDecorationLine = "underline"
	TextDecorationLineOverline    TextDecorationLine = "overline"
	TextDecorationLineLineThrough TextDecorationLine = "line-through"
)

func (t TextDecorationLine) String() string {
	return string(t)
}
