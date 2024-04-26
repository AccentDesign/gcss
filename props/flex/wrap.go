package flex

type Wrap string

const (
	WrapNowrap  Wrap = "nowrap"
	WrapWrap    Wrap = "wrap"
	WrapWrapRev Wrap = "wrap-reverse"
)

func (w Wrap) String() string {
	return string(w)
}
