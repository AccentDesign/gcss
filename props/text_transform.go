package props

type TextTransform string

const (
	TextTransformNone       TextTransform = "none"
	TextTransformCapitalize TextTransform = "capitalize"
	TextTransformUppercase  TextTransform = "uppercase"
	TextTransformLowercase  TextTransform = "lowercase"
)

func (t TextTransform) String() string {
	return string(t)
}
