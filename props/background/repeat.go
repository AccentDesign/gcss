package background

type Repeat string

const (
	RepeatRepeatX  Repeat = "repeat-x"
	RepeatRepeatY  Repeat = "repeat-y"
	RepeatRepeat   Repeat = "repeat"
	RepeatSpace    Repeat = "space"
	RepeatRound    Repeat = "round"
	RepeatNoRepeat Repeat = "no-repeat"
)

func (b Repeat) String() string {
	return string(b)
}
