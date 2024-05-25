package props

type Visibility string

const (
	VisibilityVisible  Visibility = "visible"
	VisibilityHidden   Visibility = "hidden"
	VisibilityCollapse Visibility = "collapse"
)

func (v Visibility) String() string {
	return string(v)
}
