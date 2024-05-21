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
	case UnitTypePx, UnitTypePercent, UnitTypeRem, UnitTypeEm:
		return formatSize(u)
	case UnitTypeAuto:
		return "auto"
	case UnitTypeInherit:
		return "inherit"
	case UnitTypeInitial:
		return "initial"
	}
	return ""
}

func formatSize(u Unit) string {
	format := getFormat(u.Type)
	switch v := u.Size.(type) {
	case int:
		return fmt.Sprintf(format, v)
	case float64:
		return fmt.Sprintf(format, v)
	default:
		return ""
	}
}

func getFormat(unitType UnitType) string {
	switch unitType {
	case UnitTypePx:
		return "%dpx"
	case UnitTypePercent:
		return "%.2f%%"
	case UnitTypeRem:
		return "%.3frem"
	case UnitTypeEm:
		return "%.3fem"
	default:
		return ""
	}
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
