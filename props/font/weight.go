package font

type Weight string

const (
	WeightThin       Weight = "100"
	WeightExtraLight Weight = "200"
	WeightLight      Weight = "300"
	WeightNormal     Weight = "400"
	WeightMedium     Weight = "500"
	WeightSemiBold   Weight = "600"
	WeightBold       Weight = "700"
	WeightExtraBold  Weight = "800"
	WeightBlack      Weight = "900"
)

func (a Weight) String() string {
	return string(a)
}
