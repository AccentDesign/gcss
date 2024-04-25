package props

type CaptionSide string

const (
	CaptionSideTop    CaptionSide = "top"
	CaptionSideBottom CaptionSide = "bottom"
)

func (a CaptionSide) String() string {
	return string(a)
}
