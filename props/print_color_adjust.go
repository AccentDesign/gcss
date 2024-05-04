package props

type PrintColorAdjust string

const (
	PrintColorAdjustEconomy PrintColorAdjust = "economy"
	PrintColorAdjustExact   PrintColorAdjust = "exact"
)

func (c PrintColorAdjust) String() string {
	return string(c)
}
