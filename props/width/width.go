package width

import "github.com/AccentDesign/gostyle/props/unit"

type (
	Width struct {
		Keyword string
		Unit    unit.Unit
	}
)

func (w Width) String() string {
	if w.Keyword != "" {
		return w.Keyword
	}
	return w.Unit.String()
}

func Auto() Width {
	return Width{Keyword: "auto"}
}

func FitContent() Width {
	return Width{Keyword: "fit-content"}
}
