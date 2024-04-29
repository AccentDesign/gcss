package position

type Position string

const (
	Static   Position = "static"
	Relative Position = "relative"
	Absolute Position = "absolute"
	Fixed    Position = "fixed"
	Sticky   Position = "sticky"
)
