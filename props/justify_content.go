package props

type JustifyContent string

const (
	JustifyContentNormal       JustifyContent = "normal"
	JustifyContentFlexStart    JustifyContent = "flex-start"
	JustifyContentFlexEnd      JustifyContent = "flex-end"
	JustifyContentCenter       JustifyContent = "center"
	JustifyContentSpaceBetween JustifyContent = "space-between"
	JustifyContentSpaceAround  JustifyContent = "space-around"
	JustifyContentSpaceEvenly  JustifyContent = "space-evenly"
	JustifyContentStretch      JustifyContent = "stretch"
)

func (j JustifyContent) String() string {
	return string(j)
}
