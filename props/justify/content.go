package justify

type Content string

const (
	ContentNormal       Content = "normal"
	ContentFlexStart    Content = "flex-start"
	ContentFlexEnd      Content = "flex-end"
	ContentCenter       Content = "center"
	ContentSpaceBetween Content = "space-between"
	ContentSpaceAround  Content = "space-around"
	ContentSpaceEvenly  Content = "space-evenly"
	ContentStretch      Content = "stretch"
)

func (j Content) String() string {
	return string(j)
}
