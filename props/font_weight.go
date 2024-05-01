package props

type FontWeight string

const (
	FontWeightThin       FontWeight = "100"
	FontWeightExtraLight FontWeight = "200"
	FontWeightLight      FontWeight = "300"
	FontWeightNormal     FontWeight = "400"
	FontWeightMedium     FontWeight = "500"
	FontWeightSemiBold   FontWeight = "600"
	FontWeightBold       FontWeight = "700"
	FontWeightExtraBold  FontWeight = "800"
	FontWeightBlack      FontWeight = "900"
)

func (f FontWeight) String() string {
	return string(f)
}
