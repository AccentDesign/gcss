package props

type JustifyItems string

const (
	JustifyItemsNormal    JustifyItems = "normal"
	JustifyItemsStretch   JustifyItems = "stretch"
	JustifyItemsCenter    JustifyItems = "center"
	JustifyItemsStart     JustifyItems = "start"
	JustifyItemsEnd       JustifyItems = "end"
	JustifyItemsFlexStart JustifyItems = "flex-start"
	JustifyItemsFlexEnd   JustifyItems = "flex-end"
	JustifyItemsSelfStart JustifyItems = "self-start"
	JustifyItemsSelfEnd   JustifyItems = "self-end"
	JustifyItemsLeft      JustifyItems = "left"
	JustifyItemsRight     JustifyItems = "right"
	JustifyItemsBaseline  JustifyItems = "baseline"
)

func (j JustifyItems) String() string {
	return string(j)
}
