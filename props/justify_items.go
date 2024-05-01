package props

type JustifyItems string

const (
	JustifyItemsStretch JustifyItems = "stretch"
	JustifyItemsCenter  JustifyItems = "center"
	JustifyItemsStart   JustifyItems = "start"
	JustifyItemsEnd     JustifyItems = "end"
)

func (j JustifyItems) String() string {
	return string(j)
}
