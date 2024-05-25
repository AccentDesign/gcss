package props

type VerticalAlign string

const (
	VerticalAlignBaseline   VerticalAlign = "baseline"
	VerticalAlignSub        VerticalAlign = "sub"
	VerticalAlignSuper      VerticalAlign = "super"
	VerticalAlignTextTop    VerticalAlign = "text-top"
	VerticalAlignTextBottom VerticalAlign = "text-bottom"
	VerticalAlignMiddle     VerticalAlign = "middle"
	VerticalAlignTop        VerticalAlign = "top"
	VerticalAlignBottom     VerticalAlign = "bottom"
)

func (v VerticalAlign) String() string {
	return string(v)
}
