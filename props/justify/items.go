package justify

type Items string

const (
	ItemsStretch Items = "stretch"
	ItemsCenter  Items = "center"
	ItemsStart   Items = "start"
	ItemsEnd     Items = "end"
)

func (i Items) String() string {
	return string(i)
}
