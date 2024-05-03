package props

import "strings"

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

func BackgroundRepeats(repeats ...BackgroundRepeat) BackgroundRepeat {
	var r []string
	for _, repeat := range repeats {
		r = append(r, string(repeat))
	}
	return BackgroundRepeat(strings.Join(r, ","))
}
