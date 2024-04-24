package props

type (
	Margin interface {
		String() string
	}
	MarginKeyword string
)

const (
	MarginAuto MarginKeyword = "auto"
)

func (m MarginKeyword) String() string {
	return string(m)
}
