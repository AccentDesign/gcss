package props

type FontStyle string

const (
	FontStyleNormal FontStyle = "normal"
	FontStyleItalic FontStyle = "italic"
)

func (f FontStyle) String() string {
	return string(f)
}
