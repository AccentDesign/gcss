package flex

type Direction string

const (
	DirectionColumn        Direction = "column"
	DirectionColumnReverse Direction = "column-reverse"
	DirectionRow           Direction = "row"
	DirectionRowReverse    Direction = "row-reverse"
)

func (f Direction) String() string {
	return string(f)
}
