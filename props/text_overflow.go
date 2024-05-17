package props

type TextOverflow string

const (
	TextOverflowClip     TextOverflow = "clip"
	TextOverflowEllipsis TextOverflow = "ellipsis"
)

func (t TextOverflow) String() string {
	return string(t)
}
