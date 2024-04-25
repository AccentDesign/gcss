package props

import "fmt"

type (
	UnitType int
	Unit     struct {
		Size     float64
		UnitType UnitType
	}
)

const (
	_ UnitType = iota
	UnitNone
	UnitPx
	UnitPercent
	UnitRem
)

func (u Unit) String() string {
	switch u.UnitType {
	case UnitNone:
		return fmt.Sprintf("%.0f", u.Size)
	case UnitPx:
		return fmt.Sprintf("%.0fpx", u.Size)
	case UnitPercent:
		return fmt.Sprintf("%.2f%%", u.Size)
	case UnitRem:
		return fmt.Sprintf("%.3frem", u.Size)
	}
	return ""
}
