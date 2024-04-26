package props

import "github.com/AccentDesign/gostyle/props/unit"

type (
	Width struct {
		Keyword WidthKeyword
		Unit    unit.Unit
	}
	WidthKeyword string
)

const (
	WidthAuto       WidthKeyword = "auto"
	WidthFitContent WidthKeyword = "fit-content"
)

func (w Width) String() string {
	if w.Keyword != "" {
		return string(w.Keyword)
	}
	return w.Unit.String()
}
