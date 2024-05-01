package props

type BackgroundSize struct {
	Value string
}

func (b BackgroundSize) String() string {
	return b.Value
}

func BackgroundSizeWidth(unit Unit) BackgroundSize {
	return BackgroundSize{unit.String()}
}

func BackgroundSizeDimension(width Unit, height Unit) BackgroundSize {
	return BackgroundSize{width.String() + " " + height.String()}
}

func BackgroundSizeMultiple(sizes ...BackgroundSize) BackgroundSize {
	value := ""
	for i, size := range sizes {
		if i > 0 {
			value += ", "
		}
		value += size.String()
	}
	return BackgroundSize{value}
}

func BackgroundSizeCover() BackgroundSize {
	return BackgroundSize{"cover"}
}

func BackgroundSizeContain() BackgroundSize {
	return BackgroundSize{"contain"}
}
