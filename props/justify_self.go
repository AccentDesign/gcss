package props

type JustifySelf string

const (
	JustifySelfAuto      JustifySelf = "auto"
	JustifySelfNormal    JustifySelf = "normal"
	JustifySelfStretch   JustifySelf = "stretch"
	JustifySelfCenter    JustifySelf = "center"
	JustifySelfStart     JustifySelf = "start"
	JustifySelfEnd       JustifySelf = "end"
	JustifySelfFlexStart JustifySelf = "flex-start"
	JustifySelfFlexEnd   JustifySelf = "flex-end"
	JustifySelfSelfStart JustifySelf = "self-start"
	JustifySelfSelfEnd   JustifySelf = "self-end"
	JustifySelfLeft      JustifySelf = "left"
	JustifySelfRight     JustifySelf = "right"
	JustifySelfBaseline  JustifySelf = "baseline"
)

func (j JustifySelf) String() string {
	return string(j)
}
