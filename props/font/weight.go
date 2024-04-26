package font

type Weight string

const (
	Weight100 Weight = "100"
	Weight200 Weight = "200"
	Weight300 Weight = "300"
	Weight400 Weight = "400"
	Weight500 Weight = "500"
	Weight600 Weight = "600"
	Weight700 Weight = "700"
	Weight800 Weight = "800"
	Weight900 Weight = "900"
)

func (a Weight) String() string {
	return string(a)
}
