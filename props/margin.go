package props

import "github.com/AccentDesign/gostyle/props/unit"

type (
	Margin struct {
		Keyword MarginKeyword
		Unit    unit.Unit
	}
	MarginKeyword string
)

const (
	MarginAuto MarginKeyword = "auto"
)

func (m Margin) String() string {
	if m.Keyword != "" {
		return string(m.Keyword)
	}
	return m.Unit.String()
}
