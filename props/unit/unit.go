package unit

import "fmt"

type (
	Type int
	Unit struct {
		Size float64
		Type Type
	}
)

const (
	_ Type = iota
	TypeNone
	TypePx
	TypePercent
	TypeRem
	TypeEm
	TypeAuto
)

func (u Unit) String() string {
	switch u.Type {
	case TypeNone:
		return fmt.Sprintf("%.0f", u.Size)
	case TypePx:
		return fmt.Sprintf("%.0fpx", u.Size)
	case TypePercent:
		return fmt.Sprintf("%.2f%%", u.Size)
	case TypeRem:
		return fmt.Sprintf("%.3frem", u.Size)
	case TypeEm:
		return fmt.Sprintf("%.3fem", u.Size)
	case TypeAuto:
		return "auto"
	}
	return ""
}

func None(size float64) Unit {
	return Unit{Size: size, Type: TypeNone}
}

func Px(size float64) Unit {
	return Unit{Size: size, Type: TypePx}
}

func Percent(size float64) Unit {
	return Unit{Size: size, Type: TypePercent}
}

func Rem(size float64) Unit {
	return Unit{Size: size, Type: TypeRem}
}

func Em(size float64) Unit {
	return Unit{Size: size, Type: TypeEm}
}

func Auto() Unit {
	return Unit{Type: TypeAuto}
}
