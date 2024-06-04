package props

type AlignContent string

const (
	AlignContentNormal        AlignContent = "normal"
	AlignContentStart         AlignContent = "start"
	AlignContentCenter        AlignContent = "center"
	AlignContentEnd           AlignContent = "end"
	AlignContentFlexStart     AlignContent = "flex-start"
	AlignContentFlexEnd       AlignContent = "flex-end"
	AlignContentBaseline      AlignContent = "baseline"
	AlignContentFirstBaseline AlignContent = "first baseline"
	AlignContentLastBaseline  AlignContent = "last baseline"
	AlignContentSpaceBetween  AlignContent = "space-between"
	AlignContentSpaceAround   AlignContent = "space-around"
	AlignContentSpaceEvenly   AlignContent = "space-evenly"
	AlignContentStretch       AlignContent = "stretch"
)

func (a AlignContent) String() string {
	return string(a)
}
