package props

type Position string

const (
	PositionStatic   Position = "static"
	PositionRelative Position = "relative"
	PositionAbsolute Position = "absolute"
	PositionFixed    Position = "fixed"
	PositionSticky   Position = "sticky"
)

func (p Position) String() string {
	return string(p)
}
