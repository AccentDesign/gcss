package props

type BackgroundRepeat string

const (
	BackgroundRepeatRepeatX  BackgroundRepeat = "repeat-x"
	BackgroundRepeatRepeatY  BackgroundRepeat = "repeat-y"
	BackgroundRepeatRepeat   BackgroundRepeat = "repeat"
	BackgroundRepeatSpace    BackgroundRepeat = "space"
	BackgroundRepeatRound    BackgroundRepeat = "round"
	BackgroundRepeatNoRepeat BackgroundRepeat = "no-repeat"
)

func (b BackgroundRepeat) String() string {
	return string(b)
}
