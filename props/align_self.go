package props

type AlignSelf string

const (
	AlignSelfAuto      AlignSelf = "auto"
	AlignSelfNormal    AlignSelf = "normal"
	AlignSelfStretch   AlignSelf = "stretch"
	AlignSelfCenter    AlignSelf = "center"
	AlignSelfStart     AlignSelf = "start"
	AlignSelfEnd       AlignSelf = "end"
	AlignSelfFlexStart AlignSelf = "flex-start"
	AlignSelfFlexEnd   AlignSelf = "flex-end"
	AlignSelfSelfStart AlignSelf = "self-start"
	AlignSelfSelfEnd   AlignSelf = "self-end"
	AlignSelfBaseline  AlignSelf = "baseline"
)

func (a AlignSelf) String() string {
	return string(a)
}
