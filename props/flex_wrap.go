package props

type FlexWrap string

const (
	FlexWrapNowrap  FlexWrap = "nowrap"
	FlexWrapWrap    FlexWrap = "wrap"
	FlexWrapWrapRev FlexWrap = "wrap-reverse"
)

func (f FlexWrap) String() string {
	return string(f)
}
