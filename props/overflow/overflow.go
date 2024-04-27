package overflow

type Overflow string

const (
	Visible Overflow = "visible"
	Hidden  Overflow = "hidden"
	Scroll  Overflow = "scroll"
	Auto    Overflow = "auto"
	Clip    Overflow = "clip"
)

func (o Overflow) String() string {
	return string(o)
}
