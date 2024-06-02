package props

type BorderCollapse string

const (
	BorderCollapseCollapse BorderCollapse = "collapse"
	BorderCollapseSeparate BorderCollapse = "separate"
)

func (b BorderCollapse) String() string {
	return string(b)
}
