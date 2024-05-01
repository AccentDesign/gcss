package props

type AlignItems string

const (
	AlignItemsStart    AlignItems = "flex-start"
	AlignItemsEnd      AlignItems = "flex-end"
	AlignItemsCenter   AlignItems = "center"
	AlignItemsBaseline AlignItems = "baseline"
	AlignItemsStretch  AlignItems = "stretch"
)

func (a AlignItems) String() string {
	return string(a)
}
