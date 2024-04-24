package props

type (
	Width interface {
		String() string
	}
	WidthKeyword string
)

const (
	WidthAuto WidthKeyword = "auto"
)

func (w WidthKeyword) String() string {
	return string(w)
}
