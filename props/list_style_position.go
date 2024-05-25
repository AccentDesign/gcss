package props

type ListStylePosition string

const (
	ListStylePositionInside  ListStylePosition = "inside"
	ListStylePositionOutside ListStylePosition = "outside"
)

func (l ListStylePosition) String() string {
	return string(l)
}
