package props

type (
	Width interface {
		String() string
	}
	WidthKeyword string
)

const (
	WidthAuto       WidthKeyword = "auto"
	WidthFitContent WidthKeyword = "fit-content"
)

func (w WidthKeyword) String() string {
	return string(w)
}
