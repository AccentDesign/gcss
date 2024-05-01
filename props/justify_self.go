package props

type JustifySelf string

const (
	JustifySelfStretch JustifySelf = "stretch"
	JustifySelfCenter  JustifySelf = "center"
	JustifySelfStart   JustifySelf = "start"
	JustifySelfEnd     JustifySelf = "end"
)

func (j JustifySelf) String() string {
	return string(j)
}
