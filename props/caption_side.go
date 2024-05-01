package props

type CaptionSide string

const (
	CaptionSideTop    CaptionSide = "top"
	CaptionSideBottom CaptionSide = "bottom"
)

func (c CaptionSide) String() string {
	return string(c)
}
