package props

type TextWrap string

const (
	TextWrapWrap    TextWrap = "wrap"
	TextWrapNoWrap  TextWrap = "nowrap"
	TextWrapBalance TextWrap = "balance"
)

func (t TextWrap) String() string {
	return string(t)
}
