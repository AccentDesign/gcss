package props

type AlignItems string

const (
	AlignItemsNormal    AlignItems = "normal"
	AlignItemsStretch   AlignItems = "stretch"
	AlignItemsStart     AlignItems = "start"
	AlignItemsCenter    AlignItems = "center"
	AlignItemsEnd       AlignItems = "end"
	AlignItemsFlexStart AlignItems = "flex-start"
	AlignItemsFlexEnd   AlignItems = "flex-end"
	AlignItemsSelfStart AlignItems = "self-start"
	AlignItemsSelfEnd   AlignItems = "self-end"
	AlignItemsBaseline  AlignItems = "baseline"
)

func (a AlignItems) String() string {
	return string(a)
}
