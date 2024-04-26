package align

type Items string

const (
	ItemsStart    Items = "flex-start"
	ItemsEnd      Items = "flex-end"
	ItemsCenter   Items = "center"
	ItemsBaseline Items = "baseline"
	ItemsStretch  Items = "stretch"
)

func (a Items) String() string {
	return string(a)
}
