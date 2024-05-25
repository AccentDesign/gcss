package props

type ListStyleType string

const (
	ListStyleTypeNone    ListStyleType = "none"
	ListStyleTypeDisc    ListStyleType = "disc"
	ListStyleTypeDecimal ListStyleType = "decimal"
)

func (l ListStyleType) String() string {
	return string(l)
}
