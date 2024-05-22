package props

type BoxSizing string

const (
	BoxSizingBorderBox  BoxSizing = "border-box"
	BoxSizingContentBox BoxSizing = "content-box"
)

func (b BoxSizing) String() string {
	return string(b)
}
