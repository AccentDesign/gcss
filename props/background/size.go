package background

import "github.com/AccentDesign/gostyle/props"

type Size struct {
	Value string
}

func (b Size) String() string {
	return b.Value
}

func SizeWidth(unit props.Unit) Size {
	return Size{unit.String()}
}

func SizeDimension(width props.Unit, height props.Unit) Size {
	return Size{width.String() + " " + height.String()}
}

func SizeMultiple(sizes ...Size) Size {
	value := ""
	for i, size := range sizes {
		if i > 0 {
			value += ", "
		}
		value += size.String()
	}
	return Size{value}
}

func SizeCover() Size {
	return Size{"cover"}
}

func SizeContain() Size {
	return Size{"contain"}
}
