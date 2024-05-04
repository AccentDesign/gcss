package props

type FlexWrap string

const (
	FlexWrapNoWrap      FlexWrap = "nowrap"
	FlexWrapWrap        FlexWrap = "wrap"
	FlexWrapWrapReverse FlexWrap = "wrap-reverse"
)

func (f FlexWrap) String() string {
	return string(f)
}
