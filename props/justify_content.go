package props

type JustifyContent string

const (
	JustifyContentNormal       JustifyContent = "normal"
	JustifyContentCenter       JustifyContent = "center"
	JustifyContentStart        JustifyContent = "start"
	JustifyContentEnd          JustifyContent = "end"
	JustifyContentFlexStart    JustifyContent = "flex-start"
	JustifyContentFlexEnd      JustifyContent = "flex-end"
	JustifyContentLeft         JustifyContent = "left"
	JustifyContentRight        JustifyContent = "right"
	JustifyContentSpaceBetween JustifyContent = "space-between"
	JustifyContentSpaceAround  JustifyContent = "space-around"
	JustifyContentSpaceEvenly  JustifyContent = "space-evenly"
	JustifyContentStretch      JustifyContent = "stretch"
)

func (j JustifyContent) String() string {
	return string(j)
}
