package props

type Overflow string

const (
	OverflowVisible Overflow = "visible"
	OverflowHidden  Overflow = "hidden"
	OverflowScroll  Overflow = "scroll"
	OverflowAuto    Overflow = "auto"
	OverflowClip    Overflow = "clip"
)

func (o Overflow) String() string {
	return string(o)
}
