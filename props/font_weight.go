package props

type FontWeight string

const (
	FontWeight100 FontWeight = "100"
	FontWeight200 FontWeight = "200"
	FontWeight300 FontWeight = "300"
	FontWeight400 FontWeight = "400"
	FontWeight500 FontWeight = "500"
	FontWeight600 FontWeight = "600"
	FontWeight700 FontWeight = "700"
	FontWeight800 FontWeight = "800"
	FontWeight900 FontWeight = "900"
)

func (a FontWeight) String() string {
	return string(a)
}
