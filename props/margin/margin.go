package margin

import "github.com/AccentDesign/gostyle/props/unit"

type (
	Margin struct {
		Keyword string
		Unit    unit.Unit
	}
)

func (m Margin) String() string {
	if m.Keyword != "" {
		return m.Keyword
	}
	return m.Unit.String()
}

func Auto() Margin {
	return Margin{Keyword: "auto"}
}
