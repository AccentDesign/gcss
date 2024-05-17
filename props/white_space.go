package props

type WhiteSpace string

const (
	WhiteSpaceNormal  WhiteSpace = "normal"
	WhiteSpaceNowrap  WhiteSpace = "nowrap"
	WhiteSpacePre     WhiteSpace = "pre"
	WhiteSpacePreLine WhiteSpace = "pre-line"
	WhiteSpacePreWrap WhiteSpace = "pre-wrap"
)

func (w WhiteSpace) String() string {
	return string(w)
}
