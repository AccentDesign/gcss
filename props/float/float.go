package float

type Float string

const (
	FloatNone        Float = "none"
	FloatLeft        Float = "left"
	FloatRight       Float = "right"
	FloatInlineStart Float = "inline-start"
	FloatInlineEnd   Float = "inline-end"
)

func (f Float) String() string {
	return string(f)
}
