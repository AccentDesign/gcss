package props

type FlexDirection string

const (
	FlexDirectionColumn        FlexDirection = "column"
	FlexDirectionColumnReverse FlexDirection = "column-reverse"
	FlexDirectionRow           FlexDirection = "row"
	FlexDirectionRowReverse    FlexDirection = "row-reverse"
)

func (f FlexDirection) String() string {
	return string(f)
}
