package justify

type Self string

const (
	SelfStretch Self = "stretch"
	SelfCenter  Self = "center"
	SelfStart   Self = "start"
	SelfEnd     Self = "end"
)

func (s Self) String() string {
	return string(s)
}
