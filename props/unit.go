package props

import "fmt"

type (
	UnitType int
	Unit     struct {
		Size interface{}
		Type UnitType
	}
)

const (
	_ UnitType = iota
	UnitTypeRaw
	UnitTypePx
	UnitTypePercent
	UnitTypeRem
	UnitTypeEm
	UnitTypeAuto
	UnitTypeInherit
	UnitTypeInitial
)

func (u Unit) String() string {
	switch u.Type {
	case UnitTypeRaw:
		return fmt.Sprintf("%v", u.Size)
	case UnitTypePx:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%dpx", v)
		case float64:
			return fmt.Sprintf("%.0fpx", v)
		}
	case UnitTypePercent:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%d%%", v)
		case float64:
			return fmt.Sprintf("%.2f%%", v)
		}
	case UnitTypeRem:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%drem", v)
		case float64:
			return fmt.Sprintf("%.3frem", v)
		}
	case UnitTypeEm:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%dem", v)
		case float64:
			return fmt.Sprintf("%.3fem", v)
		}
	case UnitTypeAuto:
		return "auto"
	case UnitTypeInherit:
		return "inherit"
	case UnitTypeInitial:
		return "initial"
	}
	return ""
}

func UnitRaw(size interface{}) Unit {
	return Unit{Size: size, Type: UnitTypeRaw}
}

func UnitPx(size int) Unit {
	return Unit{Size: size, Type: UnitTypePx}
}

func UnitPercent(size float64) Unit {
	return Unit{Size: size, Type: UnitTypePercent}
}

func UnitRem(size float64) Unit {
	return Unit{Size: size, Type: UnitTypeRem}
}

func UnitEm(size float64) Unit {
	return Unit{Size: size, Type: UnitTypeEm}
}

func UnitAuto() Unit {
	return Unit{Type: UnitTypeAuto}
}

func UnitInherit() Unit {
	return Unit{Type: UnitTypeInherit}
}

func UnitInitial() Unit {
	return Unit{Type: UnitTypeInitial}
}
